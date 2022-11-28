package main

import (
	"flag"
	vmInit "github.com/easysoft/zagent/cmd/vm/init"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	agentUtils "github.com/easysoft/zagent/internal/pkg/utils/common"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"
)

var (
	help    bool
	flagSet *flag.FlagSet
	runMode string
)

// @title ZAgent虚拟机API文档
// @version 1.0
// @contact.name API Support
// @contact.url https://github.com/easysoft/zv/issues
// @contact.email 462626@qq.com
func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet(consts.AppNameAgentVm, flag.ContinueOnError)

	flagSet.StringVar(&runMode, "t", consts.RunModeVm.ToString(), "")
	flagSet.StringVar(&agentConf.Inst.Server, "s", "http://127.0.0.1:55001", "")
	flagSet.StringVar(&agentConf.Inst.NodeName, "n", "", "")
	flagSet.StringVar(&agentConf.Inst.NodeIp, "i", "127.0.0.1", "")
	flagSet.IntVar(&agentConf.Inst.NodePort, "p", consts.AgentVmServicePort, "")
	flagSet.StringVar(&agentConf.Inst.Secret, "secret", "", "")
	flagSet.StringVar(&agentConf.Inst.Language, "l", "zh", "")

	flagSet.BoolVar(&help, "h", false, "")

	action := ""
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	switch action {
	case "help", "-h":
		agentUtils.PrintUsage(consts.AppNameAgentVm)

	default:
		start()
	}
}

func start() {
	_logUtils.Init(consts.AppNameAgentVm)

	if err := flagSet.Parse(os.Args[1:]); err == nil {
		agentConf.Inst.RunMode = consts.RunMode(runMode)
		vmInit.Init()
	}
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
