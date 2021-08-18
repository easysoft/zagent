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

	SeleniumService *SeleniumService `inject:""`
	AppiumService   *AppiumService   `inject:""`
	UnitService     *UnitService     `inject:""`
}

func NewRunService() *FacadeService {
	return &FacadeService{}
}

// create machine
func (s FacadeService) Create(hostId, backingId, tmplId, queueId uint) (
	result _domain.RpcResp) {

	platform := s.HostRepo.Get(hostId).Platform.ToString()

	if strings.Index(platform, consts.PlatformVm.ToString()) > -1 {
		if strings.Index(platform, consts.PlatformNative.ToString()) > -1 {
			s.CreateVmKvmNative(hostId, backingId, tmplId, queueId)
		} else if strings.Index(platform, consts.PlatformHuawei.ToString()) > -1 {
			s.CreateVmHuaweiCloud(hostId, backingId, tmplId, queueId)
		}
	} else if strings.Index(platform, consts.PlatformDocker.ToString()) > -1 {
		if strings.Index(platform, consts.PlatformHuawei.ToString()) > -1 {
			s.CreateDockerHuaweiCloud(hostId, backingId, tmplId, queueId)
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
func (s FacadeService) CreateDockerHuaweiCloud(hostId, backingId, tmplId, queueId uint) (result _domain.RpcResp) {
	result = s.HuaweiCloudDockerService.CreateRemote(hostId, backingId, tmplId, queueId)
	return
}

// run testing
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

// destory machine
func (s FacadeService) Destroy(queue model.Queue) {
	vm := s.VmRepo.GetById(queue.VmId)
	platform := s.HostRepo.Get(vm.HostId).Platform.ToString()

	if strings.Index(platform, consts.PlatformVm.ToString()) > -1 {
		if strings.Index(platform, consts.PlatformNative.ToString()) > -1 {
			s.DestroyVmKvmNative(queue)
		} else if strings.Index(platform, consts.PlatformHuawei.ToString()) > -1 {
			s.DestroyVmHuaweiCloud(queue)
		}
	} else if strings.Index(platform, consts.PlatformDocker.ToString()) > -1 {
		if strings.Index(platform, consts.PlatformHuawei.ToString()) > -1 {
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
