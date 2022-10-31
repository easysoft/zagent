package downloadUtils

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Start(task agentModel.Task, ch chan int) (pth string, status consts.TaskStatus) {
	fmt.Printf("Start to download %s ...\n", task.Url)

	targetDir := consts.FolderDownload
	if task.Md5 == "" {
		getMd5FromRemote(&task, targetDir)
	}

	pth = findSameFile(task, targetDir)
	if pth != "" {
		status = consts.Completed
		return
	}

	// start file downloads, 3 at a time
	respCh, err := grab.GetBatch(3, targetDir, task.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// start a ticker to update progress every 200ms
	t := time.NewTicker(200 * time.Millisecond)

	// monitor downloads
	completed := 0
	inProgress := 0
	responses := make([]*grab.Response, 0)

	isCanceled := false

	for completed < 1 {
		select {
		case <-ch:
			isCanceled = true
			goto ExitDownload
		default:
		}

		select {
		case resp := <-respCh:
			// a new response has been received and has started downloading
			// (nil is received once, when the channel is closed by grab)
			if resp != nil {
				responses = append(responses, resp)
			}

		case <-t.C:
			// clear lines
			if inProgress > 0 {
				fmt.Printf("\033[%dA\033[K", inProgress)
			}

			// update completed downloads
			for i, resp := range responses {
				if resp != nil && resp.IsComplete() {
					// print final result
					if resp.Err() != nil && resp.HTTPResponse.StatusCode != 416 {
						fmt.Fprintf(os.Stderr, "Error download %s: %v\n", resp.Request.URL(), resp.Err())
					} else {
						fmt.Printf("Finish %s %d / %d bytes (%d%%)\n", resp.Filename, resp.BytesComplete(), resp.Size(), int(100*resp.Progress()))
					}

					// mark completed
					responses[i] = nil
					completed++
				}
			}

			// update downloads in progress
			inProgress = 0
			for _, resp := range responses {
				if resp != nil {
					inProgress++
					fmt.Printf("Downloading %s %d / %d bytes (%d%%)\u001B[K\n", resp.Filename, resp.BytesComplete(), resp.Size(), int(100*resp.Progress()))
				}
			}
		}
	}

ExitDownload:

	t.Stop()

	if isCanceled {
		status = consts.Canceled
		fmt.Printf("Force to terminate download %s.\n", task.Url)
	} else {
		if checkMd5(task) {
			status = consts.Completed
			pth = responses[0].Filename

			fmt.Printf("Successfully download %s to %s.\n", task.Url, pth)
		} else {
			status = consts.Failed
			fmt.Printf("Failed to download %s.\n", task.Url)
		}
	}

	return
}

func checkMd5(task agentModel.Task) bool {
	expectVal := task.Md5

	cmdStr := ""
	if _commonUtils.IsWin() {
		cmdStr = "CertUtil -hashfile " + task.Path + " MD5"
	} else {
		cmdStr = "md5sum " + task.Path + " | awk '{print $1}'"
	}
	actualVal, _ := _shellUtils.ExeSysCmd(cmdStr)

	if _commonUtils.IsWin() {
		arr := strings.Split(actualVal, "\n")
		if len(arr) > 1 {
			actualVal = strings.TrimSpace(strings.Split(actualVal, "\n")[1])
		}
	}

	pass := strings.TrimSpace(actualVal) == strings.TrimSpace(expectVal)

	return pass
}

func getMd5FromRemote(task *agentModel.Task, dir string) (err error) {
	index := strings.LastIndex(task.Url, ".")
	md5FileUrl := task.Url[:index] + ".md5"

	index2 := strings.LastIndex(task.Url, "/")
	md5FilePath := filepath.Join(dir, task.Url[index2:]+".md5")

	err = _fileUtils.Download(md5FileUrl, md5FilePath)
	if err != nil {
		return
	}

	task.Md5 = _fileUtils.ReadFile(md5FilePath)

	return
}

func findSameFile(task agentModel.Task, dir string) (pth string) {
	files, _ := ioutil.ReadDir(dir)

	for _, fi := range files {
		name := fi.Name()
		extName := _fileUtils.GetExtName(fi.Name())
		if extName != ".md5" {
			continue
		}

		md5FilePath := filepath.Join(dir, name)
		md5 := _fileUtils.ReadFile(md5FilePath)

		if md5 == task.Md5 {
			pth = md5FilePath
			return
		}
	}

	return
}
