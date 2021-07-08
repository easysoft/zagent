package _const

import "os"

const (
	PthSep = string(os.PathSeparator)

	RpcPort           = 8848
	UploadDir         = "client/file/uploads"
	UploadFileMaxSize = 1000 << 20

	LanguageEN      = "en"
	LanguageZH      = "zh"
	LanguageDefault = LanguageEN

	UserTokenExpireTime = 365 * 24 * 60 * 60 * 1000
)
