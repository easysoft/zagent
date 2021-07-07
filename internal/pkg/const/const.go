package _const

import "os"

const (
	PthSep = string(os.PathSeparator)

	RpcPort   = 8848
	UploadDir = "uploads"

	LanguageEN      = "en"
	LanguageZH      = "zh"
	LanguageDefault = LanguageEN

	UserTokenExpireTime   = 365 * 24 * 60 * 60 * 1000
	WebCheckQueueInterval = 1 * 60
)
