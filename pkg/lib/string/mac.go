package _stringUtils

import (
	"fmt"
)

func AddSepForMacAddress(sep string) (ret string) {
	// 08 00 27 0A 38 2A
	ret = fmt.Sprintf("%s:%s:%s:%s:%s:%s", sep[0:2], sep[2:4], sep[4:6], sep[6:8], sep[8:10], sep[10:12])
	return
}
