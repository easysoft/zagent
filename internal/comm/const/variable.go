package consts

import "sync"

var (
	AuthToken = ""

	Verbose          = false
	ExistVncPortMap  = sync.Map{}
	ExistHttpPortMap = sync.Map{}
	ExistSshPortMap  = sync.Map{}
)
