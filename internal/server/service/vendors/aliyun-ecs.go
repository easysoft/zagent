package vendors

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	vpc "github.com/alibabacloud-go/vpc-20160428/v2/client"
	testconst "github.com/easysoft/zagent/cmd/test/_const"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"strings"
)

type AliyunEcsService struct {
}

func NewAliyunEcsService() *AliyunEcsService {
	return &AliyunEcsService{}
}

func (s AliyunEcsService) CreateInst(vmName, imageName, switchId, securityGroupId string, client *ecs.Client) (id, name string, err error) {
	regionId, _, err := s.GetRegion(client)
	if err != nil {
		return
	}

	zoneId, spec, err := s.QuerySpec(regionId, client)
	if err != nil {
		return
	}

	imageId, err := s.QueryImage(imageName, regionId, client)
	if err != nil {
		return
	}

	req := &ecs.CreateInstanceRequest{
		InstanceName:            tea.String(vmName),
		RegionId:                tea.String(regionId),
		SecurityGroupId:         tea.String(securityGroupId),
		ImageId:                 tea.String(imageId),
		ZoneId:                  tea.String(zoneId),
		VSwitchId:               tea.String(switchId),
		InstanceType:            tea.String(spec),
		InternetChargeType:      tea.String("PayByTraffic"),
		InternetMaxBandwidthOut: tea.Int32(1),
	}

	result, err := client.CreateInstance(req)
	if err != nil {
		_logUtils.Errorf("CreateInstance Resp %s, error %s", imageName, err.Error())
		return
	}

	id = *result.Body.InstanceId
	name = vmName

	return
}

func (s AliyunEcsService) StartInst(id string, ecsClient *ecs.Client) (err error) {
	startReq := &ecs.StartInstanceRequest{
		InstanceId: tea.String(id),
	}
	_, err = ecsClient.StartInstance(startReq)
	if err != nil {
		_logUtils.Errorf("StartInstance error %s", err.Error())
		return
	}

	return
}

func (s AliyunEcsService) AllocateIp(id string, ecsClient *ecs.Client) (ip string, err error) {
	ipReq := &ecs.AllocatePublicIpAddressRequest{
		InstanceId: tea.String(id),
	}

	resp, err := ecsClient.AllocatePublicIpAddress(ipReq)
	if err != nil {
		_logUtils.Errorf("AllocatePublicIpAddress error %s", err.Error())
		return
	}

	ip = *resp.Body.IpAddress

	return
}

func (s AliyunEcsService) RemoveInst(id string, ecsClient *ecs.Client) (err error) {
	req := &ecs.DeleteInstanceRequest{
		InstanceId:            tea.String(id),
		Force:                 tea.Bool(true),
		TerminateSubscription: tea.Bool(true),
	}
	_, err = ecsClient.DeleteInstance(req)
	if err != nil {
		_logUtils.Errorf("DeleteInstance %s error %s", id, err.Error())
		return
	}

	return
}

func (s AliyunEcsService) QueryInst(id, regionId string, client *ecs.Client) (status, macAddress string, err error) {
	arr := []string{id}
	jsn, _ := json.Marshal(arr)

	req := &ecs.DescribeInstancesRequest{
		InstanceIds: tea.String(string(jsn)),
		RegionId:    tea.String(regionId),
	}

	resp, err := client.DescribeInstances(req)
	if err != nil {
		_logUtils.Errorf("DescribeInstances %s error %s", id, err.Error())
		return
	}

	if resp.Body.Instances != nil {
		status = *resp.Body.Instances.Instance[0].Status
	}

	if resp.Body.Instances.Instance[0].NetworkInterfaces != nil {
		macAddress = *resp.Body.Instances.Instance[0].NetworkInterfaces.NetworkInterface[0].MacAddress
	}

	return
}
func (s AliyunEcsService) QueryVncUrl(id, vncPassword, regionId string, isWindows bool, client *ecs.Client) (url string, err error) {
	req := &ecs.DescribeInstanceVncUrlRequest{
		InstanceId: tea.String(id),
		RegionId:   tea.String(regionId),
	}

	resp, err := client.DescribeInstanceVncUrl(req)
	if err != nil {
		_logUtils.Errorf("DescribeInstances %s error %s", id, err.Error())
		return
	}

	url = fmt.Sprintf(testconst.ALIYUN_ECS_URL_VNC, *resp.Body.VncUrl, id, isWindows, vncPassword)

	return
}
func (s AliyunEcsService) QueryVncPassword(id, regionId string, client *ecs.Client) (password string, err error) {
	password = "P2s" + _stringUtils.NewUuid()[:3]

	req := &ecs.ModifyInstanceVncPasswdRequest{
		InstanceId:  tea.String(id),
		RegionId:    tea.String(regionId),
		VncPassword: tea.String(password),
	}

	_, err = client.ModifyInstanceVncPasswd(req)
	if err != nil {
		_logUtils.Errorf("DescribeInstanceVncPasswd %s error %s", id, err.Error())
		return
	}

	return
}

