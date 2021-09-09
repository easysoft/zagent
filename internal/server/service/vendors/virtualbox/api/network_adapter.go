package virtualboxapi

import (
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/srv"
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
