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

func (c *Client) GetVirtualMachine(uuid string) (*VirtualMachine, error) {
	vmResponse := virtualMachineResponse{}

	_, err := c.DoRequest("GET", fmt.Sprintf("/virtual_machines/%v", uuid), nil, &vmResponse)
	if err != nil {
		return nil, err
	}

	return &vmResponse.VirtualMachine, nil
}

func (c *Client) CreateVirtualMachine(name string, zoneUuid string, packageUuid string, templateUuid string) (*VirtualMachine, error) {
	vmResponse := virtualMachineResponse{}

	params := map[string]interface{}{
		"virtual_machine[name]":          name,
		"virtual_machine[zone_uuid]":     zoneUuid,
		"virtual_machine[package_uuid]":  packageUuid,
		"virtual_machine[template_uuid]": templateUuid,
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
