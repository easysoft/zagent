package _domain

type RpcReq struct {
	ComputerIp   string
	ComputerPort int

	ApiPath   string
	ApiMethod string
	Data      interface{}
}
