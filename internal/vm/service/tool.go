package vmAgentService

import (
	"errors"
	"fmt"
	v1 "github.com/easysoft/zagent/cmd/vm/router/v1"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"github.com/mholt/archiver/v3"
	"path/filepath"
	"strconv"
	"strings"
)

type ToolService struct {
}

func NewToolService() *ToolService {
	return &ToolService{}
}

func (s *ToolService) Setup(req v1.VmServiceInstallReq) (ret v1.VmServiceInstallResp, err error) {
	oldVersionStr, oldVersionNum := s.getOldVersion(req.Name)

	newVersionStr, newVersionNum := s.DownloadVersion(req.Name)

	pass, err := s.downloadTool(req.Name, newVersionStr)
	if pass && err == nil {
		//restartTool(app, startService)

		s.updateVersionFile(req.Name, newVersionStr)

		if newVersionNum > oldVersionNum {
			_logUtils.Infof("upgrade %s from %s to %s", req.Name, oldVersionStr, newVersionStr)
		}
	}

	return
}

func (s *ToolService) GetToolPath(name string) {
	dir := s.getToolDir(name)

	pth := filepath.Join(dir, "ztf")
	if _commonUtils.IsWin() {
		pth += ".exe"
	}
}

func (s *ToolService) DownloadVersion(name string) (versionStr string, versionNum float64) {
	dir := s.getToolDir(name)

	versionFile := filepath.Join(dir, "version.txt")
	versionUrl := fmt.Sprintf(consts.VersionDownloadUrl, name)
	_fileUtils.Download(_fileUtils.AddTimeParam(versionUrl), versionFile)

	versionStr = strings.TrimSpace(_fileUtils.ReadFile(versionFile))
	versionNum = convertVersion(versionStr)

	return
}

func (s *ToolService) downloadTool(name string, version string) (pass bool, err error) {
	dir := s.getToolDir(name)

	os := _commonUtils.GetOs()
	if _commonUtils.IsWin() {
		os = fmt.Sprintf("win%d", strconv.IntSize)
	}
	url := fmt.Sprintf(consts.PackageDownloadUrl, name, version, os, name)

	targetPath := filepath.Join(dir, "download", version+".zip")
	extractDir := filepath.Join(dir, version)

	_, err = _fileUtils.DownloadAdv(_fileUtils.AddTimeParam(url), targetPath)
	if err != nil {
		return
	}

	md5Url := url + ".md5"
	md5Path := targetPath + ".md5"
	_, err = _fileUtils.DownloadAdv(_fileUtils.AddTimeParam(md5Url), md5Path)
	if err != nil {
		return
	}

	pass = checkMd5(targetPath, md5Path)
	if !pass {
		err = errors.New("check md5 failed")
		return
	}

	_fileUtils.RemoveFile(extractDir)
	_fileUtils.MkDirIfNeeded(extractDir)
	err = archiver.Unarchive(targetPath, extractDir)

	if err != nil {
		return
	}

	return
}

func checkMd5(filePth, md5Pth string) (pass bool) {
	if !_fileUtils.FileExist(filePth) {
		return false
	}

	expectVal := _fileUtils.ReadFile(md5Pth)
	actualVal, _ := _fileUtils.GetMd5(filePth)

	return strings.TrimSpace(actualVal) == strings.TrimSpace(expectVal)
}

func (s *ToolService) getOldVersion(name string) (versionStr string, versionNum float64) {
	dir := s.getToolDir(name)

	versionFile := filepath.Join(dir, "version.txt")

	versionStr = strings.TrimSpace(_fileUtils.ReadFile(versionFile))
	versionNum = convertVersion(versionStr)

	return
}

func (s *ToolService) getToolDir(name string) (dir string) {
	if name == "ztf" {
		dir = agentConf.Inst.DirZtf
	} else if name == "zd" {
		dir = agentConf.Inst.DirZd
	}

	return
}

func (s *ToolService) updateVersionFile(name string, str string) {
	dir := s.getToolDir(name)
	pth := filepath.Join(dir, "version.txt")

	_fileUtils.WriteFile(pth, str)
}

func convertVersion(str string) (version float64) {
	arr := strings.Split(str, ".")
	if len(arr) > 2 { // ignore 3th
		str = strings.Join(arr[:2], ".")
	}

	version, _ = strconv.ParseFloat(str, 64)

	return
}
