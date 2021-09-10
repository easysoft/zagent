package main

import (
	"github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/api"
	"log"
	"testing"
)

const (
	name = "test-win10-01"
)

func TestVirtualBox(t *testing.T) {
	url := "http://192.168.0.56:18083"

	client := virtualboxapi.NewVirtualBox("aaron", "P2ssw0rd", url, false, "")
	err := client.Logon()
	if err != nil {
		log.Printf("Unable to log on to vboxwebsrv: %v\n", err)
	}

	machines, err := client.GetMachines()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	log.Printf("%#v\n", machines)

	templ, err := client.FindMachine("win10-pro-x64-zh_cn")
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

	newMachineId, err := client.CreateMachine(name, osTypeId)
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
	err = newMachine.SetCPUCount(2)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	err = newMachine.SetMemorySize(4 * 1024)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	adpt, err := newMachine.GetNetworkAdapter(0)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	err = adpt.SetBridge("br0")
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	macAddress, err := adpt.GetMACAddress()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	log.Printf("machine mac address %s", macAddress)

	err = newMachine.SaveSettings()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	err = newMachine.Register()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	machine, err := client.FindMachine(name)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	machineState, err := machine.GetMachineState()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	log.Printf("%#v\n", machineState)

	session, err := client.GetSession()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	progress, err = machine.Launch(session.ManagedObjectId)
	err = progress.WaitForCompletion(10000)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}
