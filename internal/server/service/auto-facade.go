package serverService

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"strings"
)

type FacadeService struct {
	VmRepo   *repo.VmRepo   `inject:""`
	HostRepo *repo.HostRepo `inject:""`

	KvmNativeService         *NativeKvmService         `inject:""`
	HuaweiCloudVmService     *HuaweiCloudVmService     `inject:""`
	HuaweiCloudDockerService *HuaweiCloudDockerService `inject:""`

	AliyunVmService *AliyunVmService `inject:""`

	SeleniumService *SeleniumService `inject:""`
	AppiumService   *AppiumService   `inject:""`
	UnitService     *UnitService     `inject:""`
}

func NewRunService() *FacadeService {
	return &FacadeService{}
}

// Create create machine
func (s FacadeService) Create(hostId, backingId, tmplId, queueId uint) (
	result _domain.RpcResp) {

	platform := s.HostRepo.Get(hostId).Platform.ToString()

	if s.IsVm(platform) {
		if s.IsNative(platform) {
			s.CreateVmKvmNative(hostId, backingId, tmplId, queueId)
		} else if s.IsHuaweiCloud(platform) {
			s.CreateVmHuaweiCloud(hostId, backingId, tmplId, queueId)
		}
	} else if s.IsDocker(platform) {
		if s.IsHuaweiCloud(platform) {
			s.CreateDockerHuaweiCloud(hostId, queueId)
		} else if s.IsAliyun(platform) {
			//s.CreateDockerAliyun(hostId, queueId)
		}
	}

	return
}
func (s FacadeService) CreateVmKvmNative(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {
	result = s.KvmNativeService.CreateRemote(hostId, backingId, tmplId, queueId)
	return
}

func (s FacadeService) CreateVmHuaweiCloud(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {
	result = s.HuaweiCloudVmService.CreateRemote(hostId, backingId, tmplId, queueId)
	return
}
func (s FacadeService) CreateVmAliyun(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {
	result = s.HuaweiCloudVmService.CreateRemote(hostId, backingId, tmplId, queueId)
	return
}

func (s FacadeService) CreateDockerHuaweiCloud(hostId, queueId uint) (result _domain.RpcResp) {
	result = s.HuaweiCloudDockerService.CreateRemote(hostId, queueId)
	return
}

// RunTest run testing
func (s FacadeService) RunTest(queue model.Queue, host model.Host) (result _domain.RpcResp) {
	if queue.BuildType == consts.SeleniumTest {
		s.RunSeleniumTest(queue)
	} else if queue.BuildType == consts.AppiumTest {
		s.RunAppiumTest(queue)
	} else if queue.BuildType == consts.UnitTest {
		s.RunUnitTest(queue, host)
	}

	return
}
func (s FacadeService) RunSeleniumTest(queue model.Queue) (result _domain.RpcResp) {
	result = s.SeleniumService.RunRemote(queue)
	return
}
func (s FacadeService) RunAppiumTest(queue model.Queue) (result _domain.RpcResp) {
	result = s.AppiumService.RunRemote(queue)
	return
}
func (s FacadeService) RunUnitTest(queue model.Queue, host model.Host) (result _domain.RpcResp) {
	result = s.UnitService.RunRemote(queue, host)
	return
}

// Destroy destory machine
func (s FacadeService) Destroy(queue model.Queue) {
	vm := s.VmRepo.GetById(queue.VmId)
	platform := s.HostRepo.Get(vm.HostId).Platform.ToString()

	if s.IsVm(platform) {
		if s.IsNative(platform) {
			s.DestroyVmKvmNative(queue)
		} else if s.IsHuaweiCloud(platform) {
			s.DestroyVmHuaweiCloud(queue)
		}
	} else if s.IsDocker(platform) {
		if s.IsHuaweiCloud(platform) {
			s.DestroyDockerHuaweiCloud(queue)
		}
	}
}
func (s FacadeService) DestroyVmKvmNative(queue model.Queue) (result _domain.RpcResp) {
	s.KvmNativeService.DestroyRemote(queue.VmId, queue.ID)
	return
}
func (s FacadeService) DestroyVmHuaweiCloud(queue model.Queue) (result _domain.RpcResp) {
	s.HuaweiCloudVmService.DestroyRemote(queue.VmId, queue.ID)
	return
}
func (s FacadeService) DestroyDockerHuaweiCloud(queue model.Queue) (result _domain.RpcResp) {
	s.HuaweiCloudDockerService.DestroyRemote(queue.VmId, queue.ID)
	return
}

func (s FacadeService) IsVm(platform string) (result bool) {
	result = strings.Index(platform, consts.PlatformVm.ToString()) > -1
	return
}
func (s FacadeService) IsDocker(platform string) (result bool) {
	result = strings.Index(platform, consts.PlatformDocker.ToString()) > -1
	return
}
func (s FacadeService) IsNative(platform string) (result bool) {
	result = strings.Index(platform, consts.PlatformNative.ToString()) > -1
	return
}
func (s FacadeService) IsHuaweiCloud(platform string) (result bool) {
	result = strings.Index(platform, consts.PlatformHuawei.ToString()) > -1
	return
}
func (s FacadeService) IsAliyun(platform string) (result bool) {
	result = strings.Index(platform, consts.PlatformAli.ToString()) > -1
	return
}
