package commonService

import (
	"context"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"time"
)

func AsyncCall() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*1))
	defer cancel()
	go func(ctx context.Context) {
		// 发送HTTP请求
	}(ctx)

	select {
	case <-ctx.Done():
		_logUtils.Infof("---async call completed---")
		return
	case <-time.After(time.Duration(time.Millisecond * 1)):
		_logUtils.Infof("---async call timeout---")
		return
	}
}
