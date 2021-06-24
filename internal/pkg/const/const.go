package _const

import "os"

const (
	PthSep = string(os.PathSeparator)

	RpcPort   = 8848
	UploadDir = "uploads"

	LanguageEN      = "en"
	LanguageZH      = "zh"
	LanguageDefault = LanguageEN

	RegisterExpireTime = 5  // min
	WaitForExecTime    = 60 // min

	RetryTime          = 3
	AgentRunTime       = 20 // sec
	AgentCheckInterval = 10 // sec

	SepOfMacAddress = ":"

	UserTokenExpireTime   = 365 * 24 * 60 * 60 * 1000
	WebCheckQueueInterval = 5
)
