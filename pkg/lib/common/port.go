package _commonUtils

import (
	consts "github.com/easysoft/zv/internal/pkg/const"
	"sync"
)

func GetHttpPort() (ret int) {
	ret = GetValidPort(58000, 58099, &consts.ExistHttpPortMap)
	return
}
func GetSshPort() (ret int) {
	ret = GetValidPort(52200, 52299, &consts.ExistSshPortMap)
	return
}
func GetVncPort() (ret int) {
	ret = GetValidPort(59000, 59099, &consts.ExistVncPortMap)
	return
}

func RemoveHttpPort(port int) {
	consts.ExistHttpPortMap.Delete(port)
	return
}
func RemoveSshPort(port int) {
	consts.ExistSshPortMap.Delete(port)
	return
}
func RemoveVncPort(port int) {
	consts.ExistVncPortMap.Delete(port)
	return
}

func GetValidPort(start int, end int, existPortMap *sync.Map) (ret int) {
	newPort := 0

	for i := 0; i < 99; i++ {
		port := start + i
		if port > end {
			break
		}

		if _, ok := (*existPortMap).Load(port); !ok {
			newPort = port
			break
		}
	}

	if newPort > 0 {
		ret = newPort
		(*existPortMap).Store(ret, true)
	}
	return
}
