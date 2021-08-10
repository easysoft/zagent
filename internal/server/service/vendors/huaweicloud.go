package vendors

import (
	"github.com/alibabacloud-go/tea/tea"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	imgRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/region"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	vpcRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/region"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	"strings"

	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	ecsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	ims "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2"
	imgModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/model"
)

type HuaweiCloudService struct {
}

func NewHuaweiCloudService() *HuaweiCloudService {
	return &HuaweiCloudService{}
}

func (s HuaweiCloudService) CreateInst(instName, img string,
	ecsClient *ecs.EcsClient, imgClient *ims.ImsClient, vpcClient *vpc.VpcClient) (id, name string, err error) {
	specId, _, _ := s.QuerySpec(ecsClient)
	imageId, _, _ := s.QueryImage(img, imgClient)
	vpcId, _, _ := s.QueryVpc(vpcClient)
	subnetId, _, _ := s.QuerySubNet(vpcClient)

	request := &ecsModel.CreateServersRequest{
		Body: &ecsModel.CreateServersRequestBody{
			Server: &ecsModel.PrePaidServer{
				Name:      instName,
				ImageRef:  imageId,
				FlavorRef: specId,
				Vpcid:     vpcId,
				Nics: []ecsModel.PrePaidServerNic{
					{SubnetId: subnetId},
				},
				Publicip: &ecsModel.PrePaidServerPublicip{
					Eip: &ecsModel.PrePaidServerEip{
						Iptype: "5_bgp",
						Bandwidth: &ecsModel.PrePaidServerEipBandwidth{
							Sharetype: ecsModel.GetPrePaidServerEipBandwidthSharetypeEnum().PER,
							Size:      tea.Int32(1),
						},
					},
				},
				RootVolume: &ecsModel.PrePaidServerRootVolume{
					Volumetype: ecsModel.GetPrePaidServerRootVolumeVolumetypeEnum().SAS,
					Size:       tea.Int32(50),
				},
			},
		},
	}

	response, err := ecsClient.CreateServers(request)

	if err != nil {
		_logUtils.Printf("CreateServers error %s", err.Error())
	}

	id = (*response.ServerIds)[0]

	return
}

func (s HuaweiCloudService) RemoveInst(id string, ecsClient *ecs.EcsClient) (err error) {
	request := &ecsModel.DeleteServersRequest{
		Body: &ecsModel.DeleteServersRequestBody{
			Servers: []ecsModel.ServerId{{
				Id: id,
			}},
		},
	}
	response, err := ecsClient.DeleteServers(request)

	if response.HttpStatusCode != 200 || err != nil {
		_logUtils.Printf("DeleteServers response %s, error %s", response.String(), err.Error())
	}

	return
}
func (s HuaweiCloudService) QuerySpec(client *ecs.EcsClient) (id, name string, err error) {
	request := &ecsModel.ListFlavorsRequest{}
	response, err := client.ListFlavors(request)
	if err != nil {
		_logUtils.Printf("ListFlavors error %s", err.Error())
	}

	for _, item := range *response.Flavors {
		id = item.Id
		name = item.Name
		cpus := item.Vcpus
		ram := item.Ram

		if cpus == "" && ram == 4000 {
			return
		}
	}

	return
}

func (s HuaweiCloudService) QueryImage(keywords string, client *ims.ImsClient) (id, name string, err error) {
	request := &imgModel.ListImagesRequest{}
	response, err := client.ListImages(request)

	if err != nil {
		_logUtils.Printf("ListFlavors error %s", err.Error())
	}

	for _, item := range *response.Images {
		id = item.Id
		name = item.Name

		if strings.Index(name, keywords) > -1 {
			return
		}
	}

	return
}

func (s HuaweiCloudService) QueryVpc(client *vpc.VpcClient) (id, name string, err error) {
	request := &model.ListVpcsRequest{}
	response, err := client.ListVpcs(request)

	if err != nil {
		_logUtils.Printf("ListVpcs error %s", err.Error())
	}

	for _, item := range *response.Vpcs {
		id = item.Id
		name = item.Name

		return // first one
	}

	return
}

func (s HuaweiCloudService) QuerySubNet(client *vpc.VpcClient) (id, name string, err error) {
	request := &model.ListSubnetsRequest{}
	response, err := client.ListSubnets(request)

	if err != nil {
		_logUtils.Printf("ListVpcs error %s", err.Error())
	}

	for _, item := range *response.Subnets {
		id = item.Id
		name = item.Name

		return // first one
	}

	return
}

func (s HuaweiCloudService) QueryVm(id string, client *ecs.EcsClient) (name, status, ip, mac string, err error) {
	request := &ecsModel.ShowServerRequest{
		ServerId: id,
	}
	response, err := client.ShowServer(request)
	if err != nil {
		_logUtils.Printf("ShowServer error %s", err.Error())
		return
	}

	name = response.Server.Name
	status = response.Server.Status

	for _, items := range response.Server.Addresses {
		item := items[0]
		ip = item.Addr
		mac = *item.OSEXTIPSMACmacAddr

		break
	}

	return
}

func (s HuaweiCloudService) QueryVnc(id string, client *ecs.EcsClient) (url string, err error) {
	request := &ecsModel.ShowServerRemoteConsoleRequest{
		ServerId: id,
		Body: &ecsModel.ShowServerRemoteConsoleRequestBody{
			RemoteConsole: &ecsModel.GetServerRemoteConsoleOption{
				Protocol: ecsModel.GetGetServerRemoteConsoleOptionProtocolEnum().VNC,
				Type:     ecsModel.GetGetServerRemoteConsoleOptionTypeEnum().NOVNC,
			},
		},
	}
	response, err := client.ShowServerRemoteConsole(request)
	if err != nil {
		_logUtils.Printf("ShowServer error %s", err.Error())
		return
	}

	url = response.RemoteConsole.Url

	return
}

func (s HuaweiCloudService) CreateEcsClient(ak, sk, regionId string) (
	client *ecs.EcsClient, err error) {

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client = ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(ecsRegion.ValueOf(regionId)).
			WithCredential(auth).
			Build())

	return
}

func (s HuaweiCloudService) CreateImgClient(ak, sk, regionId string) (
	client *ims.ImsClient, err error) {

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client = ims.NewImsClient(
		ims.ImsClientBuilder().
			WithRegion(imgRegion.ValueOf(regionId)).
			WithCredential(auth).
			Build())

	return
}

func (s HuaweiCloudService) CreateVpcClient(ak, sk, regionId string) (
	client *vpc.VpcClient, err error) {

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client = vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			WithRegion(vpcRegion.ValueOf(regionId)).
			WithCredential(auth).
			Build())

	return
}
