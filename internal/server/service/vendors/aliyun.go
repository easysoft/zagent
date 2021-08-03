package vendors

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"strings"
)

type AliyunService struct {
}

func NewAliyunService() *AliyunService {
	return &AliyunService{}
}

func (s AliyunService) CreateClient(accessKeyId *string, accessKeySecret *string) (
	result *ecs.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}

	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	result = &ecs.Client{}
	result, err = ecs.NewClient(config)
	return result, err
}

func (s AliyunService) GetRegions(accessKeyId, accessKeySecret string) (id, name string, err error) {
	client, err := s.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		return
	}

	describeRegionsRequest := &ecs.DescribeRegionsRequest{
		InstanceChargeType: tea.String("PostPaid"),
		ResourceType:       tea.String("instance"),
	}

	result, err := client.DescribeRegions(describeRegionsRequest)
	if err != nil {
		return
	}

	for _, item := range result.Body.Regions.Region {
		name = *item.LocalName
		_logUtils.Infof("region: %s, %s", *item.RegionId, *item.LocalName)
		if strings.Index(name, "华东") > -1 {
			id = *item.RegionId

			return
		}
	}

	return
}
