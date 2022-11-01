package virtualboxapi

import (
	virtualboxsrv "github.com/easysoft/zagent/internal/pkg/vendors/virtualbox/srv"
	_stringUtils "github.com/easysoft/zagent/pkg/lib/string"
	"strings"
)

type NetworkAdapter struct {
	virtualbox      *VirtualBox
	managedObjectId string
}

func (na *NetworkAdapter) GetMACAddress() (string, error) {
	request := virtualboxsrv.INetworkAdaptergetMACAddress{This: na.managedObjectId}

	response, err := na.virtualbox.INetworkAdaptergetMACAddress(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	response.Returnval = strings.ToLower(_stringUtils.AddSepForMacAddress(response.Returnval))

	return response.Returnval, nil
}

func (na *NetworkAdapter) SetBridge(bri string) (err error) {
	attachmentType := virtualboxsrv.NetworkAttachmentTypeBridged

	// set to bridge net
	request1 := virtualboxsrv.INetworkAdaptersetAttachmentType{
		This:           na.managedObjectId,
		AttachmentType: &attachmentType,
	}
	_, err = na.virtualbox.INetworkAdaptersetAttachmentType(&request1)
	if err != nil {
		return
	}

	// set to bridge interface
	request2 := virtualboxsrv.INetworkAdaptersetBridgedInterface{
		This:             na.managedObjectId,
		BridgedInterface: bri,
	}
	_, err = na.virtualbox.INetworkAdaptersetBridgedInterface(&request2)
	if err != nil {
		return
	}

	return
}
