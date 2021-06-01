package _fileUtils

import (
	"bytes"
	"fmt"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func Upload(url string, files []string, extraParams map[string]string) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	for _, file := range files {
		fw, _ := bodyWriter.CreateFormFile("file", file)
		f, _ := os.Open(file)
		defer f.Close()
		io.Copy(fw, f)
	}

	for key, value := range extraParams {
		_ = bodyWriter.WriteField(key, value)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuffer)
	defer resp.Body.Close()

	if err != nil {
		e := "fail to upload files, err: " + err.Error()
		_logUtils.Error(e)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e := "fail to parse upload resp, err: " + err.Error()
		_logUtils.Error(e)
	}

	_logUtils.Info("upload status " + resp.Status + ", resp is " + string(respBody))
}

func Download(url string, dst string) {
	fmt.Printf("DownloadToFile From: %s.\n", url)
	if d, err := HTTPDownload(url); err == nil {
		_logUtils.Info(fmt.Sprintf("downloaded %s.\n", url))
		if WriteDownloadFile(dst, d) == nil {
			_logUtils.Info(fmt.Sprintf("saved %s as %s\n", url, dst))
		}
	}
}

func HTTPDownload(uri string) ([]byte, error) {
	_logUtils.Info(fmt.Sprintf("HTTPDownload From: %s.\n", uri))
	res, err := http.Get(uri)
	if err != nil {
		_logUtils.Error(err.Error())
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_logUtils.Error(err.Error())
	}
	_logUtils.Info(fmt.Sprintf("ReadFile: Size of download: %d\n", len(d)))
	return d, err
}

func WriteDownloadFile(dst string, d []byte) error {
	_logUtils.Info(fmt.Sprintf("WriteFile: Size of download: %d\n", len(d)))
	err := ioutil.WriteFile(dst, d, 0444)
	if err != nil {
		_logUtils.Error(err.Error())
	}
	return err
}
