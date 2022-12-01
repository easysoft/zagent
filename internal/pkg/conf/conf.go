package agentConf

import (
	"path/filepath"

	consts "github.com/easysoft/zagent/internal/pkg/const"
	netUtils "github.com/easysoft/zagent/internal/pkg/utils/net"
	_const "github.com/easysoft/zagent/pkg/const"
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	_i118Utils "github.com/easysoft/zagent/pkg/lib/i118"
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

	if app == consts.AppNameAgentHost {
		Inst.WebsockifyPort, _ = netUtils.GetUsedPortByKeyword("websockify", consts.WebsockifyPort)
	} else if app == consts.AppNameAgentVm {
		Inst.ZtfPort, _ = netUtils.GetUsedPortByKeyword("ztf", consts.ZtfServicePort)
		Inst.ZdPort, _ = netUtils.GetUsedPortByKeyword("zd", consts.ZdServicePort)
	}

	home, _ := _fileUtils.GetUserHome()
	Inst.WorkDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(home, consts.AppName))

	if Inst.RunMode == consts.RunModeHost {
		consts.WorkDir = Inst.WorkDir
		consts.DownloadDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.WorkDir, consts.FolderDownload))
		consts.NovncDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.WorkDir, consts.FolderNovnc))
		consts.WebsockifyDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.WorkDir, consts.FolderWebsockify))

		Inst.DirKvm = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.WorkDir, consts.FolderKvm))
		Inst.DirIso = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, consts.FolderIso))
		Inst.DirBaking = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, consts.FolderBacking))
		Inst.DirImage = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.DirKvm, consts.FolderImage))

		_fileUtils.MkDirIfNeeded(consts.DownloadDir)
		_fileUtils.MkDirIfNeeded(Inst.DirIso)
		_fileUtils.MkDirIfNeeded(Inst.DirBaking)
		_fileUtils.MkDirIfNeeded(Inst.DirImage)

	} else if Inst.RunMode == consts.RunModeVm {
		if Inst.NodeIp == "" {
			ip, _ := _commonUtils.GetIp()
			Inst.NodeIp = ip.String()
		}

		consts.WorkDir = Inst.WorkDir

		Inst.DirZtf = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.WorkDir, consts.FolderZtf))
		Inst.DirZd = _fileUtils.AddPathSepIfNeeded(filepath.Join(Inst.WorkDir, consts.FolderZd))

		_fileUtils.MkDirIfNeeded(filepath.Join(Inst.DirZtf, "download"))
		_fileUtils.MkDirIfNeeded(filepath.Join(Inst.DirZd, "download"))
	}
}

type Config struct {
	// fot libvirt testing only
	Host string
	User string

	RunMode        consts.RunMode
	Server         string
	NodeIp         string
	NodePort       int
	WebsockifyPort int
	ZtfPort        int
	ZdPort         int
	MacAddress     string

	Secret   string
	Language string
	NodeName string
	WorkDir  string
	LogDir   string

	DirKvm    string
	DirIso    string
	DirBaking string
	DirImage  string

	DirZtf string
	DirZd  string

	DB DBConfig
}

type DBConfig struct {
	Prefix string
}
