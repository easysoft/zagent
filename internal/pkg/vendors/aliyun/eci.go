package aliyun

import (
	"strings"

	eci "github.com/alibabacloud-go/eci-20180808/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/easysoft/zagent/internal/pkg/domain"
)

type AliyunEciService struct {
	AliyunCommService *AliyunCommService `inject:""`
}

func NewAliyunEciService() *AliyunEciService {
	return &AliyunEciService{}
}

func (s AliyunEciService) CreateInst(groupName, image, imageCacheId string, cmd []string,
	eipId, switchId, securityGroupId, regionId string,
	eciClient *eci.Client) (
	id string, err error) {

	args := []*string{tea.String("-c"), tea.String(strings.Join(cmd, " && "))}

	container := &eci.CreateContainerGroupRequestContainer{
		Name:       tea.String(groupName),
		Image:      tea.String(image),
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
		EipInstanceId:      tea.String(eipId),
		VSwitchId:          tea.String(switchId),
		SecurityGroupId:    tea.String(securityGroupId),
	}

	if imageCacheId != "" {
		req.ImageSnapshotId = tea.String(imageCacheId)
	}

	resp, err := eciClient.CreateContainerGroup(req)
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
