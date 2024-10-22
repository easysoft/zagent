package _const

import "os"

const (
	PthSep = string(os.PathSeparator)

	UploadDir         = "down" + PthSep + "upload"
	UploadFileMaxSize = 1000 << 20

	LanguageEN      = "en"
	LanguageZH      = "zh"
	LanguageDefault = LanguageEN

	UserTokenExpireTime = 365 * 24 * 60 * 60 * 1000

	Authorization = "Authorization"
	Bearer        = "Bearer"
)

var (
	IsRelease = false
)
