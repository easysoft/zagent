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

func (s AliyunEciService) CreateInst(groupName, imageName, image string, cmd []string, regionId string, client *eci.Client) (
	id string, err error) {

	commands := make([]*string, 0)
	for _, item := range cmd {
		commands = append(commands, tea.String(item))
	}

	container := &eci.CreateContainerGroupRequestContainer{
		Image:   tea.String(image),
		Name:    tea.String(imageName),
		Command: commands,
	}

	req := &eci.CreateContainerGroupRequest{
		ContainerGroupName: tea.String(groupName),
		RegionId:           tea.String(regionId),
		Cpu:                tea.Float32(2),
		Memory:             tea.Float32(4),
		Container:          []*eci.CreateContainerGroupRequestContainer{container},
	}

	resp, err := client.CreateContainerGroup(req)
	id = *resp.Body.ContainerGroupId

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
