package aliyun

import (
	"fmt"
	"github.com/easysoft/zagent/cmd/test/const"
	consts "github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"strings"
	"testing"
)

func TestHuaweiCloudCci(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)

	srv := vendors.NewHuaweiCloudCommService()
	client, _ := srv.CreateIamClient(testconst.HUAWEI_CLOUD_KEY, testconst.HUAWEI_CLOUD_Secret, testconst.HUAWEI_CLOUD_REGION)

	token, _ := srv.GetIamToken(client)

	image := "swr.cn-east-3.myhuaweicloud.com/tester-im/maven-testng:1.0"
	name := "maven-testng"
	cmd := []string{
		"/bin/bash",
		"-c",
		strings.Join([]string{
			"sleep 30",
			"rm -rf ci_test_testng",
			"git clone https://gitee.com/ngtesting/ci_test_testng.git; cd ci_test_testng",
			"mvn clean package > logs.txt",
			"sleep 600",
		}, "; "),
	}

	reqCreate := domain.CciReqCreate{
		APIVersion: "batch/v1",
		Kind:       "Job",
		Metadata: domain.CciMetadata{
			Name: name,
		},
		SpecTempl: domain.CciSpecTempl{
			Template: domain.CciTemplate{
				Metadata: domain.CciMetadata{
					Name: name,
				},
				Spec: domain.CciSpec{
					Containers: []domain.CciContainers{
						{
							Name:  name,
							Image: image,
							Resources: domain.CciResources{
								Limits: domain.CciLimits{
									CPU:    "2000m",
									Memory: "4096Mi",
								},
								Requests: domain.CciRequests{
									CPU:    "2000m",
									Memory: "4096Mi",
								},
							},
							Command: cmd,
						}},
					ImagePullSecrets: []domain.ImagePullSecrets{
						{
							Name: "imagepull-secret",
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	createUrl := fmt.Sprintf(testconst.HuaweiCloudUrlJobCreate,
		testconst.HUAWEI_CLOUD_REGION, testconst.HUAWEI_CLOUD_NAMEAPACE)
	resp, success := srv.Post(createUrl, reqCreate, nil, map[string]string{"X-Auth-Token": token})
	_logUtils.Infof("%#v, %#v", resp, success)

	reqDestroy := domain.CciReqDestroy{
		Kind:              "DeleteOptions",
		APIVersion:        "v1",
		PropagationPolicy: "Orphan",
	}

	destroyUrl := fmt.Sprintf(testconst.HuaweiCloudUrlJobDestroy,
		testconst.HUAWEI_CLOUD_REGION, testconst.HUAWEI_CLOUD_NAMEAPACE, name)
	resp, success = srv.Delete(destroyUrl, reqDestroy, nil, map[string]string{"X-Auth-Token": token})
	_logUtils.Infof("%#v, %#v", resp, success)
}
