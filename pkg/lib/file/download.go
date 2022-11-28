package _fileUtils

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func DownloadAdv(url, filePath string) (dur int64, err error) {
	startTime := time.Now()

	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(filePath, url)

	// start download
	resp := client.Do(req)

	_logUtils.Infof("%v, downloading %v...\n", resp.HTTPResponse.Status, req.URL())

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Print("\033[G\033[K")

			fmt.Printf("  transferred %d / %d bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress())

			fmt.Print("\033[A")

		case <-resp.Done:
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		_logUtils.Infof("download failed: %v\n", err)
	}

	_logUtils.Infof("saved to %v \n", resp.Filename)

	endTime := time.Now()

	dur = endTime.Unix() - startTime.Unix()

	return
}

func Download(url string, dst string) (err error) {
	fmt.Printf("DownloadToFile From: %s to %s.\n", url, dst)

	MkDirIfNeeded(filepath.Dir(dst))

	var data []byte
	data, err = HTTPDownload(url)
	if err != nil {
		return
	}

	err = WriteDownloadFile(dst, data)
	if err != nil {
		_logUtils.Infof("save %s to %s successfully", url, dst)
	}

	return
}

func HTTPDownload(url string) (bytes []byte, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		_logUtils.Infof("download %s error %s", url, err.Error())
		return
	}
	if resp.StatusCode != 200 {
		_logUtils.Infof("download %s server return %s", url, resp.Status)
		return
	}

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_logUtils.Error(err.Error())
	}
	return d, err
}

func WriteDownloadFile(dst string, d []byte) (err error) {
	err = RemoveFile(dst)
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	err = ioutil.WriteFile(dst, d, 0777)
	if err != nil {
		_logUtils.Error(err.Error())
	}

	return
}

func AddTimeParam(url string) (ret string) {
	if strings.Index(url, "?") > 0 {
		url += fmt.Sprintf("&ts=%d", time.Now().Unix())
	} else {
		url += fmt.Sprintf("?ts=%d", time.Now().Unix())
	}

	return
}
