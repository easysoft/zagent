package hostAgentService

import (
	"fmt"
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	agentConf "github.com/easysoft/zv/internal/pkg/agent/conf"
	natHelper "github.com/easysoft/zv/internal/pkg/agent/utils/nat"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
	"github.com/jinzhu/copier"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type SetupService struct {
	syncMap sync.Map
}

func NewSetupService() *SetupService {
	srv := SetupService{}

	srv.GenWebsockifyTokens()
	srv.LaunchWebsockifyService()
	srv.LaunchNoVNCService()

	return &srv
}

func (s *SetupService) AddVmPortMap(req v1.VmPortMapReq) (resp v1.VmPortMapResp, err error) {
	resp.HostPort, err = natHelper.GetValidPort()
	if err != nil {
		return
	}

	natHelper.ForwardPort(req.VmIp, req.VmPort, resp.HostPort, req.Type)

	copier.CopyWithOption(&resp, req, copier.Option{DeepCopy: true})

	return
}

func (s *SetupService) RemoveVmPortMap(req v1.VmPortMapReq) (resp v1.VmPortMapResp, err error) {
	resp.HostPort, err = natHelper.GetValidPort()
	if err != nil {
		return
	}

	natHelper.RemoveForward(req.VmIp, req.VmPort)

	return
}

func (s *SetupService) LaunchWebsockifyService() (ret v1.VncTokenResp) {
	exePath := filepath.Join(agentConf.Inst.DirKvm, "websockify/run")
	logPath := filepath.Join(agentConf.Inst.DirKvm, "websockify/nohup.log")

	cmd := fmt.Sprintf("nohup %s --token-plugin TokenFile --token-source %s 6080 > %s 2>&1 &",
		exePath, agentConf.Inst.DirToken, logPath)

	_shellUtils.KillProcess("websockify")
	_shellUtils.ExeShell(cmd)

	return
}

func (s *SetupService) GetToken(port string) (ret v1.VncTokenResp) {
	obj, ok := s.syncMap.Load(port)

	if !ok {
		return
	}

	ret = obj.(v1.VncTokenResp)

	return
}

func (s *SetupService) GenWebsockifyTokens() {
	port := 5901
	for port <= 6000 {
		portStr := strconv.Itoa(port)

		// uuid: 192.168.1.215:5901
		content := fmt.Sprintf("%s: %s:%s", _stringUtils.Uuid(), agentConf.Inst.NodeIp, portStr)

		pth := filepath.Join(agentConf.Inst.DirToken, portStr+".txt")
		_fileUtils.WriteFile(pth, content)

		arr := strings.Split(content, ":")
		result := v1.VncTokenResp{
			Token: arr[0],
			Ip:    arr[1],
			Port:  arr[2],
		}
		s.syncMap.Store(portStr, result)

		port++
	}
}

func (s *SetupService) LaunchNoVNCService() {
	noVNCPath := filepath.Join(agentConf.Inst.DirKvm, "noVNC")
	logPath := filepath.Join(agentConf.Inst.DirKvm, "noVNC/nohup.log")

	cmd := fmt.Sprintf("nohup light-server -s  %s > %s 2>&1 &",
		noVNCPath, logPath)

	_shellUtils.KillProcess("light-server")
	_shellUtils.ExeShell(cmd)

	return
}
