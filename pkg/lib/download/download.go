package downloadUtils

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"os"
	"sync"
	"time"
)

const (
	keyNotStart   = "not_start"
	keyInProgress = "in_progress"
	keyCompleted  = "completed"
)

var (
	syncMap sync.Map
)

func Download(task DownloadTask) {
	fmt.Printf("Start to download %s ...\n", task.Url)

	// start file downloads, 3 at a time
	respCh, err := grab.GetBatch(3, ".", task.Url)
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

	for completed < 1 {
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
						fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n", resp.Request.URL(), resp.Err())
					} else {
						fmt.Printf("Finished %s %d / %d bytes (%d%%)\n", resp.Filename, resp.BytesComplete(), resp.Size(), int(100*resp.Progress()))
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

	t.Stop()

	fmt.Printf("Success to download %s.\n", task.Url)
}

func InitTasks() {
	syncMap.Store(keyNotStart, make([]DownloadTask, 0))
	syncMap.Store(keyInProgress, make([]DownloadTask, 0))
	syncMap.Store(keyCompleted, make([]DownloadTask, 0))
}

func AddTasks(urls []string) {
	val, _ := syncMap.Load(keyNotStart)
	list := val.([]DownloadTask)

	for _, item := range urls {
		list = append(list,
			DownloadTask{
				Url: item,
			})
	}

	syncMap.Store(keyNotStart, list)

	return
}

func StartTask() {
	notStartVal, _ := syncMap.Load(keyNotStart)
	notStartList := notStartVal.([]DownloadTask)

	if len(notStartList) == 0 {
		_logUtils.Infof("no download task to run")
		return
	}

	taskToStart := notStartList[0]

	notStartList = notStartList[1:]
	notStartVal = notStartList
	syncMap.Store(keyNotStart, notStartList)
	syncMap.Store(keyInProgress, []DownloadTask{taskToStart})

	Download(taskToStart)
}

func CompleteTask() {
	inProgressVal, _ := syncMap.Load(keyInProgress)
	completedVal, _ := syncMap.Load(keyCompleted)

	inProgressList := inProgressVal.([]DownloadTask)
	completedList := completedVal.([]DownloadTask)

	if len(inProgressList) == 0 {
		return
	}

	one := inProgressList[0]
	completedList = append(completedList, one)

	syncMap.Store(keyInProgress, make([]DownloadTask, 0))
	syncMap.Store(keyCompleted, completedList)

	return
}

func GetTask() (ret DownloadTask) {
	val, _ := syncMap.Load(keyNotStart)

	list := val.([]DownloadTask)

	if len(list) > 0 {
		ret = list[0]

		list = list[1:]
		syncMap.Store(keyNotStart, list)
	}

	return
}

func IsRunning() (ret bool) {
	inProgressVal, _ := syncMap.Load(keyInProgress)
	inProgressList := inProgressVal.([]DownloadTask)

	ret = len(inProgressList) > 0

	return
}

func ClearTask() {
	InitTasks()
	return
}

type DownloadTask struct {
	Url string
}
