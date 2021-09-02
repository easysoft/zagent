package serverService

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
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

	ecsUrl := fmt.Sprintf(serverConst.ALIYUN_ECS_URL, host.CloudRegion)
	eciClient, _ := s.AliyunCommService.CreateEciClient(serverConst.ALIYUN_ECI_URL, host.CloudKey, host.CloudSecret)
	ecsClient, _ := s.AliyunCommService.CreateEcsClient(ecsUrl, host.CloudKey, host.CloudSecret)
	vpcClient, _ := s.AliyunCommService.CreateVpcClient(ecsUrl, host.CloudKey, host.CloudSecret)

	switchId, _, _ := s.AliyunCommService.GetSwitch(host.VpcId, host.CloudRegion, vpcClient)
	securityGroupId, _ := s.AliyunCommService.QuerySecurityGroupByVpc(host.VpcId, host.CloudRegion, ecsClient)
	eipId, _ := s.AliyunCommService.CreateEip(host.CloudRegion, vpcClient)

	if eipId == "" || switchId == "" || securityGroupId == "" {
		msg := fmt.Sprintf("eipId (%s), switchId (%s) or securityGroupId (%s) is empty, cancel.",
			eipId, switchId, securityGroupId)
		_logUtils.Infof(msg)
		result.Fail(msg)
		return
	}

	url := fmt.Sprintf(serverConst.ALIYUN_ECS_URL, host.CloudRegion)
	vpcClient, err := s.AliyunCommService.CreateVpcClient(url, host.CloudKey, host.CloudSecret)

	cmd := []string{
		"/bin/bash",
		"-c",
		strings.Join(strings.Split(queue.BuildCommands, "\n"), "; "),
	}

	image := queue.DockerImage
	jobName := queue.TaskName + "-" + _stringUtils.NewUuid()

	id, err := s.AliyunEciService.CreateInst(jobName, image, "", cmd,
		eipId, switchId, securityGroupId, host.CloudRegion, eciClient)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	vm := model.Vm{
		HostId:      host.ID,
		HostName:    host.Name,
		Status:      consts.VmCreated,
		CloudInstId: id,
		CloudEipId:  eipId,
	}
	s.VmRepo.Save(&vm)

	return
}
func (s AliyunDockerService) DestroyRemote(vmId, queueId uint) {
	vm := s.VmRepo.GetById(vmId)
	jobName := vm.CloudInstId

	host := s.HostRepo.Get(vm.HostId)

	client, err1 := s.AliyunCommService.CreateEciClient(serverConst.ALIYUN_ECI_URL, host.CloudKey, host.CloudSecret)
	ecsUrl := fmt.Sprintf(serverConst.ALIYUN_ECS_URL, host.CloudRegion)
	vpcClient, err2 := s.AliyunCommService.CreateVpcClient(ecsUrl, host.CloudKey, host.CloudSecret)

	var status consts.VmStatus
	if err1 == nil && err2 == nil {
		_, err := s.AliyunEciService.Destroy(jobName, host.CloudRegion, client)
		if err == nil {
			err = s.AliyunCommService.DestroyEip(vm.CloudEipId, host.CloudRegion, vpcClient)
		} else {
			status = consts.VmFailDestroy
		}
	} else {
		status = consts.VmFailDestroy
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CloudInstId}, status)
	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
