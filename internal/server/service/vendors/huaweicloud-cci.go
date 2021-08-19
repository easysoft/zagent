package vendors

import (
	"encoding/json"
	"fmt"
	testconst "github.com/easysoft/zagent/cmd/test/const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/model"
	"strings"
)

type HuaweiCloudCciService struct {
	HuaweiCloudCommService *HuaweiCloudCommService `inject:""`
}

func NewHuaweiCloudCciService() *HuaweiCloudCciService {
	return &HuaweiCloudCciService{}
}

func (s HuaweiCloudCciService) CreateByQueue(queue model.Queue, host model.Host) (result _domain.RpcResp) {
	client, _ := s.HuaweiCloudCommService.CreateIamClient(
		testconst.HUAWEI_CLOUD_KEY, testconst.HUAWEI_CLOUD_Secret, testconst.HUAWEI_CLOUD_REGION)
	token, _ := s.HuaweiCloudCommService.GetIamToken(client)
	cmd := []string{
		"/bin/bash",
		"-c",
		strings.Join(strings.Split(queue.BuildCommands, "\n"), "; "),
	}

	image := queue.DockerImage
	jobName := queue.TaskName + "-" + _stringUtils.NewUuid()
	region := host.CloudRegion
	namespace := host.CloudNamespace

	resp, success := s.Create(image, jobName, cmd, token, region, namespace)
	if success {
		result.Pass("")
	} else {
		bytes, _ := json.Marshal(resp)
		result.Fail(string(bytes))
	}

	return
}

func (s HuaweiCloudCciService) Create(image string, jobName string, cmd []string,
	token string, region string, namespace string) (ret domain.CciRepsCreate, success bool) {

	reqCreate := domain.CciReqCreate{
		APIVersion: "batch/v1",
		Kind:       "Job",
		Metadata: domain.CciMetadata{
			Name: jobName,
		},
		SpecTempl: domain.CciSpecTempl{
			Template: domain.CciTemplate{
				Metadata: domain.CciMetadata{
					Name: jobName,
				},
				Spec: domain.CciSpec{
					Containers: []domain.CciContainers{
						{
							Name:  jobName,
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

	createUrl := fmt.Sprintf(testconst.HuaweiCloudUrlJobCreate, region, namespace)
	var resp []byte
	resp, success = s.HuaweiCloudCommService.post(createUrl, reqCreate, nil, map[string]string{"X-Auth-Token": token})

	json.Unmarshal(resp, &ret)
	name := ret.Metadata.Name
	_logUtils.Infof("%s#v", name)

	return
}

func (s HuaweiCloudCciService) Destroy(jobName string, token string, region string, nameapace string) (
	ret domain.CciRepsDestroy, success bool) {

	reqDestroy := domain.CciReqDestroy{
		Kind:              "DeleteOptions",
		APIVersion:        "v1",
		PropagationPolicy: "Orphan",
	}

	destroyUrl := fmt.Sprintf(testconst.HuaweiCloudUrlJobDestroy,
		testconst.HUAWEI_CLOUD_REGION, testconst.HUAWEI_CLOUD_NAMEAPACE, jobName)

	var resp []byte
	resp, success = s.HuaweiCloudCommService.delete(destroyUrl, reqDestroy, nil, map[string]string{"X-Auth-Token": token})

	json.Unmarshal(resp, ret)
	_logUtils.Infof("%#v", ret)

	return
}
