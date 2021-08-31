package vendors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
)

type HuaweiCloudCciService struct {
	HuaweiCloudCommService *HuaweiCloudCommService `inject:""`
}

func NewHuaweiCloudCciService() *HuaweiCloudCciService {
	return &HuaweiCloudCciService{}
}

func (s HuaweiCloudCciService) Create(image string, jobName string, cmd []string,
	token string, region string, namespace string) (ret domain.CciRepsCreate, err error) {

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

	createUrl := fmt.Sprintf(serverConst.HuaweiCloudUrlJobCreate, region, namespace)
	resp, success := s.HuaweiCloudCommService.post(createUrl, reqCreate, nil, map[string]string{"X-Auth-Token": token})

	if !success {
		err = errors.New("fail to create huaweicloud cci")
	}

	json.Unmarshal(resp, &ret)
	name := ret.Metadata.Name
	_logUtils.Infof("%s#v", name)

	return
}

func (s HuaweiCloudCciService) Destroy(jobName, token, region, namespace string) (
	ret domain.CciRepsDestroy, err error) {

	reqDestroy := domain.CciReqDestroy{
		Kind:              "DeleteOptions",
		APIVersion:        "v1",
		PropagationPolicy: "Orphan",
	}

	destroyUrl := fmt.Sprintf(serverConst.HuaweiCloudUrlJobDestroy, region, namespace, jobName)

	resp, success := s.HuaweiCloudCommService.delete(destroyUrl, reqDestroy, nil, map[string]string{"X-Auth-Token": token})

	if !success {
		err = errors.New("fail to destroy huaweicloud cci")
	}

	json.Unmarshal(resp, &ret)
	name := ret.Metadata.Name
	_logUtils.Infof("%s#v", name)

	return
}
