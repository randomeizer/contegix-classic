package contegixclassic

import "fmt"

// ListVirtualMachines retrieves all the VMs for the current account.
func (c *Client) ListVirtualMachines() ([]VirtualMachine, error) {
	vmResponses := []virtualMachineResponse{}

	_, err := c.DoRequest("GET", "/virtual_machines", nil, &vmResponses)
	if err != nil {
		return nil, err
	}

	virtualMachines := make([]VirtualMachine, len(vmResponses))
	for i, vmr := range vmResponses {
		virtualMachines[i] = vmr.VirtualMachine
	}
	return virtualMachines, nil
}

// GetVirtualMachine returns the details for the specified VM
func (c *Client) GetVirtualMachine(uuid string) (*VirtualMachine, error) {
	vmResponse := virtualMachineResponse{}

	_, err := c.DoRequest("GET", fmt.Sprintf("/virtual_machines/%v", uuid), nil, &vmResponse)
	if err != nil {
		return nil, err
	}

	return &vmResponse.VirtualMachine, nil
}

// CreateVirtualMachine will attempt to create a new VM with the specified details. The name must be unique within your account.
func (c *Client) CreateVirtualMachine(name string, zoneUUID string, packageUUID string, templateUUID string) (*VirtualMachine, error) {
	vmResponse := virtualMachineResponse{}

	params := map[string]interface{}{
		"virtual_machine[name]":          name,
		"virtual_machine[zone_uuid]":     zoneUUID,
		"virtual_machine[package_uuid]":  packageUUID,
		"virtual_machine[template_uuid]": templateUUID,
	}

	_, err := c.DoRequest("POST", "/virtual_machines", params, &vmResponse)
	if err != nil {
		return nil, err
	}

	return &vmResponse.VirtualMachine, nil
}

// ResizeVirtualMachine allows changing the package of a VM, which adjusts the capacity and cost.
func (c *Client) ResizeVirtualMachine(uuid string, packageUUID string) (*VirtualMachine, error) {
	vmResponse := virtualMachineResponse{}

	params := map[string]interface{}{
		"virtual_machine[package_uuid]": packageUUID,
	}

	_, err := c.DoRequest("POST", "/virtual_machines/"+uuid, params, &vmResponse)
	if err != nil {
		return nil, err
	}

	return &vmResponse.VirtualMachine, nil
}

// DeleteVirtualMachine will destroy the VM immediately and permanently.
func (c *Client) DeleteVirtualMachine(uuid string) (bool, error) {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/destroy", nil, &body)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}

// SuspendVirtualMachine will pause the VM for later use.
func (c *Client) SuspendVirtualMachine(uuid string) (bool, error) {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/suspend", nil, &body)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}

// StartVirtualMachine will start (or restart) the VM.
func (c *Client) StartVirtualMachine(uuid string) (bool, error) {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/start", nil, &body)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}

// ShutDownVirtualMachine will shut down the VM immediately.
func (c *Client) ShutDownVirtualMachine(uuid string) (bool, error) {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/shutdown", nil, &body)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}

// PowerOffVirtualMachine will shut a VM down indefinitely.
func (c *Client) PowerOffVirtualMachine(uuid string) (bool, error) {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/power_off", nil, &body)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}

type virtualMachineResponse struct {
	VirtualMachine VirtualMachine `json:"virtual_machine"`
}

// VirtualMachine represents the details of a Contegix VM Instance.
type VirtualMachine struct {
	Name          string   `json:"name"`
	UUID          string   `json:"uuid"`
	State         string   `json:"state"`
	IPAddresses   []string `json:"ip_addresses"`
	TemplateName  string   `json:"template_name"`
	TemplateUUID  string   `json:"template_uuid"`
	PackageName   string   `json:"package_name"`
	PackageUUID   string   `json:"package_uuid"`
	ZoneName      string   `json:"zone_name"`
	ZoneUUID      string   `json:"zone_uuid"`
	VMToolsStatus *string  `json:"vm_tools_status"`
}
