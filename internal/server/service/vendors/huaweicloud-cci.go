package vendors

import (
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/easysoft/zagent/internal/server/model"
)

type HuaweiCloudCciService struct {
}

func NewHuaweiCloudCciService() *HuaweiCloudCciService {
	return &HuaweiCloudCciService{}
}

func (s HuaweiCloudCciService) Run(build model.Build) (result _domain.RpcResp) {

	return
}
