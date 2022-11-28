package job

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
)

func Start(downloadTask *agentModel.Task, filePath string, ch chan int) (status consts.TaskStatus, existFile string) {
	fmt.Printf("Start to download %s ...\n", downloadTask.Url)

	startTime := time.Now()
	downloadTask.StartDate = &startTime

	targetDir := consts.DownloadDir
	if downloadTask.Md5 == "" {
		getMd5FromRemote(downloadTask, targetDir)
	}

	existFile = findSameFile(*downloadTask, targetDir)
	if existFile != "" {
		status = consts.Completed
		return
	}

	SaveTaskStatus(&TaskStatus, downloadTask.ID, 0, 0)

	// start file downloads, 3 at a time
	respCh, err := grab.GetBatch(3, targetDir, downloadTask.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

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

		default:
		}
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
					rate := resp.Progress()
					speed := GetSpeed(*downloadTask.StartDate, resp.BytesComplete()/1000)
					SaveTaskStatus(&TaskStatus, downloadTask.ID, rate, speed)

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

				rate := resp.Progress()
				speed := GetSpeed(*downloadTask.StartDate, resp.BytesComplete()/1000)
				SaveTaskStatus(&TaskStatus, downloadTask.ID, rate, speed)

				fmt.Printf("Downloading %s %d / %d bytes %f (%d%%)\u001B[K\n", resp.Filename, resp.BytesComplete(), resp.Size(), speed, int(100*rate))
			}
		}

		time.Sleep(500 * time.Millisecond)
	}

ExitDownload:

	if isCanceled {
		for index, resp := range responses {
			resp.Cancel()
			responses[index] = nil
		}

		completed++

		status = consts.Canceled
		fmt.Printf("Force to terminate download %s.\n", downloadTask.Url)
	} else {
		downloadTask.Path = filePath

		if checkMd5(*downloadTask) {
			status = consts.Completed

			if downloadTask.Md5 != "" {
				saveMd5FromRequest(downloadTask, targetDir)
			}

			fmt.Printf("Successfully download %s to %s.\n", downloadTask.Url, downloadTask.Path)
		} else {
			status = consts.Error
			fmt.Printf("Failed to download %s.\n", downloadTask.Url)
		}
	}

	downloadTask.Rate, downloadTask.Speed = GetTaskStatus(TaskStatus, downloadTask.ID)

	return
}

func checkMd5(task agentModel.Task) bool {
	expectVal := task.Md5

	if expectVal == "" {
		return true
	}

	actualVal, _ := _fileUtils.GetMd5(task.Path)

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

func saveMd5FromRequest(task *agentModel.Task, dir string) (err error) {
	index2 := strings.LastIndex(task.Url, "/")
	md5FilePath := filepath.Join(dir, task.Url[index2:]+".md5")

	_fileUtils.WriteFile(md5FilePath, task.Md5)

	return
}

func findSameFile(task agentModel.Task, dir string) (existFile string) {
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
			existFile = strings.Replace(md5FilePath, ".md5", "", -1)

			if _fileUtils.FileExist(existFile) {
				return
			} else {
				existFile = ""
			}
		}
	}

	return
}

func GetPath(task agentModel.Task) (pth string) {
	index := strings.LastIndex(task.Url, "/")
	name := task.Url[index:]

	pth = filepath.Join(consts.DownloadDir, name)

	return
}
