package main

import (
	"flag"
	"fmt"
	swaggerUtils "github.com/easysoft/zagent/internal/pkg/lib/swagger"
	serverConf "github.com/easysoft/zagent/internal/server/conf"
	_ "github.com/easysoft/zagent/res/agent-vm/docs"
	"github.com/kataras/iris/v12"
	"log"
	"os"
	"time"
)

var ()

func main() {
	ip := "0.0.0.0"
	port := flag.Int("p", 8087, "端口")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] [command]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  -p <port> default is server 8085\n")
		fmt.Fprintf(os.Stderr, "    Service Port\n")
		fmt.Fprintf(os.Stderr, "\n")
		//flag.PrintDefaults()
	}
	flag.Parse()

	app := iris.Default()
	app.Logger().SetLevel(serverConf.Inst.LogLevel)

	swaggerUtils.InitSwaggerDocs(*port, app)

	// start the service
	err := app.Run(
		iris.Addr(fmt.Sprintf("%s:%d", ip, *port)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat(time.RFC3339),
	)

	if err != nil {
		log.Fatalf("fail to start service @%s:%d", ip, *port)
		return
	}
}
