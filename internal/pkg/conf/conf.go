package agentConf

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	_const "github.com/easysoft/zv/pkg/const"
	_commonUtils "github.com/easysoft/zv/pkg/lib/common"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	_i118Utils "github.com/easysoft/zv/pkg/lib/i118"
	"path/filepath"
)

var (
	Inst = Config{}
)

func Init(app string) {
	_const.IsRelease = _commonUtils.IsRelease()
	if Inst.Language == "" {
		Inst.Language = "zh"
	}
	_i118Utils.InitI118(Inst.Language, app)

	Inst.Server = _httpUtils.AddUrlPostFixIfNeeded(Inst.Server)

	ip, macObj := _commonUtils.GetIp()
	Inst.MacAddress = macObj.String()
	if Inst.NodeIp == "" {
		Inst.NodeIp = ip.String()
	}

	home, _ := _fileUtils.GetUserHome()
	Inst.WorkDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(home, consts.AppName))

	if Inst.RunMode == consts.RunModeHost {
		Inst.DirToken = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.WorkDir, consts.FolderToken))

		Inst.DirKvm = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.WorkDir, consts.FolderKvm))
		Inst.DirIso = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, consts.AppName, consts.FolderIso))
		Inst.DirBaking = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, consts.AppName, consts.FolderBacking))
		Inst.DirImage = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, consts.AppName, consts.FolderImage))

		_fileUtils.MkDirIfNeeded(Inst.DirToken)
		_fileUtils.MkDirIfNeeded(Inst.DirIso)
		_fileUtils.MkDirIfNeeded(Inst.DirBaking)
		_fileUtils.MkDirIfNeeded(Inst.DirImage)
	}
}

type Config struct {
	// fot libvirt testing only
	Host string
	User string

	RunMode    consts.RunMode `yaml:"runMode"`
	Server     string         `yaml:"Server"`
	NodeIp     string         `yaml:"ip"`
	NodePort   int            `yaml:"port"`
	MacAddress string

	Secret   string
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
