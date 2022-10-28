package main

import (
	"fmt"
	netUtils "github.com/easysoft/zv/internal/pkg/utils/net"
)

func main() {
	port, err := netUtils.GetUsedPortByKeyword("ssh", 22)
	fmt.Println(port, err)
}
