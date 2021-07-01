package agentConf

import (
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/lib/common"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	"os/user"
	"path/filepath"
)

var (
	Inst = Config{}
)

func Init() {
	if Inst.Language == "" {
		Inst.Language = "zh"
	}
	_i118Utils.InitI118(Inst.Language, agentConst.AppName)

	Inst.Server = _httpUtils.UpdateUrl(Inst.Server)

	ip, _ := _commonUtils.GetIp()
	if Inst.NodeIp == "" {
		Inst.NodeIp = ip.String()
	}
	if Inst.NodePort == 0 {
		Inst.NodePort = _const.RpcPort
	}

	usr, _ := user.Current()
	home := usr.HomeDir
	if Inst.Host != "" {
		home = "/home/" + Inst.User
	}

	Inst.WorkDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(home, agentConst.AppName))
	Inst.DirKvm = _fileUtils.AddPathSepIfNeeded(filepath.Join(home, agentConst.FolderKvm))
	Inst.DirIso = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderIso))
	Inst.DirBase = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderBase))
	Inst.DirImage = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderImage))
	//Inst.DirTempl = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderTempl))
	//Inst.DirDef = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderDef))

	_fileUtils.MkDirIfNeeded(Inst.DirIso)
	_fileUtils.MkDirIfNeeded(Inst.DirBase)
	_fileUtils.MkDirIfNeeded(Inst.DirImage)
}

type Config struct {
	Host string
	User string

	RunMode  agentConst.RunMode `yaml:"runMode"`
	Server   string             `yaml:"Server"`
	NodeIp   string             `yaml:"ip"`
	NodePort int                `yaml:"port"`

	Language string
	NodeName string
	WorkDir  string
	LogDir   string

	DirKvm   string
	DirIso   string
	DirImage string
	DirBase  string
	//DirDef   string
	//DirTempl string
}
