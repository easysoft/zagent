package vendors

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	eci "github.com/alibabacloud-go/eci-20180808/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
)

type AliyunEciService struct {
}

func NewAliyunEciService() *AliyunEciService {
	return &AliyunEciService{}
}

func (s AliyunEciService) CreateInst(groupName, imageName, image, cmd, regionId string, client *eci.Client) (
	id, name string, err error) {

	container := &eci.CreateContainerGroupRequestContainer{
		Image:   tea.String(image),
		Name:    tea.String(imageName),
		Command: []*string{tea.String(cmd)},
	}

	req := &eci.CreateContainerGroupRequest{
		ContainerGroupName: tea.String(cmd),
		RegionId:           tea.String(regionId),
		Cpu:                tea.Float32(2),
		Memory:             tea.Float32(4),
		Container:          []*eci.CreateContainerGroupRequestContainer{container},
	}

	_, err = client.CreateContainerGroup(req)

	return
}

func (s AliyunEciService) CreateEciClient(endpoint, accessKeyId, accessKeySecret string) (
	result *eci.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	config.Endpoint = tea.String(endpoint)
	result = &eci.Client{}
	result, err = eci.NewClient(config)
	if err != nil {
		_logUtils.Errorf("CreateEciClient error %s", err.Error())
		return
	}

	return result, err
}
