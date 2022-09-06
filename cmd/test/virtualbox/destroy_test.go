package main

import (
	virtualboxapi "github.com/easysoft/zv/internal/pkg/vendors/virtualbox/api"
	virtualboxsrv "github.com/easysoft/zv/internal/pkg/vendors/virtualbox/srv"
	"log"
	"testing"
)

func TestVirtualBoxLaunch(t *testing.T) {
	url := "http://192.168.0.56:18083"

	virtualBox := virtualboxapi.NewVirtualBox("aaron", "P2ssw0rd", url, false, "")
	err := virtualBox.Logon()
	if err != nil {
		log.Printf("Unable to log on to vboxwebsrv: %v\n", err)
	}

	machine, err := virtualBox.FindMachine("win10-01")
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	machineState, err := machine.GetMachineState()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	log.Printf("%#v\n", *machineState)

	session, err := virtualBox.GetSession()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	err = machine.Lock(session, virtualboxsrv.LockTypeShared)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	console, err := session.GetConsole()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	progress, err := console.PowerDown()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
	err = progress.WaitForCompletion(10000)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	//err = machine.Unlock(session)
	//if err != nil {
	//	log.Printf("%s\n", err.Error())
	//}

	media, err := machine.Unregister()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	err = machine.DiscardSettings()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	err = machine.DeleteConfig(media)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	err = session.Release()
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}
