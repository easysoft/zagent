package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/easysoft/zv/cmd/test/_const"
	"github.com/easysoft/zv/internal/comm/const"
	"github.com/easysoft/zv/internal/pkg/lib/log"
	"os"
	"testing"
)

func TestOss(t *testing.T) {
	_logUtils.Init(consts.AppNameAgentHost)

	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", testconst.ALIYUN_KEY, testconst.ALIYUN_Secret)
	if err != nil {
		_logUtils.Printf("oss.New error %s", err.Error())
		os.Exit(-1)
	}

	bucketName := "com-deeptest"
	objectName := "tmpl-ubt20-x64-desktop-zh_cn.qcow2"
	locaFilename := "/Users/aaron/Downloads/tmpl-ubt20-x64-desktop-zh_cn.qcow2"

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
