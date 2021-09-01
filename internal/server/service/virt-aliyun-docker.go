package serverService

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	serverConst "github.com/easysoft/zagent/internal/server/utils/const"
	"strings"
)

type AliyunDockerService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`
	QueueRepo   *repo.QueueRepo   `inject:""`

	VmCommonService   *VmCommonService           `inject:""`
	HistoryService    *HistoryService            `inject:""`
	AliyunEciService  *vendors.AliyunEciService  `inject:""`
	AliyunCommService *vendors.AliyunCommService `inject:""`
}

func NewAliyunDockerService() *AliyunDockerService {
	return &AliyunDockerService{}
}

func (s AliyunDockerService) CreateRemote(hostId, queueId uint) (result _domain.RpcResp) {
	queue := s.QueueRepo.GetQueue(queueId)
	host := s.HostRepo.Get(hostId)

	eciClient, _ := s.AliyunCommService.CreateEciClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	ecsClient, _ := s.AliyunCommService.CreateEcsClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	vpcClient, _ := s.AliyunCommService.CreateVpcClient(host.CloudKey, host.CloudSecret, host.CloudRegion)

	eipId, _ := s.AliyunCommService.GetEip(host.CloudRegion, vpcClient)

	switchId, _, _ := s.AliyunCommService.GetSwitch(host.VpcId, host.CloudRegion, vpcClient)
	securityGroupId, _ := s.AliyunCommService.QuerySecurityGroupByVpc(host.VpcId, host.CloudRegion, ecsClient)

	url := fmt.Sprintf(serverConst.ALIYUN_ECS_URL, host.CloudRegion)
	vpcClient, err := s.AliyunCommService.CreateVpcClient(url, host.CloudKey, host.CloudSecret)

	cmd := []string{
		"/bin/bash",
		"-c",
		strings.Join(strings.Split(queue.BuildCommands, "\n"), "; "),
	}

	image := queue.DockerImage
	jobName := queue.TaskName + "-" + _stringUtils.NewUuid()

	id, err := s.AliyunEciService.CreateInst(jobName, jobName, image, cmd,
		eipId, switchId, securityGroupId, host.CloudRegion, eciClient)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	vm := model.Vm{
		HostId:      host.ID,
		HostName:    host.Name,
		Status:      consts.VmCreated,
		CouldInstId: id,
	}
	s.VmRepo.Save(&vm)

	return
}
func (s AliyunDockerService) DestroyRemote(vmId, queueId uint) {
	vm := s.VmRepo.GetById(vmId)
	jobName := vm.CouldInstId

	host := s.HostRepo.Get(vm.HostId)

	client, err := s.AliyunCommService.CreateEciClient(host.CloudKey, host.CloudSecret, host.CloudRegion)

	var status consts.VmStatus
	if err == nil {
		_, err = s.AliyunEciService.Destroy(jobName, host.CloudRegion, client)
	}

	if err != nil {
		status = consts.VmFailDestroy
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CouldInstId}, status)
	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
