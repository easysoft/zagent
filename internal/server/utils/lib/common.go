package serverUitls

import (
	"fmt"
	commConst "github.com/easysoft/zagent/internal/comm/const"
)

func GenVmHostName(queueId uint,
	osPlatform commConst.OsCategory, osName commConst.OsType, osLang commConst.OsLang) (ret string) {
	ret = fmt.Sprintf("queue%d-%s-%s-%s", queueId, osPlatform, osName, osLang)

	return
}
