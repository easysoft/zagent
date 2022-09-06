package virtualboxapi

import (
	"errors"
	virtualboxsrv "github.com/easysoft/zv/internal/pkg/vendors/virtualbox/srv"
)

type StorageController struct {
	virtualbox      *VirtualBox
	managedObjectId string
	Name            string
}

type StorageControllers struct {
	storageControllers []*StorageController
}

func (sc *StorageController) GetName() (string, error) {
	request := virtualboxsrv.IStorageControllergetName{This: sc.managedObjectId}

	response, err := sc.virtualbox.IStorageControllergetName(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sc *StorageController) GetPortCount() (uint32, error) {
	request := virtualboxsrv.IStorageControllergetPortCount{This: sc.managedObjectId}

	response, err := sc.virtualbox.IStorageControllergetPortCount(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sc *StorageController) GetStorageBus() (virtualboxsrv.StorageBus, error) {
	mapStorageBus := make(map[string]virtualboxsrv.StorageBus)
	mapStorageBus["SATA Controller"] = virtualboxsrv.StorageBusSATA
	mapStorageBus["IDE Controller"] = virtualboxsrv.StorageBusIDE
	mapStorageBus["SCSI"] = virtualboxsrv.StorageBusSCSI
	mapStorageBus["SAS"] = virtualboxsrv.StorageBusSAS

	scName, err := sc.GetName()
	if err != nil {
		return virtualboxsrv.StorageBusNull, err
	}

	if bus, ok := mapStorageBus[scName]; ok {
		return bus, nil
	}

	return virtualboxsrv.StorageBusNull, errors.New("bad controller controller specified")
}

func (sc *StorageController) GetMaxPortCount() (uint32, error) {
	request := virtualboxsrv.IStorageControllergetMaxPortCount{This: sc.managedObjectId}

	response, err := sc.virtualbox.IStorageControllergetMaxPortCount(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sc *StorageController) SetPortCount(count uint32) error {
	request := virtualboxsrv.IStorageControllersetPortCount{This: sc.managedObjectId, PortCount: count}

	_, err := sc.virtualbox.IStorageControllersetPortCount(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	return nil
}

func (sc *StorageController) GetNextAvailablePort(m *Machine) (int32, error) {
	c, err := sc.GetMaxPortCount()
	if err != nil {
		return 0, err
	}

	ams, err := m.GetMediumAttachmentsOfController(sc.Name)
	if err != nil {
		return 0, nil
	}

	portMap := make(map[int32]bool)
	for _, am := range ams {
		portMap[am.Port] = true
	}

	intArr := make([]int32, c)
	for i := range intArr {
		if _, isUsed := portMap[int32(i)]; !isUsed {
			return int32(i), nil
		}
	}
	return 0, errors.New("no available ports")
}

func (sc *StorageController) Release() error {
	return sc.virtualbox.Release(sc.managedObjectId)
}
