package agentService

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	uuid "github.com/satori/go.uuid"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (s *CommonService) SetBuildWorkDir(build *commDomain.IntfTest) {
	build.WorkDir = agentConf.Inst.WorkDir + uuid.NewV4().String() + _const.PthSep
	_fileUtils.MkDirIfNeeded(build.WorkDir)
}
