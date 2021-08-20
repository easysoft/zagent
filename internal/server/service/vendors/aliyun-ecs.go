package vendors

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	testconst "github.com/easysoft/zagent/cmd/test/_const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"strings"
)

type AliyunEcsService struct {
}

func NewAliyunEcsService() *AliyunEcsService {
	return &AliyunEcsService{}
}

func (s AliyunEcsService) CreateInst(vmName, imageName string, client *ecs.Client) (id, name string, err error) {
	regionId, _, err := s.GetRegion(client)
	if err != nil {
		return
	}

	zoneId, err := s.QuerySpec(regionId, client)
	if err != nil {
		return
	}

	imageId, err := s.QueryImage(imageName, regionId, client)

	req := &ecs.CreateInstanceRequest{
		InstanceName:       tea.String(vmName),
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
	name = vmName

	return
}

func (s AliyunEcsService) RemoveInst(id string, ecsClient *ecs.Client) (err error) {
	req := &ecs.DeleteInstanceRequest{
		InstanceId: tea.String(id),
	}
	_, err = ecsClient.DeleteInstance(req)
	if err != nil {
		_logUtils.Errorf("DeleteInstance %s error %s", id, err.Error())
		return
	}

	return
}

func (s AliyunEcsService) QueryInst(id string, client *ecs.Client) (nodeIp, macAddress string, err error) {
	req := &ecs.DescribeInstancesRequest{
		InstanceIds: tea.String(id),
	}

	resp, err := client.DescribeInstances(req)
	if err != nil {
		_logUtils.Errorf("DescribeInstances %s error %s", id, err.Error())
		return
	}

	nodeIp = *resp.Body.Instances.Instance[0].NetworkInterfaces.NetworkInterface[0].PrimaryIpAddress
	macAddress = *resp.Body.Instances.Instance[0].NetworkInterfaces.NetworkInterface[0].MacAddress

	return
}
func (s AliyunEcsService) QueryVnc(id string, client *ecs.Client) (url string, err error) {
	req := &ecs.DescribeInstanceVncUrlRequest{
		InstanceId: tea.String(id),
	}

	resp, err := client.DescribeInstanceVncUrl(req)
	if err != nil {
		_logUtils.Errorf("DescribeInstances %s error %s", id, err.Error())
		return
	}

	url = *resp.Body.VncUrl

	return
}

func (s AliyunEcsService) QueryImage(imageName, regionId string, client *ecs.Client) (imageId string, err error) {
	req := &ecs.DescribeImagesRequest{
		RegionId:  tea.String(regionId),
		ImageName: tea.String(imageName),
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

func (s AliyunEcsService) QuerySpec(regionId string, client *ecs.Client) (zoneId string, err error) {
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

func (s AliyunEcsService) GetRegion(client *ecs.Client) (id, name string, err error) {
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
		id = *item.RegionId
		//_logUtils.Infof("region: %s, %s", *item.RegionId, *item.LocalName)
		if strings.Index(id, testconst.ALIYUN_REGION) > -1 {
			id = *item.RegionId

			return
		}
	}

	return
}

func (s AliyunEcsService) CreateClient(endpoint, accessKeyId, accessKeySecret string) (
	result *ecs.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	config.Endpoint = tea.String(endpoint)
	result = &ecs.Client{}
	result, err = ecs.NewClient(config)
	if err != nil {
		_logUtils.Errorf("CreateEcsClient error %s", err.Error())
		return
	}

	return result, err
}
