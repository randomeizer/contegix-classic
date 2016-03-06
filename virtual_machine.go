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
