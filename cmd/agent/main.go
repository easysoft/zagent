package main

import (
	"flag"
	"github.com/easysoft/zagent/cmd/agent/router"
	"github.com/easysoft/zagent/cmd/agent/server"
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentUtils "github.com/easysoft/zagent/internal/agent/utils/common"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"
)

var (
	help     bool
	flagSet  *flag.FlagSet
	platform string
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet(agentConst.AppName, flag.ContinueOnError)

	flagSet.StringVar(&agentConf.Inst.Server, "s", "", "")
	flagSet.StringVar(&agentConf.Inst.NodeName, "n", "", "")
	flagSet.StringVar(&agentConf.Inst.NodeIp, "i", "", "")
	flagSet.IntVar(&agentConf.Inst.NodePort, "p", 0, "")
	flagSet.StringVar(&agentConf.Inst.Language, "l", "zh", "")

	flagSet.BoolVar(&help, "h", false, "")

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}

	switch os.Args[1] {
	case "help", "-h":
		agentUtils.PrintUsage()

	default:
		start()
	}
}

func start() {
	_logUtils.Init(agentConst.AppName)

	if err := flagSet.Parse(os.Args[1:]); err == nil {
		agentConf.Init()
		server.Init(router.NewRouter())
	}
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
