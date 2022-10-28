package main

import (
	"fmt"

	natHelper "github.com/easysoft/zv/internal/pkg/utils/nat"
)

func main() {
	port, err := natHelper.GetThreadPort("ssh")
	fmt.Println(port, err)
}
