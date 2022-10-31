package virtualboxapi

import (
	virtualboxsrv "github.com/easysoft/zagent/internal/pkg/vendors/virtualbox/srv"
)

type MediumAttachment struct {
	*virtualboxsrv.IMediumAttachment
	virtualbox      *VirtualBox
	managedObjectId string
}

func (m *MediumAttachment) GetMedium() (*Medium, error) {
	return &Medium{virtualbox: m.virtualbox, managedObjectId: m.Medium}, nil
}
