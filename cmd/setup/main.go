package main

import (
	"fmt"

	natHelper "github.com/easysoft/zv/internal/pkg/utils/nat"
)

func main() {
	port, err := natHelper.GetUsedPortByKw("ssh")
	fmt.Println(port, err)
}
