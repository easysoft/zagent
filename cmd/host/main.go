package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	hostInit "github.com/easysoft/zagent/cmd/host/init"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	agentUtils "github.com/easysoft/zagent/internal/pkg/utils/common"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"github.com/fatih/color"
)

var (
	help    bool
	flagSet *flag.FlagSet
	runMode string
)

// @title ZAgent宿主机API文档
// @version 1.0
// @contact.name Support
// @contact.email 462626@qq.com
// @contact.url https://github.com/easysoft/zv/issues
func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet(consts.AppNameAgentHost, flag.ContinueOnError)

	flagSet.StringVar(&runMode, "t", consts.RunModeHost.ToString(), "")
	flagSet.StringVar(&agentConf.Inst.Server, "s", "http://127.0.0.1:8085", "")
	flagSet.StringVar(&agentConf.Inst.NodeName, "n", "", "")

	flagSet.StringVar(&agentConf.Inst.NodeIp, "i", "127.0.0.1", "")
	flagSet.IntVar(&agentConf.Inst.NodePort, "p", consts.AgentHostServicePort, "")

	flagSet.StringVar(&agentConf.Inst.Secret, "secret", "", "")
	flagSet.StringVar(&agentConf.Inst.Language, "l", "zh", "")

	flagSet.BoolVar(&help, "h", false, "")

	action := ""
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	switch action {
	case "help", "-h":
		agentUtils.PrintUsage(consts.AppNameAgentHost)

	default:
		start()
	}
}

func start() {
	_logUtils.Init(consts.AppNameAgentHost)

	if err := flagSet.Parse(os.Args[1:]); err == nil {
		agentConf.Inst.RunMode = consts.RunMode(runMode)
		hostInit.Init()
	}
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
