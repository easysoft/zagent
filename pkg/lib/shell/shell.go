package _shellUtils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	_commonUtils "github.com/easysoft/zagent/pkg/lib/common"
	_i118Utils "github.com/easysoft/zagent/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
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
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		err = errors.New(err.Error() + ", " + stderr.String())
		_logUtils.Error(_i118Utils.Sprintf("fail_to_exec_command", cmdStr, cmd.Dir, ret))
	} else {
		ret = out.String()
	}

	pid = cmd.Process.Pid

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

	stderr, err := cmd.StderrPipe()
	if err != nil {
		_logUtils.Error(_i118Utils.Sprintf("fail_to_exec_command", cmdStr, cmd.Dir, err.Error()))
		return output, err
	}

	cmd.Start()

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

	buf := new(bytes.Buffer)
	buf.ReadFrom(stderr)

	output = append(output, buf.String())

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

func KillProcessByName(app string) (string, error) {
	var cmd *exec.Cmd

	cmdStr := ""
	if _commonUtils.IsWin() {
		// tasklist | findstr ztf.exe
		cmdStr = fmt.Sprintf(`taskkill.exe /f /im %s.exe`, app)

		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmdStr = fmt.Sprintf(`ps -ef | grep '%s' | grep -v "grep" | awk '{print $2}' | xargs kill -9`, app)

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
