package main

import (
	"flag"
	vmInit "github.com/easysoft/zagent/cmd/agent-vm/init"
	"github.com/easysoft/zagent/internal/agent/conf"
	"github.com/easysoft/zagent/internal/agent/utils/common"
	"github.com/easysoft/zagent/internal/agent/utils/const"
	consts "github.com/easysoft/zagent/internal/comm/const"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/lib/log"
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

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet(consts.AppNameAgent, flag.ContinueOnError)

	flagSet.StringVar(&runMode, "t", agentConst.Vm.ToString(), "")
	flagSet.StringVar(&agentConf.Inst.Server, "s", "http://192.168.0.107:8085", "")
	flagSet.StringVar(&agentConf.Inst.NodeName, "n", "", "")
	flagSet.StringVar(&agentConf.Inst.NodeIp, "i", "", "")
	flagSet.IntVar(&agentConf.Inst.NodePort, "p", _const.RpcPort, "")
	flagSet.StringVar(&agentConf.Inst.Language, "l", "zh", "")

	flagSet.BoolVar(&help, "h", false, "")

	action := ""
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	switch action {
	case "help", "-h":
		agentUtils.PrintUsage()

	default:
		start()
	}
}

func start() {
	_logUtils.Init(consts.AppNameAgent)

	if err := flagSet.Parse(os.Args[1:]); err == nil {
		agentConf.Inst.RunMode = agentConst.RunMode(runMode)
		vmInit.Init()
	}
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
