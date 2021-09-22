package vmwareService

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	vmx "github.com/johlandabee/govmx"
)

func GetVM(c *Client, id string) (*Vm, error) {
	var vms []Vm
	var vm Vm
	// If you want see the path of the VM it's necessary getting all VMs
	// because the API of VmWare Workstation doesn't permit see this the another way
	response, err := c.httpRequest("vms", "GET", bytes.Buffer{})
	if err != nil {
		log.Fatalf("[WSAPICLI] Fi: wsapitools.go Fu: GetVM Message: The request at the server API failed %s", err)
		return nil, err
	}
	err = json.NewDecoder(response).Decode(&vms)
	if err != nil {
		log.Fatalf("[WSAPICLI][ERROR] Fi: wsapitools.go Fu: GetVM Message: I can't read the json structure %s", err)
		return nil, err
	}
	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: GetVM Obj: List of VMs %#v\n", vms)
	for tempvm, value := range vms {
		if value.IdVM == id {
			vm = vms[tempvm]
		}
	}

	name, desc, err := GetNameDescription(vm.Path)
	if err != nil {
		return nil, err
	}
	vm.Denomination = name
	vm.Description = desc

	response, err = c.httpRequest("vms/"+id, "GET", bytes.Buffer{})
	if err != nil {
		log.Fatalf("[WSAPICLI] Fi: wsapitools.go Fu: GetVM Message: The request at the server API failed %s", err)
		return nil, err
	}

	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: GetVM Obj:Body of VM %#v\n", response)
	err = json.NewDecoder(response).Decode(&vm)
	if err != nil {
		log.Fatalf("[WSAPICLI] Fi: wsapitools.go Fu: GetVM Message: I can't read the json structure %s", err)
		return nil, err
	}

	response, err = c.httpRequest("vms/"+id+"/power", "GET", bytes.Buffer{})
	if err != nil {
		log.Fatalf("[WSAPICLI] Fi: wsapitools.go Fu: GetVM Message: The request at the server API failed %s", err)
		return nil, err
	}
	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: GetVM Obj:Body of power %#v\n", response)

	err = json.NewDecoder(response).Decode(&vm)
	if err != nil {
		log.Fatalf("[WSAPICLI] Fi: wsapitools.go Fu: GetVM Message: I can't read the json structure %s", err)
		return nil, err
	}

	return &vm, nil
}

func GetNameDescription(path string) (name, desc string, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("[WSAPICLI][ERROR] Fi: wsapitools.go Fu: GetNameDescription Message: Failed opening file %s, please make sure the config file exists", err)
		return
	}

	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: GetNameDescription Obj: Data File %#v\n", string(data))

	vm := new(vmx.VirtualMachine)
	err = vmx.Unmarshal(data, vm)
	if err != nil {
		log.Fatalf("[WSAPICLI][ERROR] Fi: wsapitools.go Fu: GetNameDescription Obj: %#v", err)
		return
	}

	name = vm.DisplayName
	desc = vm.Annotation

	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: GetNameDescription Obj: VM  %#v\n", vm)
	return
}

func SetNameDescription(path string, name string, desc string) error {
	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: SetNameDescription Message: parameters %#v, %#v, %#v", path, name, desc)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("[WSAPICLI] Fi: wsapitools.go Fu: SetNameDescription Message: Failed opening file %s, please make sure the config file exists", err)
		return err
	}

	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: SetNameDescription Obj: File object %#v\n", string(data))

	vm := new(vmx.VirtualMachine)
	err = vmx.Unmarshal(data, vm)
	if err != nil {
		log.Fatalf("[WSAPICLI] Fi: wsapitools.go Fu: GetNameDescription Obj: %#v", err)
		return err
	}
	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: GetNameDescription Obj: VM %#v\n", vm)

	vm.DisplayName = name
	vm.Annotation = desc
	data, err = vmx.Marshal(vm)
	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: SetNameDescription Obj: Data File %#v\n", string(data))
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		log.Fatalf("[WSAPICLI] Fi: wsapitools.go Fu: SetNameDescription Message: Failed writing in file %s, please make sure the config file exists", err)
		return err
	}
	// en este punto tambien tienes que cambiar el nombre del fihero cuando se cambia la denominacion
	return err
}

//func (c *Client) SetParameter(id string, param string, value string) error {
//	requestBody := new(bytes.Buffer)
//	request, err := json.Marshal(map[string]string{
//		"name":  param,
//		"value": value,
//	})
//	if err != nil {
//		return err
//	}
//	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: SetParameter Obj:Request %#v\n", request)
//	requestBody.Write(request)
//	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: SetParameter Obj:Request Body %#v\n", requestBody.String())
//	response, err := c.httpRequest("/vms/"+id+"/configparams", "PUT", *requestBody)
//	if err != nil {
//		return err
//	}
//	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: SetParameter Obj:response raw %#v\n", response)
//	responseBody := new(bytes.Buffer)
//	_, err = responseBody.ReadFrom(response)
//	if err != nil {
//		log.Printf("[WSAPICLI][ERROR] Fi: wsapitools.go Fu: SetParameter Obj:Response Error %#v\n", err)
//		return err
//	}
//	log.Printf("[WSAPICLI] Fi: wsapitools.go Fu: SetParameter Obj:Response Body %#v\n", responseBody.String())
//	// err = json.NewDecoder(responseBody).Decode(&vm)
//	if err != nil {
//		return err
//	}
//	return err
//}
