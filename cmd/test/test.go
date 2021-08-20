package main

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	_const "github.com/easysoft/zagent/cmd/test/_const"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient(tea.String(_const.ALIYUN_KEY), tea.String(_const.ALIYUN_Secret))
	if _err != nil {
		return _err
	}

	describeImagesRequest := &ecs20140526.DescribeImagesRequest{
		RegionId:     tea.String("cn-hangzhou"),
		OSType:       tea.String("windows"),
		Architecture: tea.String("x86_64"),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.DescribeImages(describeImagesRequest)
	if _err != nil {
		return _err
	}

	fmt.Printf("found %d images\n", len(result.Body.Images.Image))

	return _err
}

//func main() {
//	err := _main(tea.StringSlice(os.Args[1:]))
//	if err != nil {
//		panic(err)
//	}
//}
