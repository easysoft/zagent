package virtualService

import (
	"fmt"
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	agentConf "github.com/easysoft/zv/internal/pkg/conf"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_shellUtils "github.com/easysoft/zv/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zv/pkg/lib/string"
	"path/filepath"
	"strconv"
	"sync"
)

type NoVncService struct {
	syncMap sync.Map
}

func NewNoVncService() *NoVncService {
	srv := NoVncService{}

	srv.GenWebsockifyTokens()
	srv.LaunchWebsockifyService()
	srv.LaunchNoVNCService()

	return &srv
}

func (s *NoVncService) LaunchWebsockifyService() (ret v1.VncTokenResp) {
	exePath := filepath.Join(agentConf.Inst.WorkDir, "websockify/run")
	logPath := filepath.Join(agentConf.Inst.WorkDir, "websockify/nohup.log")

	cmd := fmt.Sprintf("nohup %s --token-plugin TokenFile --token-source %s %d > %s 2>&1 &",
		exePath, agentConf.Inst.DirToken, consts.WebsockifyPort, logPath)

	_shellUtils.KillProcess("websockify")
	_shellUtils.ExeShell(cmd)

	return
}

func (s *NoVncService) LaunchNoVNCService() {
	noVNCPath := filepath.Join(agentConf.Inst.WorkDir, "novnc")
	logPath := filepath.Join(agentConf.Inst.WorkDir, "novnc/nohup.log")

	cmd := fmt.Sprintf("nohup light-server -s %s -p %d > %s 2>&1 &", noVNCPath, consts.NoVncServer, logPath)

	_shellUtils.KillProcess("light-server")
	_shellUtils.ExeShell(cmd)

	return
}

func (s *NoVncService) GetToken(port string) (ret v1.VncTokenResp) {
	obj, ok := s.syncMap.Load(port)

	if !ok {
		return
	}

	ret = obj.(v1.VncTokenResp)

	return
}

func (s *NoVncService) GenWebsockifyTokens() {
	port := consts.VncPortStart
	for port <= consts.VncPortEnd {
		portStr := strconv.Itoa(port)
		token := _stringUtils.Uuid()
		ip := agentConf.Inst.NodeIp

		// uuid: 192.168.1.215:51800
		content := fmt.Sprintf("%s:%s:%s", token, ip, portStr)

		pth := filepath.Join(agentConf.Inst.DirToken, portStr+".txt")
		_fileUtils.WriteFile(pth, content)

		result := v1.VncTokenResp{
			Token: token,
			Ip:    ip,
			Port:  portStr,
		}
		s.syncMap.Store(portStr, result)

		port++
	}
}
