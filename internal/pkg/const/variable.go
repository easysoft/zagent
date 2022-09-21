package consts

import (
	"sync"
	"time"
)

var (
	AuthSecret  = ""
	AuthToken   = ""
	ExpiredDate = time.Now()

	Verbose          = false
	ExistVncPortMap  = sync.Map{}
	ExistHttpPortMap = sync.Map{}
	ExistSshPortMap  = sync.Map{}
)
