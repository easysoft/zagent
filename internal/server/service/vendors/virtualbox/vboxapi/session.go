package vboxapi

import (
	"github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/vboxwebsrv"
)

type Session struct {
	virtualbox      *VirtualBox
	ManagedObjectId string
}

func (s *Session) GetConsole() (*Console, error) {
	request := vboxwebsrv.ISessiongetConsole{This: s.ManagedObjectId}
	response, err := s.virtualbox.ISessiongetConsole(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Console{managedObjectID: response.Returnval, virtualbox: s.virtualbox}, nil
}

func (s *Session) UnlockMachine() error {
	request := vboxwebsrv.ISessionunlockMachine{This: s.ManagedObjectId}
	_, err := s.virtualbox.ISessionunlockMachine(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}

func (s *Session) LockMachine(m *Machine, l vboxwebsrv.LockType) error {
	request := vboxwebsrv.IMachinelockMachine{
		This:     m.managedObjectId,
		Session:  s.ManagedObjectId,
		LockType: &l,
	}
	_, err := s.virtualbox.IMachinelockMachine(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}

func (s *Session) GetMachine() (*Machine, error) {
	request := vboxwebsrv.ISessiongetMachine{This: s.ManagedObjectId}
	response, err := s.virtualbox.ISessiongetMachine(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Machine{managedObjectId: response.Returnval, virtualbox: s.virtualbox}, nil
}

func (s *Session) Release() error {
	return s.virtualbox.Release(s.ManagedObjectId)
}
