package agentConf

import (
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	consts "github.com/easysoft/zagent/internal/comm/const"
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

func Init(app string) {
	if Inst.Language == "" {
		Inst.Language = "zh"
	}
	_i118Utils.InitI118(Inst.Language, app)

	Inst.Server = _httpUtils.UpdateUrl(Inst.Server)

	ip, macObj := _commonUtils.GetIp()
	Inst.MacAddress = macObj.String()
	if Inst.NodeIp == "" {
		Inst.NodeIp = ip.String()
	}

	usr, _ := user.Current()
	home := usr.HomeDir
	if Inst.Host != "" {
		home = "/home/" + Inst.User
	}

	Inst.WorkDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(home, consts.AppNameAgent))

	if Inst.RunMode == agentConst.Host {
		Inst.DirKvm = _fileUtils.AddPathSepIfNeeded(filepath.Join(home, agentConst.FolderKvm))
		Inst.DirIso = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderIso))
		Inst.DirBaking = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderBacking))
		Inst.DirImage = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderImage))
		Inst.DirToken = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, agentConst.FolderToken))

		_fileUtils.MkDirIfNeeded(Inst.DirIso)
		_fileUtils.MkDirIfNeeded(Inst.DirBaking)
		_fileUtils.MkDirIfNeeded(Inst.DirImage)
	}
}

type Config struct {
	// fot libvirt testing only
	Host string
	User string

	RunMode    agentConst.RunMode `yaml:"runMode"`
	Server     string             `yaml:"Server"`
	NodeIp     string             `yaml:"ip"`
	NodePort   int                `yaml:"port"`
	MacAddress string

	Language string
	NodeName string
	WorkDir  string
	LogDir   string

	DirKvm    string
	DirIso    string
	DirBaking string
	DirImage  string
	DirToken  string

	DB DBConfig
}

type DBConfig struct {
	Prefix string
}
