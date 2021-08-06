package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	_const "github.com/easysoft/zagent/cmd/test/const"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"os"
	"testing"
)

func TestOss(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", _const.ALIYUN_KEY, _const.ALIYUN_Secret)
	if err != nil {
		_logUtils.Printf("oss.New error %s", err.Error())
		os.Exit(-1)
	}

	bucketName := "tester-im"
	objectName := "x64-pro-zh_cn.qcow2"
	locaFilename := "/Users/aaron/Downloads/x64-pro-zh_cn.qcow2"

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		_logUtils.Printf("client.Bucket error %s", err.Error())
		os.Exit(-1)
	}
	// 将本地文件分片，且分片数量指定为3。
	chunks, err := oss.SplitFileByPartNum(locaFilename, 8)
	if err != nil {
		_logUtils.Printf("SplitFileByPartNum error %s", err.Error())
	}

	fd, err := os.Open(locaFilename)
	defer fd.Close()
	if err != nil {
		_logUtils.Printf("Open error %s", err.Error())
		os.Exit(-1)
	}

	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	imur, err := bucket.InitiateMultipartUpload(objectName, storageType)
	if err != nil {
		_logUtils.Printf("InitiateMultipartUpload error %s", err.Error())
		os.Exit(-1)
	}

	var parts []oss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			_logUtils.Printf("UploadPart error %s", err.Error())
			os.Exit(-1)
		}
		parts = append(parts, part)
	}

	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	cmur, err := bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		_logUtils.Printf("CompleteMultipartUpload error %s", err.Error())
		os.Exit(-1)
	}

	_logUtils.Printf("cmur: %#v", cmur)
}
