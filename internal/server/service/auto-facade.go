package serverService

import (
	consts "github.com/easysoft/zv/internal/comm/const"
	_domain "github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	serverUitls "github.com/easysoft/zv/internal/server/utils/lib"
)

type FacadeService struct {
	VmRepo   *repo.VmRepo   `inject:""`
	HostRepo *repo.HostRepo `inject:""`

	KvmNativeService         *NativeKvmService         `inject:""`
	VirtualboxCloudVmService *VirtualboxCloudVmService `inject:""`
	VmWareCloudVmService     *VmWareCloudVmService     `inject:""`

	HuaweiCloudVmService     *HuaweiCloudVmService     `inject:""`
	HuaweiCloudDockerService *HuaweiCloudDockerService `inject:""`

	AliyunVmService     *AliyunVmService     `inject:""`
	AliyunDockerService *AliyunDockerService `inject:""`

	SeleniumService *SeleniumService `inject:""`
	AppiumService   *AppiumService   `inject:""`
	UnitService     *UnitService     `inject:""`
}

func NewRunService() *FacadeService {
	return &FacadeService{}
}

// Create create machine
func (s FacadeService) Create(hostId, backingId, tmplId, queueId uint) (
	result _domain.RemoteResp) {
	host := s.HostRepo.Get(hostId)
	capabilities := host.Capabilities.ToString()
	vendor := host.Vendor.ToString()

	if serverUitls.IsVm(capabilities) {
		if serverUitls.IsNative(vendor) {
			result = s.CreateVmKvmNative(hostId, backingId, tmplId, queueId)
		} else if serverUitls.IsVirtualBox(vendor) {
			result = s.CreateVmVirtualBox(hostId, backingId, queueId)
		} else if serverUitls.IsVmWare(vendor) {
			result = s.CreateVmVmWare(hostId, backingId, queueId)
		} else if serverUitls.IsHuaweiCloud(vendor) {
			result = s.CreateVmHuaweiCloud(hostId, backingId, queueId)
		} else if serverUitls.IsAliyun(vendor) {
			result = s.CreateVmAliyun(hostId, backingId, queueId)
		}
	} else if serverUitls.IsDocker(capabilities) {
		if serverUitls.IsHuaweiCloud(vendor) {
			result = s.CreateDockerHuaweiCloud(hostId, queueId)
		} else if serverUitls.IsAliyun(vendor) {
			result = s.CreateDockerAliyun(hostId, queueId)
		}
	}

	return
}
func (s FacadeService) CreateVmKvmNative(hostId, backingId, tmplId, queueId uint) (result _domain.RemoteResp) {
	result = s.KvmNativeService.CreateRemote(hostId, backingId, tmplId, queueId)
	return
}

func (s FacadeService) CreateVmVirtualBox(hostId, backingId, queueId uint) (result _domain.RemoteResp) {
	result = s.VirtualboxCloudVmService.CreateRemote(hostId, backingId, queueId)
	return
}
func (s FacadeService) CreateVmVmWare(hostId, backingId, queueId uint) (result _domain.RemoteResp) {
	result = s.VmWareCloudVmService.CreateRemote(hostId, backingId, queueId)
	return
}
func (s FacadeService) CreateVmHuaweiCloud(hostId, backingId, queueId uint) (result _domain.RemoteResp) {
	result = s.HuaweiCloudVmService.CreateRemote(hostId, backingId, queueId)
	return
}
func (s FacadeService) CreateVmAliyun(hostId, backingId, queueId uint) (result _domain.RemoteResp) {
	result = s.AliyunVmService.CreateRemote(hostId, backingId, queueId)
	return
}

func (s FacadeService) CreateDockerHuaweiCloud(hostId, queueId uint) (result _domain.RemoteResp) {
	result = s.HuaweiCloudDockerService.CreateRemote(hostId, queueId)
	return
}
func (s FacadeService) CreateDockerAliyun(hostId, queueId uint) (result _domain.RemoteResp) {
	result = s.AliyunDockerService.CreateRemote(hostId, queueId)
	return
}

// RunTest run testing
func (s FacadeService) RunTest(queue model.Queue, host model.Host) (result _domain.RemoteResp) {
	if queue.BuildType == consts.SeleniumTest {
		s.RunSeleniumTest(queue)
	} else if queue.BuildType == consts.AppiumTest {
		s.RunAppiumTest(queue)
	} else if queue.BuildType == consts.UnitTest {
		s.RunUnitTest(queue, host)
	}

	return
}

// Destroy destory machine
func (s FacadeService) Destroy(queue model.Queue) {
	vm := s.VmRepo.GetById(queue.VmId)
	platform := s.HostRepo.Get(vm.HostId).Capabilities.ToString()

	if serverUitls.IsVm(platform) {
		if serverUitls.IsNative(platform) {
			s.DestroyVmKvmNative(queue)
		} else if serverUitls.IsVirtualBox(platform) {
			s.DestroyVmVirtualBox(queue)
		} else if serverUitls.IsVmWare(platform) {
			s.DestroyVmVmWare(queue)
		} else if serverUitls.IsHuaweiCloud(platform) {
			s.DestroyVmHuaweiCloud(queue)
		} else if serverUitls.IsAliyun(platform) {
			s.DestroyVmAliyun(queue)
		}
	} else if serverUitls.IsDocker(platform) {
		if serverUitls.IsHuaweiCloud(platform) {
			s.DestroyDockerHuaweiCloud(queue)
		} else if serverUitls.IsAliyun(platform) {
			s.DestroyDockerAliyun(queue)
		}
	}
}

func (s FacadeService) RunSeleniumTest(queue model.Queue) (result _domain.RemoteResp) {
	result = s.SeleniumService.RunRemote(queue)
	return
}
func (s FacadeService) RunAppiumTest(queue model.Queue) (result _domain.RemoteResp) {
	result = s.AppiumService.RunRemote(queue)
	return
}
func (s FacadeService) RunUnitTest(queue model.Queue, host model.Host) (result _domain.RemoteResp) {
	result = s.UnitService.RunRemote(queue, host)
	return
}

func (s FacadeService) DestroyVmKvmNative(queue model.Queue) (result _domain.RemoteResp) {
	s.KvmNativeService.DestroyRemote(queue.VmId, queue.ID)
	return
}
func (s FacadeService) DestroyVmVirtualBox(queue model.Queue) (result _domain.RemoteResp) {
	s.VirtualboxCloudVmService.DestroyRemote(queue.VmId, queue.ID)
	return
}
func (s FacadeService) DestroyVmVmWare(queue model.Queue) (result _domain.RemoteResp) {
	s.VmWareCloudVmService.DestroyRemote(queue.VmId, queue.ID)
	return
}
func (s FacadeService) DestroyVmHuaweiCloud(queue model.Queue) (result _domain.RemoteResp) {
	s.HuaweiCloudVmService.DestroyRemote(queue.VmId, queue.ID)
	return
}
func (s FacadeService) DestroyVmAliyun(queue model.Queue) (result _domain.RemoteResp) {
	s.AliyunVmService.DestroyRemote(queue.VmId, queue.ID)
	return
}
func (s FacadeService) DestroyDockerHuaweiCloud(queue model.Queue) (result _domain.RemoteResp) {
	s.HuaweiCloudDockerService.DestroyRemote(queue.VmId, queue.ID)
	return
}
func (s FacadeService) DestroyDockerAliyun(queue model.Queue) (result _domain.RemoteResp) {
	s.AliyunDockerService.DestroyRemote(queue.VmId, queue.ID)
	return
}
