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
	"sync"
	"time"
)

var (
	TaskMap sync.Map
)

func Start(task agentModel.Task, filePath string, ch chan int) (status consts.TaskStatus, existFile string) {
	fmt.Printf("Start to download %s ...\n", task.Url)

	startTime := time.Now()
	task.StartTime = &startTime

	targetDir := consts.DownloadDir
	if task.Md5 == "" {
		getMd5FromRemote(&task, targetDir)
	}

	existFile = findSameFile(task, targetDir)
	if existFile != "" {
		status = consts.Completed
		return
	}

	SaveTaskStatus(&TaskMap, task.ID, 0, 0)

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
						rate := resp.Progress()
						speed := GetSpeed(*task.StartTime, resp.BytesComplete()/1000)
						SaveTaskStatus(&TaskMap, task.ID, rate, speed)

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
					speed := GetSpeed(*task.StartTime, resp.BytesComplete()/1000)
					SaveTaskStatus(&TaskMap, task.ID, rate, speed)

					fmt.Printf("Downloading %s %d / %d bytes (%d%%)\u001B[K\n", resp.Filename, resp.BytesComplete(), resp.Size(), int(100*rate))
				}
			}
		}
	}

ExitDownload:

	t.Stop()

	if isCanceled {
		if len(responses) > 0 {
			responses[0] = nil
		}

		completed++

		status = consts.Canceled
		fmt.Printf("Force to terminate download %s.\n", task.Url)
	} else {
		task.Path = filePath

		if checkMd5(task) {
			status = consts.Completed

			if task.Md5 != "" {
				saveMd5FromRequest(&task, targetDir)
			}

			fmt.Printf("Successfully download %s to %s.\n", task.Url, task.Path)
		} else {
			status = consts.Error
			fmt.Printf("Failed to download %s.\n", task.Url)
		}
	}

	task.CompletionRate, task.Speed = GetTaskStatus(TaskMap, task.ID)

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

func GetSpeed(startTime time.Time, sumKByte int64) (ret float64) {
	sec := time.Now().Unix() - startTime.Unix()

	ret = float64(sumKByte) / float64(sec)

	return
}

func SaveTaskStatus(cache *sync.Map, id uint, rate, speed float64) {
	valObj, ok := cache.Load(id)

	if !ok {
		valObj = map[string]float64{
			"rate":  rate,
			"speed": speed,
		}
		cache.Store(id, valObj)

		return
	}

	valMap := valObj.(map[string]float64)

	if rate > valMap["rate"] {
		valMap["rate"] = rate
	}
	if speed > 0 {
		valMap["speed"] = speed
	}

	cache.Store(id, valMap)

	return
}

func GetTaskStatus(cache sync.Map, id uint) (rate, speed float64) {
	valObj, ok := cache.Load(id)

	if !ok {
		return
	}

	valMap := valObj.(map[string]float64)

	rate = valMap["rate"]
	speed = valMap["speed"]

	return
}
