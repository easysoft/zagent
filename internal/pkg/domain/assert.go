package domain

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
)

type VmAssert struct {
	ID uint

	OsCategory consts.OsCategory
	OsType     consts.OsType
	OsVersion  string
	OsLang     consts.OsLang
}

type VmHost struct {
	HostId      uint
	VmBackingId uint
}
