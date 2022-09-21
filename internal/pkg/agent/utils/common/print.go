package agentUtils

import (
	"fmt"
	consts "github.com/easysoft/zv/internal/pkg/const"
	_commonUtils "github.com/easysoft/zv/pkg/lib/common"
	_i118Utils "github.com/easysoft/zv/pkg/lib/i118"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"github.com/easysoft/zv/res/host"
	"github.com/easysoft/zv/res/vm"
	"github.com/fatih/color"
	"io/ioutil"
	"path/filepath"
)

var (
	usageFile = filepath.Join("res", "doc", "usage.txt")
)

func PrintUsage(agentName string) {
	_logUtils.PrintColor(_i118Utils.Sprintf("usage"), color.FgCyan)

	usage := ReadResData(usageFile, agentName)

	app := consts.AppName + "-" + agentName
	if _commonUtils.IsWin() {
		app += ".exe"
	}
	usage = fmt.Sprintf(usage, app)
	fmt.Printf("%s\n", usage)
}

func ReadResData(path, agentName string) string {
	isRelease := _commonUtils.IsRelease()

	var jsonStr string
	if isRelease {
		var data []byte
		if agentName == "host" {
			data, _ = hostRes.Asset(path)
		} else if agentName == "vm" {
			data, _ = vmRes.Asset(path)
		}
		jsonStr = string(data)

	} else {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			jsonStr = "fail to read " + path
		} else {
			str := string(buf)
			jsonStr = _commonUtils.RemoveBlankLine(str)
		}
	}

	return jsonStr
}
