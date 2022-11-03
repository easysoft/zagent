package _fileUtils

import (
	"fmt"
	_i118Utils "github.com/easysoft/zagent/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func Download(url string, dst string) (err error) {
	fmt.Printf("DownloadToFile From: %s to %s.\n", url, dst)

	MkDirIfNeeded(filepath.Dir(dst))

	var data []byte
	data, err = HTTPDownload(url)
	if err == nil && len(string(data)) == 32 {
		_logUtils.Info(_i118Utils.Sprintf("file_downloaded", url))

		err = WriteDownloadFile(dst, data)
		if err == nil {
			_logUtils.Info(_i118Utils.Sprintf("file_download_saved", url, dst))
		}
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

func WriteDownloadFile(dst string, d []byte) error {
	err := ioutil.WriteFile(dst, d, 0444)
	if err != nil {
		_logUtils.Error(err.Error())
	}
	return err
}
