package vboxapi

import "github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/vboxwebsrv"

type MediumAttachment struct {
	*vboxwebsrv.IMediumAttachment
	virtualbox      *VirtualBox
	managedObjectId string
}

func (m *MediumAttachment) GetMedium() (*Medium, error) {
	return &Medium{virtualbox: m.virtualbox, managedObjectId: m.Medium}, nil
}
