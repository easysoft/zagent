package hostAgentService

import (
	"fmt"
	"path/filepath"
	"time"

	agentModel "github.com/easysoft/zagent/internal/host/model"
	hostRepo "github.com/easysoft/zagent/internal/host/repo"
	kvmService "github.com/easysoft/zagent/internal/host/service/kvm"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/job"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	"github.com/gofrs/uuid"
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

		s.TaskRepo.UpdateStatus(po.ID, targetBakingFilePath, 0, "", consts.Inprogress, true, false)

		xmlDesc, finalStatus := s.ExportVm(po, targetBakingFilePath)

		s.TaskRepo.UpdateStatus(po.ID, targetBakingFilePath, 0, xmlDesc, finalStatus, false, true)

		po, _ = s.TaskRepo.Get(po.ID)
		s.TaskService.SubmitResult(po)

		if ch != nil {
			close(ch)
		}
	}()
}

func (s *ExportService) ExportVm(task agentModel.Task, targetBakingFilePath string) (xml string, status consts.TaskStatus) {
	vmName := task.Vm

	dom, err := s.LibvirtService.GetVm(vmName)
	if err != nil {
		status = consts.Failed
		return
	}

	xml, err = dom.GetXMLDesc(0)
	if err != nil {
		status = consts.Failed
		return
	}

	vmDiskPath, err := s.QemuService.GetDisk(dom)
	if err != nil {
		status = consts.Failed
		return
	}

	backingSize := s.QemuService.GetBackingFileSize(vmDiskPath)

	bizErr := s.LibvirtService.TryThenForceDestroyVmByName(vmName)
	if bizErr != nil {
		status = consts.Failed
		return
	}

	uuidStr := uuid.Must(uuid.NewV4()).String()
	srcVmDiskSize, _ := _fileUtils.GetFileSize(vmDiskPath)
	srcVmDiskSize += backingSize

	_fileUtils.RemoveFile(targetBakingFilePath)

	// check disk size
	completed := false
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for !completed {
			size, _ := _fileUtils.GetFileSize(targetBakingFilePath)
			rate := float64(size) / float64(srcVmDiskSize)

			job.SaveTaskStatus(&job.TaskStatus, task.ID, rate, 0)

			fmt.Print("\033[G\033[K")
			fmt.Printf("Converting %s %d / %d bytes (%d%%)\u001B[K\n", vmDiskPath, size, srcVmDiskSize, int(100*rate))
			fmt.Print("\033[A")

			<-ticker.C
		}

		_logUtils.Infof("complete converting vm to image")
	}()

	ch := make(chan string)
	go func() {
		cmd := fmt.Sprintf(consts.CmdExportVm, vmDiskPath, targetBakingFilePath)
		_, e := _shellUtils.ExeShell(cmd)

		msg := consts.Completed.ToString()
		if e != nil {
			msg = consts.Error.ToString()
		}

		ch <- msg
	}()

	select {
	case val := <-ch:
		status = consts.TaskStatus(val)
		completed = true

	case <-time.After(consts.ExportVmTimeout * time.Second):
		status = consts.Timeout
		_shellUtils.KillProcessByUUID(uuidStr)

		completed = true
		return
	}

	err = s.LibvirtService.BootVmByName(vmName)

	return
}
