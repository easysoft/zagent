package _shellUtils

import (
	"bufio"
	"bytes"
	"fmt"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/lib/common"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

func ExeSysCmd(cmdStr string) (string, error) {
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	output := out.String()

	return output, err
}

func ExeShell(cmdStr string) (string, error) {
	return ExeShellInDir(cmdStr, "")
}

func ExeShellInDir(cmdStr string, dir string) (ret string, err error) {
	ret, err, _ = ExeShellInDirWithPid(cmdStr, dir)
	return
}

func ExeShellWithPid(cmdStr string) (string, error, int) {
	return ExeShellInDirWithPid(cmdStr, "")
}

func ExeShellInDirWithPid(cmdStr string, dir string) (ret string, err error, pid int) {
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}
	if dir != "" {
		cmd.Dir = dir
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		_logUtils.Error(_i118Utils.Sprintf("fail_to_exec_command", cmdStr, cmd.Dir, err.Error()))
	}

	pid = cmd.Process.Pid
	ret = _stringUtils.TrimAll(out.String())
	return
}

func ExeShellWithOutput(cmdStr string) ([]string, error) {
	return ExeShellWithOutputInDir(cmdStr, "")
}

func ExeShellWithOutputInDir(cmdStr string, dir string) ([]string, error) {
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	if dir != "" {
		cmd.Dir = dir
	}

	output := make([]string, 0)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		_logUtils.Error(_i118Utils.Sprintf("fail_to_exec_command", cmdStr, cmd.Dir, err.Error()))
		return output, err
	}

	cmd.Start()

	if err != nil {
		_logUtils.Error(_i118Utils.Sprintf("fail_to_exec_command", cmdStr, cmd.Dir, err.Error()))
		return output, err
	}

	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		_logUtils.Info(strings.TrimRight(line, "\n"))
		output = append(output, line)
	}

	cmd.Wait()

	return output, nil
}

func GetProcess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if _commonUtils.IsWin() {
		tmpl = `tasklist`
		cmdStr = fmt.Sprintf(tmpl)

		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		tmpl = `ps -ef | grep "%s" | grep -v "grep" | awk '{print $2}'`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	output := ""
	if _commonUtils.IsWin() {
		arr := strings.Split(out.String(), "\n")
		for _, line := range arr {
			if strings.Index(line, app+".exe") > -1 {
				arr2 := regexp.MustCompile(`\s+`).Split(line, -1)
				output = arr2[1]
				break
			}
		}
	} else {
		output = out.String()
	}

	return output, err
}

func KillProcess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if _commonUtils.IsWin() {
		// tasklist | findstr ztf.exe
		tmpl = `taskkill.exe /f /im %s.exe`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		tmpl = `ps -ef | grep '%s' | grep -v "grep" | awk '{print $2}' | xargs kill -9`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	output := out.String()

	return output, err
}

func KillProcessById(pid int) {
	cmdStr := fmt.Sprintf("kill -9 %d", pid)
	ExeShell(cmdStr)
}
