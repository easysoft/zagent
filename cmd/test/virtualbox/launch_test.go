package main

import (
	"github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/api"
	"log"
	"testing"
)

func TestVirtualBox(t *testing.T) {
	url := "http://192.168.0.56:18083"

	virtualBox := virtualboxapi.NewVirtualBox("aaron", "P2ssw0rd", url, false, "")
	err := virtualBox.Logon()
	if err != nil {
		log.Printf("Unable to log on to vboxwebsrv: %v\n", err)
	}

	machines, err := virtualBox.GetMachines()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	log.Printf("%#v\n", machines)

	templ, err := virtualBox.FindMachine("win10-pro-x64-zh_cn")
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	osTypeId, err := templ.GetOsTypeId()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	snapshot, err := templ.FindSnapshot()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	snapshotMachine, err := snapshot.FindSnapshotMachine()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	newMachineId, err := virtualBox.CreateMachine("win10-01", osTypeId)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	progress, newMachine, err := snapshotMachine.CloneTo(newMachineId)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	err = progress.WaitForCompletion(10000)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	err = newMachine.SaveSettings()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	err = newMachine.Register()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	machine, err := virtualBox.FindMachine("win10-01")
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	machineState, err := machine.GetMachineState()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	log.Printf("%#v\n", machineState)

	session, err := virtualBox.GetSession()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	progress, err = machine.Launch(session.ManagedObjectId)
	err = progress.WaitForCompletion(10000)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}
