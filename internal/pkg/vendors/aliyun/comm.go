package aliyun

import (
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	eci "github.com/alibabacloud-go/eci-20180808/v2/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	vpc "github.com/alibabacloud-go/vpc-20160428/v2/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ens"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
)

type AliyunCommService struct {
}

func NewAliyunCommService() *AliyunCommService {
	return &AliyunCommService{}
}

func (s AliyunCommService) QuerySecurityGroupByVpc(vpcId, regionId string, client *ecs.Client) (id string, err error) {
	req := &ecs.DescribeSecurityGroupsRequest{
		RegionId:    tea.String(regionId),
		VpcId:       tea.String(vpcId),
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

func (s AliyunCommService) GetSwitch(vpcId, regionId string, vpcClient *vpc.Client) (switchId, name string, err error) {
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

func (s AliyunCommService) GetRegion(region string, client *ecs.Client) (id, name string, err error) {
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
		if strings.Index(id, region) > -1 {
			id = *item.RegionId

			return
		}
	}

	return
}

func (s AliyunCommService) CreateEip(region string, client *vpc.Client) (id string, err error) {
	req := &vpc.AllocateEipAddressRequest{
		RegionId:  tea.String(region),
		Bandwidth: tea.String("1"),
	}

	resp, err := client.AllocateEipAddress(req)
	id = *resp.Body.AllocationId

	return
}
func (s AliyunCommService) DestroyEip(id, region string, client *vpc.Client) (err error) {
	req := &vpc.ReleaseEipAddressRequest{
		AllocationId: tea.String(id),
		RegionId:     tea.String(region),
	}

	_, err = client.ReleaseEipAddress(req)

	return
}

func (s AliyunCommService) GetAvailableEip(region string, client *vpc.Client) (id string, err error) {
	describeEipAddressesRequest := &vpc.DescribeEipAddressesRequest{
		RegionId: tea.String(region),
		Status:   tea.String("Available"),
	}

	resp, err := client.DescribeEipAddresses(describeEipAddressesRequest)
	if err != nil {
		_logUtils.Errorf("DescribeEipAddresses error %s", err.Error())
		return
	}

	if len(resp.Body.EipAddresses.EipAddress) == 0 {
		_logUtils.Errorf("DescribeEipAddresses, can not find any EIP instances in 'Available' status, all are 'InUse' status?")
		return
	}

	id = *resp.Body.EipAddresses.EipAddress[0].AllocationId
	return
}

func (s AliyunCommService) CreateEcsClient(endpoint, accessKeyId, accessKeySecret string) (
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

func (s AliyunCommService) CreateVpcClient(endpoint, accessKeyId, accessKeySecret string) (
	result *vpc.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	config.Endpoint = tea.String(endpoint)
	result = &vpc.Client{}
	result, err = vpc.NewClient(config)

	if err != nil {
		_logUtils.Errorf("CreateVpcClient error %s", err.Error())
		return
	}

	return result, err
}

func (s AliyunCommService) CreateEciClient(endpoint, accessKeyId, accessKeySecret string) (
	result *eci.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	config.Endpoint = tea.String(endpoint)
	result = &eci.Client{}
	result, err = eci.NewClient(config)
	if err != nil {
		_logUtils.Errorf("CreateEciClient error %s", err.Error())
		return
	}

	return result, err
}

func (s AliyunCommService) CreateEnsClient(url string, accessKeyId string, accessKeySecret string) (client *ens.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String(url)
	client = &ens.Client{}
	//client, _err = ens.NewClient(config)
	return client, err
}
