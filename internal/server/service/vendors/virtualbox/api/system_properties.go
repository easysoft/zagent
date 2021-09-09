package virtualboxapi

import (
	"github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/srv"
)

type SystemProperties struct {
	virtualbox      *VirtualBox
	managedObjectId string
}

func (sp *SystemProperties) GetMaxNetworkAdapters(chipset *virtualboxsrv.ChipsetType) (uint32, error) {
	request := virtualboxsrv.ISystemPropertiesgetMaxNetworkAdapters{This: sp.managedObjectId, Chipset: chipset}

	response, err := sp.virtualbox.ISystemPropertiesgetMaxNetworkAdapters(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sp *SystemProperties) GetMaxDevicesPerPortForStorageBus(bus virtualboxsrv.StorageBus) (uint32, error) {
	request := virtualboxsrv.ISystemPropertiesgetMaxDevicesPerPortForStorageBus{This: sp.managedObjectId, Bus: &bus}
	response, err := sp.virtualbox.ISystemPropertiesgetMaxDevicesPerPortForStorageBus(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sp *SystemProperties) GetMinPortCountForStorageBus(bus virtualboxsrv.StorageBus) (uint32, error) {
	request := virtualboxsrv.ISystemPropertiesgetMinPortCountForStorageBus{This: sp.managedObjectId, Bus: &bus}
	response, err := sp.virtualbox.ISystemPropertiesgetMinPortCountForStorageBus(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sp *SystemProperties) Release() error {
	return sp.virtualbox.Release(sp.managedObjectId)
}
