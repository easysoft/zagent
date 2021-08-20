package vendors

import (
	"encoding/json"
	"fmt"
	"github.com/easysoft/zagent/cmd/test/_const"
	"github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
)

type HuaweiCloudCciService struct {
	HuaweiCloudCommService *HuaweiCloudCommService `inject:""`
}

func NewHuaweiCloudCciService() *HuaweiCloudCciService {
	return &HuaweiCloudCciService{}
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

func (s HuaweiCloudCciService) Destroy(jobName, token, region, namespace string) (
	ret domain.CciRepsDestroy, success bool) {

	reqDestroy := domain.CciReqDestroy{
		Kind:              "DeleteOptions",
		APIVersion:        "v1",
		PropagationPolicy: "Orphan",
	}

	destroyUrl := fmt.Sprintf(testconst.HuaweiCloudUrlJobDestroy, region, namespace, jobName)

	var resp []byte
	resp, success = s.HuaweiCloudCommService.delete(destroyUrl, reqDestroy, nil, map[string]string{"X-Auth-Token": token})

	json.Unmarshal(resp, ret)
	_logUtils.Infof("%#v", ret)

	return
}
