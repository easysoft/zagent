package virtualService

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	_stringUtils "github.com/easysoft/zagent/pkg/lib/string"
)

type NoVncService struct {
	syncMap sync.Map
}

func NewNovncService() *NoVncService {
	srv := NoVncService{}

	//srv.genWebsockifyTokens()
	srv.launchWebsockifyService()

	return &srv
}

func (s *NoVncService) GetToken(port string) (ret v1.VncTokenResp, err error) {
	token := _stringUtils.Uuid()
	ip := agentConf.Inst.NodeIp
	ret = v1.VncTokenResp{
		Token: token,
		Host:  ip,
		Port:  port,
	}
	s.syncMap.Store(token, ret)

	//delete other token
	s.syncMap.Range(func(key, value interface{}) bool {
		if key != token {
			vncResp := value.(v1.VncTokenResp)

			if vncResp.Host == ret.Host && vncResp.Port == ret.Port {
				s.syncMap.Delete(key)
			}
		}

		return true
	})

	return
}

func (s *NoVncService) GetAddressByToken(token string) (ret v1.VncTokenResp, err error) {
	obj, ok := s.syncMap.Load(token)

	if !ok {
		return
	}

	ret = obj.(v1.VncTokenResp)

	return
}

//func (s *NoVncService) genWebsockifyTokens() {
//	port := consts.VncPortStart
//	for port <= consts.VncPortEnd {
//		portStr := strconv.Itoa(port)
//		token := _stringUtils.Uuid()
//		ip := agentConf.Inst.NodeIp
//
//		// uuid: 192.168.1.215:51800	// must has space after first :
//		content := fmt.Sprintf("%s: %s:%s", token, ip, portStr)
//
//		pth := filepath.Join(agentConf.Inst.DirToken, portStr+".txt")
//		_fileUtils.WriteFile(pth, content)
//
//		result := v1.VncTokenResp{
//			Token: token,
//			Host:    ip,
//			Port:  portStr,
//		}
//		s.syncMap.Store(portStr, result)
//
//		port++
//	}
//}

func (s *NoVncService) launchWebsockifyService() (ret v1.VncTokenResp) {
	exePath := filepath.Join(agentConf.Inst.WorkDir, "websockify/run")
	logPath := filepath.Join(agentConf.Inst.WorkDir, "websockify/nohup.log")

	cmd := fmt.Sprintf("nohup %s --token-plugin TokenFile %d > %s 2>&1 &",
		exePath, consts.WebsockifyPort, logPath)

	_shellUtils.KillProcessByName("websockify")
	_shellUtils.ExeShell(cmd)

	return
}

func (s *NoVncService) getPortByName(name string) (port int, err error) {
	cmdStr := fmt.Sprintf("virsh vncdisplay %s", name)
	out, err := exec.Command("/bin/bash", "-c", cmdStr).Output()
	if err != nil {
		return
	}

	line := strings.TrimSpace(strings.Split(string(out), "\n")[0])
	if strings.Index(line, ":") != 0 {
		err = errors.New("virsh vncdisplay cmd no putput")
		return
	}

	port, err = strconv.Atoi(line[1:])
	port += 5900

	return
}
