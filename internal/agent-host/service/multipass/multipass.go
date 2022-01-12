package multiPassService

import (
	"errors"
	"fmt"
	v1 "github.com/easysoft/zv/cmd/agent-host/router/v1"
	agentConf "github.com/easysoft/zv/internal/agent/conf"
	"github.com/easysoft/zv/internal/comm/domain"
	_fileUtils "github.com/easysoft/zv/internal/pkg/lib/file"
	_logUtils "github.com/easysoft/zv/internal/pkg/lib/log"
	_shellUtils "github.com/easysoft/zv/internal/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zv/internal/pkg/lib/string"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

const (
	mpTokenPrefix = "mp"

	cmdMpls      = "multipass ls"
	cmdMpInfo    = "multipass info %s"
	cmdMpStart   = "multipass start %s"
	cmdMpStop    = "multipass stop %s"
	cmdMpSuspend = "multipass suspend %s"
	cmdMpReboot  = "multipass restart %s"
	cmdMpDelete  = "multipass delete %s"
	cmdMpPurge   = "multipass purge"
	cmdMpStopAll = "multipass stop --all"
)

var cmdMpLaunch = "multipass launch "

type MultiPassService struct {
	syncMap sync.Map
}

func (s *MultiPassService) GetVms() (doms []domain.MultiPass, err error) {
	outRets, err := _shellUtils.ExeShellWithOutput(cmdMpls)
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	dom := domain.MultiPass{}
	var rets []string
	for i := 1; i < len(outRets)-1; i++ {
		rets = strings.Fields(outRets[i])

		dom.Name = rets[0]
		dom.State = rets[1]
		dom.IPv4 = rets[2]
		for i := 3; i < len(rets); i++ {
			dom.Image += rets[i]
		}

		doms = append(doms, dom)
	}
	return
}

func (s *MultiPassService) VmInfo(name string) (dom domain.MultiPass) {
	outRets, err := _shellUtils.ExeShellWithOutput(fmt.Sprintf(cmdMpInfo, name))
	if strings.Contains(outRets[0], "info failed:") || err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	vmInfoMap := parseOutput(outRets[:len(outRets)-1])
	dom = domain.MultiPass{
		Name:  vmInfoMap["Name"],
		State: vmInfoMap["State"],
		IPv4:  vmInfoMap["IPv4"],
	}

	return
}

func parseOutput(lines []string) (vmInfoMap map[string]string) {
	vmInfoMap = make(map[string]string, len(lines))

	for _, v := range lines {
		rets := strings.Fields(v)
		vmInfoMap[rets[0]] = rets[1]
	}
	return
}

func (s *MultiPassService) CreateVm(name, cpus, disk, mem, filePath string) (dom domain.MultiPass, err error) {
	if name != "" {
		vm := s.VmInfo(name)
		if vm.Name == "" {
			msg := "vm %s not found"
			_logUtils.Errorf(msg, name)
			err = errors.New(msg)
			return
		}
		cmdMpLaunch = cmdMpLaunch + fmt.Sprintf("-n %s ", name)
	}

	if cpus != "" {
		cmdMpLaunch = cmdMpLaunch + fmt.Sprintf("-c %s ", cpus)
	}
	if disk != "" {
		cmdMpLaunch = cmdMpLaunch + fmt.Sprintf("-d %s ", disk)
	}
	if mem != "" {
		cmdMpLaunch = cmdMpLaunch + fmt.Sprintf("-m %s ", mem)
	}
	if filePath != "" {
		filePath = filepath.Join(agentConf.Inst.DirImage, filePath)

		cmdMpLaunch = cmdMpLaunch + fmt.Sprintf("file://%s ", filePath)
	}
	outRets, err := _shellUtils.ExeShellWithOutput(cmdMpLaunch)
	if strings.Contains(outRets[0], "launch failed") == true || err != nil {
		_logUtils.Errorf(err.Error())
		return
	}

	return
}

func (s *MultiPassService) RebootVmByName(name string) (dom domain.MultiPass, err error) {
	_, err = _shellUtils.ExeShellWithOutput(fmt.Sprintf(cmdMpReboot, name))
	if s.VmInfo(name).State != "Running" && err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	return
}

func (s *MultiPassService) DestroyVm(name string) (dom domain.MultiPass, err error) {
	_, err = _shellUtils.ExeShellWithOutput(fmt.Sprintf(cmdMpDelete+"&&"+cmdMpPurge, name))
	if s.VmInfo(name).Name == "" && err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	return
}

func (s *MultiPassService) SuspendVmByName(name string) (dom domain.MultiPass, err error) {
	_, err = _shellUtils.ExeShellWithOutput(fmt.Sprintf(cmdMpSuspend, name))
	if s.VmInfo(name).State != "Suspended" && err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	return
}

func (s *MultiPassService) ResumeVmByName(name string) (dom domain.MultiPass, err error) {
	_, err = _shellUtils.ExeShellWithOutput(fmt.Sprintf(cmdMpStart, name))
	if s.VmInfo(name).State != "Running" && err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	return
}

func (s *MultiPassService) GenWebsockifyTokens() {
	vms, _ := s.GetVms()
	port := 5901
	for _, v := range vms {
		if v.State != "Running" {
			continue
		}
		portStr := strconv.Itoa(port)

		// uuid: 192.168.1.215:5901
		content := fmt.Sprintf("%s: %s:%s", _stringUtils.NewUuid(), agentConf.Inst.NodeIp, portStr)

		pth := filepath.Join(agentConf.Inst.DirToken, mpTokenPrefix+portStr+".txt")
		_fileUtils.WriteFile(pth, content)

		arr := strings.Split(content, ":")
		result := v1.VncTokenResp{
			Token: arr[0],
			Ip:    v.IPv4,
			Port:  arr[2],
		}
		s.syncMap.Store(mpTokenPrefix+portStr, result)

		port++
	}
}

func (s *MultiPassService) GetToken(port string) (ret v1.VncTokenResp) {
	s.GenWebsockifyTokens()
	obj, ok := s.syncMap.Load(mpTokenPrefix + port)

	if !ok {
		return
	}

	ret = obj.(v1.VncTokenResp)

	return
}
