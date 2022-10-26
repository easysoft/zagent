package downloadUtils

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	agentModel "github.com/easysoft/zv/internal/host/model"
	"os"
	"time"
)

func Start(task agentModel.Download, ch chan int) {
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

	isTerminate := false

	for completed < 1 {
		select {
		case <-ch:
			isTerminate = true
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

	if isTerminate {
		fmt.Printf("Force to terminate download %s.\n", task.Url)
	} else {
		fmt.Printf("Success to download %s.\n", task.Url)
	}
}
