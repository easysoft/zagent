package agentConf

import (
	"github.com/easysoft/zagent/internal/agent/domain"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/libs/common"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/libs/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/libs/http"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/libs/i118"
	"os/user"
	"path/filepath"
)

var (
	Inst = domain.Config{}
)

func Init() {
	_i118Utils.InitI118(Inst.Language, "agent")

	Inst.Server = _httpUtils.UpdateUrl(Inst.Server)

	ip, _ := _commonUtils.GetIp()
	if Inst.NodeIp == "" {

		Inst.NodeIp = ip.String()
	}
	if Inst.NodePort == 0 {
		Inst.NodePort = _const.RpcPort
	}

	usr, _ := user.Current()
	Inst.WorkDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(usr.HomeDir, agentConst.AppName))

	Inst.DirKvm = _fileUtils.AddPathSepIfNeeded(filepath.Join(usr.HomeDir, agentConst.FolderKvm))
	Inst.DirIso = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderIso))
	Inst.DirImage = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderImage))
	Inst.DirTempl = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderTempl))
	Inst.DirDef = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderDef))
}
