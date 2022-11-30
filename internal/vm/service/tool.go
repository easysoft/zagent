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
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
	"github.com/mholt/archiver/v3"
	"log"
	"os"
	"os/exec"
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
	s.stopTool(req.Name)

	oldVersionStr, oldVersionNum := s.getOldVersion(req.Name)
	newVersionStr, newVersionNum := s.getVersion(req.Name, req.Version)

	pass, err := s.downloadTool(req.Name, newVersionStr)
	if !pass || err != nil {
		return
	}

	s.startTool(req.Name, newVersionStr, req.Secret, req.Server, req.Ip)

	s.updateVersionFile(req.Name, newVersionStr)
	if newVersionNum > oldVersionNum {
		_logUtils.Infof("replace %s from %s to %s", req.Name, oldVersionStr, newVersionStr)
	}

	return
}

func (s *ToolService) stopTool(name string) (err error) {
	uuid := ""
	if name == "ztf" {
		uuid = consts.ZtfUuid
	} else if name == "zd" {
		uuid = consts.ZdUuid
	}

	_shellUtils.KillProcessByUUID(uuid)

	return
}

func (s *ToolService) startTool(name, version, secret, server, ip string) (err error) {
	pth := s.getToolPath(name, version)

	uuid := ""
	if name == "ztf" {
		uuid = consts.ZtfUuid
	} else if name == "zd" {
		uuid = consts.ZdUuid
	}

	s.startProcess(name, pth, uuid, ip, server, secret)

	return
}

func (s *ToolService) downloadTool(name string, version string) (pass bool, err error) {
	dir := s.getToolDir(name)

	os := _commonUtils.GetOs()
	if _commonUtils.IsWin() {
		os = fmt.Sprintf("win%d", strconv.IntSize)
	}

	zipName := name
	if name == "ztf" {
		zipName += "-server"
	}
	url := fmt.Sprintf(consts.PackageDownloadUrl, name, version, os, zipName)

	targetPath := filepath.Join(dir, "download", version+".zip")
	extractDir := filepath.Join(dir, version)

	_fileUtils.RemoveFile(targetPath)
	_, err = _fileUtils.DownloadAdv(_fileUtils.AddTimeParam(url), targetPath)
	if err != nil {
		return
	}

	md5Url := url + ".md5"
	md5Path := targetPath + ".md5"

	_fileUtils.RemoveFile(md5Path)
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

func (s *ToolService) getToolPath(name, version string) (pth string) {
	dir := s.getToolDir(name)

	exeName := name
	if name == "ztf" {
		exeName += "-server"
	}

	pth = filepath.Join(dir, version, exeName)
	if _commonUtils.IsWin() {
		pth += ".exe"
	}

	return
}

func (s *ToolService) getVersion(name, version string) (versionStr string, versionNum float64) {
	dir := s.getToolDir(name)

	if version == "" {
		versionFile := filepath.Join(dir, "version.txt")
		versionUrl := fmt.Sprintf(consts.VersionDownloadUrl, name)
		_fileUtils.Download(_fileUtils.AddTimeParam(versionUrl), versionFile)

		versionStr = strings.TrimSpace(_fileUtils.ReadFile(versionFile))
	} else {
		versionStr = version
	}

	versionNum = s.convertVersion(versionStr)

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
	versionNum = s.convertVersion(versionStr)

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

func (s *ToolService) convertVersion(str string) (version float64) {
	arr := strings.Split(str, ".")
	if len(arr) > 2 { // ignore 3th
		str = strings.Join(arr[:2], ".")
	}

	version, _ = strconv.ParseFloat(str, 64)

	return
}

func (s *ToolService) startProcess(name, execPath, uuid, ip, server, secret string) (out string, err error) {
	execDir := _fileUtils.GetAbsolutePath(filepath.Dir(execPath))

	cmdStr := ""
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		if name == "ztf" {
			tmpl := `start cmd /c %s -uuid %s -secret %s -s %s -i %s -p %d ^1^> %snohup.%s.log ^2^>^&^1`
			cmdStr = fmt.Sprintf(tmpl, execPath, uuid, secret, server, ip, consts.ZtfServicePort, consts.WorkDir, name)
		} else if name == "zd" { // set root for workdir
			tmpl := `start cmd /c %s -uuid %s -b %s -p %d ^1^> %snohup.%s.log ^2^>^&^1`
			cmdStr = fmt.Sprintf(tmpl, execPath, uuid, ip, consts.ZdServicePort, consts.WorkDir, name)
		}

		cmd = exec.Command("cmd", "/C", cmdStr)

	} else {
		if name == "ztf" {
			cmd = exec.Command("nohup", execPath, "-uuid", uuid,
				"-secret", secret, "-s", server, "-i", ip, "-p", strconv.Itoa(consts.ZtfServicePort))
		} else if name == "zd" {
			cmd = exec.Command("nohup", execPath, "-uuid", uuid, "-b", ip, "-p", strconv.Itoa(consts.ZdServicePort))
		}

		log := filepath.Join(consts.WorkDir, "nohup."+name+".log")
		f, _ := os.Create(log)

		cmd.Stdout = f
		cmd.Stderr = f
	}

	log.Println("launch tool by using cmd " + cmd.String())

	cmd.Dir = execDir
	err = cmd.Start()

	return
}
