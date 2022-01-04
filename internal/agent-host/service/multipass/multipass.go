package multiPassService

import (
	"github.com/easysoft/zv/internal/comm/domain"
	_logUtils "github.com/easysoft/zv/internal/pkg/lib/log"
	_shellUtils "github.com/easysoft/zv/internal/pkg/lib/shell"
	"strings"
)

type MultiPassService struct {
}

func (s *MultiPassService) ListVm() (doms []domain.MultiPass, err error) {
	outRets, err := _shellUtils.ExeShellWithOutput("multipass ls")
	if err != nil {
		_logUtils.Errorf(err.Error())
		return
	}
	dom := domain.MultiPass{}
	var rets []string
	for i := 1; i < len(outRets)-1; i++ {
		rets = strings.Fields(outRets[i])

		dom.Name = rets[0]
		dom.State = rets[1]
		dom.IPv4 = rets[2]
		dom.Image = rets[3] + rets[4]
		doms = append(doms, dom)
	}
	return
}
