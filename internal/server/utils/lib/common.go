package serverUitls

import (
	"fmt"
	"github.com/easysoft/zagent/internal/comm/const"
)

func GenVmHostName(queueId uint,
	osPlatform consts.OsCategory, osName consts.OsType, osLang consts.OsLang) (ret string) {
	ret = fmt.Sprintf("queue%d-%s-%s-%s", queueId, osPlatform, osName, osLang)

	return
}
