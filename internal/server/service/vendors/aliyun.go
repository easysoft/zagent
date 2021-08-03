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

func (s AliyunService) CreateInst(osType, arch string, client *ecs.Client) (id, name string, err error) {
	regionId, _, err := s.GetRegion(client)
	if err != nil {
		return
	}

	zoneId, err := s.QuerySpec(regionId, client)
	if err != nil {
		return
	}

	imageId, imageName, err := s.QueryImage("windows", "x86_64", regionId, client)

	req := &ecs.CreateInstanceRequest{
		RegionId:           tea.String(regionId),
		ImageId:            tea.String(imageId),
		InstanceType:       tea.String(zoneId),
		InternetChargeType: tea.String("PayByTraffic"),
	}

	result, err := client.CreateInstance(req)
	if err != nil {
		_logUtils.Errorf("CreateInstance image %s error %s", imageName, err.Error())
		return
	}

	id = *result.Body.InstanceId

	return
}

func (s AliyunService) QueryImage(osType, architecture, regionId string, client *ecs.Client) (
	imageId, imageName string, err error) {
	req := &ecs.DescribeImagesRequest{
		RegionId: tea.String(regionId),
		//OSType: tea.String(osType),
		//Architecture: tea.String(architecture),
	}

	result, err := client.DescribeImages(req)
	if err != nil {
		_logUtils.Errorf("DescribeImages error %s", err.Error())
		return
	} else if len(result.Body.Images.Image) == 0 {
		_logUtils.Errorf("DescribeImages found %d images, request %#v", len(result.Body.Images.Image), req)
	}

	for _, item := range result.Body.Images.Image {
		imageId = *item.ImageId
		imageName = *item.ImageName
		_logUtils.Infof("region: %s, %s", imageId, imageName)

		return
	}

	return
}

func (s AliyunService) QuerySpec(regionId string, client *ecs.Client) (zoneId string, err error) {
	req := &ecs.DescribeAvailableResourceRequest{
		RegionId:            tea.String(regionId),
		InstanceChargeType:  tea.String("PostPaid"),
		DestinationResource: tea.String("InstanceType"),
		Cores:               tea.Int32(2),
		Memory:              tea.Float32(4),
		ResourceType:        tea.String("instance"),
	}

	result, err := client.DescribeAvailableResource(req)
	if err != nil {
		_logUtils.Errorf("DescribeAvailableResource error %s", err.Error())
		return
	}

	for _, item := range result.Body.AvailableZones.AvailableZone {
		zoneId = *item.ZoneId
		_logUtils.Infof("region: %s, %s", *item.ZoneId, *item.Status)

		return
	}

	return
}

func (s AliyunService) GetRegion(client *ecs.Client) (id, name string, err error) {
	describeRegionsRequest := &ecs.DescribeRegionsRequest{
		InstanceChargeType: tea.String("PostPaid"),
		ResourceType:       tea.String("instance"),
	}

	result, err := client.DescribeRegions(describeRegionsRequest)
	if err != nil {
		_logUtils.Errorf("DescribeRegions error %s", err.Error())
		return
	}

	for _, item := range result.Body.Regions.Region {
		name = *item.LocalName
		//_logUtils.Infof("region: %s, %s", *item.RegionId, *item.LocalName)
		if strings.Index(name, "华东") > -1 {
			id = *item.RegionId

			return
		}
	}

	return
}

func (s AliyunService) CreateClient(endpoint, accessKeyId, accessKeySecret string) (
	result *ecs.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	config.Endpoint = tea.String(endpoint)
	result = &ecs.Client{}
	result, err = ecs.NewClient(config)
	if err != nil {
		_logUtils.Errorf("CreateClient error %s", err.Error())
		return
	}

	return result, err
}
