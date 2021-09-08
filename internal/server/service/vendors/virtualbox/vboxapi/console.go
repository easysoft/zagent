package vboxapi

import "github.com/easysoft/zagent/internal/server/service/vendors/virtualbox/vboxwebsrv"

// Console is a VirtualBox console object
type Console struct {
	virtualbox      *VirtualBox
	managedObjectID string
}

// PowerDown starts forcibly powering off the controlled VM.
// It returns a Progress and any error encountered.
func (c *Console) PowerDown() (*Progress, error) {
	request := vboxwebsrv.IConsolepowerDown{This: c.managedObjectID}

	response, err := c.virtualbox.IConsolepowerDown(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Progress{virtualbox: c.virtualbox, managedObjectId: response.Returnval}, nil
}

// PowerUp starts powering on the controlled VM.
// It returns a Progress and any error encountered.
func (c *Console) PowerUp() (*Progress, error) {
	request := vboxwebsrv.IConsolepowerUp{This: c.managedObjectID}

	response, err := c.virtualbox.IConsolepowerUp(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Progress{virtualbox: c.virtualbox, managedObjectId: response.Returnval}, nil
}

// func (console *Console) PowerDown() (Progress, error) {
// 	var progress Progress
// 	result := C.GoVboxConsolePowerDown(console.cconsole, &progress.cprogress)
// 	if C.GoVboxFAILED(result) != 0 || progress.cprogress == nil {
// 		return progress, errors.NewVirtualBox(
// 			fmt.Sprintf("Failed to power down VM via IConsole: %x", result))
// 	}
// 	return progress, nil
// }
