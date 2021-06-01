package _domain

import _const "github.com/easysoft/zagent/internal/pkg/const"

type Node struct {
	HostName   string
	WorkDir    string
	PublicIp   string
	PublicPort int

	Status _const.NodeStatus
}
