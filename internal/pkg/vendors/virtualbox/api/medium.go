package virtualboxapi

import (
	virtualboxsrv "github.com/easysoft/zagent/internal/pkg/vendors/virtualbox/srv"
)

type Medium struct {
	virtualbox      *VirtualBox
	managedObjectId string
	Location        string
	Name            string
	DeviceType      virtualboxsrv.DeviceType
	Description     string
	LogicalSize     int64
	Size            int64
	Format          string
	MediumFormat    string
	HostDrive       bool
	Children        []string
	Parent          string
	ID              string
	MachineIDs      []string
	SnapshotIDs     []string
}

func (m *Medium) CreateBaseStorage(logicalSize int64, variant []*virtualboxsrv.MediumVariant) (*Progress, error) {
	request := virtualboxsrv.IMediumcreateBaseStorage{This: m.managedObjectId, LogicalSize: logicalSize, Variant: variant}

	response, err := m.virtualbox.IMediumcreateBaseStorage(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Progress{virtualbox: m.virtualbox, managedObjectId: response.Returnval}, nil
}

func (m *Medium) DeleteStorage() (*Progress, error) {
	request := virtualboxsrv.IMediumdeleteStorage{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumdeleteStorage(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Progress{virtualbox: m.virtualbox, managedObjectId: response.Returnval}, nil
}

func (m *Medium) Release() error {
	return m.virtualbox.Release(m.managedObjectId)
}

func (m *Medium) GetLocation() (string, error) {
	request := virtualboxsrv.IMediumgetLocation{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetLocation(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetName() (string, error) {
	request := virtualboxsrv.IMediumgetName{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetName(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetDeviceType() (*virtualboxsrv.DeviceType, error) {
	request := virtualboxsrv.IMediumgetDeviceType{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetDeviceType(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetDescription() (string, error) {
	request := virtualboxsrv.IMediumgetDescription{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetDescription(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetSize() (int64, error) {
	request := virtualboxsrv.IMediumgetSize{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetSize(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetLogicalSize() (int64, error) {
	request := virtualboxsrv.IMediumgetLogicalSize{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetLogicalSize(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetState() (*virtualboxsrv.MediumState, error) {
	request := virtualboxsrv.IMediumgetState{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetState(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetFormat() (string, error) {
	request := virtualboxsrv.IMediumgetFormat{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetFormat(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetMediumFormat() (string, error) {
	request := virtualboxsrv.IMediumgetMediumFormat{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetMediumFormat(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetHostDrive() (bool, error) {
	request := virtualboxsrv.IMediumgetHostDrive{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetHostDrive(&request)
	if err != nil {
		return false, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetParent() (string, error) {
	request := virtualboxsrv.IMediumgetParent{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetParent(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetChildren() ([]string, error) {
	request := virtualboxsrv.IMediumgetChildren{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetChildren(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) DetachMachines() error {
	for _, mid := range m.MachineIDs {
		machine, err := m.virtualbox.FindMachine(mid)
		if err != nil {
			return err
		}
		defer machine.Release()

		if err := machine.DetachDevice(m); err != nil {
			return err
		}
	}
	return nil
}

func (m *Medium) GetID() (string, error) {
	request := virtualboxsrv.IMediumgetId{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetId(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetSnapshotIDs() ([]string, error) {
	request := virtualboxsrv.IMediumgetSnapshotIds{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetSnapshotIds(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) GetMachineIDs() ([]string, error) {
	request := virtualboxsrv.IMediumgetMachineIds{This: m.managedObjectId}

	response, err := m.virtualbox.IMediumgetMachineIds(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Medium) Get() (*Medium, error) {
	var err error
	m.Location, err = m.GetLocation()
	if err != nil {
		return nil, err
	}
	m.Name, err = m.GetName()
	if err != nil {
		return nil, err
	}

	m.Description, err = m.GetDescription()
	if err != nil {
		return nil, err
	}

	m.Size, err = m.GetSize()
	if err != nil {
		return nil, err
	}

	m.LogicalSize, err = m.GetLogicalSize()
	if err != nil {
		return nil, err
	}

	dt, err := m.GetDeviceType()
	if err != nil {
		return nil, err
	}
	m.DeviceType = *dt

	m.Format, err = m.GetFormat()
	if err != nil {
		return nil, err
	}

	m.MediumFormat, err = m.GetMediumFormat()
	if err != nil {
		return nil, err
	}

	m.HostDrive, err = m.GetHostDrive()
	if err != nil {
		return nil, err
	}

	m.Children, err = m.GetChildren()
	if err != nil {
		return nil, err
	}

	m.Parent, err = m.GetParent()
	if err != nil {
		return nil, err
	}

	m.ID, err = m.GetID()
	if err != nil {
		return nil, err
	}

	m.MachineIDs, err = m.GetMachineIDs()
	if err != nil {
		return nil, err
	}

	m.SnapshotIDs, err = m.GetSnapshotIDs()
	if err != nil {
		return nil, err
	}

	return m, nil

}

func (m *Medium) GetIDName() (*Medium, error) {
	var err error
	m.ID, err = m.GetID()
	if err != nil {
		return nil, err
	}
	m.Name, err = m.GetName()
	if err != nil {
		return nil, err
	}
	return m, nil
}
