package hostAgentService

import (
	"fmt"
	agentModel "github.com/easysoft/zagent/internal/host/model"
	hostRepo "github.com/easysoft/zagent/internal/host/repo"
	kvmService "github.com/easysoft/zagent/internal/host/service/kvm"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	"path/filepath"
)

type ExportService struct {
	TaskService    *TaskService               `inject:""`
	LibvirtService *kvmService.LibvirtService `inject:""`
	QemuService    *kvmService.QemuService    `inject:""`

	TaskRepo *hostRepo.TaskRepo `inject:""`
}

func NewExportService() *ExportService {
	return &ExportService{}
}

func (s *ExportService) StartTask(po agentModel.Task) {
	ch := make(chan int, 1)

	go func() {
		targetBakingFilePath := filepath.Join(agentConf.Inst.DirBaking, po.Backing)

		s.TaskRepo.UpdateStatus(po.ID, targetBakingFilePath, 0, "", consts.InProgress, true, false)

		xmlDesc, finalStatus := s.ExportVm(po, targetBakingFilePath)

		s.TaskRepo.UpdateStatus(po.ID, targetBakingFilePath, 0, xmlDesc, finalStatus, false, true)

		po, _ = s.TaskRepo.Get(po.ID)
		s.TaskService.SubmitResult(po)

		if ch != nil {
			close(ch)
		}
	}()
}

func (s *ExportService) ExportVm(po agentModel.Task, targetBakingFilePath string) (xml string, status consts.TaskStatus) {
	vmName := po.Vm

	dom, err := s.LibvirtService.GetVm(vmName)
	if err != nil {
		return
	}

	xml, err = dom.GetXMLDesc(0)
	if err != nil {
		return
	}

	vmDiskPath, err := s.QemuService.GetDisk(dom)
	if err != nil {
		return
	}

	err = s.LibvirtService.SafeDestroyVmByName(vmName)
	if err != nil {
		return
	}

	cmd := fmt.Sprintf(consts.CmdExportVm, vmDiskPath, targetBakingFilePath)
	_shellUtils.ExeShell(cmd)

	s.LibvirtService.BootVmByName(vmName)

	return
}
