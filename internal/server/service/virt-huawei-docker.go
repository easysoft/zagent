package serverService

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/easysoft/zagent/internal/server/service/vendors/huaweicloud"
	"strings"
)

type HuaweiCloudDockerService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	QueueRepo   *repo.QueueRepo   `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`

	VmCommonService        *VmCommonService `inject:""`
	HistoryService         *HistoryService  `inject:""`
	HuaweiCloudCommService *huaweicloud.HuaweiCloudCommService
	HuaweiCloudCciService  *huaweicloud.HuaweiCloudCciService `inject:""`
}

func (s HuaweiCloudDockerService) CreateRemote(hostId, queueId uint) (result _domain.RpcResp) {
	queue := s.QueueRepo.GetQueue(queueId)
	host := s.HostRepo.Get(hostId)

	client, _ := s.HuaweiCloudCommService.CreateIamClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	token, _ := s.HuaweiCloudCommService.GetIamToken(
		host.CloudUser, host.CloudIamUser, host.CloudIamPassword,
		client)
	cmd := []string{
		"/bin/bash",
		"-c",
		strings.Join(strings.Split(queue.BuildCommands, "\n"), "; "),
	}

	jobName := queue.TaskName + "-" + _stringUtils.NewUuid()
	image := queue.DockerImage
	region := host.CloudRegion
	namespace := host.CloudNamespace

	resp, err := s.HuaweiCloudCciService.Create(jobName, image, cmd, token, region, namespace)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	vm := model.Vm{
		HostId:      host.ID,
		HostName:    host.Name,
		Status:      consts.VmCreated,
		CloudInstId: resp.Metadata.Name,
	}
	s.VmRepo.Save(&vm)

	return
}

func (s HuaweiCloudDockerService) DestroyRemote(vmId, queueId uint) {
	vm := s.VmRepo.GetById(vmId)
	jobName := vm.CloudInstId

	host := s.HostRepo.Get(vm.HostId)

	client, err := s.HuaweiCloudCommService.CreateIamClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	token, _ := s.HuaweiCloudCommService.GetIamToken(
		host.CloudUser, host.CloudIamUser, host.CloudIamPassword,
		client)

	var status consts.VmStatus
	if err == nil {
		_, err = s.HuaweiCloudCciService.Destroy(jobName, token, host.CloudRegion, host.CloudNamespace)
	}

	if err != nil {
		status = consts.VmFailDestroy
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CloudInstId}, status)
	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
