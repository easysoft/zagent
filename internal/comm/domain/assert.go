package domain

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

type VmAssert struct {
	ID uint

	OsCategory commConst.OsCategory
	OsType     commConst.OsType
	OsVersion  string
	OsLang     commConst.OsLang
}

type VmHost struct {
	HostId      uint
	VmBackingId uint
}
