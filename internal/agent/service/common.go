package agentService

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/libs/file"
	uuid "github.com/satori/go.uuid"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (s *CommonService) SetBuildWorkDir(build *_domain.BuildTo) {
	build.WorkDir = agentConf.Inst.WorkDir + uuid.NewV4().String() + _const.PthSep
	_fileUtils.MkDirIfNeeded(build.WorkDir)
}
