package main

import (
	"flag"
	"fmt"

	_const "github.com/easysoft/zagent/pkg/const"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
)

var (
	softList     string
	action       string
	forceInstall bool
	check        bool
)

func main() {

	flag.StringVar(&softList, "s", "", "")
	flag.BoolVar(&forceInstall, "r", false, "")
	flag.BoolVar(&check, "c", false, "")
	flag.Parse()

	_logUtils.Init("")

	resDir := fmt.Sprintf("res%s", _const.PthSep)
	fmt.Println(softList, forceInstall, check, resDir)

	if check {
		cmd := fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s) -c`, "https://raw.githubusercontent.com/easysoft/zenagent/main/res/setup/zagent.sh")

		_shellUtils.ExeShellWithOutput(cmd)
	} else if softList != "" {
		cmd := fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s) -s %s`, "https://raw.githubusercontent.com/easysoft/zenagent/main/res/setup/zagent.sh", softList)
		if forceInstall {
			cmd = fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s) -s %s -r`, "https://raw.githubusercontent.com/easysoft/zenagent/main/res/setup/zagent.sh", softList)
		}

		_shellUtils.ExeShellWithOutput(cmd)
	} else {
		cmd := fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s)`, "https://raw.githubusercontent.com/easysoft/zenagent/main/res/setup/zagent.sh")

		if forceInstall {
			cmd = fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s) -r`, "https://raw.githubusercontent.com/easysoft/zenagent/main/res/setup/zagent.sh")
		}

		_shellUtils.ExeShellWithOutput(cmd)
	}
}
