package virtualboxsrv

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

func NewVboxPortType(url string, tls bool, auth *BasicAuth) *VboxPortType {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &VboxPortType{
		client: client,
	}
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetResultCode(request *IVirtualBoxErrorInfogetResultCode) (*IVirtualBoxErrorInfogetResultCodeResponse, error) {
	response := new(IVirtualBoxErrorInfogetResultCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetResultDetail(request *IVirtualBoxErrorInfogetResultDetail) (*IVirtualBoxErrorInfogetResultDetailResponse, error) {
	response := new(IVirtualBoxErrorInfogetResultDetailResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetInterfaceID(request *IVirtualBoxErrorInfogetInterfaceID) (*IVirtualBoxErrorInfogetInterfaceIDResponse, error) {
	response := new(IVirtualBoxErrorInfogetInterfaceIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetComponent(request *IVirtualBoxErrorInfogetComponent) (*IVirtualBoxErrorInfogetComponentResponse, error) {
	response := new(IVirtualBoxErrorInfogetComponentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetText(request *IVirtualBoxErrorInfogetText) (*IVirtualBoxErrorInfogetTextResponse, error) {
	response := new(IVirtualBoxErrorInfogetTextResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetNext(request *IVirtualBoxErrorInfogetNext) (*IVirtualBoxErrorInfogetNextResponse, error) {
	response := new(IVirtualBoxErrorInfogetNextResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetNetworkName(request *INATNetworkgetNetworkName) (*INATNetworkgetNetworkNameResponse, error) {
	response := new(INATNetworkgetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetNetworkName(request *INATNetworksetNetworkName) (*INATNetworksetNetworkNameResponse, error) {
	response := new(INATNetworksetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetEnabled(request *INATNetworkgetEnabled) (*INATNetworkgetEnabledResponse, error) {
	response := new(INATNetworkgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetEnabled(request *INATNetworksetEnabled) (*INATNetworksetEnabledResponse, error) {
	response := new(INATNetworksetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetNetwork(request *INATNetworkgetNetwork) (*INATNetworkgetNetworkResponse, error) {
	response := new(INATNetworkgetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetNetwork(request *INATNetworksetNetwork) (*INATNetworksetNetworkResponse, error) {
	response := new(INATNetworksetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetGateway(request *INATNetworkgetGateway) (*INATNetworkgetGatewayResponse, error) {
	response := new(INATNetworkgetGatewayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetIPv6Enabled(request *INATNetworkgetIPv6Enabled) (*INATNetworkgetIPv6EnabledResponse, error) {
	response := new(INATNetworkgetIPv6EnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetIPv6Enabled(request *INATNetworksetIPv6Enabled) (*INATNetworksetIPv6EnabledResponse, error) {
	response := new(INATNetworksetIPv6EnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetIPv6Prefix(request *INATNetworkgetIPv6Prefix) (*INATNetworkgetIPv6PrefixResponse, error) {
	response := new(INATNetworkgetIPv6PrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetIPv6Prefix(request *INATNetworksetIPv6Prefix) (*INATNetworksetIPv6PrefixResponse, error) {
	response := new(INATNetworksetIPv6PrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetAdvertiseDefaultIPv6RouteEnabled(request *INATNetworkgetAdvertiseDefaultIPv6RouteEnabled) (*INATNetworkgetAdvertiseDefaultIPv6RouteEnabledResponse, error) {
	response := new(INATNetworkgetAdvertiseDefaultIPv6RouteEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetAdvertiseDefaultIPv6RouteEnabled(request *INATNetworksetAdvertiseDefaultIPv6RouteEnabled) (*INATNetworksetAdvertiseDefaultIPv6RouteEnabledResponse, error) {
	response := new(INATNetworksetAdvertiseDefaultIPv6RouteEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetNeedDhcpServer(request *INATNetworkgetNeedDhcpServer) (*INATNetworkgetNeedDhcpServerResponse, error) {
	response := new(INATNetworkgetNeedDhcpServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetNeedDhcpServer(request *INATNetworksetNeedDhcpServer) (*INATNetworksetNeedDhcpServerResponse, error) {
	response := new(INATNetworksetNeedDhcpServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetEventSource(request *INATNetworkgetEventSource) (*INATNetworkgetEventSourceResponse, error) {
	response := new(INATNetworkgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetPortForwardRules4(request *INATNetworkgetPortForwardRules4) (*INATNetworkgetPortForwardRules4Response, error) {
	response := new(INATNetworkgetPortForwardRules4Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetLocalMappings(request *INATNetworkgetLocalMappings) (*INATNetworkgetLocalMappingsResponse, error) {
	response := new(INATNetworkgetLocalMappingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetLoopbackIp6(request *INATNetworkgetLoopbackIp6) (*INATNetworkgetLoopbackIp6Response, error) {
	response := new(INATNetworkgetLoopbackIp6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetLoopbackIp6(request *INATNetworksetLoopbackIp6) (*INATNetworksetLoopbackIp6Response, error) {
	response := new(INATNetworksetLoopbackIp6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetPortForwardRules6(request *INATNetworkgetPortForwardRules6) (*INATNetworkgetPortForwardRules6Response, error) {
	response := new(INATNetworkgetPortForwardRules6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkaddLocalMapping(request *INATNetworkaddLocalMapping) (*INATNetworkaddLocalMappingResponse, error) {
	response := new(INATNetworkaddLocalMappingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkaddPortForwardRule(request *INATNetworkaddPortForwardRule) (*INATNetworkaddPortForwardRuleResponse, error) {
	response := new(INATNetworkaddPortForwardRuleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkremovePortForwardRule(request *INATNetworkremovePortForwardRule) (*INATNetworkremovePortForwardRuleResponse, error) {
	response := new(INATNetworkremovePortForwardRuleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkstart(request *INATNetworkstart) (*INATNetworkstartResponse, error) {
	response := new(INATNetworkstartResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkstop(request *INATNetworkstop) (*INATNetworkstopResponse, error) {
	response := new(INATNetworkstopResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetEventSource(request *IDHCPServergetEventSource) (*IDHCPServergetEventSourceResponse, error) {
	response := new(IDHCPServergetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetEnabled(request *IDHCPServergetEnabled) (*IDHCPServergetEnabledResponse, error) {
	response := new(IDHCPServergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServersetEnabled(request *IDHCPServersetEnabled) (*IDHCPServersetEnabledResponse, error) {
	response := new(IDHCPServersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetIPAddress(request *IDHCPServergetIPAddress) (*IDHCPServergetIPAddressResponse, error) {
	response := new(IDHCPServergetIPAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetNetworkMask(request *IDHCPServergetNetworkMask) (*IDHCPServergetNetworkMaskResponse, error) {
	response := new(IDHCPServergetNetworkMaskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetNetworkName(request *IDHCPServergetNetworkName) (*IDHCPServergetNetworkNameResponse, error) {
	response := new(IDHCPServergetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetLowerIP(request *IDHCPServergetLowerIP) (*IDHCPServergetLowerIPResponse, error) {
	response := new(IDHCPServergetLowerIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetUpperIP(request *IDHCPServergetUpperIP) (*IDHCPServergetUpperIPResponse, error) {
	response := new(IDHCPServergetUpperIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetGlobalOptions(request *IDHCPServergetGlobalOptions) (*IDHCPServergetGlobalOptionsResponse, error) {
	response := new(IDHCPServergetGlobalOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetVmConfigs(request *IDHCPServergetVmConfigs) (*IDHCPServergetVmConfigsResponse, error) {
	response := new(IDHCPServergetVmConfigsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServeraddGlobalOption(request *IDHCPServeraddGlobalOption) (*IDHCPServeraddGlobalOptionResponse, error) {
	response := new(IDHCPServeraddGlobalOptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServeraddVmSlotOption(request *IDHCPServeraddVmSlotOption) (*IDHCPServeraddVmSlotOptionResponse, error) {
	response := new(IDHCPServeraddVmSlotOptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServerremoveVmSlotOptions(request *IDHCPServerremoveVmSlotOptions) (*IDHCPServerremoveVmSlotOptionsResponse, error) {
	response := new(IDHCPServerremoveVmSlotOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetVmSlotOptions(request *IDHCPServergetVmSlotOptions) (*IDHCPServergetVmSlotOptionsResponse, error) {
	response := new(IDHCPServergetVmSlotOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetMacOptions(request *IDHCPServergetMacOptions) (*IDHCPServergetMacOptionsResponse, error) {
	response := new(IDHCPServergetMacOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServersetConfiguration(request *IDHCPServersetConfiguration) (*IDHCPServersetConfigurationResponse, error) {
	response := new(IDHCPServersetConfigurationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServerstart(request *IDHCPServerstart) (*IDHCPServerstartResponse, error) {
	response := new(IDHCPServerstartResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServerstop(request *IDHCPServerstop) (*IDHCPServerstopResponse, error) {
	response := new(IDHCPServerstopResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetVersion(request *IVirtualBoxgetVersion) (*IVirtualBoxgetVersionResponse, error) {
	response := new(IVirtualBoxgetVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetVersionNormalized(request *IVirtualBoxgetVersionNormalized) (*IVirtualBoxgetVersionNormalizedResponse, error) {
	response := new(IVirtualBoxgetVersionNormalizedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetRevision(request *IVirtualBoxgetRevision) (*IVirtualBoxgetRevisionResponse, error) {
	response := new(IVirtualBoxgetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetPackageType(request *IVirtualBoxgetPackageType) (*IVirtualBoxgetPackageTypeResponse, error) {
	response := new(IVirtualBoxgetPackageTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetAPIVersion(request *IVirtualBoxgetAPIVersion) (*IVirtualBoxgetAPIVersionResponse, error) {
	response := new(IVirtualBoxgetAPIVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetAPIRevision(request *IVirtualBoxgetAPIRevision) (*IVirtualBoxgetAPIRevisionResponse, error) {
	response := new(IVirtualBoxgetAPIRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHomeFolder(request *IVirtualBoxgetHomeFolder) (*IVirtualBoxgetHomeFolderResponse, error) {
	response := new(IVirtualBoxgetHomeFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSettingsFilePath(request *IVirtualBoxgetSettingsFilePath) (*IVirtualBoxgetSettingsFilePathResponse, error) {
	response := new(IVirtualBoxgetSettingsFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHost(request *IVirtualBoxgetHost) (*IVirtualBoxgetHostResponse, error) {
	response := new(IVirtualBoxgetHostResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSystemProperties(request *IVirtualBoxgetSystemProperties) (*IVirtualBoxgetSystemPropertiesResponse, error) {
	response := new(IVirtualBoxgetSystemPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachines(request *IVirtualBoxgetMachines) (*IVirtualBoxgetMachinesResponse, error) {
	response := new(IVirtualBoxgetMachinesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachineGroups(request *IVirtualBoxgetMachineGroups) (*IVirtualBoxgetMachineGroupsResponse, error) {
	response := new(IVirtualBoxgetMachineGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHardDisks(request *IVirtualBoxgetHardDisks) (*IVirtualBoxgetHardDisksResponse, error) {
	response := new(IVirtualBoxgetHardDisksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetDVDImages(request *IVirtualBoxgetDVDImages) (*IVirtualBoxgetDVDImagesResponse, error) {
	response := new(IVirtualBoxgetDVDImagesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetFloppyImages(request *IVirtualBoxgetFloppyImages) (*IVirtualBoxgetFloppyImagesResponse, error) {
	response := new(IVirtualBoxgetFloppyImagesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetProgressOperations(request *IVirtualBoxgetProgressOperations) (*IVirtualBoxgetProgressOperationsResponse, error) {
	response := new(IVirtualBoxgetProgressOperationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetGuestOSTypes(request *IVirtualBoxgetGuestOSTypes) (*IVirtualBoxgetGuestOSTypesResponse, error) {
	response := new(IVirtualBoxgetGuestOSTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSharedFolders(request *IVirtualBoxgetSharedFolders) (*IVirtualBoxgetSharedFoldersResponse, error) {
	response := new(IVirtualBoxgetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetPerformanceCollector(request *IVirtualBoxgetPerformanceCollector) (*IVirtualBoxgetPerformanceCollectorResponse, error) {
	response := new(IVirtualBoxgetPerformanceCollectorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetDHCPServers(request *IVirtualBoxgetDHCPServers) (*IVirtualBoxgetDHCPServersResponse, error) {
	response := new(IVirtualBoxgetDHCPServersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetNATNetworks(request *IVirtualBoxgetNATNetworks) (*IVirtualBoxgetNATNetworksResponse, error) {
	response := new(IVirtualBoxgetNATNetworksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetEventSource(request *IVirtualBoxgetEventSource) (*IVirtualBoxgetEventSourceResponse, error) {
	response := new(IVirtualBoxgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetInternalNetworks(request *IVirtualBoxgetInternalNetworks) (*IVirtualBoxgetInternalNetworksResponse, error) {
	response := new(IVirtualBoxgetInternalNetworksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetGenericNetworkDrivers(request *IVirtualBoxgetGenericNetworkDrivers) (*IVirtualBoxgetGenericNetworkDriversResponse, error) {
	response := new(IVirtualBoxgetGenericNetworkDriversResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcomposeMachineFilename(request *IVirtualBoxcomposeMachineFilename) (*IVirtualBoxcomposeMachineFilenameResponse, error) {
	response := new(IVirtualBoxcomposeMachineFilenameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateMachine(request *IVirtualBoxcreateMachine) (*IVirtualBoxcreateMachineResponse, error) {
	response := new(IVirtualBoxcreateMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenMachine(request *IVirtualBoxopenMachine) (*IVirtualBoxopenMachineResponse, error) {
	response := new(IVirtualBoxopenMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxregisterMachine(request *IVirtualBoxregisterMachine) (*IVirtualBoxregisterMachineResponse, error) {
	response := new(IVirtualBoxregisterMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindMachine(request *IVirtualBoxfindMachine) (*IVirtualBoxfindMachineResponse, error) {
	response := new(IVirtualBoxfindMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachinesByGroups(request *IVirtualBoxgetMachinesByGroups) (*IVirtualBoxgetMachinesByGroupsResponse, error) {
	response := new(IVirtualBoxgetMachinesByGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachineStates(request *IVirtualBoxgetMachineStates) (*IVirtualBoxgetMachineStatesResponse, error) {
	response := new(IVirtualBoxgetMachineStatesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateAppliance(request *IVirtualBoxcreateAppliance) (*IVirtualBoxcreateApplianceResponse, error) {
	response := new(IVirtualBoxcreateApplianceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateMedium(request *IVirtualBoxcreateMediumReq) (*IVirtualBoxcreateMediumResponse, error) {
	response := new(IVirtualBoxcreateMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenMedium(request *IVirtualBoxopenMedium) (*IVirtualBoxopenMediumResponse, error) {
	response := new(IVirtualBoxopenMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetGuestOSType(request *IVirtualBoxgetGuestOSType) (*IVirtualBoxgetGuestOSTypeResponse, error) {
	response := new(IVirtualBoxgetGuestOSTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateSharedFolder(request *IVirtualBoxcreateSharedFolder) (*IVirtualBoxcreateSharedFolderResponse, error) {
	response := new(IVirtualBoxcreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxremoveSharedFolder(request *IVirtualBoxremoveSharedFolder) (*IVirtualBoxremoveSharedFolderResponse, error) {
	response := new(IVirtualBoxremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetExtraDataKeys(request *IVirtualBoxgetExtraDataKeys) (*IVirtualBoxgetExtraDataKeysResponse, error) {
	response := new(IVirtualBoxgetExtraDataKeysResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetExtraData(request *IVirtualBoxgetExtraData) (*IVirtualBoxgetExtraDataResponse, error) {
	response := new(IVirtualBoxgetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxsetExtraData(request *IVirtualBoxsetExtraData) (*IVirtualBoxsetExtraDataResponse, error) {
	response := new(IVirtualBoxsetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxsetSettingsSecret(request *IVirtualBoxsetSettingsSecret) (*IVirtualBoxsetSettingsSecretResponse, error) {
	response := new(IVirtualBoxsetSettingsSecretResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateDHCPServer(request *IVirtualBoxcreateDHCPServer) (*IVirtualBoxcreateDHCPServerResponse, error) {
	response := new(IVirtualBoxcreateDHCPServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindDHCPServerByNetworkName(request *IVirtualBoxfindDHCPServerByNetworkName) (*IVirtualBoxfindDHCPServerByNetworkNameResponse, error) {
	response := new(IVirtualBoxfindDHCPServerByNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxremoveDHCPServer(request *IVirtualBoxremoveDHCPServer) (*IVirtualBoxremoveDHCPServerResponse, error) {
	response := new(IVirtualBoxremoveDHCPServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateNATNetwork(request *IVirtualBoxcreateNATNetwork) (*IVirtualBoxcreateNATNetworkResponse, error) {
	response := new(IVirtualBoxcreateNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindNATNetworkByName(request *IVirtualBoxfindNATNetworkByName) (*IVirtualBoxfindNATNetworkByNameResponse, error) {
	response := new(IVirtualBoxfindNATNetworkByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxremoveNATNetwork(request *IVirtualBoxremoveNATNetwork) (*IVirtualBoxremoveNATNetworkResponse, error) {
	response := new(IVirtualBoxremoveNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcheckFirmwarePresent(request *IVirtualBoxcheckFirmwarePresent) (*IVirtualBoxcheckFirmwarePresentResponse, error) {
	response := new(IVirtualBoxcheckFirmwarePresentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorergetPath(request *IVFSExplorergetPath) (*IVFSExplorergetPathResponse, error) {
	response := new(IVFSExplorergetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorergetType(request *IVFSExplorergetType) (*IVFSExplorergetTypeResponse, error) {
	response := new(IVFSExplorergetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerupdate(request *IVFSExplorerupdate) (*IVFSExplorerupdateResponse, error) {
	response := new(IVFSExplorerupdateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorercd(request *IVFSExplorercd) (*IVFSExplorercdResponse, error) {
	response := new(IVFSExplorercdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorercdUp(request *IVFSExplorercdUp) (*IVFSExplorercdUpResponse, error) {
	response := new(IVFSExplorercdUpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerentryList(request *IVFSExplorerentryList) (*IVFSExplorerentryListResponse, error) {
	response := new(IVFSExplorerentryListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerexists(request *IVFSExplorerexists) (*IVFSExplorerexistsResponse, error) {
	response := new(IVFSExplorerexistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerremove(request *IVFSExplorerremove) (*IVFSExplorerremoveResponse, error) {
	response := new(IVFSExplorerremoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetPath(request *IAppliancegetPath) (*IAppliancegetPathResponse, error) {
	response := new(IAppliancegetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetDisks(request *IAppliancegetDisks) (*IAppliancegetDisksResponse, error) {
	response := new(IAppliancegetDisksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetVirtualSystemDescriptions(request *IAppliancegetVirtualSystemDescriptions) (*IAppliancegetVirtualSystemDescriptionsResponse, error) {
	response := new(IAppliancegetVirtualSystemDescriptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetMachines(request *IAppliancegetMachines) (*IAppliancegetMachinesResponse, error) {
	response := new(IAppliancegetMachinesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceread(request *IApplianceread) (*IAppliancereadResponse, error) {
	response := new(IAppliancereadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceinterpret(request *IApplianceinterpret) (*IApplianceinterpretResponse, error) {
	response := new(IApplianceinterpretResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceimportMachines(request *IApplianceimportMachines) (*IApplianceimportMachinesResponse, error) {
	response := new(IApplianceimportMachinesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancecreateVFSExplorer(request *IAppliancecreateVFSExplorer) (*IAppliancecreateVFSExplorerResponse, error) {
	response := new(IAppliancecreateVFSExplorerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancewrite(request *IAppliancewrite) (*IAppliancewriteResponse, error) {
	response := new(IAppliancewriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetWarnings(request *IAppliancegetWarnings) (*IAppliancegetWarningsResponse, error) {
	response := new(IAppliancegetWarningsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetPasswordIds(request *IAppliancegetPasswordIds) (*IAppliancegetPasswordIdsResponse, error) {
	response := new(IAppliancegetPasswordIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetMediumIdsForPasswordId(request *IAppliancegetMediumIdsForPasswordId) (*IAppliancegetMediumIdsForPasswordIdResponse, error) {
	response := new(IAppliancegetMediumIdsForPasswordIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceaddPasswords(request *IApplianceaddPasswords) (*IApplianceaddPasswordsResponse, error) {
	response := new(IApplianceaddPasswordsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetCount(request *IVirtualSystemDescriptiongetCount) (*IVirtualSystemDescriptiongetCountResponse, error) {
	response := new(IVirtualSystemDescriptiongetCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetDescription(request *IVirtualSystemDescriptiongetDescription) (*IVirtualSystemDescriptiongetDescriptionResponse, error) {
	response := new(IVirtualSystemDescriptiongetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetDescriptionByType(request *IVirtualSystemDescriptiongetDescriptionByType) (*IVirtualSystemDescriptiongetDescriptionByTypeResponse, error) {
	response := new(IVirtualSystemDescriptiongetDescriptionByTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetValuesByType(request *IVirtualSystemDescriptiongetValuesByType) (*IVirtualSystemDescriptiongetValuesByTypeResponse, error) {
	response := new(IVirtualSystemDescriptiongetValuesByTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptionsetFinalValues(request *IVirtualSystemDescriptionsetFinalValues) (*IVirtualSystemDescriptionsetFinalValuesResponse, error) {
	response := new(IVirtualSystemDescriptionsetFinalValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptionaddDescription(request *IVirtualSystemDescriptionaddDescription) (*IVirtualSystemDescriptionaddDescriptionResponse, error) {
	response := new(IVirtualSystemDescriptionaddDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoFadeIn(request *IBIOSSettingsgetLogoFadeIn) (*IBIOSSettingsgetLogoFadeInResponse, error) {
	response := new(IBIOSSettingsgetLogoFadeInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoFadeIn(request *IBIOSSettingssetLogoFadeIn) (*IBIOSSettingssetLogoFadeInResponse, error) {
	response := new(IBIOSSettingssetLogoFadeInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoFadeOut(request *IBIOSSettingsgetLogoFadeOut) (*IBIOSSettingsgetLogoFadeOutResponse, error) {
	response := new(IBIOSSettingsgetLogoFadeOutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoFadeOut(request *IBIOSSettingssetLogoFadeOut) (*IBIOSSettingssetLogoFadeOutResponse, error) {
	response := new(IBIOSSettingssetLogoFadeOutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoDisplayTime(request *IBIOSSettingsgetLogoDisplayTime) (*IBIOSSettingsgetLogoDisplayTimeResponse, error) {
	response := new(IBIOSSettingsgetLogoDisplayTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoDisplayTime(request *IBIOSSettingssetLogoDisplayTime) (*IBIOSSettingssetLogoDisplayTimeResponse, error) {
	response := new(IBIOSSettingssetLogoDisplayTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoImagePath(request *IBIOSSettingsgetLogoImagePath) (*IBIOSSettingsgetLogoImagePathResponse, error) {
	response := new(IBIOSSettingsgetLogoImagePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoImagePath(request *IBIOSSettingssetLogoImagePath) (*IBIOSSettingssetLogoImagePathResponse, error) {
	response := new(IBIOSSettingssetLogoImagePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetBootMenuMode(request *IBIOSSettingsgetBootMenuMode) (*IBIOSSettingsgetBootMenuModeResponse, error) {
	response := new(IBIOSSettingsgetBootMenuModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetBootMenuMode(request *IBIOSSettingssetBootMenuMode) (*IBIOSSettingssetBootMenuModeResponse, error) {
	response := new(IBIOSSettingssetBootMenuModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetACPIEnabled(request *IBIOSSettingsgetACPIEnabled) (*IBIOSSettingsgetACPIEnabledResponse, error) {
	response := new(IBIOSSettingsgetACPIEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetACPIEnabled(request *IBIOSSettingssetACPIEnabled) (*IBIOSSettingssetACPIEnabledResponse, error) {
	response := new(IBIOSSettingssetACPIEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetIOAPICEnabled(request *IBIOSSettingsgetIOAPICEnabled) (*IBIOSSettingsgetIOAPICEnabledResponse, error) {
	response := new(IBIOSSettingsgetIOAPICEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetIOAPICEnabled(request *IBIOSSettingssetIOAPICEnabled) (*IBIOSSettingssetIOAPICEnabledResponse, error) {
	response := new(IBIOSSettingssetIOAPICEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetTimeOffset(request *IBIOSSettingsgetTimeOffset) (*IBIOSSettingsgetTimeOffsetResponse, error) {
	response := new(IBIOSSettingsgetTimeOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetTimeOffset(request *IBIOSSettingssetTimeOffset) (*IBIOSSettingssetTimeOffsetResponse, error) {
	response := new(IBIOSSettingssetTimeOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetPXEDebugEnabled(request *IBIOSSettingsgetPXEDebugEnabled) (*IBIOSSettingsgetPXEDebugEnabledResponse, error) {
	response := new(IBIOSSettingsgetPXEDebugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetPXEDebugEnabled(request *IBIOSSettingssetPXEDebugEnabled) (*IBIOSSettingssetPXEDebugEnabledResponse, error) {
	response := new(IBIOSSettingssetPXEDebugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetNonVolatileStorageFile(request *IBIOSSettingsgetNonVolatileStorageFile) (*IBIOSSettingsgetNonVolatileStorageFileResponse, error) {
	response := new(IBIOSSettingsgetNonVolatileStorageFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressgetBus(request *IPCIAddressgetBus) (*IPCIAddressgetBusResponse, error) {
	response := new(IPCIAddressgetBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddresssetBus(request *IPCIAddresssetBus) (*IPCIAddresssetBusResponse, error) {
	response := new(IPCIAddresssetBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressgetDevice(request *IPCIAddressgetDevice) (*IPCIAddressgetDeviceResponse, error) {
	response := new(IPCIAddressgetDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddresssetDevice(request *IPCIAddresssetDevice) (*IPCIAddresssetDeviceResponse, error) {
	response := new(IPCIAddresssetDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressgetDevFunction(request *IPCIAddressgetDevFunction) (*IPCIAddressgetDevFunctionResponse, error) {
	response := new(IPCIAddressgetDevFunctionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddresssetDevFunction(request *IPCIAddresssetDevFunction) (*IPCIAddresssetDevFunctionResponse, error) {
	response := new(IPCIAddresssetDevFunctionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressasLong(request *IPCIAddressasLong) (*IPCIAddressasLongResponse, error) {
	response := new(IPCIAddressasLongResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressfromLong(request *IPCIAddressfromLong) (*IPCIAddressfromLongResponse, error) {
	response := new(IPCIAddressfromLongResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetParent(request *IMachinegetParent) (*IMachinegetParentResponse, error) {
	response := new(IMachinegetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIcon(request *IMachinegetIcon) (*IMachinegetIconResponse, error) {
	response := new(IMachinegetIconResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIcon(request *IMachinesetIcon) (*IMachinesetIconResponse, error) {
	response := new(IMachinesetIconResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccessible(request *IMachinegetAccessible) (*IMachinegetAccessibleResponse, error) {
	response := new(IMachinegetAccessibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccessError(request *IMachinegetAccessError) (*IMachinegetAccessErrorResponse, error) {
	response := new(IMachinegetAccessErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetName(request *IMachinegetName) (*IMachinegetNameResponse, error) {
	response := new(IMachinegetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetName(request *IMachinesetName) (*IMachinesetNameResponse, error) {
	response := new(IMachinesetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetDescription(request *IMachinegetDescription) (*IMachinegetDescriptionResponse, error) {
	response := new(IMachinegetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetDescription(request *IMachinesetDescription) (*IMachinesetDescriptionResponse, error) {
	response := new(IMachinesetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetId(request *IMachinegetId) (*IMachinegetIdResponse, error) {
	response := new(IMachinegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGroups(request *IMachinegetGroups) (*IMachinegetGroupsResponse, error) {
	response := new(IMachinegetGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGroups(request *IMachinesetGroups) (*IMachinesetGroupsResponse, error) {
	response := new(IMachinesetGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetOSTypeId(request *IMachinegetOSTypeId) (*IMachinegetOSTypeIdResponse, error) {
	response := new(IMachinegetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetOSTypeId(request *IMachinesetOSTypeId) (*IMachinesetOSTypeIdResponse, error) {
	response := new(IMachinesetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHardwareVersion(request *IMachinegetHardwareVersion) (*IMachinegetHardwareVersionResponse, error) {
	response := new(IMachinegetHardwareVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHardwareVersion(request *IMachinesetHardwareVersion) (*IMachinesetHardwareVersionResponse, error) {
	response := new(IMachinesetHardwareVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHardwareUUID(request *IMachinegetHardwareUUID) (*IMachinegetHardwareUUIDResponse, error) {
	response := new(IMachinegetHardwareUUIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHardwareUUID(request *IMachinesetHardwareUUID) (*IMachinesetHardwareUUIDResponse, error) {
	response := new(IMachinesetHardwareUUIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUCount(request *IMachinegetCPUCount) (*IMachinegetCPUCountResponse, error) {
	response := new(IMachinegetCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUCount(request *IMachinesetCPUCount) (*IMachinesetCPUCountResponse, error) {
	response := new(IMachinesetCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUHotPlugEnabled(request *IMachinegetCPUHotPlugEnabled) (*IMachinegetCPUHotPlugEnabledResponse, error) {
	response := new(IMachinegetCPUHotPlugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUHotPlugEnabled(request *IMachinesetCPUHotPlugEnabled) (*IMachinesetCPUHotPlugEnabledResponse, error) {
	response := new(IMachinesetCPUHotPlugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUExecutionCap(request *IMachinegetCPUExecutionCap) (*IMachinegetCPUExecutionCapResponse, error) {
	response := new(IMachinegetCPUExecutionCapResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUExecutionCap(request *IMachinesetCPUExecutionCap) (*IMachinesetCPUExecutionCapResponse, error) {
	response := new(IMachinesetCPUExecutionCapResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUIDPortabilityLevel(request *IMachinegetCPUIDPortabilityLevel) (*IMachinegetCPUIDPortabilityLevelResponse, error) {
	response := new(IMachinegetCPUIDPortabilityLevelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUIDPortabilityLevel(request *IMachinesetCPUIDPortabilityLevel) (*IMachinesetCPUIDPortabilityLevelResponse, error) {
	response := new(IMachinesetCPUIDPortabilityLevelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMemorySize(request *IMachinegetMemorySize) (*IMachinegetMemorySizeResponse, error) {
	response := new(IMachinegetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMemorySize(request *IMachinesetMemorySize) (*IMachinesetMemorySizeResponse, error) {
	response := new(IMachinesetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMemoryBalloonSize(request *IMachinegetMemoryBalloonSize) (*IMachinegetMemoryBalloonSizeResponse, error) {
	response := new(IMachinegetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMemoryBalloonSize(request *IMachinesetMemoryBalloonSize) (*IMachinesetMemoryBalloonSizeResponse, error) {
	response := new(IMachinesetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetPageFusionEnabled(request *IMachinegetPageFusionEnabled) (*IMachinegetPageFusionEnabledResponse, error) {
	response := new(IMachinegetPageFusionEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetPageFusionEnabled(request *IMachinesetPageFusionEnabled) (*IMachinesetPageFusionEnabledResponse, error) {
	response := new(IMachinesetPageFusionEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGraphicsControllerType(request *IMachinegetGraphicsControllerType) (*IMachinegetGraphicsControllerTypeResponse, error) {
	response := new(IMachinegetGraphicsControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGraphicsControllerType(request *IMachinesetGraphicsControllerType) (*IMachinesetGraphicsControllerTypeResponse, error) {
	response := new(IMachinesetGraphicsControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVRAMSize(request *IMachinegetVRAMSize) (*IMachinegetVRAMSizeResponse, error) {
	response := new(IMachinegetVRAMSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVRAMSize(request *IMachinesetVRAMSize) (*IMachinesetVRAMSizeResponse, error) {
	response := new(IMachinesetVRAMSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccelerate3DEnabled(request *IMachinegetAccelerate3DEnabled) (*IMachinegetAccelerate3DEnabledResponse, error) {
	response := new(IMachinegetAccelerate3DEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAccelerate3DEnabled(request *IMachinesetAccelerate3DEnabled) (*IMachinesetAccelerate3DEnabledResponse, error) {
	response := new(IMachinesetAccelerate3DEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccelerate2DVideoEnabled(request *IMachinegetAccelerate2DVideoEnabled) (*IMachinegetAccelerate2DVideoEnabledResponse, error) {
	response := new(IMachinegetAccelerate2DVideoEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAccelerate2DVideoEnabled(request *IMachinesetAccelerate2DVideoEnabled) (*IMachinesetAccelerate2DVideoEnabledResponse, error) {
	response := new(IMachinesetAccelerate2DVideoEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMonitorCount(request *IMachinegetMonitorCount) (*IMachinegetMonitorCountResponse, error) {
	response := new(IMachinegetMonitorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMonitorCount(request *IMachinesetMonitorCount) (*IMachinesetMonitorCountResponse, error) {
	response := new(IMachinesetMonitorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureEnabled(request *IMachinegetVideoCaptureEnabled) (*IMachinegetVideoCaptureEnabledResponse, error) {
	response := new(IMachinegetVideoCaptureEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureEnabled(request *IMachinesetVideoCaptureEnabled) (*IMachinesetVideoCaptureEnabledResponse, error) {
	response := new(IMachinesetVideoCaptureEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureScreens(request *IMachinegetVideoCaptureScreens) (*IMachinegetVideoCaptureScreensResponse, error) {
	response := new(IMachinegetVideoCaptureScreensResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureScreens(request *IMachinesetVideoCaptureScreens) (*IMachinesetVideoCaptureScreensResponse, error) {
	response := new(IMachinesetVideoCaptureScreensResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureFile(request *IMachinegetVideoCaptureFile) (*IMachinegetVideoCaptureFileResponse, error) {
	response := new(IMachinegetVideoCaptureFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureFile(request *IMachinesetVideoCaptureFile) (*IMachinesetVideoCaptureFileResponse, error) {
	response := new(IMachinesetVideoCaptureFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureWidth(request *IMachinegetVideoCaptureWidth) (*IMachinegetVideoCaptureWidthResponse, error) {
	response := new(IMachinegetVideoCaptureWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureWidth(request *IMachinesetVideoCaptureWidth) (*IMachinesetVideoCaptureWidthResponse, error) {
	response := new(IMachinesetVideoCaptureWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureHeight(request *IMachinegetVideoCaptureHeight) (*IMachinegetVideoCaptureHeightResponse, error) {
	response := new(IMachinegetVideoCaptureHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureHeight(request *IMachinesetVideoCaptureHeight) (*IMachinesetVideoCaptureHeightResponse, error) {
	response := new(IMachinesetVideoCaptureHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureRate(request *IMachinegetVideoCaptureRate) (*IMachinegetVideoCaptureRateResponse, error) {
	response := new(IMachinegetVideoCaptureRateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureRate(request *IMachinesetVideoCaptureRate) (*IMachinesetVideoCaptureRateResponse, error) {
	response := new(IMachinesetVideoCaptureRateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureFPS(request *IMachinegetVideoCaptureFPS) (*IMachinegetVideoCaptureFPSResponse, error) {
	response := new(IMachinegetVideoCaptureFPSResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureFPS(request *IMachinesetVideoCaptureFPS) (*IMachinesetVideoCaptureFPSResponse, error) {
	response := new(IMachinesetVideoCaptureFPSResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureMaxTime(request *IMachinegetVideoCaptureMaxTime) (*IMachinegetVideoCaptureMaxTimeResponse, error) {
	response := new(IMachinegetVideoCaptureMaxTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureMaxTime(request *IMachinesetVideoCaptureMaxTime) (*IMachinesetVideoCaptureMaxTimeResponse, error) {
	response := new(IMachinesetVideoCaptureMaxTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureMaxFileSize(request *IMachinegetVideoCaptureMaxFileSize) (*IMachinegetVideoCaptureMaxFileSizeResponse, error) {
	response := new(IMachinegetVideoCaptureMaxFileSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureMaxFileSize(request *IMachinesetVideoCaptureMaxFileSize) (*IMachinesetVideoCaptureMaxFileSizeResponse, error) {
	response := new(IMachinesetVideoCaptureMaxFileSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureOptions(request *IMachinegetVideoCaptureOptions) (*IMachinegetVideoCaptureOptionsResponse, error) {
	response := new(IMachinegetVideoCaptureOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureOptions(request *IMachinesetVideoCaptureOptions) (*IMachinesetVideoCaptureOptionsResponse, error) {
	response := new(IMachinesetVideoCaptureOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetBIOSSettings(request *IMachinegetBIOSSettings) (*IMachinegetBIOSSettingsResponse, error) {
	response := new(IMachinegetBIOSSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFirmwareType(request *IMachinegetFirmwareType) (*IMachinegetFirmwareTypeResponse, error) {
	response := new(IMachinegetFirmwareTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFirmwareType(request *IMachinesetFirmwareType) (*IMachinesetFirmwareTypeResponse, error) {
	response := new(IMachinesetFirmwareTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetPointingHIDType(request *IMachinegetPointingHIDType) (*IMachinegetPointingHIDTypeResponse, error) {
	response := new(IMachinegetPointingHIDTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetPointingHIDType(request *IMachinesetPointingHIDType) (*IMachinesetPointingHIDTypeResponse, error) {
	response := new(IMachinesetPointingHIDTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetKeyboardHIDType(request *IMachinegetKeyboardHIDType) (*IMachinegetKeyboardHIDTypeResponse, error) {
	response := new(IMachinegetKeyboardHIDTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetKeyboardHIDType(request *IMachinesetKeyboardHIDType) (*IMachinesetKeyboardHIDTypeResponse, error) {
	response := new(IMachinesetKeyboardHIDTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHPETEnabled(request *IMachinegetHPETEnabled) (*IMachinegetHPETEnabledResponse, error) {
	response := new(IMachinegetHPETEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHPETEnabled(request *IMachinesetHPETEnabled) (*IMachinesetHPETEnabledResponse, error) {
	response := new(IMachinesetHPETEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetChipsetType(request *IMachinegetChipsetType) (*IMachinegetChipsetTypeResponse, error) {
	response := new(IMachinegetChipsetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetChipsetType(request *IMachinesetChipsetType) (*IMachinesetChipsetTypeResponse, error) {
	response := new(IMachinesetChipsetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSnapshotFolder(request *IMachinegetSnapshotFolder) (*IMachinegetSnapshotFolderResponse, error) {
	response := new(IMachinegetSnapshotFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetSnapshotFolder(request *IMachinesetSnapshotFolder) (*IMachinesetSnapshotFolderResponse, error) {
	response := new(IMachinesetSnapshotFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVRDEServer(request *IMachinegetVRDEServer) (*IMachinegetVRDEServerResponse, error) {
	response := new(IMachinegetVRDEServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetEmulatedUSBCardReaderEnabled(request *IMachinegetEmulatedUSBCardReaderEnabled) (*IMachinegetEmulatedUSBCardReaderEnabledResponse, error) {
	response := new(IMachinegetEmulatedUSBCardReaderEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetEmulatedUSBCardReaderEnabled(request *IMachinesetEmulatedUSBCardReaderEnabled) (*IMachinesetEmulatedUSBCardReaderEnabledResponse, error) {
	response := new(IMachinesetEmulatedUSBCardReaderEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMediumAttachments(request *IMachinegetMediumAttachments) (*IMachinegetMediumAttachmentsResponse, error) {
	response := new(IMachinegetMediumAttachmentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBControllers(request *IMachinegetUSBControllers) (*IMachinegetUSBControllersResponse, error) {
	response := new(IMachinegetUSBControllersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBDeviceFilters(request *IMachinegetUSBDeviceFilters) (*IMachinegetUSBDeviceFiltersResponse, error) {
	response := new(IMachinegetUSBDeviceFiltersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAudioAdapter(request *IMachinegetAudioAdapter) (*IMachinegetAudioAdapterResponse, error) {
	response := new(IMachinegetAudioAdapterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllers(request *IMachinegetStorageControllers) (*IMachinegetStorageControllersResponse, error) {
	response := new(IMachinegetStorageControllersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSettingsFilePath(request *IMachinegetSettingsFilePath) (*IMachinegetSettingsFilePathResponse, error) {
	response := new(IMachinegetSettingsFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSettingsModified(request *IMachinegetSettingsModified) (*IMachinegetSettingsModifiedResponse, error) {
	response := new(IMachinegetSettingsModifiedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionState(request *IMachinegetSessionState) (*IMachinegetSessionStateResponse, error) {
	response := new(IMachinegetSessionStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionName(request *IMachinegetSessionName) (*IMachinegetSessionNameResponse, error) {
	response := new(IMachinegetSessionNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionPID(request *IMachinegetSessionPID) (*IMachinegetSessionPIDResponse, error) {
	response := new(IMachinegetSessionPIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetState(request *IMachinegetState) (*IMachinegetStateResponse, error) {
	response := new(IMachinegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetLastStateChange(request *IMachinegetLastStateChange) (*IMachinegetLastStateChangeResponse, error) {
	response := new(IMachinegetLastStateChangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStateFilePath(request *IMachinegetStateFilePath) (*IMachinegetStateFilePathResponse, error) {
	response := new(IMachinegetStateFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetLogFolder(request *IMachinegetLogFolder) (*IMachinegetLogFolderResponse, error) {
	response := new(IMachinegetLogFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCurrentSnapshot(request *IMachinegetCurrentSnapshot) (*IMachinegetCurrentSnapshotResponse, error) {
	response := new(IMachinegetCurrentSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSnapshotCount(request *IMachinegetSnapshotCount) (*IMachinegetSnapshotCountResponse, error) {
	response := new(IMachinegetSnapshotCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCurrentStateModified(request *IMachinegetCurrentStateModified) (*IMachinegetCurrentStateModifiedResponse, error) {
	response := new(IMachinegetCurrentStateModifiedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSharedFolders(request *IMachinegetSharedFolders) (*IMachinegetSharedFoldersResponse, error) {
	response := new(IMachinegetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetClipboardMode(request *IMachinegetClipboardMode) (*IMachinegetClipboardModeResponse, error) {
	response := new(IMachinegetClipboardModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetClipboardMode(request *IMachinesetClipboardMode) (*IMachinesetClipboardModeResponse, error) {
	response := new(IMachinesetClipboardModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetDnDMode(request *IMachinegetDnDMode) (*IMachinegetDnDModeResponse, error) {
	response := new(IMachinegetDnDModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetDnDMode(request *IMachinesetDnDMode) (*IMachinesetDnDModeResponse, error) {
	response := new(IMachinesetDnDModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterEnabled(request *IMachinegetTeleporterEnabled) (*IMachinegetTeleporterEnabledResponse, error) {
	response := new(IMachinegetTeleporterEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterEnabled(request *IMachinesetTeleporterEnabled) (*IMachinesetTeleporterEnabledResponse, error) {
	response := new(IMachinesetTeleporterEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterPort(request *IMachinegetTeleporterPort) (*IMachinegetTeleporterPortResponse, error) {
	response := new(IMachinegetTeleporterPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterPort(request *IMachinesetTeleporterPort) (*IMachinesetTeleporterPortResponse, error) {
	response := new(IMachinesetTeleporterPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterAddress(request *IMachinegetTeleporterAddress) (*IMachinegetTeleporterAddressResponse, error) {
	response := new(IMachinegetTeleporterAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterAddress(request *IMachinesetTeleporterAddress) (*IMachinesetTeleporterAddressResponse, error) {
	response := new(IMachinesetTeleporterAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterPassword(request *IMachinegetTeleporterPassword) (*IMachinegetTeleporterPasswordResponse, error) {
	response := new(IMachinegetTeleporterPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterPassword(request *IMachinesetTeleporterPassword) (*IMachinesetTeleporterPasswordResponse, error) {
	response := new(IMachinesetTeleporterPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetParavirtProvider(request *IMachinegetParavirtProvider) (*IMachinegetParavirtProviderResponse, error) {
	response := new(IMachinegetParavirtProviderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetParavirtProvider(request *IMachinesetParavirtProvider) (*IMachinesetParavirtProviderResponse, error) {
	response := new(IMachinesetParavirtProviderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultToleranceState(request *IMachinegetFaultToleranceState) (*IMachinegetFaultToleranceStateResponse, error) {
	response := new(IMachinegetFaultToleranceStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultToleranceState(request *IMachinesetFaultToleranceState) (*IMachinesetFaultToleranceStateResponse, error) {
	response := new(IMachinesetFaultToleranceStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultTolerancePort(request *IMachinegetFaultTolerancePort) (*IMachinegetFaultTolerancePortResponse, error) {
	response := new(IMachinegetFaultTolerancePortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultTolerancePort(request *IMachinesetFaultTolerancePort) (*IMachinesetFaultTolerancePortResponse, error) {
	response := new(IMachinesetFaultTolerancePortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultToleranceAddress(request *IMachinegetFaultToleranceAddress) (*IMachinegetFaultToleranceAddressResponse, error) {
	response := new(IMachinegetFaultToleranceAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultToleranceAddress(request *IMachinesetFaultToleranceAddress) (*IMachinesetFaultToleranceAddressResponse, error) {
	response := new(IMachinesetFaultToleranceAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultTolerancePassword(request *IMachinegetFaultTolerancePassword) (*IMachinegetFaultTolerancePasswordResponse, error) {
	response := new(IMachinegetFaultTolerancePasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultTolerancePassword(request *IMachinesetFaultTolerancePassword) (*IMachinesetFaultTolerancePasswordResponse, error) {
	response := new(IMachinesetFaultTolerancePasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultToleranceSyncInterval(request *IMachinegetFaultToleranceSyncInterval) (*IMachinegetFaultToleranceSyncIntervalResponse, error) {
	response := new(IMachinegetFaultToleranceSyncIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultToleranceSyncInterval(request *IMachinesetFaultToleranceSyncInterval) (*IMachinesetFaultToleranceSyncIntervalResponse, error) {
	response := new(IMachinesetFaultToleranceSyncIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetRTCUseUTC(request *IMachinegetRTCUseUTC) (*IMachinegetRTCUseUTCResponse, error) {
	response := new(IMachinegetRTCUseUTCResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetRTCUseUTC(request *IMachinesetRTCUseUTC) (*IMachinesetRTCUseUTCResponse, error) {
	response := new(IMachinesetRTCUseUTCResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIOCacheEnabled(request *IMachinegetIOCacheEnabled) (*IMachinegetIOCacheEnabledResponse, error) {
	response := new(IMachinegetIOCacheEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIOCacheEnabled(request *IMachinesetIOCacheEnabled) (*IMachinesetIOCacheEnabledResponse, error) {
	response := new(IMachinesetIOCacheEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIOCacheSize(request *IMachinegetIOCacheSize) (*IMachinegetIOCacheSizeResponse, error) {
	response := new(IMachinegetIOCacheSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIOCacheSize(request *IMachinesetIOCacheSize) (*IMachinesetIOCacheSizeResponse, error) {
	response := new(IMachinesetIOCacheSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetPCIDeviceAssignments(request *IMachinegetPCIDeviceAssignments) (*IMachinegetPCIDeviceAssignmentsResponse, error) {
	response := new(IMachinegetPCIDeviceAssignmentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetBandwidthControl(request *IMachinegetBandwidthControl) (*IMachinegetBandwidthControlResponse, error) {
	response := new(IMachinegetBandwidthControlResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTracingEnabled(request *IMachinegetTracingEnabled) (*IMachinegetTracingEnabledResponse, error) {
	response := new(IMachinegetTracingEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTracingEnabled(request *IMachinesetTracingEnabled) (*IMachinesetTracingEnabledResponse, error) {
	response := new(IMachinesetTracingEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTracingConfig(request *IMachinegetTracingConfig) (*IMachinegetTracingConfigResponse, error) {
	response := new(IMachinegetTracingConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTracingConfig(request *IMachinesetTracingConfig) (*IMachinesetTracingConfigResponse, error) {
	response := new(IMachinesetTracingConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAllowTracingToAccessVM(request *IMachinegetAllowTracingToAccessVM) (*IMachinegetAllowTracingToAccessVMResponse, error) {
	response := new(IMachinegetAllowTracingToAccessVMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAllowTracingToAccessVM(request *IMachinesetAllowTracingToAccessVM) (*IMachinesetAllowTracingToAccessVMResponse, error) {
	response := new(IMachinesetAllowTracingToAccessVMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAutostartEnabled(request *IMachinegetAutostartEnabled) (*IMachinegetAutostartEnabledResponse, error) {
	response := new(IMachinegetAutostartEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAutostartEnabled(request *IMachinesetAutostartEnabled) (*IMachinesetAutostartEnabledResponse, error) {
	response := new(IMachinesetAutostartEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAutostartDelay(request *IMachinegetAutostartDelay) (*IMachinegetAutostartDelayResponse, error) {
	response := new(IMachinegetAutostartDelayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAutostartDelay(request *IMachinesetAutostartDelay) (*IMachinesetAutostartDelayResponse, error) {
	response := new(IMachinesetAutostartDelayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAutostopType(request *IMachinegetAutostopType) (*IMachinegetAutostopTypeResponse, error) {
	response := new(IMachinegetAutostopTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAutostopType(request *IMachinesetAutostopType) (*IMachinesetAutostopTypeResponse, error) {
	response := new(IMachinesetAutostopTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetDefaultFrontend(request *IMachinegetDefaultFrontend) (*IMachinegetDefaultFrontendResponse, error) {
	response := new(IMachinegetDefaultFrontendResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetDefaultFrontend(request *IMachinesetDefaultFrontend) (*IMachinesetDefaultFrontendResponse, error) {
	response := new(IMachinesetDefaultFrontendResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBProxyAvailable(request *IMachinegetUSBProxyAvailable) (*IMachinegetUSBProxyAvailableResponse, error) {
	response := new(IMachinegetUSBProxyAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVMProcessPriority(request *IMachinegetVMProcessPriority) (*IMachinegetVMProcessPriorityResponse, error) {
	response := new(IMachinegetVMProcessPriorityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVMProcessPriority(request *IMachinesetVMProcessPriority) (*IMachinesetVMProcessPriorityResponse, error) {
	response := new(IMachinesetVMProcessPriorityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinelockMachine(request *IMachinelockMachine) (*IMachinelockMachineResponse, error) {
	response := new(IMachinelockMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinelaunchVMProcess(request *IMachinelaunchVMProcess) (*IMachinelaunchVMProcessResponse, error) {
	response := new(IMachinelaunchVMProcessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetBootOrder(request *IMachinesetBootOrder) (*IMachinesetBootOrderResponse, error) {
	response := new(IMachinesetBootOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetBootOrder(request *IMachinegetBootOrder) (*IMachinegetBootOrderResponse, error) {
	response := new(IMachinegetBootOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineattachDevice(request *IMachineattachDevice) (*IMachineattachDeviceResponse, error) {
	response := new(IMachineattachDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineattachDeviceWithoutMedium(request *IMachineattachDeviceWithoutMedium) (*IMachineattachDeviceWithoutMediumResponse, error) {
	response := new(IMachineattachDeviceWithoutMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedetachDevice(request *IMachinedetachDevice) (*IMachinedetachDeviceResponse, error) {
	response := new(IMachinedetachDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinepassthroughDevice(request *IMachinepassthroughDevice) (*IMachinepassthroughDeviceResponse, error) {
	response := new(IMachinepassthroughDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinetemporaryEjectDevice(request *IMachinetemporaryEjectDevice) (*IMachinetemporaryEjectDeviceResponse, error) {
	response := new(IMachinetemporaryEjectDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinenonRotationalDevice(request *IMachinenonRotationalDevice) (*IMachinenonRotationalDeviceResponse, error) {
	response := new(IMachinenonRotationalDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAutoDiscardForDevice(request *IMachinesetAutoDiscardForDevice) (*IMachinesetAutoDiscardForDeviceResponse, error) {
	response := new(IMachinesetAutoDiscardForDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHotPluggableForDevice(request *IMachinesetHotPluggableForDevice) (*IMachinesetHotPluggableForDeviceResponse, error) {
	response := new(IMachinesetHotPluggableForDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetBandwidthGroupForDevice(request *IMachinesetBandwidthGroupForDevice) (*IMachinesetBandwidthGroupForDeviceResponse, error) {
	response := new(IMachinesetBandwidthGroupForDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetNoBandwidthGroupForDevice(request *IMachinesetNoBandwidthGroupForDevice) (*IMachinesetNoBandwidthGroupForDeviceResponse, error) {
	response := new(IMachinesetNoBandwidthGroupForDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineunmountMedium(request *IMachineunmountMedium) (*IMachineunmountMediumResponse, error) {
	response := new(IMachineunmountMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinemountMedium(request *IMachinemountMedium) (*IMachinemountMediumResponse, error) {
	response := new(IMachinemountMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMedium(request *IMachinegetMedium) (*IMachinegetMediumResponse, error) {
	response := new(IMachinegetMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMediumAttachmentsOfController(request *IMachinegetMediumAttachmentsOfController) (*IMachinegetMediumAttachmentsOfControllerResponse, error) {
	response := new(IMachinegetMediumAttachmentsOfControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMediumAttachment(request *IMachinegetMediumAttachment) (*IMachinegetMediumAttachmentResponse, error) {
	response := new(IMachinegetMediumAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineattachHostPCIDevice(request *IMachineattachHostPCIDevice) (*IMachineattachHostPCIDeviceResponse, error) {
	response := new(IMachineattachHostPCIDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedetachHostPCIDevice(request *IMachinedetachHostPCIDevice) (*IMachinedetachHostPCIDeviceResponse, error) {
	response := new(IMachinedetachHostPCIDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetNetworkAdapter(request *IMachinegetNetworkAdapter) (*IMachinegetNetworkAdapterResponse, error) {
	response := new(IMachinegetNetworkAdapterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineaddStorageController(request *IMachineaddStorageController) (*IMachineaddStorageControllerResponse, error) {
	response := new(IMachineaddStorageControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllerByName(request *IMachinegetStorageControllerByName) (*IMachinegetStorageControllerByNameResponse, error) {
	response := new(IMachinegetStorageControllerByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllerByInstance(request *IMachinegetStorageControllerByInstance) (*IMachinegetStorageControllerByInstanceResponse, error) {
	response := new(IMachinegetStorageControllerByInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveStorageController(request *IMachineremoveStorageController) (*IMachineremoveStorageControllerResponse, error) {
	response := new(IMachineremoveStorageControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetStorageControllerBootable(request *IMachinesetStorageControllerBootable) (*IMachinesetStorageControllerBootableResponse, error) {
	response := new(IMachinesetStorageControllerBootableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineaddUSBController(request *IMachineaddUSBController) (*IMachineaddUSBControllerResponse, error) {
	response := new(IMachineaddUSBControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveUSBController(request *IMachineremoveUSBController) (*IMachineremoveUSBControllerResponse, error) {
	response := new(IMachineremoveUSBControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBControllerByName(request *IMachinegetUSBControllerByName) (*IMachinegetUSBControllerByNameResponse, error) {
	response := new(IMachinegetUSBControllerByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBControllerCountByType(request *IMachinegetUSBControllerCountByType) (*IMachinegetUSBControllerCountByTypeResponse, error) {
	response := new(IMachinegetUSBControllerCountByTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSerialPort(request *IMachinegetSerialPort) (*IMachinegetSerialPortResponse, error) {
	response := new(IMachinegetSerialPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetParallelPort(request *IMachinegetParallelPort) (*IMachinegetParallelPortResponse, error) {
	response := new(IMachinegetParallelPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetExtraDataKeys(request *IMachinegetExtraDataKeys) (*IMachinegetExtraDataKeysResponse, error) {
	response := new(IMachinegetExtraDataKeysResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetExtraData(request *IMachinegetExtraData) (*IMachinegetExtraDataResponse, error) {
	response := new(IMachinegetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetExtraData(request *IMachinesetExtraData) (*IMachinesetExtraDataResponse, error) {
	response := new(IMachinesetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUProperty(request *IMachinegetCPUProperty) (*IMachinegetCPUPropertyResponse, error) {
	response := new(IMachinegetCPUPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUProperty(request *IMachinesetCPUProperty) (*IMachinesetCPUPropertyResponse, error) {
	response := new(IMachinesetCPUPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUIDLeaf(request *IMachinegetCPUIDLeaf) (*IMachinegetCPUIDLeafResponse, error) {
	response := new(IMachinegetCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUIDLeaf(request *IMachinesetCPUIDLeaf) (*IMachinesetCPUIDLeafResponse, error) {
	response := new(IMachinesetCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveCPUIDLeaf(request *IMachineremoveCPUIDLeaf) (*IMachineremoveCPUIDLeafResponse, error) {
	response := new(IMachineremoveCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveAllCPUIDLeaves(request *IMachineremoveAllCPUIDLeaves) (*IMachineremoveAllCPUIDLeavesResponse, error) {
	response := new(IMachineremoveAllCPUIDLeavesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHWVirtExProperty(request *IMachinegetHWVirtExProperty) (*IMachinegetHWVirtExPropertyResponse, error) {
	response := new(IMachinegetHWVirtExPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHWVirtExProperty(request *IMachinesetHWVirtExProperty) (*IMachinesetHWVirtExPropertyResponse, error) {
	response := new(IMachinesetHWVirtExPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetSettingsFilePath(request *IMachinesetSettingsFilePath) (*IMachinesetSettingsFilePathResponse, error) {
	response := new(IMachinesetSettingsFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesaveSettings(request *IMachinesaveSettings) (*IMachinesaveSettingsResponse, error) {
	response := new(IMachinesaveSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinediscardSettings(request *IMachinediscardSettings) (*IMachinediscardSettingsResponse, error) {
	response := new(IMachinediscardSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineunregister(request *IMachineunregister) (*IMachineunregisterResponse, error) {
	response := new(IMachineunregisterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedeleteConfig(request *IMachinedeleteConfig) (*IMachinedeleteConfigResponse, error) {
	response := new(IMachinedeleteConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineexportTo(request *IMachineexportTo) (*IMachineexportToResponse, error) {
	response := new(IMachineexportToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinefindSnapshot(request *IMachinefindSnapshot) (*IMachinefindSnapshotResponse, error) {
	response := new(IMachinefindSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinecreateSharedFolder(request *IMachinecreateSharedFolder) (*IMachinecreateSharedFolderResponse, error) {
	response := new(IMachinecreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveSharedFolder(request *IMachineremoveSharedFolder) (*IMachineremoveSharedFolderResponse, error) {
	response := new(IMachineremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinecanShowConsoleWindow(request *IMachinecanShowConsoleWindow) (*IMachinecanShowConsoleWindowResponse, error) {
	response := new(IMachinecanShowConsoleWindowResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineshowConsoleWindow(request *IMachineshowConsoleWindow) (*IMachineshowConsoleWindowResponse, error) {
	response := new(IMachineshowConsoleWindowResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestProperty(request *IMachinegetGuestProperty) (*IMachinegetGuestPropertyResponse, error) {
	response := new(IMachinegetGuestPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestPropertyValue(request *IMachinegetGuestPropertyValue) (*IMachinegetGuestPropertyValueResponse, error) {
	response := new(IMachinegetGuestPropertyValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestPropertyTimestamp(request *IMachinegetGuestPropertyTimestamp) (*IMachinegetGuestPropertyTimestampResponse, error) {
	response := new(IMachinegetGuestPropertyTimestampResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGuestProperty(request *IMachinesetGuestProperty) (*IMachinesetGuestPropertyResponse, error) {
	response := new(IMachinesetGuestPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGuestPropertyValue(request *IMachinesetGuestPropertyValue) (*IMachinesetGuestPropertyValueResponse, error) {
	response := new(IMachinesetGuestPropertyValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedeleteGuestProperty(request *IMachinedeleteGuestProperty) (*IMachinedeleteGuestPropertyResponse, error) {
	response := new(IMachinedeleteGuestPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineenumerateGuestProperties(request *IMachineenumerateGuestProperties) (*IMachineenumerateGuestPropertiesResponse, error) {
	response := new(IMachineenumerateGuestPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequerySavedGuestScreenInfo(request *IMachinequerySavedGuestScreenInfo) (*IMachinequerySavedGuestScreenInfoResponse, error) {
	response := new(IMachinequerySavedGuestScreenInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadSavedThumbnailToArray(request *IMachinereadSavedThumbnailToArray) (*IMachinereadSavedThumbnailToArrayResponse, error) {
	response := new(IMachinereadSavedThumbnailToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequerySavedScreenshotInfo(request *IMachinequerySavedScreenshotInfo) (*IMachinequerySavedScreenshotInfoResponse, error) {
	response := new(IMachinequerySavedScreenshotInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadSavedScreenshotToArray(request *IMachinereadSavedScreenshotToArray) (*IMachinereadSavedScreenshotToArrayResponse, error) {
	response := new(IMachinereadSavedScreenshotToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinehotPlugCPU(request *IMachinehotPlugCPU) (*IMachinehotPlugCPUResponse, error) {
	response := new(IMachinehotPlugCPUResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinehotUnplugCPU(request *IMachinehotUnplugCPU) (*IMachinehotUnplugCPUResponse, error) {
	response := new(IMachinehotUnplugCPUResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUStatus(request *IMachinegetCPUStatus) (*IMachinegetCPUStatusResponse, error) {
	response := new(IMachinegetCPUStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetEffectiveParavirtProvider(request *IMachinegetEffectiveParavirtProvider) (*IMachinegetEffectiveParavirtProviderResponse, error) {
	response := new(IMachinegetEffectiveParavirtProviderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequeryLogFilename(request *IMachinequeryLogFilename) (*IMachinequeryLogFilenameResponse, error) {
	response := new(IMachinequeryLogFilenameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadLog(request *IMachinereadLog) (*IMachinereadLogResponse, error) {
	response := new(IMachinereadLogResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinecloneTo(request *IMachinecloneTo) (*IMachinecloneToResponse, error) {
	response := new(IMachinecloneToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesaveState(request *IMachinesaveState) (*IMachinesaveStateResponse, error) {
	response := new(IMachinesaveStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineadoptSavedState(request *IMachineadoptSavedState) (*IMachineadoptSavedStateResponse, error) {
	response := new(IMachineadoptSavedStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinediscardSavedState(request *IMachinediscardSavedState) (*IMachinediscardSavedStateResponse, error) {
	response := new(IMachinediscardSavedStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinetakeSnapshot(request *IMachinetakeSnapshot) (*IMachinetakeSnapshotResponse, error) {
	response := new(IMachinetakeSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedeleteSnapshot(request *IMachinedeleteSnapshot) (*IMachinedeleteSnapshotResponse, error) {
	response := new(IMachinedeleteSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedeleteSnapshotAndAllChildren(request *IMachinedeleteSnapshotAndAllChildren) (*IMachinedeleteSnapshotAndAllChildrenResponse, error) {
	response := new(IMachinedeleteSnapshotAndAllChildrenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedeleteSnapshotRange(request *IMachinedeleteSnapshotRange) (*IMachinedeleteSnapshotRangeResponse, error) {
	response := new(IMachinedeleteSnapshotRangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinerestoreSnapshot(request *IMachinerestoreSnapshot) (*IMachinerestoreSnapshotResponse, error) {
	response := new(IMachinerestoreSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineapplyDefaults(request *IMachineapplyDefaults) (*IMachineapplyDefaultsResponse, error) {
	response := new(IMachineapplyDefaultsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEmulatedUSBgetWebcams(request *IEmulatedUSBgetWebcams) (*IEmulatedUSBgetWebcamsResponse, error) {
	response := new(IEmulatedUSBgetWebcamsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEmulatedUSBwebcamAttach(request *IEmulatedUSBwebcamAttach) (*IEmulatedUSBwebcamAttachResponse, error) {
	response := new(IEmulatedUSBwebcamAttachResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEmulatedUSBwebcamDetach(request *IEmulatedUSBwebcamDetach) (*IEmulatedUSBwebcamDetachResponse, error) {
	response := new(IEmulatedUSBwebcamDetachResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetMachine(request *IConsolegetMachine) (*IConsolegetMachineResponse, error) {
	response := new(IConsolegetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetState(request *IConsolegetState) (*IConsolegetStateResponse, error) {
	response := new(IConsolegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetGuest(request *IConsolegetGuest) (*IConsolegetGuestResponse, error) {
	response := new(IConsolegetGuestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetKeyboard(request *IConsolegetKeyboard) (*IConsolegetKeyboardResponse, error) {
	response := new(IConsolegetKeyboardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetMouse(request *IConsolegetMouse) (*IConsolegetMouseResponse, error) {
	response := new(IConsolegetMouseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetDisplay(request *IConsolegetDisplay) (*IConsolegetDisplayResponse, error) {
	response := new(IConsolegetDisplayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetDebugger(request *IConsolegetDebugger) (*IConsolegetDebuggerResponse, error) {
	response := new(IConsolegetDebuggerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetUSBDevices(request *IConsolegetUSBDevices) (*IConsolegetUSBDevicesResponse, error) {
	response := new(IConsolegetUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetRemoteUSBDevices(request *IConsolegetRemoteUSBDevices) (*IConsolegetRemoteUSBDevicesResponse, error) {
	response := new(IConsolegetRemoteUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetSharedFolders(request *IConsolegetSharedFolders) (*IConsolegetSharedFoldersResponse, error) {
	response := new(IConsolegetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetVRDEServerInfo(request *IConsolegetVRDEServerInfo) (*IConsolegetVRDEServerInfoResponse, error) {
	response := new(IConsolegetVRDEServerInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetEventSource(request *IConsolegetEventSource) (*IConsolegetEventSourceResponse, error) {
	response := new(IConsolegetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetAttachedPCIDevices(request *IConsolegetAttachedPCIDevices) (*IConsolegetAttachedPCIDevicesResponse, error) {
	response := new(IConsolegetAttachedPCIDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetUseHostClipboard(request *IConsolegetUseHostClipboard) (*IConsolegetUseHostClipboardResponse, error) {
	response := new(IConsolegetUseHostClipboardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolesetUseHostClipboard(request *IConsolesetUseHostClipboard) (*IConsolesetUseHostClipboardResponse, error) {
	response := new(IConsolesetUseHostClipboardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetEmulatedUSB(request *IConsolegetEmulatedUSB) (*IConsolegetEmulatedUSBResponse, error) {
	response := new(IConsolegetEmulatedUSBResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerUp(request *IConsolepowerUp) (*IConsolepowerUpResponse, error) {
	response := new(IConsolepowerUpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerUpPaused(request *IConsolepowerUpPaused) (*IConsolepowerUpPausedResponse, error) {
	response := new(IConsolepowerUpPausedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerDown(request *IConsolepowerDown) (*IConsolepowerDownResponse, error) {
	response := new(IConsolepowerDownResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolereset(request *IConsolereset) (*IConsoleresetResponse, error) {
	response := new(IConsoleresetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepause(request *IConsolepause) (*IConsolepauseResponse, error) {
	response := new(IConsolepauseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleresume(request *IConsoleresume) (*IConsoleresumeResponse, error) {
	response := new(IConsoleresumeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerButton(request *IConsolepowerButton) (*IConsolepowerButtonResponse, error) {
	response := new(IConsolepowerButtonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolesleepButton(request *IConsolesleepButton) (*IConsolesleepButtonResponse, error) {
	response := new(IConsolesleepButtonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetPowerButtonHandled(request *IConsolegetPowerButtonHandled) (*IConsolegetPowerButtonHandledResponse, error) {
	response := new(IConsolegetPowerButtonHandledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetGuestEnteredACPIMode(request *IConsolegetGuestEnteredACPIMode) (*IConsolegetGuestEnteredACPIModeResponse, error) {
	response := new(IConsolegetGuestEnteredACPIModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetDeviceActivity(request *IConsolegetDeviceActivity) (*IConsolegetDeviceActivityResponse, error) {
	response := new(IConsolegetDeviceActivityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleattachUSBDevice(request *IConsoleattachUSBDevice) (*IConsoleattachUSBDeviceResponse, error) {
	response := new(IConsoleattachUSBDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoledetachUSBDevice(request *IConsoledetachUSBDevice) (*IConsoledetachUSBDeviceResponse, error) {
	response := new(IConsoledetachUSBDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolefindUSBDeviceByAddress(request *IConsolefindUSBDeviceByAddress) (*IConsolefindUSBDeviceByAddressResponse, error) {
	response := new(IConsolefindUSBDeviceByAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolefindUSBDeviceById(request *IConsolefindUSBDeviceById) (*IConsolefindUSBDeviceByIdResponse, error) {
	response := new(IConsolefindUSBDeviceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolecreateSharedFolder(request *IConsolecreateSharedFolder) (*IConsolecreateSharedFolderResponse, error) {
	response := new(IConsolecreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleremoveSharedFolder(request *IConsoleremoveSharedFolder) (*IConsoleremoveSharedFolderResponse, error) {
	response := new(IConsoleremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleteleport(request *IConsoleteleport) (*IConsoleteleportResponse, error) {
	response := new(IConsoleteleportResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleaddDiskEncryptionPassword(request *IConsoleaddDiskEncryptionPassword) (*IConsoleaddDiskEncryptionPasswordResponse, error) {
	response := new(IConsoleaddDiskEncryptionPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleaddDiskEncryptionPasswords(request *IConsoleaddDiskEncryptionPasswords) (*IConsoleaddDiskEncryptionPasswordsResponse, error) {
	response := new(IConsoleaddDiskEncryptionPasswordsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleremoveDiskEncryptionPassword(request *IConsoleremoveDiskEncryptionPassword) (*IConsoleremoveDiskEncryptionPasswordResponse, error) {
	response := new(IConsoleremoveDiskEncryptionPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleclearAllDiskEncryptionPasswords(request *IConsoleclearAllDiskEncryptionPasswords) (*IConsoleclearAllDiskEncryptionPasswordsResponse, error) {
	response := new(IConsoleclearAllDiskEncryptionPasswordsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetName(request *IHostNetworkInterfacegetName) (*IHostNetworkInterfacegetNameResponse, error) {
	response := new(IHostNetworkInterfacegetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetShortName(request *IHostNetworkInterfacegetShortName) (*IHostNetworkInterfacegetShortNameResponse, error) {
	response := new(IHostNetworkInterfacegetShortNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetId(request *IHostNetworkInterfacegetId) (*IHostNetworkInterfacegetIdResponse, error) {
	response := new(IHostNetworkInterfacegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetNetworkName(request *IHostNetworkInterfacegetNetworkName) (*IHostNetworkInterfacegetNetworkNameResponse, error) {
	response := new(IHostNetworkInterfacegetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetDHCPEnabled(request *IHostNetworkInterfacegetDHCPEnabled) (*IHostNetworkInterfacegetDHCPEnabledResponse, error) {
	response := new(IHostNetworkInterfacegetDHCPEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPAddress(request *IHostNetworkInterfacegetIPAddress) (*IHostNetworkInterfacegetIPAddressResponse, error) {
	response := new(IHostNetworkInterfacegetIPAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetNetworkMask(request *IHostNetworkInterfacegetNetworkMask) (*IHostNetworkInterfacegetNetworkMaskResponse, error) {
	response := new(IHostNetworkInterfacegetNetworkMaskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6Supported(request *IHostNetworkInterfacegetIPV6Supported) (*IHostNetworkInterfacegetIPV6SupportedResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6SupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6Address(request *IHostNetworkInterfacegetIPV6Address) (*IHostNetworkInterfacegetIPV6AddressResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6AddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6NetworkMaskPrefixLength(request *IHostNetworkInterfacegetIPV6NetworkMaskPrefixLength) (*IHostNetworkInterfacegetIPV6NetworkMaskPrefixLengthResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6NetworkMaskPrefixLengthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetHardwareAddress(request *IHostNetworkInterfacegetHardwareAddress) (*IHostNetworkInterfacegetHardwareAddressResponse, error) {
	response := new(IHostNetworkInterfacegetHardwareAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetMediumType(request *IHostNetworkInterfacegetMediumType) (*IHostNetworkInterfacegetMediumTypeResponse, error) {
	response := new(IHostNetworkInterfacegetMediumTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetStatus(request *IHostNetworkInterfacegetStatus) (*IHostNetworkInterfacegetStatusResponse, error) {
	response := new(IHostNetworkInterfacegetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetInterfaceType(request *IHostNetworkInterfacegetInterfaceType) (*IHostNetworkInterfacegetInterfaceTypeResponse, error) {
	response := new(IHostNetworkInterfacegetInterfaceTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableStaticIPConfig(request *IHostNetworkInterfaceenableStaticIPConfig) (*IHostNetworkInterfaceenableStaticIPConfigResponse, error) {
	response := new(IHostNetworkInterfaceenableStaticIPConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableStaticIPConfigV6(request *IHostNetworkInterfaceenableStaticIPConfigV6) (*IHostNetworkInterfaceenableStaticIPConfigV6Response, error) {
	response := new(IHostNetworkInterfaceenableStaticIPConfigV6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableDynamicIPConfig(request *IHostNetworkInterfaceenableDynamicIPConfig) (*IHostNetworkInterfaceenableDynamicIPConfigResponse, error) {
	response := new(IHostNetworkInterfaceenableDynamicIPConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceDHCPRediscover(request *IHostNetworkInterfaceDHCPRediscover) (*IHostNetworkInterfaceDHCPRediscoverResponse, error) {
	response := new(IHostNetworkInterfaceDHCPRediscoverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostVideoInputDevicegetName(request *IHostVideoInputDevicegetName) (*IHostVideoInputDevicegetNameResponse, error) {
	response := new(IHostVideoInputDevicegetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostVideoInputDevicegetPath(request *IHostVideoInputDevicegetPath) (*IHostVideoInputDevicegetPathResponse, error) {
	response := new(IHostVideoInputDevicegetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostVideoInputDevicegetAlias(request *IHostVideoInputDevicegetAlias) (*IHostVideoInputDevicegetAliasResponse, error) {
	response := new(IHostVideoInputDevicegetAliasResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetDVDDrives(request *IHostgetDVDDrives) (*IHostgetDVDDrivesResponse, error) {
	response := new(IHostgetDVDDrivesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetFloppyDrives(request *IHostgetFloppyDrives) (*IHostgetFloppyDrivesResponse, error) {
	response := new(IHostgetFloppyDrivesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUSBDevices(request *IHostgetUSBDevices) (*IHostgetUSBDevicesResponse, error) {
	response := new(IHostgetUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUSBDeviceFilters(request *IHostgetUSBDeviceFilters) (*IHostgetUSBDeviceFiltersResponse, error) {
	response := new(IHostgetUSBDeviceFiltersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetNetworkInterfaces(request *IHostgetNetworkInterfaces) (*IHostgetNetworkInterfacesResponse, error) {
	response := new(IHostgetNetworkInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetNameServers(request *IHostgetNameServers) (*IHostgetNameServersResponse, error) {
	response := new(IHostgetNameServersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetDomainName(request *IHostgetDomainName) (*IHostgetDomainNameResponse, error) {
	response := new(IHostgetDomainNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetSearchStrings(request *IHostgetSearchStrings) (*IHostgetSearchStringsResponse, error) {
	response := new(IHostgetSearchStringsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCount(request *IHostgetProcessorCount) (*IHostgetProcessorCountResponse, error) {
	response := new(IHostgetProcessorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorOnlineCount(request *IHostgetProcessorOnlineCount) (*IHostgetProcessorOnlineCountResponse, error) {
	response := new(IHostgetProcessorOnlineCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCoreCount(request *IHostgetProcessorCoreCount) (*IHostgetProcessorCoreCountResponse, error) {
	response := new(IHostgetProcessorCoreCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorOnlineCoreCount(request *IHostgetProcessorOnlineCoreCount) (*IHostgetProcessorOnlineCoreCountResponse, error) {
	response := new(IHostgetProcessorOnlineCoreCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetMemorySize(request *IHostgetMemorySize) (*IHostgetMemorySizeResponse, error) {
	response := new(IHostgetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetMemoryAvailable(request *IHostgetMemoryAvailable) (*IHostgetMemoryAvailableResponse, error) {
	response := new(IHostgetMemoryAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetOperatingSystem(request *IHostgetOperatingSystem) (*IHostgetOperatingSystemResponse, error) {
	response := new(IHostgetOperatingSystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetOSVersion(request *IHostgetOSVersion) (*IHostgetOSVersionResponse, error) {
	response := new(IHostgetOSVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUTCTime(request *IHostgetUTCTime) (*IHostgetUTCTimeResponse, error) {
	response := new(IHostgetUTCTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetAcceleration3DAvailable(request *IHostgetAcceleration3DAvailable) (*IHostgetAcceleration3DAvailableResponse, error) {
	response := new(IHostgetAcceleration3DAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetVideoInputDevices(request *IHostgetVideoInputDevices) (*IHostgetVideoInputDevicesResponse, error) {
	response := new(IHostgetVideoInputDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorSpeed(request *IHostgetProcessorSpeed) (*IHostgetProcessorSpeedResponse, error) {
	response := new(IHostgetProcessorSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorFeature(request *IHostgetProcessorFeature) (*IHostgetProcessorFeatureResponse, error) {
	response := new(IHostgetProcessorFeatureResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorDescription(request *IHostgetProcessorDescription) (*IHostgetProcessorDescriptionResponse, error) {
	response := new(IHostgetProcessorDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCPUIDLeaf(request *IHostgetProcessorCPUIDLeaf) (*IHostgetProcessorCPUIDLeafResponse, error) {
	response := new(IHostgetProcessorCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostcreateHostOnlyNetworkInterface(request *IHostcreateHostOnlyNetworkInterface) (*IHostcreateHostOnlyNetworkInterfaceResponse, error) {
	response := new(IHostcreateHostOnlyNetworkInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostremoveHostOnlyNetworkInterface(request *IHostremoveHostOnlyNetworkInterface) (*IHostremoveHostOnlyNetworkInterfaceResponse, error) {
	response := new(IHostremoveHostOnlyNetworkInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostcreateUSBDeviceFilter(request *IHostcreateUSBDeviceFilter) (*IHostcreateUSBDeviceFilterResponse, error) {
	response := new(IHostcreateUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostinsertUSBDeviceFilter(request *IHostinsertUSBDeviceFilter) (*IHostinsertUSBDeviceFilterResponse, error) {
	response := new(IHostinsertUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostremoveUSBDeviceFilter(request *IHostremoveUSBDeviceFilter) (*IHostremoveUSBDeviceFilterResponse, error) {
	response := new(IHostremoveUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostDVDDrive(request *IHostfindHostDVDDrive) (*IHostfindHostDVDDriveResponse, error) {
	response := new(IHostfindHostDVDDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostFloppyDrive(request *IHostfindHostFloppyDrive) (*IHostfindHostFloppyDriveResponse, error) {
	response := new(IHostfindHostFloppyDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfaceByName(request *IHostfindHostNetworkInterfaceByName) (*IHostfindHostNetworkInterfaceByNameResponse, error) {
	response := new(IHostfindHostNetworkInterfaceByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfaceById(request *IHostfindHostNetworkInterfaceById) (*IHostfindHostNetworkInterfaceByIdResponse, error) {
	response := new(IHostfindHostNetworkInterfaceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfacesOfType(request *IHostfindHostNetworkInterfacesOfType) (*IHostfindHostNetworkInterfacesOfTypeResponse, error) {
	response := new(IHostfindHostNetworkInterfacesOfTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindUSBDeviceById(request *IHostfindUSBDeviceById) (*IHostfindUSBDeviceByIdResponse, error) {
	response := new(IHostfindUSBDeviceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindUSBDeviceByAddress(request *IHostfindUSBDeviceByAddress) (*IHostfindUSBDeviceByAddressResponse, error) {
	response := new(IHostfindUSBDeviceByAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgenerateMACAddress(request *IHostgenerateMACAddress) (*IHostgenerateMACAddressResponse, error) {
	response := new(IHostgenerateMACAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestRAM(request *ISystemPropertiesgetMinGuestRAM) (*ISystemPropertiesgetMinGuestRAMResponse, error) {
	response := new(ISystemPropertiesgetMinGuestRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestRAM(request *ISystemPropertiesgetMaxGuestRAM) (*ISystemPropertiesgetMaxGuestRAMResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestVRAM(request *ISystemPropertiesgetMinGuestVRAM) (*ISystemPropertiesgetMinGuestVRAMResponse, error) {
	response := new(ISystemPropertiesgetMinGuestVRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestVRAM(request *ISystemPropertiesgetMaxGuestVRAM) (*ISystemPropertiesgetMaxGuestVRAMResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestVRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestCPUCount(request *ISystemPropertiesgetMinGuestCPUCount) (*ISystemPropertiesgetMinGuestCPUCountResponse, error) {
	response := new(ISystemPropertiesgetMinGuestCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestCPUCount(request *ISystemPropertiesgetMaxGuestCPUCount) (*ISystemPropertiesgetMaxGuestCPUCountResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestMonitors(request *ISystemPropertiesgetMaxGuestMonitors) (*ISystemPropertiesgetMaxGuestMonitorsResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestMonitorsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetInfoVDSize(request *ISystemPropertiesgetInfoVDSize) (*ISystemPropertiesgetInfoVDSizeResponse, error) {
	response := new(ISystemPropertiesgetInfoVDSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetSerialPortCount(request *ISystemPropertiesgetSerialPortCount) (*ISystemPropertiesgetSerialPortCountResponse, error) {
	response := new(ISystemPropertiesgetSerialPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetParallelPortCount(request *ISystemPropertiesgetParallelPortCount) (*ISystemPropertiesgetParallelPortCountResponse, error) {
	response := new(ISystemPropertiesgetParallelPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxBootPosition(request *ISystemPropertiesgetMaxBootPosition) (*ISystemPropertiesgetMaxBootPositionResponse, error) {
	response := new(ISystemPropertiesgetMaxBootPositionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetRawModeSupported(request *ISystemPropertiesgetRawModeSupported) (*ISystemPropertiesgetRawModeSupportedResponse, error) {
	response := new(ISystemPropertiesgetRawModeSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetExclusiveHwVirt(request *ISystemPropertiesgetExclusiveHwVirt) (*ISystemPropertiesgetExclusiveHwVirtResponse, error) {
	response := new(ISystemPropertiesgetExclusiveHwVirtResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetExclusiveHwVirt(request *ISystemPropertiessetExclusiveHwVirt) (*ISystemPropertiessetExclusiveHwVirtResponse, error) {
	response := new(ISystemPropertiessetExclusiveHwVirtResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultMachineFolder(request *ISystemPropertiesgetDefaultMachineFolder) (*ISystemPropertiesgetDefaultMachineFolderResponse, error) {
	response := new(ISystemPropertiesgetDefaultMachineFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultMachineFolder(request *ISystemPropertiessetDefaultMachineFolder) (*ISystemPropertiessetDefaultMachineFolderResponse, error) {
	response := new(ISystemPropertiessetDefaultMachineFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetLoggingLevel(request *ISystemPropertiesgetLoggingLevel) (*ISystemPropertiesgetLoggingLevelResponse, error) {
	response := new(ISystemPropertiesgetLoggingLevelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetLoggingLevel(request *ISystemPropertiessetLoggingLevel) (*ISystemPropertiessetLoggingLevelResponse, error) {
	response := new(ISystemPropertiessetLoggingLevelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMediumFormats(request *ISystemPropertiesgetMediumFormats) (*ISystemPropertiesgetMediumFormatsResponse, error) {
	response := new(ISystemPropertiesgetMediumFormatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultHardDiskFormat(request *ISystemPropertiesgetDefaultHardDiskFormat) (*ISystemPropertiesgetDefaultHardDiskFormatResponse, error) {
	response := new(ISystemPropertiesgetDefaultHardDiskFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultHardDiskFormat(request *ISystemPropertiessetDefaultHardDiskFormat) (*ISystemPropertiessetDefaultHardDiskFormatResponse, error) {
	response := new(ISystemPropertiessetDefaultHardDiskFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpaceWarning(request *ISystemPropertiesgetFreeDiskSpaceWarning) (*ISystemPropertiesgetFreeDiskSpaceWarningResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpaceWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpaceWarning(request *ISystemPropertiessetFreeDiskSpaceWarning) (*ISystemPropertiessetFreeDiskSpaceWarningResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpaceWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpacePercentWarning(request *ISystemPropertiesgetFreeDiskSpacePercentWarning) (*ISystemPropertiesgetFreeDiskSpacePercentWarningResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpacePercentWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpacePercentWarning(request *ISystemPropertiessetFreeDiskSpacePercentWarning) (*ISystemPropertiessetFreeDiskSpacePercentWarningResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpacePercentWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpaceError(request *ISystemPropertiesgetFreeDiskSpaceError) (*ISystemPropertiesgetFreeDiskSpaceErrorResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpaceErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpaceError(request *ISystemPropertiessetFreeDiskSpaceError) (*ISystemPropertiessetFreeDiskSpaceErrorResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpaceErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpacePercentError(request *ISystemPropertiesgetFreeDiskSpacePercentError) (*ISystemPropertiesgetFreeDiskSpacePercentErrorResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpacePercentErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpacePercentError(request *ISystemPropertiessetFreeDiskSpacePercentError) (*ISystemPropertiessetFreeDiskSpacePercentErrorResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpacePercentErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetVRDEAuthLibrary(request *ISystemPropertiesgetVRDEAuthLibrary) (*ISystemPropertiesgetVRDEAuthLibraryResponse, error) {
	response := new(ISystemPropertiesgetVRDEAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetVRDEAuthLibrary(request *ISystemPropertiessetVRDEAuthLibrary) (*ISystemPropertiessetVRDEAuthLibraryResponse, error) {
	response := new(ISystemPropertiessetVRDEAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetWebServiceAuthLibrary(request *ISystemPropertiesgetWebServiceAuthLibrary) (*ISystemPropertiesgetWebServiceAuthLibraryResponse, error) {
	response := new(ISystemPropertiesgetWebServiceAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetWebServiceAuthLibrary(request *ISystemPropertiessetWebServiceAuthLibrary) (*ISystemPropertiessetWebServiceAuthLibraryResponse, error) {
	response := new(ISystemPropertiessetWebServiceAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultVRDEExtPack(request *ISystemPropertiesgetDefaultVRDEExtPack) (*ISystemPropertiesgetDefaultVRDEExtPackResponse, error) {
	response := new(ISystemPropertiesgetDefaultVRDEExtPackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultVRDEExtPack(request *ISystemPropertiessetDefaultVRDEExtPack) (*ISystemPropertiessetDefaultVRDEExtPackResponse, error) {
	response := new(ISystemPropertiessetDefaultVRDEExtPackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetLogHistoryCount(request *ISystemPropertiesgetLogHistoryCount) (*ISystemPropertiesgetLogHistoryCountResponse, error) {
	response := new(ISystemPropertiesgetLogHistoryCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetLogHistoryCount(request *ISystemPropertiessetLogHistoryCount) (*ISystemPropertiessetLogHistoryCountResponse, error) {
	response := new(ISystemPropertiessetLogHistoryCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultAudioDriver(request *ISystemPropertiesgetDefaultAudioDriver) (*ISystemPropertiesgetDefaultAudioDriverResponse, error) {
	response := new(ISystemPropertiesgetDefaultAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetAutostartDatabasePath(request *ISystemPropertiesgetAutostartDatabasePath) (*ISystemPropertiesgetAutostartDatabasePathResponse, error) {
	response := new(ISystemPropertiesgetAutostartDatabasePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetAutostartDatabasePath(request *ISystemPropertiessetAutostartDatabasePath) (*ISystemPropertiessetAutostartDatabasePathResponse, error) {
	response := new(ISystemPropertiessetAutostartDatabasePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultAdditionsISO(request *ISystemPropertiesgetDefaultAdditionsISO) (*ISystemPropertiesgetDefaultAdditionsISOResponse, error) {
	response := new(ISystemPropertiesgetDefaultAdditionsISOResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultAdditionsISO(request *ISystemPropertiessetDefaultAdditionsISO) (*ISystemPropertiessetDefaultAdditionsISOResponse, error) {
	response := new(ISystemPropertiessetDefaultAdditionsISOResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultFrontend(request *ISystemPropertiesgetDefaultFrontend) (*ISystemPropertiesgetDefaultFrontendResponse, error) {
	response := new(ISystemPropertiesgetDefaultFrontendResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultFrontend(request *ISystemPropertiessetDefaultFrontend) (*ISystemPropertiessetDefaultFrontendResponse, error) {
	response := new(ISystemPropertiessetDefaultFrontendResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetScreenShotFormats(request *ISystemPropertiesgetScreenShotFormats) (*ISystemPropertiesgetScreenShotFormatsResponse, error) {
	response := new(ISystemPropertiesgetScreenShotFormatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxNetworkAdapters(request *ISystemPropertiesgetMaxNetworkAdapters) (*ISystemPropertiesgetMaxNetworkAdaptersResponse, error) {
	response := new(ISystemPropertiesgetMaxNetworkAdaptersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxNetworkAdaptersOfType(request *ISystemPropertiesgetMaxNetworkAdaptersOfType) (*ISystemPropertiesgetMaxNetworkAdaptersOfTypeResponse, error) {
	response := new(ISystemPropertiesgetMaxNetworkAdaptersOfTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxDevicesPerPortForStorageBus(request *ISystemPropertiesgetMaxDevicesPerPortForStorageBus) (*ISystemPropertiesgetMaxDevicesPerPortForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxDevicesPerPortForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinPortCountForStorageBus(request *ISystemPropertiesgetMinPortCountForStorageBus) (*ISystemPropertiesgetMinPortCountForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMinPortCountForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxPortCountForStorageBus(request *ISystemPropertiesgetMaxPortCountForStorageBus) (*ISystemPropertiesgetMaxPortCountForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxPortCountForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxInstancesOfStorageBus(request *ISystemPropertiesgetMaxInstancesOfStorageBus) (*ISystemPropertiesgetMaxInstancesOfStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxInstancesOfStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDeviceTypesForStorageBus(request *ISystemPropertiesgetDeviceTypesForStorageBus) (*ISystemPropertiesgetDeviceTypesForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetDeviceTypesForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultIoCacheSettingForStorageController(request *ISystemPropertiesgetDefaultIoCacheSettingForStorageController) (*ISystemPropertiesgetDefaultIoCacheSettingForStorageControllerResponse, error) {
	response := new(ISystemPropertiesgetDefaultIoCacheSettingForStorageControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetStorageControllerHotplugCapable(request *ISystemPropertiesgetStorageControllerHotplugCapable) (*ISystemPropertiesgetStorageControllerHotplugCapableResponse, error) {
	response := new(ISystemPropertiesgetStorageControllerHotplugCapableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxInstancesOfUSBControllerType(request *ISystemPropertiesgetMaxInstancesOfUSBControllerType) (*ISystemPropertiesgetMaxInstancesOfUSBControllerTypeResponse, error) {
	response := new(ISystemPropertiesgetMaxInstancesOfUSBControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDBasegetFormats(request *IDnDBasegetFormats) (*IDnDBasegetFormatsResponse, error) {
	response := new(IDnDBasegetFormatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDBasegetProtocolVersion(request *IDnDBasegetProtocolVersion) (*IDnDBasegetProtocolVersionResponse, error) {
	response := new(IDnDBasegetProtocolVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDBaseisFormatSupported(request *IDnDBaseisFormatSupported) (*IDnDBaseisFormatSupportedResponse, error) {
	response := new(IDnDBaseisFormatSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDBaseaddFormats(request *IDnDBaseaddFormats) (*IDnDBaseaddFormatsResponse, error) {
	response := new(IDnDBaseaddFormatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDBaseremoveFormats(request *IDnDBaseremoveFormats) (*IDnDBaseremoveFormatsResponse, error) {
	response := new(IDnDBaseremoveFormatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDSourcedragIsPending(request *IDnDSourcedragIsPending) (*IDnDSourcedragIsPendingResponse, error) {
	response := new(IDnDSourcedragIsPendingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDSourcedrop(request *IDnDSourcedrop) (*IDnDSourcedropResponse, error) {
	response := new(IDnDSourcedropResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDSourcereceiveData(request *IDnDSourcereceiveData) (*IDnDSourcereceiveDataResponse, error) {
	response := new(IDnDSourcereceiveDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestDnDSourcegetMidlDoesNotLikeEmptyInterfaces(request *IGuestDnDSourcegetMidlDoesNotLikeEmptyInterfaces) (*IGuestDnDSourcegetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IGuestDnDSourcegetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDTargetenter(request *IDnDTargetenter) (*IDnDTargetenterResponse, error) {
	response := new(IDnDTargetenterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDTargetmove(request *IDnDTargetmove) (*IDnDTargetmoveResponse, error) {
	response := new(IDnDTargetmoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDTargetleave(request *IDnDTargetleave) (*IDnDTargetleaveResponse, error) {
	response := new(IDnDTargetleaveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDTargetdrop(request *IDnDTargetdrop) (*IDnDTargetdropResponse, error) {
	response := new(IDnDTargetdropResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDTargetsendData(request *IDnDTargetsendData) (*IDnDTargetsendDataResponse, error) {
	response := new(IDnDTargetsendDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDTargetcancel(request *IDnDTargetcancel) (*IDnDTargetcancelResponse, error) {
	response := new(IDnDTargetcancelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestDnDTargetgetMidlDoesNotLikeEmptyInterfaces(request *IGuestDnDTargetgetMidlDoesNotLikeEmptyInterfaces) (*IGuestDnDTargetgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IGuestDnDTargetgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetUser(request *IGuestSessiongetUser) (*IGuestSessiongetUserResponse, error) {
	response := new(IGuestSessiongetUserResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetDomain(request *IGuestSessiongetDomain) (*IGuestSessiongetDomainResponse, error) {
	response := new(IGuestSessiongetDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetName(request *IGuestSessiongetName) (*IGuestSessiongetNameResponse, error) {
	response := new(IGuestSessiongetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetId(request *IGuestSessiongetId) (*IGuestSessiongetIdResponse, error) {
	response := new(IGuestSessiongetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetTimeout(request *IGuestSessiongetTimeout) (*IGuestSessiongetTimeoutResponse, error) {
	response := new(IGuestSessiongetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsetTimeout(request *IGuestSessionsetTimeout) (*IGuestSessionsetTimeoutResponse, error) {
	response := new(IGuestSessionsetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetProtocolVersion(request *IGuestSessiongetProtocolVersion) (*IGuestSessiongetProtocolVersionResponse, error) {
	response := new(IGuestSessiongetProtocolVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetStatus(request *IGuestSessiongetStatus) (*IGuestSessiongetStatusResponse, error) {
	response := new(IGuestSessiongetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetEnvironmentChanges(request *IGuestSessiongetEnvironmentChanges) (*IGuestSessiongetEnvironmentChangesResponse, error) {
	response := new(IGuestSessiongetEnvironmentChangesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsetEnvironmentChanges(request *IGuestSessionsetEnvironmentChanges) (*IGuestSessionsetEnvironmentChangesResponse, error) {
	response := new(IGuestSessionsetEnvironmentChangesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetEnvironmentBase(request *IGuestSessiongetEnvironmentBase) (*IGuestSessiongetEnvironmentBaseResponse, error) {
	response := new(IGuestSessiongetEnvironmentBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetProcesses(request *IGuestSessiongetProcesses) (*IGuestSessiongetProcessesResponse, error) {
	response := new(IGuestSessiongetProcessesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetPathStyle(request *IGuestSessiongetPathStyle) (*IGuestSessiongetPathStyleResponse, error) {
	response := new(IGuestSessiongetPathStyleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetCurrentDirectory(request *IGuestSessiongetCurrentDirectory) (*IGuestSessiongetCurrentDirectoryResponse, error) {
	response := new(IGuestSessiongetCurrentDirectoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsetCurrentDirectory(request *IGuestSessionsetCurrentDirectory) (*IGuestSessionsetCurrentDirectoryResponse, error) {
	response := new(IGuestSessionsetCurrentDirectoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetDirectories(request *IGuestSessiongetDirectories) (*IGuestSessiongetDirectoriesResponse, error) {
	response := new(IGuestSessiongetDirectoriesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetFiles(request *IGuestSessiongetFiles) (*IGuestSessiongetFilesResponse, error) {
	response := new(IGuestSessiongetFilesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetEventSource(request *IGuestSessiongetEventSource) (*IGuestSessiongetEventSourceResponse, error) {
	response := new(IGuestSessiongetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionclose(request *IGuestSessionclose) (*IGuestSessioncloseResponse, error) {
	response := new(IGuestSessioncloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryCopy(request *IGuestSessiondirectoryCopy) (*IGuestSessiondirectoryCopyResponse, error) {
	response := new(IGuestSessiondirectoryCopyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryCopyFromGuest(request *IGuestSessiondirectoryCopyFromGuest) (*IGuestSessiondirectoryCopyFromGuestResponse, error) {
	response := new(IGuestSessiondirectoryCopyFromGuestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryCopyToGuest(request *IGuestSessiondirectoryCopyToGuest) (*IGuestSessiondirectoryCopyToGuestResponse, error) {
	response := new(IGuestSessiondirectoryCopyToGuestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryCreate(request *IGuestSessiondirectoryCreate) (*IGuestSessiondirectoryCreateResponse, error) {
	response := new(IGuestSessiondirectoryCreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryCreateTemp(request *IGuestSessiondirectoryCreateTemp) (*IGuestSessiondirectoryCreateTempResponse, error) {
	response := new(IGuestSessiondirectoryCreateTempResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryExists(request *IGuestSessiondirectoryExists) (*IGuestSessiondirectoryExistsResponse, error) {
	response := new(IGuestSessiondirectoryExistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryOpen(request *IGuestSessiondirectoryOpen) (*IGuestSessiondirectoryOpenResponse, error) {
	response := new(IGuestSessiondirectoryOpenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryRemove(request *IGuestSessiondirectoryRemove) (*IGuestSessiondirectoryRemoveResponse, error) {
	response := new(IGuestSessiondirectoryRemoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryRemoveRecursive(request *IGuestSessiondirectoryRemoveRecursive) (*IGuestSessiondirectoryRemoveRecursiveResponse, error) {
	response := new(IGuestSessiondirectoryRemoveRecursiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionenvironmentScheduleSet(request *IGuestSessionenvironmentScheduleSet) (*IGuestSessionenvironmentScheduleSetResponse, error) {
	response := new(IGuestSessionenvironmentScheduleSetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionenvironmentScheduleUnset(request *IGuestSessionenvironmentScheduleUnset) (*IGuestSessionenvironmentScheduleUnsetResponse, error) {
	response := new(IGuestSessionenvironmentScheduleUnsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionenvironmentGetBaseVariable(request *IGuestSessionenvironmentGetBaseVariable) (*IGuestSessionenvironmentGetBaseVariableResponse, error) {
	response := new(IGuestSessionenvironmentGetBaseVariableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionenvironmentDoesBaseVariableExist(request *IGuestSessionenvironmentDoesBaseVariableExist) (*IGuestSessionenvironmentDoesBaseVariableExistResponse, error) {
	response := new(IGuestSessionenvironmentDoesBaseVariableExistResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileCopy(request *IGuestSessionfileCopy) (*IGuestSessionfileCopyResponse, error) {
	response := new(IGuestSessionfileCopyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileCopyFromGuest(request *IGuestSessionfileCopyFromGuest) (*IGuestSessionfileCopyFromGuestResponse, error) {
	response := new(IGuestSessionfileCopyFromGuestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileCopyToGuest(request *IGuestSessionfileCopyToGuest) (*IGuestSessionfileCopyToGuestResponse, error) {
	response := new(IGuestSessionfileCopyToGuestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileCreateTemp(request *IGuestSessionfileCreateTemp) (*IGuestSessionfileCreateTempResponse, error) {
	response := new(IGuestSessionfileCreateTempResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileExists(request *IGuestSessionfileExists) (*IGuestSessionfileExistsResponse, error) {
	response := new(IGuestSessionfileExistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileOpen(request *IGuestSessionfileOpen) (*IGuestSessionfileOpenResponse, error) {
	response := new(IGuestSessionfileOpenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileOpenEx(request *IGuestSessionfileOpenEx) (*IGuestSessionfileOpenExResponse, error) {
	response := new(IGuestSessionfileOpenExResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileQuerySize(request *IGuestSessionfileQuerySize) (*IGuestSessionfileQuerySizeResponse, error) {
	response := new(IGuestSessionfileQuerySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfsObjExists(request *IGuestSessionfsObjExists) (*IGuestSessionfsObjExistsResponse, error) {
	response := new(IGuestSessionfsObjExistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfsObjQueryInfo(request *IGuestSessionfsObjQueryInfo) (*IGuestSessionfsObjQueryInfoResponse, error) {
	response := new(IGuestSessionfsObjQueryInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfsObjRemove(request *IGuestSessionfsObjRemove) (*IGuestSessionfsObjRemoveResponse, error) {
	response := new(IGuestSessionfsObjRemoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfsObjRename(request *IGuestSessionfsObjRename) (*IGuestSessionfsObjRenameResponse, error) {
	response := new(IGuestSessionfsObjRenameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfsObjMove(request *IGuestSessionfsObjMove) (*IGuestSessionfsObjMoveResponse, error) {
	response := new(IGuestSessionfsObjMoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfsObjSetACL(request *IGuestSessionfsObjSetACL) (*IGuestSessionfsObjSetACLResponse, error) {
	response := new(IGuestSessionfsObjSetACLResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionprocessCreate(request *IGuestSessionprocessCreate) (*IGuestSessionprocessCreateResponse, error) {
	response := new(IGuestSessionprocessCreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionprocessCreateEx(request *IGuestSessionprocessCreateEx) (*IGuestSessionprocessCreateExResponse, error) {
	response := new(IGuestSessionprocessCreateExResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionprocessGet(request *IGuestSessionprocessGet) (*IGuestSessionprocessGetResponse, error) {
	response := new(IGuestSessionprocessGetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsymlinkCreate(request *IGuestSessionsymlinkCreate) (*IGuestSessionsymlinkCreateResponse, error) {
	response := new(IGuestSessionsymlinkCreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsymlinkExists(request *IGuestSessionsymlinkExists) (*IGuestSessionsymlinkExistsResponse, error) {
	response := new(IGuestSessionsymlinkExistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsymlinkRead(request *IGuestSessionsymlinkRead) (*IGuestSessionsymlinkReadResponse, error) {
	response := new(IGuestSessionsymlinkReadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionwaitFor(request *IGuestSessionwaitFor) (*IGuestSessionwaitForResponse, error) {
	response := new(IGuestSessionwaitForResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionwaitForArray(request *IGuestSessionwaitForArray) (*IGuestSessionwaitForArrayResponse, error) {
	response := new(IGuestSessionwaitForArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetArguments(request *IProcessgetArguments) (*IProcessgetArgumentsResponse, error) {
	response := new(IProcessgetArgumentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetEnvironment(request *IProcessgetEnvironment) (*IProcessgetEnvironmentResponse, error) {
	response := new(IProcessgetEnvironmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetEventSource(request *IProcessgetEventSource) (*IProcessgetEventSourceResponse, error) {
	response := new(IProcessgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetExecutablePath(request *IProcessgetExecutablePath) (*IProcessgetExecutablePathResponse, error) {
	response := new(IProcessgetExecutablePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetExitCode(request *IProcessgetExitCode) (*IProcessgetExitCodeResponse, error) {
	response := new(IProcessgetExitCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetName(request *IProcessgetName) (*IProcessgetNameResponse, error) {
	response := new(IProcessgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetPID(request *IProcessgetPID) (*IProcessgetPIDResponse, error) {
	response := new(IProcessgetPIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetStatus(request *IProcessgetStatus) (*IProcessgetStatusResponse, error) {
	response := new(IProcessgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcesswaitFor(request *IProcesswaitFor) (*IProcesswaitForResponse, error) {
	response := new(IProcesswaitForResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcesswaitForArray(request *IProcesswaitForArray) (*IProcesswaitForArrayResponse, error) {
	response := new(IProcesswaitForArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessread(request *IProcessread) (*IProcessreadResponse, error) {
	response := new(IProcessreadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcesswrite(request *IProcesswrite) (*IProcesswriteResponse, error) {
	response := new(IProcesswriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcesswriteArray(request *IProcesswriteArray) (*IProcesswriteArrayResponse, error) {
	response := new(IProcesswriteArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessterminate(request *IProcessterminate) (*IProcessterminateResponse, error) {
	response := new(IProcessterminateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessgetMidlDoesNotLikeEmptyInterfaces(request *IGuestProcessgetMidlDoesNotLikeEmptyInterfaces) (*IGuestProcessgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IGuestProcessgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDirectorygetDirectoryName(request *IDirectorygetDirectoryName) (*IDirectorygetDirectoryNameResponse, error) {
	response := new(IDirectorygetDirectoryNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDirectorygetFilter(request *IDirectorygetFilter) (*IDirectorygetFilterResponse, error) {
	response := new(IDirectorygetFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDirectoryclose(request *IDirectoryclose) (*IDirectorycloseResponse, error) {
	response := new(IDirectorycloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDirectoryread(request *IDirectoryread) (*IDirectoryreadResponse, error) {
	response := new(IDirectoryreadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestDirectorygetMidlDoesNotLikeEmptyInterfaces(request *IGuestDirectorygetMidlDoesNotLikeEmptyInterfaces) (*IGuestDirectorygetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IGuestDirectorygetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetEventSource(request *IFilegetEventSource) (*IFilegetEventSourceResponse, error) {
	response := new(IFilegetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetId(request *IFilegetId) (*IFilegetIdResponse, error) {
	response := new(IFilegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetInitialSize(request *IFilegetInitialSize) (*IFilegetInitialSizeResponse, error) {
	response := new(IFilegetInitialSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetOffset(request *IFilegetOffset) (*IFilegetOffsetResponse, error) {
	response := new(IFilegetOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetStatus(request *IFilegetStatus) (*IFilegetStatusResponse, error) {
	response := new(IFilegetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetFileName(request *IFilegetFileName) (*IFilegetFileNameResponse, error) {
	response := new(IFilegetFileNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetCreationMode(request *IFilegetCreationMode) (*IFilegetCreationModeResponse, error) {
	response := new(IFilegetCreationModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetOpenAction(request *IFilegetOpenAction) (*IFilegetOpenActionResponse, error) {
	response := new(IFilegetOpenActionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetAccessMode(request *IFilegetAccessMode) (*IFilegetAccessModeResponse, error) {
	response := new(IFilegetAccessModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFileclose(request *IFileclose) (*IFilecloseResponse, error) {
	response := new(IFilecloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilequeryInfo(request *IFilequeryInfo) (*IFilequeryInfoResponse, error) {
	response := new(IFilequeryInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilequerySize(request *IFilequerySize) (*IFilequerySizeResponse, error) {
	response := new(IFilequerySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFileread(request *IFileread) (*IFilereadResponse, error) {
	response := new(IFilereadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilereadAt(request *IFilereadAt) (*IFilereadAtResponse, error) {
	response := new(IFilereadAtResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFileseek(request *IFileseek) (*IFileseekResponse, error) {
	response := new(IFileseekResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilesetACL(request *IFilesetACL) (*IFilesetACLResponse, error) {
	response := new(IFilesetACLResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilesetSize(request *IFilesetSize) (*IFilesetSizeResponse, error) {
	response := new(IFilesetSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilewrite(request *IFilewrite) (*IFilewriteResponse, error) {
	response := new(IFilewriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilewriteAt(request *IFilewriteAt) (*IFilewriteAtResponse, error) {
	response := new(IFilewriteAtResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFilegetMidlDoesNotLikeEmptyInterfaces(request *IGuestFilegetMidlDoesNotLikeEmptyInterfaces) (*IGuestFilegetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IGuestFilegetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetAccessTime(request *IFsObjInfogetAccessTime) (*IFsObjInfogetAccessTimeResponse, error) {
	response := new(IFsObjInfogetAccessTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetAllocatedSize(request *IFsObjInfogetAllocatedSize) (*IFsObjInfogetAllocatedSizeResponse, error) {
	response := new(IFsObjInfogetAllocatedSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetBirthTime(request *IFsObjInfogetBirthTime) (*IFsObjInfogetBirthTimeResponse, error) {
	response := new(IFsObjInfogetBirthTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetChangeTime(request *IFsObjInfogetChangeTime) (*IFsObjInfogetChangeTimeResponse, error) {
	response := new(IFsObjInfogetChangeTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetDeviceNumber(request *IFsObjInfogetDeviceNumber) (*IFsObjInfogetDeviceNumberResponse, error) {
	response := new(IFsObjInfogetDeviceNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetFileAttributes(request *IFsObjInfogetFileAttributes) (*IFsObjInfogetFileAttributesResponse, error) {
	response := new(IFsObjInfogetFileAttributesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetGenerationId(request *IFsObjInfogetGenerationId) (*IFsObjInfogetGenerationIdResponse, error) {
	response := new(IFsObjInfogetGenerationIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetGID(request *IFsObjInfogetGID) (*IFsObjInfogetGIDResponse, error) {
	response := new(IFsObjInfogetGIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetGroupName(request *IFsObjInfogetGroupName) (*IFsObjInfogetGroupNameResponse, error) {
	response := new(IFsObjInfogetGroupNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetHardLinks(request *IFsObjInfogetHardLinks) (*IFsObjInfogetHardLinksResponse, error) {
	response := new(IFsObjInfogetHardLinksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetModificationTime(request *IFsObjInfogetModificationTime) (*IFsObjInfogetModificationTimeResponse, error) {
	response := new(IFsObjInfogetModificationTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetName(request *IFsObjInfogetName) (*IFsObjInfogetNameResponse, error) {
	response := new(IFsObjInfogetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetNodeId(request *IFsObjInfogetNodeId) (*IFsObjInfogetNodeIdResponse, error) {
	response := new(IFsObjInfogetNodeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetNodeIdDevice(request *IFsObjInfogetNodeIdDevice) (*IFsObjInfogetNodeIdDeviceResponse, error) {
	response := new(IFsObjInfogetNodeIdDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetObjectSize(request *IFsObjInfogetObjectSize) (*IFsObjInfogetObjectSizeResponse, error) {
	response := new(IFsObjInfogetObjectSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetType(request *IFsObjInfogetType) (*IFsObjInfogetTypeResponse, error) {
	response := new(IFsObjInfogetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetUID(request *IFsObjInfogetUID) (*IFsObjInfogetUIDResponse, error) {
	response := new(IFsObjInfogetUIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetUserFlags(request *IFsObjInfogetUserFlags) (*IFsObjInfogetUserFlagsResponse, error) {
	response := new(IFsObjInfogetUserFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetUserName(request *IFsObjInfogetUserName) (*IFsObjInfogetUserNameResponse, error) {
	response := new(IFsObjInfogetUserNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFsObjInfogetMidlDoesNotLikeEmptyInterfaces(request *IGuestFsObjInfogetMidlDoesNotLikeEmptyInterfaces) (*IGuestFsObjInfogetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IGuestFsObjInfogetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetOSTypeId(request *IGuestgetOSTypeId) (*IGuestgetOSTypeIdResponse, error) {
	response := new(IGuestgetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsRunLevel(request *IGuestgetAdditionsRunLevel) (*IGuestgetAdditionsRunLevelResponse, error) {
	response := new(IGuestgetAdditionsRunLevelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsVersion(request *IGuestgetAdditionsVersion) (*IGuestgetAdditionsVersionResponse, error) {
	response := new(IGuestgetAdditionsVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsRevision(request *IGuestgetAdditionsRevision) (*IGuestgetAdditionsRevisionResponse, error) {
	response := new(IGuestgetAdditionsRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetDnDSource(request *IGuestgetDnDSource) (*IGuestgetDnDSourceResponse, error) {
	response := new(IGuestgetDnDSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetDnDTarget(request *IGuestgetDnDTarget) (*IGuestgetDnDTargetResponse, error) {
	response := new(IGuestgetDnDTargetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetEventSource(request *IGuestgetEventSource) (*IGuestgetEventSourceResponse, error) {
	response := new(IGuestgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetFacilities(request *IGuestgetFacilities) (*IGuestgetFacilitiesResponse, error) {
	response := new(IGuestgetFacilitiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetSessions(request *IGuestgetSessions) (*IGuestgetSessionsResponse, error) {
	response := new(IGuestgetSessionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetMemoryBalloonSize(request *IGuestgetMemoryBalloonSize) (*IGuestgetMemoryBalloonSizeResponse, error) {
	response := new(IGuestgetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetMemoryBalloonSize(request *IGuestsetMemoryBalloonSize) (*IGuestsetMemoryBalloonSizeResponse, error) {
	response := new(IGuestsetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetStatisticsUpdateInterval(request *IGuestgetStatisticsUpdateInterval) (*IGuestgetStatisticsUpdateIntervalResponse, error) {
	response := new(IGuestgetStatisticsUpdateIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetStatisticsUpdateInterval(request *IGuestsetStatisticsUpdateInterval) (*IGuestsetStatisticsUpdateIntervalResponse, error) {
	response := new(IGuestsetStatisticsUpdateIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestinternalGetStatistics(request *IGuestinternalGetStatistics) (*IGuestinternalGetStatisticsResponse, error) {
	response := new(IGuestinternalGetStatisticsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetFacilityStatus(request *IGuestgetFacilityStatus) (*IGuestgetFacilityStatusResponse, error) {
	response := new(IGuestgetFacilityStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsStatus(request *IGuestgetAdditionsStatus) (*IGuestgetAdditionsStatusResponse, error) {
	response := new(IGuestgetAdditionsStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetCredentials(request *IGuestsetCredentials) (*IGuestsetCredentialsResponse, error) {
	response := new(IGuestsetCredentialsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestcreateSession(request *IGuestcreateSession) (*IGuestcreateSessionResponse, error) {
	response := new(IGuestcreateSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestfindSession(request *IGuestfindSession) (*IGuestfindSessionResponse, error) {
	response := new(IGuestfindSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestupdateGuestAdditions(request *IGuestupdateGuestAdditions) (*IGuestupdateGuestAdditionsResponse, error) {
	response := new(IGuestupdateGuestAdditionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetId(request *IProgressgetId) (*IProgressgetIdResponse, error) {
	response := new(IProgressgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetDescription(request *IProgressgetDescription) (*IProgressgetDescriptionResponse, error) {
	response := new(IProgressgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetInitiator(request *IProgressgetInitiator) (*IProgressgetInitiatorResponse, error) {
	response := new(IProgressgetInitiatorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCancelable(request *IProgressgetCancelable) (*IProgressgetCancelableResponse, error) {
	response := new(IProgressgetCancelableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetPercent(request *IProgressgetPercent) (*IProgressgetPercentResponse, error) {
	response := new(IProgressgetPercentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetTimeRemaining(request *IProgressgetTimeRemaining) (*IProgressgetTimeRemainingResponse, error) {
	response := new(IProgressgetTimeRemainingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCompleted(request *IProgressgetCompleted) (*IProgressgetCompletedResponse, error) {
	response := new(IProgressgetCompletedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCanceled(request *IProgressgetCanceled) (*IProgressgetCanceledResponse, error) {
	response := new(IProgressgetCanceledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetResultCode(request *IProgressgetResultCode) (*IProgressgetResultCodeResponse, error) {
	response := new(IProgressgetResultCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetErrorInfo(request *IProgressgetErrorInfo) (*IProgressgetErrorInfoResponse, error) {
	response := new(IProgressgetErrorInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationCount(request *IProgressgetOperationCount) (*IProgressgetOperationCountResponse, error) {
	response := new(IProgressgetOperationCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperation(request *IProgressgetOperation) (*IProgressgetOperationResponse, error) {
	response := new(IProgressgetOperationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationDescription(request *IProgressgetOperationDescription) (*IProgressgetOperationDescriptionResponse, error) {
	response := new(IProgressgetOperationDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationPercent(request *IProgressgetOperationPercent) (*IProgressgetOperationPercentResponse, error) {
	response := new(IProgressgetOperationPercentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationWeight(request *IProgressgetOperationWeight) (*IProgressgetOperationWeightResponse, error) {
	response := new(IProgressgetOperationWeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetTimeout(request *IProgressgetTimeout) (*IProgressgetTimeoutResponse, error) {
	response := new(IProgressgetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetTimeout(request *IProgresssetTimeout) (*IProgresssetTimeoutResponse, error) {
	response := new(IProgresssetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetCurrentOperationProgress(request *IProgresssetCurrentOperationProgress) (*IProgresssetCurrentOperationProgressResponse, error) {
	response := new(IProgresssetCurrentOperationProgressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetNextOperation(request *IProgresssetNextOperation) (*IProgresssetNextOperationResponse, error) {
	response := new(IProgresssetNextOperationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresswaitForCompletion(request *IProgresswaitForCompletion) (*IProgresswaitForCompletionResponse, error) {
	response := new(IProgresswaitForCompletionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresswaitForOperationCompletion(request *IProgresswaitForOperationCompletion) (*IProgresswaitForOperationCompletionResponse, error) {
	response := new(IProgresswaitForOperationCompletionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresswaitForAsyncProgressCompletion(request *IProgresswaitForAsyncProgressCompletion) (*IProgresswaitForAsyncProgressCompletionResponse, error) {
	response := new(IProgresswaitForAsyncProgressCompletionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresscancel(request *IProgresscancel) (*IProgresscancelResponse, error) {
	response := new(IProgresscancelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetId(request *ISnapshotgetId) (*ISnapshotgetIdResponse, error) {
	response := new(ISnapshotgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetName(request *ISnapshotgetName) (*ISnapshotgetNameResponse, error) {
	response := new(ISnapshotgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotsetName(request *ISnapshotsetName) (*ISnapshotsetNameResponse, error) {
	response := new(ISnapshotsetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetDescription(request *ISnapshotgetDescription) (*ISnapshotgetDescriptionResponse, error) {
	response := new(ISnapshotgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotsetDescription(request *ISnapshotsetDescription) (*ISnapshotsetDescriptionResponse, error) {
	response := new(ISnapshotsetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetTimeStamp(request *ISnapshotgetTimeStamp) (*ISnapshotgetTimeStampResponse, error) {
	response := new(ISnapshotgetTimeStampResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetOnline(request *ISnapshotgetOnline) (*ISnapshotgetOnlineResponse, error) {
	response := new(ISnapshotgetOnlineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetMachine(request *ISnapshotgetMachine) (*ISnapshotgetMachineResponse, error) {
	response := new(ISnapshotgetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetParent(request *ISnapshotgetParent) (*ISnapshotgetParentResponse, error) {
	response := new(ISnapshotgetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetChildren(request *ISnapshotgetChildren) (*ISnapshotgetChildrenResponse, error) {
	response := new(ISnapshotgetChildrenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetChildrenCount(request *ISnapshotgetChildrenCount) (*ISnapshotgetChildrenCountResponse, error) {
	response := new(ISnapshotgetChildrenCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetId(request *IMediumgetId) (*IMediumgetIdResponse, error) {
	response := new(IMediumgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetDescription(request *IMediumgetDescription) (*IMediumgetDescriptionResponse, error) {
	response := new(IMediumgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetDescription(request *IMediumsetDescription) (*IMediumsetDescriptionResponse, error) {
	response := new(IMediumsetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetState(request *IMediumgetState) (*IMediumgetStateResponse, error) {
	response := new(IMediumgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetVariant(request *IMediumgetVariant) (*IMediumgetVariantResponse, error) {
	response := new(IMediumgetVariantResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLocation(request *IMediumgetLocation) (*IMediumgetLocationResponse, error) {
	response := new(IMediumgetLocationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetName(request *IMediumgetName) (*IMediumgetNameResponse, error) {
	response := new(IMediumgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetDeviceType(request *IMediumgetDeviceType) (*IMediumgetDeviceTypeResponse, error) {
	response := new(IMediumgetDeviceTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetHostDrive(request *IMediumgetHostDrive) (*IMediumgetHostDriveResponse, error) {
	response := new(IMediumgetHostDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetSize(request *IMediumgetSize) (*IMediumgetSizeResponse, error) {
	response := new(IMediumgetSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetFormat(request *IMediumgetFormat) (*IMediumgetFormatResponse, error) {
	response := new(IMediumgetFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetMediumFormat(request *IMediumgetMediumFormat) (*IMediumgetMediumFormatResponse, error) {
	response := new(IMediumgetMediumFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetType(request *IMediumgetType) (*IMediumgetTypeResponse, error) {
	response := new(IMediumgetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetType(request *IMediumsetType) (*IMediumsetTypeResponse, error) {
	response := new(IMediumsetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetAllowedTypes(request *IMediumgetAllowedTypes) (*IMediumgetAllowedTypesResponse, error) {
	response := new(IMediumgetAllowedTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetParent(request *IMediumgetParent) (*IMediumgetParentResponse, error) {
	response := new(IMediumgetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetChildren(request *IMediumgetChildren) (*IMediumgetChildrenResponse, error) {
	response := new(IMediumgetChildrenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetBase(request *IMediumgetBase) (*IMediumgetBaseResponse, error) {
	response := new(IMediumgetBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetReadOnly(request *IMediumgetReadOnly) (*IMediumgetReadOnlyResponse, error) {
	response := new(IMediumgetReadOnlyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLogicalSize(request *IMediumgetLogicalSize) (*IMediumgetLogicalSizeResponse, error) {
	response := new(IMediumgetLogicalSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetAutoReset(request *IMediumgetAutoReset) (*IMediumgetAutoResetResponse, error) {
	response := new(IMediumgetAutoResetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetAutoReset(request *IMediumsetAutoReset) (*IMediumsetAutoResetResponse, error) {
	response := new(IMediumsetAutoResetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLastAccessError(request *IMediumgetLastAccessError) (*IMediumgetLastAccessErrorResponse, error) {
	response := new(IMediumgetLastAccessErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetMachineIds(request *IMediumgetMachineIds) (*IMediumgetMachineIdsResponse, error) {
	response := new(IMediumgetMachineIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetIds(request *IMediumsetIds) (*IMediumsetIdsResponse, error) {
	response := new(IMediumsetIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumrefreshState(request *IMediumrefreshState) (*IMediumrefreshStateResponse, error) {
	response := new(IMediumrefreshStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetSnapshotIds(request *IMediumgetSnapshotIds) (*IMediumgetSnapshotIdsResponse, error) {
	response := new(IMediumgetSnapshotIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumlockRead(request *IMediumlockRead) (*IMediumlockReadResponse, error) {
	response := new(IMediumlockReadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumlockWrite(request *IMediumlockWrite) (*IMediumlockWriteResponse, error) {
	response := new(IMediumlockWriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumclose(request *IMediumclose) (*IMediumcloseResponse, error) {
	response := new(IMediumcloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetProperty(request *IMediumgetProperty) (*IMediumgetPropertyResponse, error) {
	response := new(IMediumgetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetProperty(request *IMediumsetProperty) (*IMediumsetPropertyResponse, error) {
	response := new(IMediumsetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetProperties(request *IMediumgetProperties) (*IMediumgetPropertiesResponse, error) {
	response := new(IMediumgetPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetProperties(request *IMediumsetProperties) (*IMediumsetPropertiesResponse, error) {
	response := new(IMediumsetPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcreateBaseStorage(request *IMediumcreateBaseStorage) (*IMediumcreateBaseStorageResponse, error) {
	response := new(IMediumcreateBaseStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumdeleteStorage(request *IMediumdeleteStorage) (*IMediumdeleteStorageResponse, error) {
	response := new(IMediumdeleteStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcreateDiffStorage(request *IMediumcreateDiffStorage) (*IMediumcreateDiffStorageResponse, error) {
	response := new(IMediumcreateDiffStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediummergeTo(request *IMediummergeTo) (*IMediummergeToResponse, error) {
	response := new(IMediummergeToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcloneTo(request *IMediumcloneTo) (*IMediumcloneToResponse, error) {
	response := new(IMediumcloneToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcloneToBase(request *IMediumcloneToBase) (*IMediumcloneToBaseResponse, error) {
	response := new(IMediumcloneToBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetLocation(request *IMediumsetLocation) (*IMediumsetLocationResponse, error) {
	response := new(IMediumsetLocationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcompact(request *IMediumcompact) (*IMediumcompactResponse, error) {
	response := new(IMediumcompactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumresize(request *IMediumresize) (*IMediumresizeResponse, error) {
	response := new(IMediumresizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumreset(request *IMediumreset) (*IMediumresetResponse, error) {
	response := new(IMediumresetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumchangeEncryption(request *IMediumchangeEncryption) (*IMediumchangeEncryptionResponse, error) {
	response := new(IMediumchangeEncryptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetEncryptionSettings(request *IMediumgetEncryptionSettings) (*IMediumgetEncryptionSettingsResponse, error) {
	response := new(IMediumgetEncryptionSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcheckEncryptionPassword(request *IMediumcheckEncryptionPassword) (*IMediumcheckEncryptionPasswordResponse, error) {
	response := new(IMediumcheckEncryptionPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetId(request *IMediumFormatgetId) (*IMediumFormatgetIdResponse, error) {
	response := new(IMediumFormatgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetName(request *IMediumFormatgetName) (*IMediumFormatgetNameResponse, error) {
	response := new(IMediumFormatgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetCapabilities(request *IMediumFormatgetCapabilities) (*IMediumFormatgetCapabilitiesResponse, error) {
	response := new(IMediumFormatgetCapabilitiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatdescribeFileExtensions(request *IMediumFormatdescribeFileExtensions) (*IMediumFormatdescribeFileExtensionsResponse, error) {
	response := new(IMediumFormatdescribeFileExtensionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatdescribeProperties(request *IMediumFormatdescribeProperties) (*IMediumFormatdescribePropertiesResponse, error) {
	response := new(IMediumFormatdescribePropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ITokenabandon(request *ITokenabandon) (*ITokenabandonResponse, error) {
	response := new(ITokenabandonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ITokendummy(request *ITokendummy) (*ITokendummyResponse, error) {
	response := new(ITokendummyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardgetKeyboardLEDs(request *IKeyboardgetKeyboardLEDs) (*IKeyboardgetKeyboardLEDsResponse, error) {
	response := new(IKeyboardgetKeyboardLEDsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardgetEventSource(request *IKeyboardgetEventSource) (*IKeyboardgetEventSourceResponse, error) {
	response := new(IKeyboardgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputScancode(request *IKeyboardputScancode) (*IKeyboardputScancodeResponse, error) {
	response := new(IKeyboardputScancodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputScancodes(request *IKeyboardputScancodes) (*IKeyboardputScancodesResponse, error) {
	response := new(IKeyboardputScancodesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputCAD(request *IKeyboardputCAD) (*IKeyboardputCADResponse, error) {
	response := new(IKeyboardputCADResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardreleaseKeys(request *IKeyboardreleaseKeys) (*IKeyboardreleaseKeysResponse, error) {
	response := new(IKeyboardreleaseKeysResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapegetVisible(request *IMousePointerShapegetVisible) (*IMousePointerShapegetVisibleResponse, error) {
	response := new(IMousePointerShapegetVisibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapegetAlpha(request *IMousePointerShapegetAlpha) (*IMousePointerShapegetAlphaResponse, error) {
	response := new(IMousePointerShapegetAlphaResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapegetHotX(request *IMousePointerShapegetHotX) (*IMousePointerShapegetHotXResponse, error) {
	response := new(IMousePointerShapegetHotXResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapegetHotY(request *IMousePointerShapegetHotY) (*IMousePointerShapegetHotYResponse, error) {
	response := new(IMousePointerShapegetHotYResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapegetWidth(request *IMousePointerShapegetWidth) (*IMousePointerShapegetWidthResponse, error) {
	response := new(IMousePointerShapegetWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapegetHeight(request *IMousePointerShapegetHeight) (*IMousePointerShapegetHeightResponse, error) {
	response := new(IMousePointerShapegetHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapegetShape(request *IMousePointerShapegetShape) (*IMousePointerShapegetShapeResponse, error) {
	response := new(IMousePointerShapegetShapeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetAbsoluteSupported(request *IMousegetAbsoluteSupported) (*IMousegetAbsoluteSupportedResponse, error) {
	response := new(IMousegetAbsoluteSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetRelativeSupported(request *IMousegetRelativeSupported) (*IMousegetRelativeSupportedResponse, error) {
	response := new(IMousegetRelativeSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetMultiTouchSupported(request *IMousegetMultiTouchSupported) (*IMousegetMultiTouchSupportedResponse, error) {
	response := new(IMousegetMultiTouchSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetNeedsHostCursor(request *IMousegetNeedsHostCursor) (*IMousegetNeedsHostCursorResponse, error) {
	response := new(IMousegetNeedsHostCursorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetPointerShape(request *IMousegetPointerShape) (*IMousegetPointerShapeResponse, error) {
	response := new(IMousegetPointerShapeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetEventSource(request *IMousegetEventSource) (*IMousegetEventSourceResponse, error) {
	response := new(IMousegetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputMouseEvent(request *IMouseputMouseEvent) (*IMouseputMouseEventResponse, error) {
	response := new(IMouseputMouseEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputMouseEventAbsolute(request *IMouseputMouseEventAbsolute) (*IMouseputMouseEventAbsoluteResponse, error) {
	response := new(IMouseputMouseEventAbsoluteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputEventMultiTouch(request *IMouseputEventMultiTouch) (*IMouseputEventMultiTouchResponse, error) {
	response := new(IMouseputEventMultiTouchResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputEventMultiTouchString(request *IMouseputEventMultiTouchString) (*IMouseputEventMultiTouchStringResponse, error) {
	response := new(IMouseputEventMultiTouchStringResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetWidth(request *IFramebuffergetWidth) (*IFramebuffergetWidthResponse, error) {
	response := new(IFramebuffergetWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetHeight(request *IFramebuffergetHeight) (*IFramebuffergetHeightResponse, error) {
	response := new(IFramebuffergetHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetBitsPerPixel(request *IFramebuffergetBitsPerPixel) (*IFramebuffergetBitsPerPixelResponse, error) {
	response := new(IFramebuffergetBitsPerPixelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetBytesPerLine(request *IFramebuffergetBytesPerLine) (*IFramebuffergetBytesPerLineResponse, error) {
	response := new(IFramebuffergetBytesPerLineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetPixelFormat(request *IFramebuffergetPixelFormat) (*IFramebuffergetPixelFormatResponse, error) {
	response := new(IFramebuffergetPixelFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetHeightReduction(request *IFramebuffergetHeightReduction) (*IFramebuffergetHeightReductionResponse, error) {
	response := new(IFramebuffergetHeightReductionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetOverlay(request *IFramebuffergetOverlay) (*IFramebuffergetOverlayResponse, error) {
	response := new(IFramebuffergetOverlayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetCapabilities(request *IFramebuffergetCapabilities) (*IFramebuffergetCapabilitiesResponse, error) {
	response := new(IFramebuffergetCapabilitiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffernotifyUpdate(request *IFramebuffernotifyUpdate) (*IFramebuffernotifyUpdateResponse, error) {
	response := new(IFramebuffernotifyUpdateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffernotifyUpdateImage(request *IFramebuffernotifyUpdateImage) (*IFramebuffernotifyUpdateImageResponse, error) {
	response := new(IFramebuffernotifyUpdateImageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffernotifyChange(request *IFramebuffernotifyChange) (*IFramebuffernotifyChangeResponse, error) {
	response := new(IFramebuffernotifyChangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffervideoModeSupported(request *IFramebuffervideoModeSupported) (*IFramebuffervideoModeSupportedResponse, error) {
	response := new(IFramebuffervideoModeSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffernotify3DEvent(request *IFramebuffernotify3DEvent) (*IFramebuffernotify3DEventResponse, error) {
	response := new(IFramebuffernotify3DEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaygetX(request *IFramebufferOverlaygetX) (*IFramebufferOverlaygetXResponse, error) {
	response := new(IFramebufferOverlaygetXResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaygetY(request *IFramebufferOverlaygetY) (*IFramebufferOverlaygetYResponse, error) {
	response := new(IFramebufferOverlaygetYResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaygetVisible(request *IFramebufferOverlaygetVisible) (*IFramebufferOverlaygetVisibleResponse, error) {
	response := new(IFramebufferOverlaygetVisibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaysetVisible(request *IFramebufferOverlaysetVisible) (*IFramebufferOverlaysetVisibleResponse, error) {
	response := new(IFramebufferOverlaysetVisibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaygetAlpha(request *IFramebufferOverlaygetAlpha) (*IFramebufferOverlaygetAlphaResponse, error) {
	response := new(IFramebufferOverlaygetAlphaResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaysetAlpha(request *IFramebufferOverlaysetAlpha) (*IFramebufferOverlaysetAlphaResponse, error) {
	response := new(IFramebufferOverlaysetAlphaResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaymove(request *IFramebufferOverlaymove) (*IFramebufferOverlaymoveResponse, error) {
	response := new(IFramebufferOverlaymoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaygetScreenResolution(request *IDisplaygetScreenResolution) (*IDisplaygetScreenResolutionResponse, error) {
	response := new(IDisplaygetScreenResolutionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayattachFramebuffer(request *IDisplayattachFramebuffer) (*IDisplayattachFramebufferResponse, error) {
	response := new(IDisplayattachFramebufferResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaydetachFramebuffer(request *IDisplaydetachFramebuffer) (*IDisplaydetachFramebufferResponse, error) {
	response := new(IDisplaydetachFramebufferResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayqueryFramebuffer(request *IDisplayqueryFramebuffer) (*IDisplayqueryFramebufferResponse, error) {
	response := new(IDisplayqueryFramebufferResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaysetVideoModeHint(request *IDisplaysetVideoModeHint) (*IDisplaysetVideoModeHintResponse, error) {
	response := new(IDisplaysetVideoModeHintResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaysetSeamlessMode(request *IDisplaysetSeamlessMode) (*IDisplaysetSeamlessModeResponse, error) {
	response := new(IDisplaysetSeamlessModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaytakeScreenShotToArray(request *IDisplaytakeScreenShotToArray) (*IDisplaytakeScreenShotToArrayResponse, error) {
	response := new(IDisplaytakeScreenShotToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayinvalidateAndUpdate(request *IDisplayinvalidateAndUpdate) (*IDisplayinvalidateAndUpdateResponse, error) {
	response := new(IDisplayinvalidateAndUpdateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayinvalidateAndUpdateScreen(request *IDisplayinvalidateAndUpdateScreen) (*IDisplayinvalidateAndUpdateScreenResponse, error) {
	response := new(IDisplayinvalidateAndUpdateScreenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayviewportChanged(request *IDisplayviewportChanged) (*IDisplayviewportChangedResponse, error) {
	response := new(IDisplayviewportChangedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaynotifyScaleFactorChange(request *IDisplaynotifyScaleFactorChange) (*IDisplaynotifyScaleFactorChangeResponse, error) {
	response := new(IDisplaynotifyScaleFactorChangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaynotifyHiDPIOutputPolicyChange(request *IDisplaynotifyHiDPIOutputPolicyChange) (*IDisplaynotifyHiDPIOutputPolicyChangeResponse, error) {
	response := new(IDisplaynotifyHiDPIOutputPolicyChangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetAdapterType(request *INetworkAdaptergetAdapterType) (*INetworkAdaptergetAdapterTypeResponse, error) {
	response := new(INetworkAdaptergetAdapterTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetAdapterType(request *INetworkAdaptersetAdapterType) (*INetworkAdaptersetAdapterTypeResponse, error) {
	response := new(INetworkAdaptersetAdapterTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetSlot(request *INetworkAdaptergetSlot) (*INetworkAdaptergetSlotResponse, error) {
	response := new(INetworkAdaptergetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetEnabled(request *INetworkAdaptergetEnabled) (*INetworkAdaptergetEnabledResponse, error) {
	response := new(INetworkAdaptergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetEnabled(request *INetworkAdaptersetEnabled) (*INetworkAdaptersetEnabledResponse, error) {
	response := new(INetworkAdaptersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetMACAddress(request *INetworkAdaptergetMACAddress) (*INetworkAdaptergetMACAddressResponse, error) {
	response := new(INetworkAdaptergetMACAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetMACAddress(request *INetworkAdaptersetMACAddress) (*INetworkAdaptersetMACAddressResponse, error) {
	response := new(INetworkAdaptersetMACAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetAttachmentType(request *INetworkAdaptergetAttachmentType) (*INetworkAdaptergetAttachmentTypeResponse, error) {
	response := new(INetworkAdaptergetAttachmentTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetAttachmentType(request *INetworkAdaptersetAttachmentType) (*INetworkAdaptersetAttachmentTypeResponse, error) {
	response := new(INetworkAdaptersetAttachmentTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetBridgedInterface(request *INetworkAdaptergetBridgedInterface) (*INetworkAdaptergetBridgedInterfaceResponse, error) {
	response := new(INetworkAdaptergetBridgedInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetBridgedInterface(request *INetworkAdaptersetBridgedInterface) (*INetworkAdaptersetBridgedInterfaceResponse, error) {
	response := new(INetworkAdaptersetBridgedInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetHostOnlyInterface(request *INetworkAdaptergetHostOnlyInterface) (*INetworkAdaptergetHostOnlyInterfaceResponse, error) {
	response := new(INetworkAdaptergetHostOnlyInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetHostOnlyInterface(request *INetworkAdaptersetHostOnlyInterface) (*INetworkAdaptersetHostOnlyInterfaceResponse, error) {
	response := new(INetworkAdaptersetHostOnlyInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetInternalNetwork(request *INetworkAdaptergetInternalNetwork) (*INetworkAdaptergetInternalNetworkResponse, error) {
	response := new(INetworkAdaptergetInternalNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetInternalNetwork(request *INetworkAdaptersetInternalNetwork) (*INetworkAdaptersetInternalNetworkResponse, error) {
	response := new(INetworkAdaptersetInternalNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetNATNetwork(request *INetworkAdaptergetNATNetwork) (*INetworkAdaptergetNATNetworkResponse, error) {
	response := new(INetworkAdaptergetNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetNATNetwork(request *INetworkAdaptersetNATNetwork) (*INetworkAdaptersetNATNetworkResponse, error) {
	response := new(INetworkAdaptersetNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetGenericDriver(request *INetworkAdaptergetGenericDriver) (*INetworkAdaptergetGenericDriverResponse, error) {
	response := new(INetworkAdaptergetGenericDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetGenericDriver(request *INetworkAdaptersetGenericDriver) (*INetworkAdaptersetGenericDriverResponse, error) {
	response := new(INetworkAdaptersetGenericDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetCableConnected(request *INetworkAdaptergetCableConnected) (*INetworkAdaptergetCableConnectedResponse, error) {
	response := new(INetworkAdaptergetCableConnectedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetCableConnected(request *INetworkAdaptersetCableConnected) (*INetworkAdaptersetCableConnectedResponse, error) {
	response := new(INetworkAdaptersetCableConnectedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetLineSpeed(request *INetworkAdaptergetLineSpeed) (*INetworkAdaptergetLineSpeedResponse, error) {
	response := new(INetworkAdaptergetLineSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetLineSpeed(request *INetworkAdaptersetLineSpeed) (*INetworkAdaptersetLineSpeedResponse, error) {
	response := new(INetworkAdaptersetLineSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetPromiscModePolicy(request *INetworkAdaptergetPromiscModePolicy) (*INetworkAdaptergetPromiscModePolicyResponse, error) {
	response := new(INetworkAdaptergetPromiscModePolicyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetPromiscModePolicy(request *INetworkAdaptersetPromiscModePolicy) (*INetworkAdaptersetPromiscModePolicyResponse, error) {
	response := new(INetworkAdaptersetPromiscModePolicyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetTraceEnabled(request *INetworkAdaptergetTraceEnabled) (*INetworkAdaptergetTraceEnabledResponse, error) {
	response := new(INetworkAdaptergetTraceEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetTraceEnabled(request *INetworkAdaptersetTraceEnabled) (*INetworkAdaptersetTraceEnabledResponse, error) {
	response := new(INetworkAdaptersetTraceEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetTraceFile(request *INetworkAdaptergetTraceFile) (*INetworkAdaptergetTraceFileResponse, error) {
	response := new(INetworkAdaptergetTraceFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetTraceFile(request *INetworkAdaptersetTraceFile) (*INetworkAdaptersetTraceFileResponse, error) {
	response := new(INetworkAdaptersetTraceFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetNATEngine(request *INetworkAdaptergetNATEngine) (*INetworkAdaptergetNATEngineResponse, error) {
	response := new(INetworkAdaptergetNATEngineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetBootPriority(request *INetworkAdaptergetBootPriority) (*INetworkAdaptergetBootPriorityResponse, error) {
	response := new(INetworkAdaptergetBootPriorityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetBootPriority(request *INetworkAdaptersetBootPriority) (*INetworkAdaptersetBootPriorityResponse, error) {
	response := new(INetworkAdaptersetBootPriorityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetBandwidthGroup(request *INetworkAdaptergetBandwidthGroup) (*INetworkAdaptergetBandwidthGroupResponse, error) {
	response := new(INetworkAdaptergetBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetBandwidthGroup(request *INetworkAdaptersetBandwidthGroup) (*INetworkAdaptersetBandwidthGroupResponse, error) {
	response := new(INetworkAdaptersetBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetProperty(request *INetworkAdaptergetProperty) (*INetworkAdaptergetPropertyResponse, error) {
	response := new(INetworkAdaptergetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetProperty(request *INetworkAdaptersetProperty) (*INetworkAdaptersetPropertyResponse, error) {
	response := new(INetworkAdaptersetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetProperties(request *INetworkAdaptergetProperties) (*INetworkAdaptergetPropertiesResponse, error) {
	response := new(INetworkAdaptergetPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetSlot(request *ISerialPortgetSlot) (*ISerialPortgetSlotResponse, error) {
	response := new(ISerialPortgetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetEnabled(request *ISerialPortgetEnabled) (*ISerialPortgetEnabledResponse, error) {
	response := new(ISerialPortgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetEnabled(request *ISerialPortsetEnabled) (*ISerialPortsetEnabledResponse, error) {
	response := new(ISerialPortsetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetIOBase(request *ISerialPortgetIOBase) (*ISerialPortgetIOBaseResponse, error) {
	response := new(ISerialPortgetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetIOBase(request *ISerialPortsetIOBase) (*ISerialPortsetIOBaseResponse, error) {
	response := new(ISerialPortsetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetIRQ(request *ISerialPortgetIRQ) (*ISerialPortgetIRQResponse, error) {
	response := new(ISerialPortgetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetIRQ(request *ISerialPortsetIRQ) (*ISerialPortsetIRQResponse, error) {
	response := new(ISerialPortsetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetHostMode(request *ISerialPortgetHostMode) (*ISerialPortgetHostModeResponse, error) {
	response := new(ISerialPortgetHostModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetHostMode(request *ISerialPortsetHostMode) (*ISerialPortsetHostModeResponse, error) {
	response := new(ISerialPortsetHostModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetServer(request *ISerialPortgetServer) (*ISerialPortgetServerResponse, error) {
	response := new(ISerialPortgetServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetServer(request *ISerialPortsetServer) (*ISerialPortsetServerResponse, error) {
	response := new(ISerialPortsetServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetPath(request *ISerialPortgetPath) (*ISerialPortgetPathResponse, error) {
	response := new(ISerialPortgetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetPath(request *ISerialPortsetPath) (*ISerialPortsetPathResponse, error) {
	response := new(ISerialPortsetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetSlot(request *IParallelPortgetSlot) (*IParallelPortgetSlotResponse, error) {
	response := new(IParallelPortgetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetEnabled(request *IParallelPortgetEnabled) (*IParallelPortgetEnabledResponse, error) {
	response := new(IParallelPortgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetEnabled(request *IParallelPortsetEnabled) (*IParallelPortsetEnabledResponse, error) {
	response := new(IParallelPortsetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetIOBase(request *IParallelPortgetIOBase) (*IParallelPortgetIOBaseResponse, error) {
	response := new(IParallelPortgetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetIOBase(request *IParallelPortsetIOBase) (*IParallelPortsetIOBaseResponse, error) {
	response := new(IParallelPortsetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetIRQ(request *IParallelPortgetIRQ) (*IParallelPortgetIRQResponse, error) {
	response := new(IParallelPortgetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetIRQ(request *IParallelPortsetIRQ) (*IParallelPortsetIRQResponse, error) {
	response := new(IParallelPortsetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetPath(request *IParallelPortgetPath) (*IParallelPortgetPathResponse, error) {
	response := new(IParallelPortgetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetPath(request *IParallelPortsetPath) (*IParallelPortsetPathResponse, error) {
	response := new(IParallelPortsetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetSingleStep(request *IMachineDebuggergetSingleStep) (*IMachineDebuggergetSingleStepResponse, error) {
	response := new(IMachineDebuggergetSingleStepResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetSingleStep(request *IMachineDebuggersetSingleStep) (*IMachineDebuggersetSingleStepResponse, error) {
	response := new(IMachineDebuggersetSingleStepResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetRecompileUser(request *IMachineDebuggergetRecompileUser) (*IMachineDebuggergetRecompileUserResponse, error) {
	response := new(IMachineDebuggergetRecompileUserResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetRecompileUser(request *IMachineDebuggersetRecompileUser) (*IMachineDebuggersetRecompileUserResponse, error) {
	response := new(IMachineDebuggersetRecompileUserResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetRecompileSupervisor(request *IMachineDebuggergetRecompileSupervisor) (*IMachineDebuggergetRecompileSupervisorResponse, error) {
	response := new(IMachineDebuggergetRecompileSupervisorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetRecompileSupervisor(request *IMachineDebuggersetRecompileSupervisor) (*IMachineDebuggersetRecompileSupervisorResponse, error) {
	response := new(IMachineDebuggersetRecompileSupervisorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetExecuteAllInIEM(request *IMachineDebuggergetExecuteAllInIEM) (*IMachineDebuggergetExecuteAllInIEMResponse, error) {
	response := new(IMachineDebuggergetExecuteAllInIEMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetExecuteAllInIEM(request *IMachineDebuggersetExecuteAllInIEM) (*IMachineDebuggersetExecuteAllInIEMResponse, error) {
	response := new(IMachineDebuggersetExecuteAllInIEMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetPATMEnabled(request *IMachineDebuggergetPATMEnabled) (*IMachineDebuggergetPATMEnabledResponse, error) {
	response := new(IMachineDebuggergetPATMEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetPATMEnabled(request *IMachineDebuggersetPATMEnabled) (*IMachineDebuggersetPATMEnabledResponse, error) {
	response := new(IMachineDebuggersetPATMEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetCSAMEnabled(request *IMachineDebuggergetCSAMEnabled) (*IMachineDebuggergetCSAMEnabledResponse, error) {
	response := new(IMachineDebuggergetCSAMEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetCSAMEnabled(request *IMachineDebuggersetCSAMEnabled) (*IMachineDebuggersetCSAMEnabledResponse, error) {
	response := new(IMachineDebuggersetCSAMEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogEnabled(request *IMachineDebuggergetLogEnabled) (*IMachineDebuggergetLogEnabledResponse, error) {
	response := new(IMachineDebuggergetLogEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetLogEnabled(request *IMachineDebuggersetLogEnabled) (*IMachineDebuggersetLogEnabledResponse, error) {
	response := new(IMachineDebuggersetLogEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogDbgFlags(request *IMachineDebuggergetLogDbgFlags) (*IMachineDebuggergetLogDbgFlagsResponse, error) {
	response := new(IMachineDebuggergetLogDbgFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogDbgGroups(request *IMachineDebuggergetLogDbgGroups) (*IMachineDebuggergetLogDbgGroupsResponse, error) {
	response := new(IMachineDebuggergetLogDbgGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogDbgDestinations(request *IMachineDebuggergetLogDbgDestinations) (*IMachineDebuggergetLogDbgDestinationsResponse, error) {
	response := new(IMachineDebuggergetLogDbgDestinationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogRelFlags(request *IMachineDebuggergetLogRelFlags) (*IMachineDebuggergetLogRelFlagsResponse, error) {
	response := new(IMachineDebuggergetLogRelFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogRelGroups(request *IMachineDebuggergetLogRelGroups) (*IMachineDebuggergetLogRelGroupsResponse, error) {
	response := new(IMachineDebuggergetLogRelGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogRelDestinations(request *IMachineDebuggergetLogRelDestinations) (*IMachineDebuggergetLogRelDestinationsResponse, error) {
	response := new(IMachineDebuggergetLogRelDestinationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetHWVirtExEnabled(request *IMachineDebuggergetHWVirtExEnabled) (*IMachineDebuggergetHWVirtExEnabledResponse, error) {
	response := new(IMachineDebuggergetHWVirtExEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetHWVirtExNestedPagingEnabled(request *IMachineDebuggergetHWVirtExNestedPagingEnabled) (*IMachineDebuggergetHWVirtExNestedPagingEnabledResponse, error) {
	response := new(IMachineDebuggergetHWVirtExNestedPagingEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetHWVirtExVPIDEnabled(request *IMachineDebuggergetHWVirtExVPIDEnabled) (*IMachineDebuggergetHWVirtExVPIDEnabledResponse, error) {
	response := new(IMachineDebuggergetHWVirtExVPIDEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetHWVirtExUXEnabled(request *IMachineDebuggergetHWVirtExUXEnabled) (*IMachineDebuggergetHWVirtExUXEnabledResponse, error) {
	response := new(IMachineDebuggergetHWVirtExUXEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetOSName(request *IMachineDebuggergetOSName) (*IMachineDebuggergetOSNameResponse, error) {
	response := new(IMachineDebuggergetOSNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetOSVersion(request *IMachineDebuggergetOSVersion) (*IMachineDebuggergetOSVersionResponse, error) {
	response := new(IMachineDebuggergetOSVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetPAEEnabled(request *IMachineDebuggergetPAEEnabled) (*IMachineDebuggergetPAEEnabledResponse, error) {
	response := new(IMachineDebuggergetPAEEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetVirtualTimeRate(request *IMachineDebuggergetVirtualTimeRate) (*IMachineDebuggergetVirtualTimeRateResponse, error) {
	response := new(IMachineDebuggergetVirtualTimeRateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetVirtualTimeRate(request *IMachineDebuggersetVirtualTimeRate) (*IMachineDebuggersetVirtualTimeRateResponse, error) {
	response := new(IMachineDebuggersetVirtualTimeRateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdumpGuestCore(request *IMachineDebuggerdumpGuestCore) (*IMachineDebuggerdumpGuestCoreResponse, error) {
	response := new(IMachineDebuggerdumpGuestCoreResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdumpHostProcessCore(request *IMachineDebuggerdumpHostProcessCore) (*IMachineDebuggerdumpHostProcessCoreResponse, error) {
	response := new(IMachineDebuggerdumpHostProcessCoreResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerinfo(request *IMachineDebuggerinfo) (*IMachineDebuggerinfoResponse, error) {
	response := new(IMachineDebuggerinfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerinjectNMI(request *IMachineDebuggerinjectNMI) (*IMachineDebuggerinjectNMIResponse, error) {
	response := new(IMachineDebuggerinjectNMIResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggermodifyLogGroups(request *IMachineDebuggermodifyLogGroups) (*IMachineDebuggermodifyLogGroupsResponse, error) {
	response := new(IMachineDebuggermodifyLogGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggermodifyLogFlags(request *IMachineDebuggermodifyLogFlags) (*IMachineDebuggermodifyLogFlagsResponse, error) {
	response := new(IMachineDebuggermodifyLogFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggermodifyLogDestinations(request *IMachineDebuggermodifyLogDestinations) (*IMachineDebuggermodifyLogDestinationsResponse, error) {
	response := new(IMachineDebuggermodifyLogDestinationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerreadPhysicalMemory(request *IMachineDebuggerreadPhysicalMemory) (*IMachineDebuggerreadPhysicalMemoryResponse, error) {
	response := new(IMachineDebuggerreadPhysicalMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerwritePhysicalMemory(request *IMachineDebuggerwritePhysicalMemory) (*IMachineDebuggerwritePhysicalMemoryResponse, error) {
	response := new(IMachineDebuggerwritePhysicalMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerreadVirtualMemory(request *IMachineDebuggerreadVirtualMemory) (*IMachineDebuggerreadVirtualMemoryResponse, error) {
	response := new(IMachineDebuggerreadVirtualMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerwriteVirtualMemory(request *IMachineDebuggerwriteVirtualMemory) (*IMachineDebuggerwriteVirtualMemoryResponse, error) {
	response := new(IMachineDebuggerwriteVirtualMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerloadPlugIn(request *IMachineDebuggerloadPlugIn) (*IMachineDebuggerloadPlugInResponse, error) {
	response := new(IMachineDebuggerloadPlugInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerunloadPlugIn(request *IMachineDebuggerunloadPlugIn) (*IMachineDebuggerunloadPlugInResponse, error) {
	response := new(IMachineDebuggerunloadPlugInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdetectOS(request *IMachineDebuggerdetectOS) (*IMachineDebuggerdetectOSResponse, error) {
	response := new(IMachineDebuggerdetectOSResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerqueryOSKernelLog(request *IMachineDebuggerqueryOSKernelLog) (*IMachineDebuggerqueryOSKernelLogResponse, error) {
	response := new(IMachineDebuggerqueryOSKernelLogResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetRegister(request *IMachineDebuggergetRegister) (*IMachineDebuggergetRegisterResponse, error) {
	response := new(IMachineDebuggergetRegisterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetRegisters(request *IMachineDebuggergetRegisters) (*IMachineDebuggergetRegistersResponse, error) {
	response := new(IMachineDebuggergetRegistersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetRegister(request *IMachineDebuggersetRegister) (*IMachineDebuggersetRegisterResponse, error) {
	response := new(IMachineDebuggersetRegisterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetRegisters(request *IMachineDebuggersetRegisters) (*IMachineDebuggersetRegistersResponse, error) {
	response := new(IMachineDebuggersetRegistersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdumpGuestStack(request *IMachineDebuggerdumpGuestStack) (*IMachineDebuggerdumpGuestStackResponse, error) {
	response := new(IMachineDebuggerdumpGuestStackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerresetStats(request *IMachineDebuggerresetStats) (*IMachineDebuggerresetStatsResponse, error) {
	response := new(IMachineDebuggerresetStatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdumpStats(request *IMachineDebuggerdumpStats) (*IMachineDebuggerdumpStatsResponse, error) {
	response := new(IMachineDebuggerdumpStatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetStats(request *IMachineDebuggergetStats) (*IMachineDebuggergetStatsResponse, error) {
	response := new(IMachineDebuggergetStatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersgetDeviceFilters(request *IUSBDeviceFiltersgetDeviceFilters) (*IUSBDeviceFiltersgetDeviceFiltersResponse, error) {
	response := new(IUSBDeviceFiltersgetDeviceFiltersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFilterscreateDeviceFilter(request *IUSBDeviceFilterscreateDeviceFilter) (*IUSBDeviceFilterscreateDeviceFilterResponse, error) {
	response := new(IUSBDeviceFilterscreateDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersinsertDeviceFilter(request *IUSBDeviceFiltersinsertDeviceFilter) (*IUSBDeviceFiltersinsertDeviceFilterResponse, error) {
	response := new(IUSBDeviceFiltersinsertDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersremoveDeviceFilter(request *IUSBDeviceFiltersremoveDeviceFilter) (*IUSBDeviceFiltersremoveDeviceFilterResponse, error) {
	response := new(IUSBDeviceFiltersremoveDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetName(request *IUSBControllergetName) (*IUSBControllergetNameResponse, error) {
	response := new(IUSBControllergetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllersetName(request *IUSBControllersetName) (*IUSBControllersetNameResponse, error) {
	response := new(IUSBControllersetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetType(request *IUSBControllergetType) (*IUSBControllergetTypeResponse, error) {
	response := new(IUSBControllergetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllersetType(request *IUSBControllersetType) (*IUSBControllersetTypeResponse, error) {
	response := new(IUSBControllersetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetUSBStandard(request *IUSBControllergetUSBStandard) (*IUSBControllergetUSBStandardResponse, error) {
	response := new(IUSBControllergetUSBStandardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetId(request *IUSBDevicegetId) (*IUSBDevicegetIdResponse, error) {
	response := new(IUSBDevicegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetVendorId(request *IUSBDevicegetVendorId) (*IUSBDevicegetVendorIdResponse, error) {
	response := new(IUSBDevicegetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetProductId(request *IUSBDevicegetProductId) (*IUSBDevicegetProductIdResponse, error) {
	response := new(IUSBDevicegetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetRevision(request *IUSBDevicegetRevision) (*IUSBDevicegetRevisionResponse, error) {
	response := new(IUSBDevicegetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetManufacturer(request *IUSBDevicegetManufacturer) (*IUSBDevicegetManufacturerResponse, error) {
	response := new(IUSBDevicegetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetProduct(request *IUSBDevicegetProduct) (*IUSBDevicegetProductResponse, error) {
	response := new(IUSBDevicegetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetSerialNumber(request *IUSBDevicegetSerialNumber) (*IUSBDevicegetSerialNumberResponse, error) {
	response := new(IUSBDevicegetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetAddress(request *IUSBDevicegetAddress) (*IUSBDevicegetAddressResponse, error) {
	response := new(IUSBDevicegetAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetPort(request *IUSBDevicegetPort) (*IUSBDevicegetPortResponse, error) {
	response := new(IUSBDevicegetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetVersion(request *IUSBDevicegetVersion) (*IUSBDevicegetVersionResponse, error) {
	response := new(IUSBDevicegetVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetPortVersion(request *IUSBDevicegetPortVersion) (*IUSBDevicegetPortVersionResponse, error) {
	response := new(IUSBDevicegetPortVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetSpeed(request *IUSBDevicegetSpeed) (*IUSBDevicegetSpeedResponse, error) {
	response := new(IUSBDevicegetSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetRemote(request *IUSBDevicegetRemote) (*IUSBDevicegetRemoteResponse, error) {
	response := new(IUSBDevicegetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetName(request *IUSBDeviceFiltergetName) (*IUSBDeviceFiltergetNameResponse, error) {
	response := new(IUSBDeviceFiltergetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetName(request *IUSBDeviceFiltersetName) (*IUSBDeviceFiltersetNameResponse, error) {
	response := new(IUSBDeviceFiltersetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetActive(request *IUSBDeviceFiltergetActive) (*IUSBDeviceFiltergetActiveResponse, error) {
	response := new(IUSBDeviceFiltergetActiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetActive(request *IUSBDeviceFiltersetActive) (*IUSBDeviceFiltersetActiveResponse, error) {
	response := new(IUSBDeviceFiltersetActiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetVendorId(request *IUSBDeviceFiltergetVendorId) (*IUSBDeviceFiltergetVendorIdResponse, error) {
	response := new(IUSBDeviceFiltergetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetVendorId(request *IUSBDeviceFiltersetVendorId) (*IUSBDeviceFiltersetVendorIdResponse, error) {
	response := new(IUSBDeviceFiltersetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetProductId(request *IUSBDeviceFiltergetProductId) (*IUSBDeviceFiltergetProductIdResponse, error) {
	response := new(IUSBDeviceFiltergetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetProductId(request *IUSBDeviceFiltersetProductId) (*IUSBDeviceFiltersetProductIdResponse, error) {
	response := new(IUSBDeviceFiltersetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetRevision(request *IUSBDeviceFiltergetRevision) (*IUSBDeviceFiltergetRevisionResponse, error) {
	response := new(IUSBDeviceFiltergetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetRevision(request *IUSBDeviceFiltersetRevision) (*IUSBDeviceFiltersetRevisionResponse, error) {
	response := new(IUSBDeviceFiltersetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetManufacturer(request *IUSBDeviceFiltergetManufacturer) (*IUSBDeviceFiltergetManufacturerResponse, error) {
	response := new(IUSBDeviceFiltergetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetManufacturer(request *IUSBDeviceFiltersetManufacturer) (*IUSBDeviceFiltersetManufacturerResponse, error) {
	response := new(IUSBDeviceFiltersetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetProduct(request *IUSBDeviceFiltergetProduct) (*IUSBDeviceFiltergetProductResponse, error) {
	response := new(IUSBDeviceFiltergetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetProduct(request *IUSBDeviceFiltersetProduct) (*IUSBDeviceFiltersetProductResponse, error) {
	response := new(IUSBDeviceFiltersetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetSerialNumber(request *IUSBDeviceFiltergetSerialNumber) (*IUSBDeviceFiltergetSerialNumberResponse, error) {
	response := new(IUSBDeviceFiltergetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetSerialNumber(request *IUSBDeviceFiltersetSerialNumber) (*IUSBDeviceFiltersetSerialNumberResponse, error) {
	response := new(IUSBDeviceFiltersetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetPort(request *IUSBDeviceFiltergetPort) (*IUSBDeviceFiltergetPortResponse, error) {
	response := new(IUSBDeviceFiltergetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetPort(request *IUSBDeviceFiltersetPort) (*IUSBDeviceFiltersetPortResponse, error) {
	response := new(IUSBDeviceFiltersetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetRemote(request *IUSBDeviceFiltergetRemote) (*IUSBDeviceFiltergetRemoteResponse, error) {
	response := new(IUSBDeviceFiltergetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetRemote(request *IUSBDeviceFiltersetRemote) (*IUSBDeviceFiltersetRemoteResponse, error) {
	response := new(IUSBDeviceFiltersetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetMaskedInterfaces(request *IUSBDeviceFiltergetMaskedInterfaces) (*IUSBDeviceFiltergetMaskedInterfacesResponse, error) {
	response := new(IUSBDeviceFiltergetMaskedInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetMaskedInterfaces(request *IUSBDeviceFiltersetMaskedInterfaces) (*IUSBDeviceFiltersetMaskedInterfacesResponse, error) {
	response := new(IUSBDeviceFiltersetMaskedInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDevicegetState(request *IHostUSBDevicegetState) (*IHostUSBDevicegetStateResponse, error) {
	response := new(IHostUSBDevicegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDeviceFiltergetAction(request *IHostUSBDeviceFiltergetAction) (*IHostUSBDeviceFiltergetActionResponse, error) {
	response := new(IHostUSBDeviceFiltergetActionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDeviceFiltersetAction(request *IHostUSBDeviceFiltersetAction) (*IHostUSBDeviceFiltersetActionResponse, error) {
	response := new(IHostUSBDeviceFiltersetActionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetEnabled(request *IAudioAdaptergetEnabled) (*IAudioAdaptergetEnabledResponse, error) {
	response := new(IAudioAdaptergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetEnabled(request *IAudioAdaptersetEnabled) (*IAudioAdaptersetEnabledResponse, error) {
	response := new(IAudioAdaptersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetEnabledIn(request *IAudioAdaptergetEnabledIn) (*IAudioAdaptergetEnabledInResponse, error) {
	response := new(IAudioAdaptergetEnabledInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetEnabledIn(request *IAudioAdaptersetEnabledIn) (*IAudioAdaptersetEnabledInResponse, error) {
	response := new(IAudioAdaptersetEnabledInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetEnabledOut(request *IAudioAdaptergetEnabledOut) (*IAudioAdaptergetEnabledOutResponse, error) {
	response := new(IAudioAdaptergetEnabledOutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetEnabledOut(request *IAudioAdaptersetEnabledOut) (*IAudioAdaptersetEnabledOutResponse, error) {
	response := new(IAudioAdaptersetEnabledOutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetAudioController(request *IAudioAdaptergetAudioController) (*IAudioAdaptergetAudioControllerResponse, error) {
	response := new(IAudioAdaptergetAudioControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetAudioController(request *IAudioAdaptersetAudioController) (*IAudioAdaptersetAudioControllerResponse, error) {
	response := new(IAudioAdaptersetAudioControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetAudioCodec(request *IAudioAdaptergetAudioCodec) (*IAudioAdaptergetAudioCodecResponse, error) {
	response := new(IAudioAdaptergetAudioCodecResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetAudioCodec(request *IAudioAdaptersetAudioCodec) (*IAudioAdaptersetAudioCodecResponse, error) {
	response := new(IAudioAdaptersetAudioCodecResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetAudioDriver(request *IAudioAdaptergetAudioDriver) (*IAudioAdaptergetAudioDriverResponse, error) {
	response := new(IAudioAdaptergetAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetAudioDriver(request *IAudioAdaptersetAudioDriver) (*IAudioAdaptersetAudioDriverResponse, error) {
	response := new(IAudioAdaptersetAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetPropertiesList(request *IAudioAdaptergetPropertiesList) (*IAudioAdaptergetPropertiesListResponse, error) {
	response := new(IAudioAdaptergetPropertiesListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetProperty(request *IAudioAdaptersetProperty) (*IAudioAdaptersetPropertyResponse, error) {
	response := new(IAudioAdaptersetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetProperty(request *IAudioAdaptergetProperty) (*IAudioAdaptergetPropertyResponse, error) {
	response := new(IAudioAdaptergetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetEnabled(request *IVRDEServergetEnabled) (*IVRDEServergetEnabledResponse, error) {
	response := new(IVRDEServergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetEnabled(request *IVRDEServersetEnabled) (*IVRDEServersetEnabledResponse, error) {
	response := new(IVRDEServersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetAuthType(request *IVRDEServergetAuthType) (*IVRDEServergetAuthTypeResponse, error) {
	response := new(IVRDEServergetAuthTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetAuthType(request *IVRDEServersetAuthType) (*IVRDEServersetAuthTypeResponse, error) {
	response := new(IVRDEServersetAuthTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetAuthTimeout(request *IVRDEServergetAuthTimeout) (*IVRDEServergetAuthTimeoutResponse, error) {
	response := new(IVRDEServergetAuthTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetAuthTimeout(request *IVRDEServersetAuthTimeout) (*IVRDEServersetAuthTimeoutResponse, error) {
	response := new(IVRDEServersetAuthTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetAllowMultiConnection(request *IVRDEServergetAllowMultiConnection) (*IVRDEServergetAllowMultiConnectionResponse, error) {
	response := new(IVRDEServergetAllowMultiConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetAllowMultiConnection(request *IVRDEServersetAllowMultiConnection) (*IVRDEServersetAllowMultiConnectionResponse, error) {
	response := new(IVRDEServersetAllowMultiConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetReuseSingleConnection(request *IVRDEServergetReuseSingleConnection) (*IVRDEServergetReuseSingleConnectionResponse, error) {
	response := new(IVRDEServergetReuseSingleConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetReuseSingleConnection(request *IVRDEServersetReuseSingleConnection) (*IVRDEServersetReuseSingleConnectionResponse, error) {
	response := new(IVRDEServersetReuseSingleConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetVRDEExtPack(request *IVRDEServergetVRDEExtPack) (*IVRDEServergetVRDEExtPackResponse, error) {
	response := new(IVRDEServergetVRDEExtPackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetVRDEExtPack(request *IVRDEServersetVRDEExtPack) (*IVRDEServersetVRDEExtPackResponse, error) {
	response := new(IVRDEServersetVRDEExtPackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetAuthLibrary(request *IVRDEServergetAuthLibrary) (*IVRDEServergetAuthLibraryResponse, error) {
	response := new(IVRDEServergetAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetAuthLibrary(request *IVRDEServersetAuthLibrary) (*IVRDEServersetAuthLibraryResponse, error) {
	response := new(IVRDEServersetAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetVRDEProperties(request *IVRDEServergetVRDEProperties) (*IVRDEServergetVRDEPropertiesResponse, error) {
	response := new(IVRDEServergetVRDEPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetVRDEProperty(request *IVRDEServersetVRDEProperty) (*IVRDEServersetVRDEPropertyResponse, error) {
	response := new(IVRDEServersetVRDEPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetVRDEProperty(request *IVRDEServergetVRDEProperty) (*IVRDEServergetVRDEPropertyResponse, error) {
	response := new(IVRDEServergetVRDEPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetState(request *ISessiongetState) (*ISessiongetStateResponse, error) {
	response := new(ISessiongetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetType(request *ISessiongetType) (*ISessiongetTypeResponse, error) {
	response := new(ISessiongetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetName(request *ISessiongetName) (*ISessiongetNameResponse, error) {
	response := new(ISessiongetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessionsetName(request *ISessionsetName) (*ISessionsetNameResponse, error) {
	response := new(ISessionsetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetMachine(request *ISessiongetMachine) (*ISessiongetMachineResponse, error) {
	response := new(ISessiongetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetConsole(request *ISessiongetConsole) (*ISessiongetConsoleResponse, error) {
	response := new(ISessiongetConsoleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessionunlockMachine(request *ISessionunlockMachine) (*ISessionunlockMachineResponse, error) {
	response := new(ISessionunlockMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetName(request *IStorageControllergetName) (*IStorageControllergetNameResponse, error) {
	response := new(IStorageControllergetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetName(request *IStorageControllersetName) (*IStorageControllersetNameResponse, error) {
	response := new(IStorageControllersetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMaxDevicesPerPortCount(request *IStorageControllergetMaxDevicesPerPortCount) (*IStorageControllergetMaxDevicesPerPortCountResponse, error) {
	response := new(IStorageControllergetMaxDevicesPerPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMinPortCount(request *IStorageControllergetMinPortCount) (*IStorageControllergetMinPortCountResponse, error) {
	response := new(IStorageControllergetMinPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMaxPortCount(request *IStorageControllergetMaxPortCount) (*IStorageControllergetMaxPortCountResponse, error) {
	response := new(IStorageControllergetMaxPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetInstance(request *IStorageControllergetInstance) (*IStorageControllergetInstanceResponse, error) {
	response := new(IStorageControllergetInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetInstance(request *IStorageControllersetInstance) (*IStorageControllersetInstanceResponse, error) {
	response := new(IStorageControllersetInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetPortCount(request *IStorageControllergetPortCount) (*IStorageControllergetPortCountResponse, error) {
	response := new(IStorageControllergetPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetPortCount(request *IStorageControllersetPortCount) (*IStorageControllersetPortCountResponse, error) {
	response := new(IStorageControllersetPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetBus(request *IStorageControllergetBus) (*IStorageControllergetBusResponse, error) {
	response := new(IStorageControllergetBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetControllerType(request *IStorageControllergetControllerType) (*IStorageControllergetControllerTypeResponse, error) {
	response := new(IStorageControllergetControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetControllerType(request *IStorageControllersetControllerType) (*IStorageControllersetControllerTypeResponse, error) {
	response := new(IStorageControllersetControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetUseHostIOCache(request *IStorageControllergetUseHostIOCache) (*IStorageControllergetUseHostIOCacheResponse, error) {
	response := new(IStorageControllergetUseHostIOCacheResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetUseHostIOCache(request *IStorageControllersetUseHostIOCache) (*IStorageControllersetUseHostIOCacheResponse, error) {
	response := new(IStorageControllersetUseHostIOCacheResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetBootable(request *IStorageControllergetBootable) (*IStorageControllergetBootableResponse, error) {
	response := new(IStorageControllergetBootableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IManagedObjectRefgetInterfaceName(request *IManagedObjectRefgetInterfaceName) (*IManagedObjectRefgetInterfaceNameResponse, error) {
	response := new(IManagedObjectRefgetInterfaceNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IManagedObjectRefrelease(request *IManagedObjectRefrelease) (*IManagedObjectRefreleaseResponse, error) {
	response := new(IManagedObjectRefreleaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagerlogon(request *IWebsessionManagerlogon) (*IWebsessionManagerlogonResponse, error) {
	response := new(IWebsessionManagerlogonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagergetSessionObject(request *IWebsessionManagergetSessionObject) (*IWebsessionManagergetSessionObjectResponse, error) {
	response := new(IWebsessionManagergetSessionObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagerlogoff(request *IWebsessionManagerlogoff) (*IWebsessionManagerlogoffResponse, error) {
	response := new(IWebsessionManagerlogoffResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMetricName(request *IPerformanceMetricgetMetricName) (*IPerformanceMetricgetMetricNameResponse, error) {
	response := new(IPerformanceMetricgetMetricNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetObject(request *IPerformanceMetricgetObject) (*IPerformanceMetricgetObjectResponse, error) {
	response := new(IPerformanceMetricgetObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetDescription(request *IPerformanceMetricgetDescription) (*IPerformanceMetricgetDescriptionResponse, error) {
	response := new(IPerformanceMetricgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetPeriod(request *IPerformanceMetricgetPeriod) (*IPerformanceMetricgetPeriodResponse, error) {
	response := new(IPerformanceMetricgetPeriodResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetCount(request *IPerformanceMetricgetCount) (*IPerformanceMetricgetCountResponse, error) {
	response := new(IPerformanceMetricgetCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetUnit(request *IPerformanceMetricgetUnit) (*IPerformanceMetricgetUnitResponse, error) {
	response := new(IPerformanceMetricgetUnitResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMinimumValue(request *IPerformanceMetricgetMinimumValue) (*IPerformanceMetricgetMinimumValueResponse, error) {
	response := new(IPerformanceMetricgetMinimumValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMaximumValue(request *IPerformanceMetricgetMaximumValue) (*IPerformanceMetricgetMaximumValueResponse, error) {
	response := new(IPerformanceMetricgetMaximumValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorgetMetricNames(request *IPerformanceCollectorgetMetricNames) (*IPerformanceCollectorgetMetricNamesResponse, error) {
	response := new(IPerformanceCollectorgetMetricNamesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorgetMetrics(request *IPerformanceCollectorgetMetrics) (*IPerformanceCollectorgetMetricsResponse, error) {
	response := new(IPerformanceCollectorgetMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorsetupMetrics(request *IPerformanceCollectorsetupMetrics) (*IPerformanceCollectorsetupMetricsResponse, error) {
	response := new(IPerformanceCollectorsetupMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorenableMetrics(request *IPerformanceCollectorenableMetrics) (*IPerformanceCollectorenableMetricsResponse, error) {
	response := new(IPerformanceCollectorenableMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectordisableMetrics(request *IPerformanceCollectordisableMetrics) (*IPerformanceCollectordisableMetricsResponse, error) {
	response := new(IPerformanceCollectordisableMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorqueryMetricsData(request *IPerformanceCollectorqueryMetricsData) (*IPerformanceCollectorqueryMetricsDataResponse, error) {
	response := new(IPerformanceCollectorqueryMetricsDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetNetwork(request *INATEnginegetNetwork) (*INATEnginegetNetworkResponse, error) {
	response := new(INATEnginegetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetNetwork(request *INATEnginesetNetwork) (*INATEnginesetNetworkResponse, error) {
	response := new(INATEnginesetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetHostIP(request *INATEnginegetHostIP) (*INATEnginegetHostIPResponse, error) {
	response := new(INATEnginegetHostIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetHostIP(request *INATEnginesetHostIP) (*INATEnginesetHostIPResponse, error) {
	response := new(INATEnginesetHostIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTFTPPrefix(request *INATEnginegetTFTPPrefix) (*INATEnginegetTFTPPrefixResponse, error) {
	response := new(INATEnginegetTFTPPrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTFTPPrefix(request *INATEnginesetTFTPPrefix) (*INATEnginesetTFTPPrefixResponse, error) {
	response := new(INATEnginesetTFTPPrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTFTPBootFile(request *INATEnginegetTFTPBootFile) (*INATEnginegetTFTPBootFileResponse, error) {
	response := new(INATEnginegetTFTPBootFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTFTPBootFile(request *INATEnginesetTFTPBootFile) (*INATEnginesetTFTPBootFileResponse, error) {
	response := new(INATEnginesetTFTPBootFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTFTPNextServer(request *INATEnginegetTFTPNextServer) (*INATEnginegetTFTPNextServerResponse, error) {
	response := new(INATEnginegetTFTPNextServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTFTPNextServer(request *INATEnginesetTFTPNextServer) (*INATEnginesetTFTPNextServerResponse, error) {
	response := new(INATEnginesetTFTPNextServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetAliasMode(request *INATEnginegetAliasMode) (*INATEnginegetAliasModeResponse, error) {
	response := new(INATEnginegetAliasModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetAliasMode(request *INATEnginesetAliasMode) (*INATEnginesetAliasModeResponse, error) {
	response := new(INATEnginesetAliasModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDNSPassDomain(request *INATEnginegetDNSPassDomain) (*INATEnginegetDNSPassDomainResponse, error) {
	response := new(INATEnginegetDNSPassDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDNSPassDomain(request *INATEnginesetDNSPassDomain) (*INATEnginesetDNSPassDomainResponse, error) {
	response := new(INATEnginesetDNSPassDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDNSProxy(request *INATEnginegetDNSProxy) (*INATEnginegetDNSProxyResponse, error) {
	response := new(INATEnginegetDNSProxyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDNSProxy(request *INATEnginesetDNSProxy) (*INATEnginesetDNSProxyResponse, error) {
	response := new(INATEnginesetDNSProxyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDNSUseHostResolver(request *INATEnginegetDNSUseHostResolver) (*INATEnginegetDNSUseHostResolverResponse, error) {
	response := new(INATEnginegetDNSUseHostResolverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDNSUseHostResolver(request *INATEnginesetDNSUseHostResolver) (*INATEnginesetDNSUseHostResolverResponse, error) {
	response := new(INATEnginesetDNSUseHostResolverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetRedirects(request *INATEnginegetRedirects) (*INATEnginegetRedirectsResponse, error) {
	response := new(INATEnginegetRedirectsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetNetworkSettings(request *INATEnginesetNetworkSettings) (*INATEnginesetNetworkSettingsResponse, error) {
	response := new(INATEnginesetNetworkSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetNetworkSettings(request *INATEnginegetNetworkSettings) (*INATEnginegetNetworkSettingsResponse, error) {
	response := new(INATEnginegetNetworkSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEngineaddRedirect(request *INATEngineaddRedirect) (*INATEngineaddRedirectResponse, error) {
	response := new(INATEngineaddRedirectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEngineremoveRedirect(request *INATEngineremoveRedirect) (*INATEngineremoveRedirectResponse, error) {
	response := new(INATEngineremoveRedirectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupgetName(request *IBandwidthGroupgetName) (*IBandwidthGroupgetNameResponse, error) {
	response := new(IBandwidthGroupgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupgetType(request *IBandwidthGroupgetType) (*IBandwidthGroupgetTypeResponse, error) {
	response := new(IBandwidthGroupgetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupgetReference(request *IBandwidthGroupgetReference) (*IBandwidthGroupgetReferenceResponse, error) {
	response := new(IBandwidthGroupgetReferenceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupgetMaxBytesPerSec(request *IBandwidthGroupgetMaxBytesPerSec) (*IBandwidthGroupgetMaxBytesPerSecResponse, error) {
	response := new(IBandwidthGroupgetMaxBytesPerSecResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupsetMaxBytesPerSec(request *IBandwidthGroupsetMaxBytesPerSec) (*IBandwidthGroupsetMaxBytesPerSecResponse, error) {
	response := new(IBandwidthGroupsetMaxBytesPerSecResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControlgetNumGroups(request *IBandwidthControlgetNumGroups) (*IBandwidthControlgetNumGroupsResponse, error) {
	response := new(IBandwidthControlgetNumGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControlcreateBandwidthGroup(request *IBandwidthControlcreateBandwidthGroup) (*IBandwidthControlcreateBandwidthGroupResponse, error) {
	response := new(IBandwidthControlcreateBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControldeleteBandwidthGroup(request *IBandwidthControldeleteBandwidthGroup) (*IBandwidthControldeleteBandwidthGroupResponse, error) {
	response := new(IBandwidthControldeleteBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControlgetBandwidthGroup(request *IBandwidthControlgetBandwidthGroup) (*IBandwidthControlgetBandwidthGroupResponse, error) {
	response := new(IBandwidthControlgetBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControlgetAllBandwidthGroups(request *IBandwidthControlgetAllBandwidthGroups) (*IBandwidthControlgetAllBandwidthGroupsResponse, error) {
	response := new(IBandwidthControlgetAllBandwidthGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourcecreateListener(request *IEventSourcecreateListener) (*IEventSourcecreateListenerResponse, error) {
	response := new(IEventSourcecreateListenerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourcecreateAggregator(request *IEventSourcecreateAggregator) (*IEventSourcecreateAggregatorResponse, error) {
	response := new(IEventSourcecreateAggregatorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceregisterListener(request *IEventSourceregisterListener) (*IEventSourceregisterListenerResponse, error) {
	response := new(IEventSourceregisterListenerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceunregisterListener(request *IEventSourceunregisterListener) (*IEventSourceunregisterListenerResponse, error) {
	response := new(IEventSourceunregisterListenerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourcefireEvent(request *IEventSourcefireEvent) (*IEventSourcefireEventResponse, error) {
	response := new(IEventSourcefireEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourcegetEvent(request *IEventSourcegetEvent) (*IEventSourcegetEventResponse, error) {
	response := new(IEventSourcegetEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceeventProcessed(request *IEventSourceeventProcessed) (*IEventSourceeventProcessedResponse, error) {
	response := new(IEventSourceeventProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventListenerhandleEvent(request *IEventListenerhandleEvent) (*IEventListenerhandleEventResponse, error) {
	response := new(IEventListenerhandleEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventgetType(request *IEventgetType) (*IEventgetTypeResponse, error) {
	response := new(IEventgetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventgetSource(request *IEventgetSource) (*IEventgetSourceResponse, error) {
	response := new(IEventgetSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventgetWaitable(request *IEventgetWaitable) (*IEventgetWaitableResponse, error) {
	response := new(IEventgetWaitableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventsetProcessed(request *IEventsetProcessed) (*IEventsetProcessedResponse, error) {
	response := new(IEventsetProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventwaitProcessed(request *IEventwaitProcessed) (*IEventwaitProcessedResponse, error) {
	response := new(IEventwaitProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IReusableEventgetGeneration(request *IReusableEventgetGeneration) (*IReusableEventgetGenerationResponse, error) {
	response := new(IReusableEventgetGenerationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IReusableEventreuse(request *IReusableEventreuse) (*IReusableEventreuseResponse, error) {
	response := new(IReusableEventreuseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineEventgetMachineId(request *IMachineEventgetMachineId) (*IMachineEventgetMachineIdResponse, error) {
	response := new(IMachineEventgetMachineIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineStateChangedEventgetState(request *IMachineStateChangedEventgetState) (*IMachineStateChangedEventgetStateResponse, error) {
	response := new(IMachineStateChangedEventgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDataChangedEventgetTemporary(request *IMachineDataChangedEventgetTemporary) (*IMachineDataChangedEventgetTemporaryResponse, error) {
	response := new(IMachineDataChangedEventgetTemporaryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumRegisteredEventgetMediumId(request *IMediumRegisteredEventgetMediumId) (*IMediumRegisteredEventgetMediumIdResponse, error) {
	response := new(IMediumRegisteredEventgetMediumIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumRegisteredEventgetMediumType(request *IMediumRegisteredEventgetMediumType) (*IMediumRegisteredEventgetMediumTypeResponse, error) {
	response := new(IMediumRegisteredEventgetMediumTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumRegisteredEventgetRegistered(request *IMediumRegisteredEventgetRegistered) (*IMediumRegisteredEventgetRegisteredResponse, error) {
	response := new(IMediumRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumConfigChangedEventgetMedium(request *IMediumConfigChangedEventgetMedium) (*IMediumConfigChangedEventgetMediumResponse, error) {
	response := new(IMediumConfigChangedEventgetMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineRegisteredEventgetRegistered(request *IMachineRegisteredEventgetRegistered) (*IMachineRegisteredEventgetRegisteredResponse, error) {
	response := new(IMachineRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessionStateChangedEventgetState(request *ISessionStateChangedEventgetState) (*ISessionStateChangedEventgetStateResponse, error) {
	response := new(ISessionStateChangedEventgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestPropertyChangedEventgetName(request *IGuestPropertyChangedEventgetName) (*IGuestPropertyChangedEventgetNameResponse, error) {
	response := new(IGuestPropertyChangedEventgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestPropertyChangedEventgetValue(request *IGuestPropertyChangedEventgetValue) (*IGuestPropertyChangedEventgetValueResponse, error) {
	response := new(IGuestPropertyChangedEventgetValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestPropertyChangedEventgetFlags(request *IGuestPropertyChangedEventgetFlags) (*IGuestPropertyChangedEventgetFlagsResponse, error) {
	response := new(IGuestPropertyChangedEventgetFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotEventgetSnapshotId(request *ISnapshotEventgetSnapshotId) (*ISnapshotEventgetSnapshotIdResponse, error) {
	response := new(ISnapshotEventgetSnapshotIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotTakenEventgetMidlDoesNotLikeEmptyInterfaces(request *ISnapshotTakenEventgetMidlDoesNotLikeEmptyInterfaces) (*ISnapshotTakenEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(ISnapshotTakenEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotDeletedEventgetMidlDoesNotLikeEmptyInterfaces(request *ISnapshotDeletedEventgetMidlDoesNotLikeEmptyInterfaces) (*ISnapshotDeletedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(ISnapshotDeletedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotRestoredEventgetMidlDoesNotLikeEmptyInterfaces(request *ISnapshotRestoredEventgetMidlDoesNotLikeEmptyInterfaces) (*ISnapshotRestoredEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(ISnapshotRestoredEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotChangedEventgetMidlDoesNotLikeEmptyInterfaces(request *ISnapshotChangedEventgetMidlDoesNotLikeEmptyInterfaces) (*ISnapshotChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(ISnapshotChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetVisible(request *IMousePointerShapeChangedEventgetVisible) (*IMousePointerShapeChangedEventgetVisibleResponse, error) {
	response := new(IMousePointerShapeChangedEventgetVisibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetAlpha(request *IMousePointerShapeChangedEventgetAlpha) (*IMousePointerShapeChangedEventgetAlphaResponse, error) {
	response := new(IMousePointerShapeChangedEventgetAlphaResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetXhot(request *IMousePointerShapeChangedEventgetXhot) (*IMousePointerShapeChangedEventgetXhotResponse, error) {
	response := new(IMousePointerShapeChangedEventgetXhotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetYhot(request *IMousePointerShapeChangedEventgetYhot) (*IMousePointerShapeChangedEventgetYhotResponse, error) {
	response := new(IMousePointerShapeChangedEventgetYhotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetWidth(request *IMousePointerShapeChangedEventgetWidth) (*IMousePointerShapeChangedEventgetWidthResponse, error) {
	response := new(IMousePointerShapeChangedEventgetWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetHeight(request *IMousePointerShapeChangedEventgetHeight) (*IMousePointerShapeChangedEventgetHeightResponse, error) {
	response := new(IMousePointerShapeChangedEventgetHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetShape(request *IMousePointerShapeChangedEventgetShape) (*IMousePointerShapeChangedEventgetShapeResponse, error) {
	response := new(IMousePointerShapeChangedEventgetShapeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseCapabilityChangedEventgetSupportsAbsolute(request *IMouseCapabilityChangedEventgetSupportsAbsolute) (*IMouseCapabilityChangedEventgetSupportsAbsoluteResponse, error) {
	response := new(IMouseCapabilityChangedEventgetSupportsAbsoluteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseCapabilityChangedEventgetSupportsRelative(request *IMouseCapabilityChangedEventgetSupportsRelative) (*IMouseCapabilityChangedEventgetSupportsRelativeResponse, error) {
	response := new(IMouseCapabilityChangedEventgetSupportsRelativeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseCapabilityChangedEventgetSupportsMultiTouch(request *IMouseCapabilityChangedEventgetSupportsMultiTouch) (*IMouseCapabilityChangedEventgetSupportsMultiTouchResponse, error) {
	response := new(IMouseCapabilityChangedEventgetSupportsMultiTouchResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseCapabilityChangedEventgetNeedsHostCursor(request *IMouseCapabilityChangedEventgetNeedsHostCursor) (*IMouseCapabilityChangedEventgetNeedsHostCursorResponse, error) {
	response := new(IMouseCapabilityChangedEventgetNeedsHostCursorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardLedsChangedEventgetNumLock(request *IKeyboardLedsChangedEventgetNumLock) (*IKeyboardLedsChangedEventgetNumLockResponse, error) {
	response := new(IKeyboardLedsChangedEventgetNumLockResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardLedsChangedEventgetCapsLock(request *IKeyboardLedsChangedEventgetCapsLock) (*IKeyboardLedsChangedEventgetCapsLockResponse, error) {
	response := new(IKeyboardLedsChangedEventgetCapsLockResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardLedsChangedEventgetScrollLock(request *IKeyboardLedsChangedEventgetScrollLock) (*IKeyboardLedsChangedEventgetScrollLockResponse, error) {
	response := new(IKeyboardLedsChangedEventgetScrollLockResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStateChangedEventgetState(request *IStateChangedEventgetState) (*IStateChangedEventgetStateResponse, error) {
	response := new(IStateChangedEventgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAdditionsStateChangedEventgetMidlDoesNotLikeEmptyInterfaces(request *IAdditionsStateChangedEventgetMidlDoesNotLikeEmptyInterfaces) (*IAdditionsStateChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IAdditionsStateChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdapterChangedEventgetNetworkAdapter(request *INetworkAdapterChangedEventgetNetworkAdapter) (*INetworkAdapterChangedEventgetNetworkAdapterResponse, error) {
	response := new(INetworkAdapterChangedEventgetNetworkAdapterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortChangedEventgetSerialPort(request *ISerialPortChangedEventgetSerialPort) (*ISerialPortChangedEventgetSerialPortResponse, error) {
	response := new(ISerialPortChangedEventgetSerialPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortChangedEventgetParallelPort(request *IParallelPortChangedEventgetParallelPort) (*IParallelPortChangedEventgetParallelPortResponse, error) {
	response := new(IParallelPortChangedEventgetParallelPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllerChangedEventgetMidlDoesNotLikeEmptyInterfaces(request *IStorageControllerChangedEventgetMidlDoesNotLikeEmptyInterfaces) (*IStorageControllerChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IStorageControllerChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumChangedEventgetMediumAttachment(request *IMediumChangedEventgetMediumAttachment) (*IMediumChangedEventgetMediumAttachmentResponse, error) {
	response := new(IMediumChangedEventgetMediumAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IClipboardModeChangedEventgetClipboardMode(request *IClipboardModeChangedEventgetClipboardMode) (*IClipboardModeChangedEventgetClipboardModeResponse, error) {
	response := new(IClipboardModeChangedEventgetClipboardModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDnDModeChangedEventgetDndMode(request *IDnDModeChangedEventgetDndMode) (*IDnDModeChangedEventgetDndModeResponse, error) {
	response := new(IDnDModeChangedEventgetDndModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ICPUChangedEventgetCPU(request *ICPUChangedEventgetCPU) (*ICPUChangedEventgetCPUResponse, error) {
	response := new(ICPUChangedEventgetCPUResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ICPUChangedEventgetAdd(request *ICPUChangedEventgetAdd) (*ICPUChangedEventgetAddResponse, error) {
	response := new(ICPUChangedEventgetAddResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ICPUExecutionCapChangedEventgetExecutionCap(request *ICPUExecutionCapChangedEventgetExecutionCap) (*ICPUExecutionCapChangedEventgetExecutionCapResponse, error) {
	response := new(ICPUExecutionCapChangedEventgetExecutionCapResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestKeyboardEventgetScancodes(request *IGuestKeyboardEventgetScancodes) (*IGuestKeyboardEventgetScancodesResponse, error) {
	response := new(IGuestKeyboardEventgetScancodesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetMode(request *IGuestMouseEventgetMode) (*IGuestMouseEventgetModeResponse, error) {
	response := new(IGuestMouseEventgetModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetX(request *IGuestMouseEventgetX) (*IGuestMouseEventgetXResponse, error) {
	response := new(IGuestMouseEventgetXResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetY(request *IGuestMouseEventgetY) (*IGuestMouseEventgetYResponse, error) {
	response := new(IGuestMouseEventgetYResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetZ(request *IGuestMouseEventgetZ) (*IGuestMouseEventgetZResponse, error) {
	response := new(IGuestMouseEventgetZResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetW(request *IGuestMouseEventgetW) (*IGuestMouseEventgetWResponse, error) {
	response := new(IGuestMouseEventgetWResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetButtons(request *IGuestMouseEventgetButtons) (*IGuestMouseEventgetButtonsResponse, error) {
	response := new(IGuestMouseEventgetButtonsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetContactCount(request *IGuestMultiTouchEventgetContactCount) (*IGuestMultiTouchEventgetContactCountResponse, error) {
	response := new(IGuestMultiTouchEventgetContactCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetXPositions(request *IGuestMultiTouchEventgetXPositions) (*IGuestMultiTouchEventgetXPositionsResponse, error) {
	response := new(IGuestMultiTouchEventgetXPositionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetYPositions(request *IGuestMultiTouchEventgetYPositions) (*IGuestMultiTouchEventgetYPositionsResponse, error) {
	response := new(IGuestMultiTouchEventgetYPositionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetContactIds(request *IGuestMultiTouchEventgetContactIds) (*IGuestMultiTouchEventgetContactIdsResponse, error) {
	response := new(IGuestMultiTouchEventgetContactIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetContactFlags(request *IGuestMultiTouchEventgetContactFlags) (*IGuestMultiTouchEventgetContactFlagsResponse, error) {
	response := new(IGuestMultiTouchEventgetContactFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetScanTime(request *IGuestMultiTouchEventgetScanTime) (*IGuestMultiTouchEventgetScanTimeResponse, error) {
	response := new(IGuestMultiTouchEventgetScanTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionEventgetSession(request *IGuestSessionEventgetSession) (*IGuestSessionEventgetSessionResponse, error) {
	response := new(IGuestSessionEventgetSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionStateChangedEventgetId(request *IGuestSessionStateChangedEventgetId) (*IGuestSessionStateChangedEventgetIdResponse, error) {
	response := new(IGuestSessionStateChangedEventgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionStateChangedEventgetStatus(request *IGuestSessionStateChangedEventgetStatus) (*IGuestSessionStateChangedEventgetStatusResponse, error) {
	response := new(IGuestSessionStateChangedEventgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionStateChangedEventgetError(request *IGuestSessionStateChangedEventgetError) (*IGuestSessionStateChangedEventgetErrorResponse, error) {
	response := new(IGuestSessionStateChangedEventgetErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionRegisteredEventgetRegistered(request *IGuestSessionRegisteredEventgetRegistered) (*IGuestSessionRegisteredEventgetRegisteredResponse, error) {
	response := new(IGuestSessionRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessEventgetProcess(request *IGuestProcessEventgetProcess) (*IGuestProcessEventgetProcessResponse, error) {
	response := new(IGuestProcessEventgetProcessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessEventgetPid(request *IGuestProcessEventgetPid) (*IGuestProcessEventgetPidResponse, error) {
	response := new(IGuestProcessEventgetPidResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessRegisteredEventgetRegistered(request *IGuestProcessRegisteredEventgetRegistered) (*IGuestProcessRegisteredEventgetRegisteredResponse, error) {
	response := new(IGuestProcessRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessStateChangedEventgetStatus(request *IGuestProcessStateChangedEventgetStatus) (*IGuestProcessStateChangedEventgetStatusResponse, error) {
	response := new(IGuestProcessStateChangedEventgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessStateChangedEventgetError(request *IGuestProcessStateChangedEventgetError) (*IGuestProcessStateChangedEventgetErrorResponse, error) {
	response := new(IGuestProcessStateChangedEventgetErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessIOEventgetHandle(request *IGuestProcessIOEventgetHandle) (*IGuestProcessIOEventgetHandleResponse, error) {
	response := new(IGuestProcessIOEventgetHandleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessIOEventgetProcessed(request *IGuestProcessIOEventgetProcessed) (*IGuestProcessIOEventgetProcessedResponse, error) {
	response := new(IGuestProcessIOEventgetProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessInputNotifyEventgetStatus(request *IGuestProcessInputNotifyEventgetStatus) (*IGuestProcessInputNotifyEventgetStatusResponse, error) {
	response := new(IGuestProcessInputNotifyEventgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessOutputEventgetData(request *IGuestProcessOutputEventgetData) (*IGuestProcessOutputEventgetDataResponse, error) {
	response := new(IGuestProcessOutputEventgetDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileEventgetFile(request *IGuestFileEventgetFile) (*IGuestFileEventgetFileResponse, error) {
	response := new(IGuestFileEventgetFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileRegisteredEventgetRegistered(request *IGuestFileRegisteredEventgetRegistered) (*IGuestFileRegisteredEventgetRegisteredResponse, error) {
	response := new(IGuestFileRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileStateChangedEventgetStatus(request *IGuestFileStateChangedEventgetStatus) (*IGuestFileStateChangedEventgetStatusResponse, error) {
	response := new(IGuestFileStateChangedEventgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileStateChangedEventgetError(request *IGuestFileStateChangedEventgetError) (*IGuestFileStateChangedEventgetErrorResponse, error) {
	response := new(IGuestFileStateChangedEventgetErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileIOEventgetOffset(request *IGuestFileIOEventgetOffset) (*IGuestFileIOEventgetOffsetResponse, error) {
	response := new(IGuestFileIOEventgetOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileIOEventgetProcessed(request *IGuestFileIOEventgetProcessed) (*IGuestFileIOEventgetProcessedResponse, error) {
	response := new(IGuestFileIOEventgetProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileOffsetChangedEventgetMidlDoesNotLikeEmptyInterfaces(request *IGuestFileOffsetChangedEventgetMidlDoesNotLikeEmptyInterfaces) (*IGuestFileOffsetChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IGuestFileOffsetChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileReadEventgetData(request *IGuestFileReadEventgetData) (*IGuestFileReadEventgetDataResponse, error) {
	response := new(IGuestFileReadEventgetDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileWriteEventgetMidlDoesNotLikeEmptyInterfaces(request *IGuestFileWriteEventgetMidlDoesNotLikeEmptyInterfaces) (*IGuestFileWriteEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IGuestFileWriteEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServerChangedEventgetMidlDoesNotLikeEmptyInterfaces(request *IVRDEServerChangedEventgetMidlDoesNotLikeEmptyInterfaces) (*IVRDEServerChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IVRDEServerChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServerInfoChangedEventgetMidlDoesNotLikeEmptyInterfaces(request *IVRDEServerInfoChangedEventgetMidlDoesNotLikeEmptyInterfaces) (*IVRDEServerInfoChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IVRDEServerInfoChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVideoCaptureChangedEventgetMidlDoesNotLikeEmptyInterfaces(request *IVideoCaptureChangedEventgetMidlDoesNotLikeEmptyInterfaces) (*IVideoCaptureChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IVideoCaptureChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllerChangedEventgetMidlDoesNotLikeEmptyInterfaces(request *IUSBControllerChangedEventgetMidlDoesNotLikeEmptyInterfaces) (*IUSBControllerChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IUSBControllerChangedEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceStateChangedEventgetDevice(request *IUSBDeviceStateChangedEventgetDevice) (*IUSBDeviceStateChangedEventgetDeviceResponse, error) {
	response := new(IUSBDeviceStateChangedEventgetDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceStateChangedEventgetAttached(request *IUSBDeviceStateChangedEventgetAttached) (*IUSBDeviceStateChangedEventgetAttachedResponse, error) {
	response := new(IUSBDeviceStateChangedEventgetAttachedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceStateChangedEventgetError(request *IUSBDeviceStateChangedEventgetError) (*IUSBDeviceStateChangedEventgetErrorResponse, error) {
	response := new(IUSBDeviceStateChangedEventgetErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISharedFolderChangedEventgetScope(request *ISharedFolderChangedEventgetScope) (*ISharedFolderChangedEventgetScopeResponse, error) {
	response := new(ISharedFolderChangedEventgetScopeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IRuntimeErrorEventgetFatal(request *IRuntimeErrorEventgetFatal) (*IRuntimeErrorEventgetFatalResponse, error) {
	response := new(IRuntimeErrorEventgetFatalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IRuntimeErrorEventgetId(request *IRuntimeErrorEventgetId) (*IRuntimeErrorEventgetIdResponse, error) {
	response := new(IRuntimeErrorEventgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IRuntimeErrorEventgetMessage(request *IRuntimeErrorEventgetMessage) (*IRuntimeErrorEventgetMessageResponse, error) {
	response := new(IRuntimeErrorEventgetMessageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceChangedEventgetListener(request *IEventSourceChangedEventgetListener) (*IEventSourceChangedEventgetListenerResponse, error) {
	response := new(IEventSourceChangedEventgetListenerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceChangedEventgetAdd(request *IEventSourceChangedEventgetAdd) (*IEventSourceChangedEventgetAddResponse, error) {
	response := new(IEventSourceChangedEventgetAddResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataChangedEventgetMachineId(request *IExtraDataChangedEventgetMachineId) (*IExtraDataChangedEventgetMachineIdResponse, error) {
	response := new(IExtraDataChangedEventgetMachineIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataChangedEventgetKey(request *IExtraDataChangedEventgetKey) (*IExtraDataChangedEventgetKeyResponse, error) {
	response := new(IExtraDataChangedEventgetKeyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataChangedEventgetValue(request *IExtraDataChangedEventgetValue) (*IExtraDataChangedEventgetValueResponse, error) {
	response := new(IExtraDataChangedEventgetValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventaddVeto(request *IVetoEventaddVeto) (*IVetoEventaddVetoResponse, error) {
	response := new(IVetoEventaddVetoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventisVetoed(request *IVetoEventisVetoed) (*IVetoEventisVetoedResponse, error) {
	response := new(IVetoEventisVetoedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventgetVetos(request *IVetoEventgetVetos) (*IVetoEventgetVetosResponse, error) {
	response := new(IVetoEventgetVetosResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventaddApproval(request *IVetoEventaddApproval) (*IVetoEventaddApprovalResponse, error) {
	response := new(IVetoEventaddApprovalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventisApproved(request *IVetoEventisApproved) (*IVetoEventisApprovedResponse, error) {
	response := new(IVetoEventisApprovedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventgetApprovals(request *IVetoEventgetApprovals) (*IVetoEventgetApprovalsResponse, error) {
	response := new(IVetoEventgetApprovalsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataCanChangeEventgetMachineId(request *IExtraDataCanChangeEventgetMachineId) (*IExtraDataCanChangeEventgetMachineIdResponse, error) {
	response := new(IExtraDataCanChangeEventgetMachineIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataCanChangeEventgetKey(request *IExtraDataCanChangeEventgetKey) (*IExtraDataCanChangeEventgetKeyResponse, error) {
	response := new(IExtraDataCanChangeEventgetKeyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataCanChangeEventgetValue(request *IExtraDataCanChangeEventgetValue) (*IExtraDataCanChangeEventgetValueResponse, error) {
	response := new(IExtraDataCanChangeEventgetValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ICanShowWindowEventgetMidlDoesNotLikeEmptyInterfaces(request *ICanShowWindowEventgetMidlDoesNotLikeEmptyInterfaces) (*ICanShowWindowEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(ICanShowWindowEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IShowWindowEventgetWinId(request *IShowWindowEventgetWinId) (*IShowWindowEventgetWinIdResponse, error) {
	response := new(IShowWindowEventgetWinIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IShowWindowEventsetWinId(request *IShowWindowEventsetWinId) (*IShowWindowEventsetWinIdResponse, error) {
	response := new(IShowWindowEventsetWinIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetSlot(request *INATRedirectEventgetSlot) (*INATRedirectEventgetSlotResponse, error) {
	response := new(INATRedirectEventgetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetRemove(request *INATRedirectEventgetRemove) (*INATRedirectEventgetRemoveResponse, error) {
	response := new(INATRedirectEventgetRemoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetName(request *INATRedirectEventgetName) (*INATRedirectEventgetNameResponse, error) {
	response := new(INATRedirectEventgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetProto(request *INATRedirectEventgetProto) (*INATRedirectEventgetProtoResponse, error) {
	response := new(INATRedirectEventgetProtoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetHostIP(request *INATRedirectEventgetHostIP) (*INATRedirectEventgetHostIPResponse, error) {
	response := new(INATRedirectEventgetHostIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetHostPort(request *INATRedirectEventgetHostPort) (*INATRedirectEventgetHostPortResponse, error) {
	response := new(INATRedirectEventgetHostPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetGuestIP(request *INATRedirectEventgetGuestIP) (*INATRedirectEventgetGuestIPResponse, error) {
	response := new(INATRedirectEventgetGuestIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetGuestPort(request *INATRedirectEventgetGuestPort) (*INATRedirectEventgetGuestPortResponse, error) {
	response := new(INATRedirectEventgetGuestPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostPCIDevicePlugEventgetPlugged(request *IHostPCIDevicePlugEventgetPlugged) (*IHostPCIDevicePlugEventgetPluggedResponse, error) {
	response := new(IHostPCIDevicePlugEventgetPluggedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostPCIDevicePlugEventgetSuccess(request *IHostPCIDevicePlugEventgetSuccess) (*IHostPCIDevicePlugEventgetSuccessResponse, error) {
	response := new(IHostPCIDevicePlugEventgetSuccessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostPCIDevicePlugEventgetAttachment(request *IHostPCIDevicePlugEventgetAttachment) (*IHostPCIDevicePlugEventgetAttachmentResponse, error) {
	response := new(IHostPCIDevicePlugEventgetAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostPCIDevicePlugEventgetMessage(request *IHostPCIDevicePlugEventgetMessage) (*IHostPCIDevicePlugEventgetMessageResponse, error) {
	response := new(IHostPCIDevicePlugEventgetMessageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVBoxSVCAvailabilityChangedEventgetAvailable(request *IVBoxSVCAvailabilityChangedEventgetAvailable) (*IVBoxSVCAvailabilityChangedEventgetAvailableResponse, error) {
	response := new(IVBoxSVCAvailabilityChangedEventgetAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupChangedEventgetBandwidthGroup(request *IBandwidthGroupChangedEventgetBandwidthGroup) (*IBandwidthGroupChangedEventgetBandwidthGroupResponse, error) {
	response := new(IBandwidthGroupChangedEventgetBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetChangeType(request *IGuestMonitorChangedEventgetChangeType) (*IGuestMonitorChangedEventgetChangeTypeResponse, error) {
	response := new(IGuestMonitorChangedEventgetChangeTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetScreenId(request *IGuestMonitorChangedEventgetScreenId) (*IGuestMonitorChangedEventgetScreenIdResponse, error) {
	response := new(IGuestMonitorChangedEventgetScreenIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetOriginX(request *IGuestMonitorChangedEventgetOriginX) (*IGuestMonitorChangedEventgetOriginXResponse, error) {
	response := new(IGuestMonitorChangedEventgetOriginXResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetOriginY(request *IGuestMonitorChangedEventgetOriginY) (*IGuestMonitorChangedEventgetOriginYResponse, error) {
	response := new(IGuestMonitorChangedEventgetOriginYResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetWidth(request *IGuestMonitorChangedEventgetWidth) (*IGuestMonitorChangedEventgetWidthResponse, error) {
	response := new(IGuestMonitorChangedEventgetWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetHeight(request *IGuestMonitorChangedEventgetHeight) (*IGuestMonitorChangedEventgetHeightResponse, error) {
	response := new(IGuestMonitorChangedEventgetHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestUserStateChangedEventgetName(request *IGuestUserStateChangedEventgetName) (*IGuestUserStateChangedEventgetNameResponse, error) {
	response := new(IGuestUserStateChangedEventgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestUserStateChangedEventgetDomain(request *IGuestUserStateChangedEventgetDomain) (*IGuestUserStateChangedEventgetDomainResponse, error) {
	response := new(IGuestUserStateChangedEventgetDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestUserStateChangedEventgetState(request *IGuestUserStateChangedEventgetState) (*IGuestUserStateChangedEventgetStateResponse, error) {
	response := new(IGuestUserStateChangedEventgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestUserStateChangedEventgetStateDetails(request *IGuestUserStateChangedEventgetStateDetails) (*IGuestUserStateChangedEventgetStateDetailsResponse, error) {
	response := new(IGuestUserStateChangedEventgetStateDetailsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageDeviceChangedEventgetStorageDevice(request *IStorageDeviceChangedEventgetStorageDevice) (*IStorageDeviceChangedEventgetStorageDeviceResponse, error) {
	response := new(IStorageDeviceChangedEventgetStorageDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageDeviceChangedEventgetRemoved(request *IStorageDeviceChangedEventgetRemoved) (*IStorageDeviceChangedEventgetRemovedResponse, error) {
	response := new(IStorageDeviceChangedEventgetRemovedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageDeviceChangedEventgetSilent(request *IStorageDeviceChangedEventgetSilent) (*IStorageDeviceChangedEventgetSilentResponse, error) {
	response := new(IStorageDeviceChangedEventgetSilentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkChangedEventgetNetworkName(request *INATNetworkChangedEventgetNetworkName) (*INATNetworkChangedEventgetNetworkNameResponse, error) {
	response := new(INATNetworkChangedEventgetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkStartStopEventgetStartEvent(request *INATNetworkStartStopEventgetStartEvent) (*INATNetworkStartStopEventgetStartEventResponse, error) {
	response := new(INATNetworkStartStopEventgetStartEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkAlterEventgetMidlDoesNotLikeEmptyInterfaces(request *INATNetworkAlterEventgetMidlDoesNotLikeEmptyInterfaces) (*INATNetworkAlterEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(INATNetworkAlterEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkCreationDeletionEventgetCreationEvent(request *INATNetworkCreationDeletionEventgetCreationEvent) (*INATNetworkCreationDeletionEventgetCreationEventResponse, error) {
	response := new(INATNetworkCreationDeletionEventgetCreationEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetEnabled(request *INATNetworkSettingEventgetEnabled) (*INATNetworkSettingEventgetEnabledResponse, error) {
	response := new(INATNetworkSettingEventgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetNetwork(request *INATNetworkSettingEventgetNetwork) (*INATNetworkSettingEventgetNetworkResponse, error) {
	response := new(INATNetworkSettingEventgetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetGateway(request *INATNetworkSettingEventgetGateway) (*INATNetworkSettingEventgetGatewayResponse, error) {
	response := new(INATNetworkSettingEventgetGatewayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabled(request *INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabled) (*INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabledResponse, error) {
	response := new(INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetNeedDhcpServer(request *INATNetworkSettingEventgetNeedDhcpServer) (*INATNetworkSettingEventgetNeedDhcpServerResponse, error) {
	response := new(INATNetworkSettingEventgetNeedDhcpServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetCreate(request *INATNetworkPortForwardEventgetCreate) (*INATNetworkPortForwardEventgetCreateResponse, error) {
	response := new(INATNetworkPortForwardEventgetCreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetIpv6(request *INATNetworkPortForwardEventgetIpv6) (*INATNetworkPortForwardEventgetIpv6Response, error) {
	response := new(INATNetworkPortForwardEventgetIpv6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetName(request *INATNetworkPortForwardEventgetName) (*INATNetworkPortForwardEventgetNameResponse, error) {
	response := new(INATNetworkPortForwardEventgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetProto(request *INATNetworkPortForwardEventgetProto) (*INATNetworkPortForwardEventgetProtoResponse, error) {
	response := new(INATNetworkPortForwardEventgetProtoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetHostIp(request *INATNetworkPortForwardEventgetHostIp) (*INATNetworkPortForwardEventgetHostIpResponse, error) {
	response := new(INATNetworkPortForwardEventgetHostIpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetHostPort(request *INATNetworkPortForwardEventgetHostPort) (*INATNetworkPortForwardEventgetHostPortResponse, error) {
	response := new(INATNetworkPortForwardEventgetHostPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetGuestIp(request *INATNetworkPortForwardEventgetGuestIp) (*INATNetworkPortForwardEventgetGuestIpResponse, error) {
	response := new(INATNetworkPortForwardEventgetGuestIpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetGuestPort(request *INATNetworkPortForwardEventgetGuestPort) (*INATNetworkPortForwardEventgetGuestPortResponse, error) {
	response := new(INATNetworkPortForwardEventgetGuestPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNameResolutionConfigurationChangeEventgetMidlDoesNotLikeEmptyInterfaces(request *IHostNameResolutionConfigurationChangeEventgetMidlDoesNotLikeEmptyInterfaces) (*IHostNameResolutionConfigurationChangeEventgetMidlDoesNotLikeEmptyInterfacesResponse, error) {
	response := new(IHostNameResolutionConfigurationChangeEventgetMidlDoesNotLikeEmptyInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
		//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	err := encoder.Encode(envelope)
	if err == nil {
		err = encoder.Flush()
	}

	// log.Println(buffer.String())
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	// log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
