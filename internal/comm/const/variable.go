package consts

import (
	"sync"
	"time"
)

var (
	AuthToken   = ""
	ExpiredDate = time.Now()

	Verbose          = false
	ExistVncPortMap  = sync.Map{}
	ExistHttpPortMap = sync.Map{}
	ExistSshPortMap  = sync.Map{}
)