func (s AliyunEcsService) QueryImage(imageName, regionId string, client *ecs.Client) (imageId string, err error) {
	req := &ecs.DescribeImagesRequest{
		RegionId:        tea.String(regionId),
		ImageName:       tea.String(imageName),
		ImageOwnerAlias: tea.String("self"),
	}

	result, err := client.DescribeImages(req)
	if err != nil {
		_logUtils.Errorf("DescribeImages error %s", err.Error())
		return
	} else if len(result.Body.Images.Image) == 0 {
		_logUtils.Errorf("DescribeImages found %d images, request %#v", len(result.Body.Images.Image), req)
		return
	}

	for _, item := range result.Body.Images.Image {
		imageId = *item.ImageId
		imageName = *item.ImageName
		_logUtils.Infof("region: %s, %s", imageId, imageName)

		return
	}

	return
}

func (s AliyunEcsService) QuerySpec(regionId string, client *ecs.Client) (zoneId, spec string, err error) {
	req := &ecs.DescribeAvailableResourceRequest{
		RegionId:            tea.String(regionId),
		InstanceChargeType:  tea.String("PostPaid"),
		DestinationResource: tea.String("InstanceType"),
		Cores:               tea.Int32(2),
		Memory:              tea.Float32(4),
		ResourceType:        tea.String("instance"),
		NetworkCategory:     tea.String("classic"),
	}

	result, err := client.DescribeAvailableResource(req)
	if err != nil {
		_logUtils.Errorf("DescribeAvailableResource error %s", err.Error())
		return
	}

	zoneId = *result.Body.AvailableZones.AvailableZone[0].ZoneId
	spec = *result.Body.AvailableZones.AvailableZone[0].
		AvailableResources.AvailableResource[0].SupportedResources.SupportedResource[0].Value

	return
}

func (s AliyunEcsService) QuerySecurityGroupByVpc(vpcId, regionId string, client *ecs.Client) (id string, err error) {
	req := &ecs.DescribeSecurityGroupsRequest{
		VpcId:       tea.String(vpcId),
		RegionId:    tea.String(regionId),
		NetworkType: tea.String("vpc"),
	}

	resp, err := client.DescribeSecurityGroups(req)
	if err != nil {
		_logUtils.Errorf("DescribeSecurityGroups %s error %s", id, err.Error())
		return
	}

	if resp.Body.SecurityGroups != nil {
		id = *resp.Body.SecurityGroups.SecurityGroup[0].SecurityGroupId
	}

	return
}

func (s AliyunEcsService) GetSwitch(vpcId, regionId string, vpcClient *vpc.Client) (switchId, name string, err error) {
	req := &vpc.DescribeVSwitchesRequest{
		RegionId: tea.String(regionId),
		VpcId:    tea.String(vpcId),
	}

	resp, err := vpcClient.DescribeVSwitches(req)

	if resp.Body.VSwitches != nil {
		switchId = *resp.Body.VSwitches.VSwitch[0].VSwitchId
		name = *resp.Body.VSwitches.VSwitch[0].VSwitchId
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

func (s AliyunEcsService) CreateEcsClient(endpoint, accessKeyId, accessKeySecret string) (
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

func (s AliyunEcsService) CreateVpcClient(endpoint, accessKeyId, accessKeySecret string) (
	result *vpc.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	config.Endpoint = tea.String("vpc.aliyuncs.com")
	result = &vpc.Client{}
	result, err = vpc.NewClient(config)

	if err != nil {
		_logUtils.Errorf("CreateVpcClient error %s", err.Error())
		return
	}

	return result, err
}
