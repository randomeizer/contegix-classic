package contegixclassic

import "fmt"

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

type CreateVirtualMachine struct {
	Name         string `form:"virtual_machine[name]"`
	ZoneUUID     string `form:"virtual_machine[zone_uuid]"`
	PackageUUID  string `form:"virtual_machine[package_uuid]"`
	TemplateUUID string `form:"virtual_machine[template_uuid]"`
}

type UpdateVirtualMachine struct {
	PackageUUID string `form:"virtual_machine[package_uuid]"`
}

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
func (c *Client) CreateVirtualMachine(details *CreateVirtualMachine) (*VirtualMachine, error) {
	vmResponse := virtualMachineResponse{}

	_, err := c.DoRequest("POST", "/virtual_machines", details, &vmResponse)
	if err != nil {
		return nil, err
	}

	return &vmResponse.VirtualMachine, nil
}

// UpdateVirtualMachine allows changing the package of a VM, which adjusts the capacity and cost.
func (c *Client) UpdateVirtualMachine(uuid string, details *UpdateVirtualMachine) (*VirtualMachine, error) {
	vmResponse := virtualMachineResponse{}

	_, err := c.DoRequest("POST", "/virtual_machines/"+uuid, details, &vmResponse)
	if err != nil {
		return nil, err
	}

	return &vmResponse.VirtualMachine, nil
}

// DeleteVirtualMachine will destroy the VM immediately and permanently.
func (c *Client) DeleteVirtualMachine(uuid string) error {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/destroy", nil, &body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Delete of Virtual Machine with a UUID of '%v' failed: %s", uuid, resp.Status)
	}
	return nil
}

// SuspendVirtualMachine will pause the VM for later use.
func (c *Client) SuspendVirtualMachine(uuid string) error {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/suspend", nil, &body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Suspend of Virtual Machine with a UUID of '%v' failed: %s", uuid, resp.Status)
	}
	return nil
}

// StartVirtualMachine will start (or restart) the VM.
func (c *Client) StartVirtualMachine(uuid string) error {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/start", nil, &body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Start of Virtual Machine with a UUID of '%v' failed: %s", uuid, resp.Status)
	}
	return nil
}

// ShutDownVirtualMachine will shut down the VM immediately.
func (c *Client) ShutDownVirtualMachine(uuid string) error {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/shutdown", nil, &body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Shut Down of Virtual Machine with a UUID of '%v' failed: %s", uuid, resp.Status)
	}
	return nil
}

// PowerOffVirtualMachine will shut a VM down indefinitely.
func (c *Client) PowerOffVirtualMachine(uuid string) error {
	var body *string

	resp, err := c.DoRequest("PUT", "/virtual_machines/"+uuid+"/power_off", nil, &body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Power Off of Virtual Machine with a UUID of '%v' failed: %s", uuid, resp.Status)
	}
	return nil
}

type virtualMachineResponse struct {
	VirtualMachine VirtualMachine `json:"virtual_machine"`
}
