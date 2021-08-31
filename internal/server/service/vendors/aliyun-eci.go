package vendors

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	eci "github.com/alibabacloud-go/eci-20180808/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/easysoft/zagent/internal/comm/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"strings"
)

type AliyunEciService struct {
}

func NewAliyunEciService() *AliyunEciService {
	return &AliyunEciService{}
}

func (s AliyunEciService) CreateInst(groupName, imageName, image string, cmd []string, regionId string, client *eci.Client) (
	id string, err error) {

	args := []*string{tea.String("-c"), tea.String(strings.Join(cmd, " && "))}

	container := &eci.CreateContainerGroupRequestContainer{
		Image:      tea.String(image),
		Name:       tea.String(imageName),
		WorkingDir: tea.String("/"),
		Command:    []*string{tea.String("/bin/bash")},
		Arg:        args,
		Tty:        tea.Bool(true),
		ReadinessProbe: &eci.CreateContainerGroupRequestContainerReadinessProbe{
			TcpSocket: &eci.CreateContainerGroupRequestContainerReadinessProbeTcpSocket{
				Port: tea.Int32(22),
			},
			HttpGet: &eci.CreateContainerGroupRequestContainerReadinessProbeHttpGet{
				Scheme: tea.String("https"),
			},
			Exec: &eci.CreateContainerGroupRequestContainerReadinessProbeExec{
				Command: []*string{tea.String("ls")},
			},
		},
		SecurityContext: &eci.CreateContainerGroupRequestContainerSecurityContext{
			Capability: &eci.CreateContainerGroupRequestContainerSecurityContextCapability{
				Add: []*string{tea.String("NET_ADMIN")},
			},
		},
		LivenessProbe: &eci.CreateContainerGroupRequestContainerLivenessProbe{
			TcpSocket: &eci.CreateContainerGroupRequestContainerLivenessProbeTcpSocket{
				Port: tea.Int32(22),
			},
			HttpGet: &eci.CreateContainerGroupRequestContainerLivenessProbeHttpGet{
				Scheme: tea.String("https"),
			},
			Exec: &eci.CreateContainerGroupRequestContainerLivenessProbeExec{
				Command: []*string{tea.String("ls")},
			},
		},
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

func (s AliyunEciService) Destroy(containerGroupId, region string, client *eci.Client) (
	ret domain.CciRepsDestroy, err error) {

	req := &eci.DeleteContainerGroupRequest{
		ContainerGroupId: tea.String(containerGroupId),
		RegionId:         tea.String(region),
	}
	_, err = client.DeleteContainerGroup(req)

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
