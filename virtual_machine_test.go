package contegixclassic

import (
	"fmt"
	. "gopkg.in/check.v1"
)

func (s *S) Test_ListVirtualMachines(c *C) {
	testServer.Response(202, nil, vmsList)

	vms, err := s.client.ListVirtualMachines()
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(vms, DeepEquals, []VirtualMachine{
		VirtualMachine{
			Name:          "TestVM",
			UUID:          "00000000000000000000000000000000",
			State:         "powered_on",
			IPAddresses:   []string{"207.223.0.1"},
			TemplateName:  "Ubuntu-810-32",
			TemplateUUID:  "a78624bd5e2d5823a003d14777fcde7b",
			PackageName:   "$25.00 - 1 CPU, 256MB RAM 10GB HDD",
			PackageUUID:   "b20caf7301eb5d55b6a6cc25f78a1366",
			ZoneName:      "us1",
			ZoneUUID:      "36fe43313fe05e028d32c821be0ab75f",
			VMToolsStatus: nil,
		},
		VirtualMachine{
			Name:          "TestVM2",
			UUID:          "00000000000000000000000000000001",
			State:         "pending_deploy",
			IPAddresses:   []string{"207.223.0.2", "207.223.0.3"},
			TemplateName:  "CentOS-5-32",
			TemplateUUID:  "3e0fbcadc76351709ee654dea38c51a5",
			PackageName:   "$25.00 - 1 CPU, 256MB RAM 10GB HDD",
			PackageUUID:   "b20caf7301eb5d55b6a6cc25f78a1366",
			ZoneName:      "us1",
			ZoneUUID:      "36fe43313fe05e028d32c821be0ab75f",
			VMToolsStatus: nil,
		},
	})
}

func (s *S) Test_GetVirtualMachine(c *C) {
	testServer.Response(202, nil, vm1)

	vm, err := s.client.GetVirtualMachine("00000000000000000000000000000000")
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(vm, DeepEquals,
		&VirtualMachine{
			Name:          "TestVM",
			UUID:          "00000000000000000000000000000000",
			State:         "powered_on",
			IPAddresses:   []string{"207.223.0.1"},
			TemplateName:  "Ubuntu-810-32",
			TemplateUUID:  "a78624bd5e2d5823a003d14777fcde7b",
			PackageName:   "$25.00 - 1 CPU, 256MB RAM 10GB HDD",
			PackageUUID:   "b20caf7301eb5d55b6a6cc25f78a1366",
			ZoneName:      "us1",
			ZoneUUID:      "36fe43313fe05e028d32c821be0ab75f",
			VMToolsStatus: nil,
		},
	)
}

func (s *S) Test_CreateVirtualMachine(c *C) {
	testServer.Response(202, nil, vm2)

	details := &CreateVirtualMachine{
		Name:         "TestVM2",
		ZoneUUID:     "36fe43313fe05e028d32c821be0ab75f",
		PackageUUID:  "b20caf7301eb5d55b6a6cc25f78a1366",
		TemplateUUID: "a78624bd5e2d5823a003d14777fcde7b",
	}

	vm, err := s.client.CreateVirtualMachine(details)
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(vm, DeepEquals,
		&VirtualMachine{
			Name:          "TestVM2",
			UUID:          "00000000000000000000000000000001",
			State:         "pending_deploy",
			IPAddresses:   []string{"207.223.0.2", "207.223.0.3"},
			TemplateName:  "CentOS-5-32",
			TemplateUUID:  "3e0fbcadc76351709ee654dea38c51a5",
			PackageName:   "$25.00 - 1 CPU, 256MB RAM 10GB HDD",
			PackageUUID:   "b20caf7301eb5d55b6a6cc25f78a1366",
			ZoneName:      "us1",
			ZoneUUID:      "36fe43313fe05e028d32c821be0ab75f",
			VMToolsStatus: nil,
		},
	)
}

func (s *S) Test_UpdateVirtualMachine(c *C) {
	testServer.Response(200, nil, vm2)

	details := &UpdateVirtualMachine{
		PackageUUID: "b20caf7301eb5d55b6a6cc25f78a1366",
	}

	vm, err := s.client.UpdateVirtualMachine("00000000000000000000000000000001", details)
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(vm, DeepEquals,
		&VirtualMachine{
			Name:          "TestVM2",
			UUID:          "00000000000000000000000000000001",
			State:         "pending_deploy",
			IPAddresses:   []string{"207.223.0.2", "207.223.0.3"},
			TemplateName:  "CentOS-5-32",
			TemplateUUID:  "3e0fbcadc76351709ee654dea38c51a5",
			PackageName:   "$25.00 - 1 CPU, 256MB RAM 10GB HDD",
			PackageUUID:   "b20caf7301eb5d55b6a6cc25f78a1366",
			ZoneName:      "us1",
			ZoneUUID:      "36fe43313fe05e028d32c821be0ab75f",
			VMToolsStatus: nil,
		},
	)
}

func (s *S) Test_DeleteVirtualMachine(c *C) {
	testServer.Response(200, nil, "")

	err := s.client.DeleteVirtualMachine("00000000000000000000000000000001")
	c.Assert(err, IsNil)
	_ = testServer.WaitRequest()

	testServer.Response(404, nil, "")

	err = s.client.DeleteVirtualMachine("00000000000000000000000000000001")
	c.Assert(err, NotNil)
	_ = testServer.WaitRequest()
}

func (s *S) Test_SuspendVirtualMachine(c *C) {
	testServer.Response(200, nil, "")

	err := s.client.SuspendVirtualMachine("00000000000000000000000000000001")
	c.Assert(err, IsNil)
	_ = testServer.WaitRequest()

	testServer.Response(404, nil, "")

	err = s.client.SuspendVirtualMachine("00000000000000000000000000000001")
	c.Assert(err, NotNil)
	_ = testServer.WaitRequest()
}

var vm1 = `{
    "virtual_machine": {
      "template_name": "Ubuntu-810-32",
      "name":"TestVM",
      "uuid":"00000000000000000000000000000000",
      "zone_name":"us1",
      "vm_tools_status":null,
      "template_uuid":"a78624bd5e2d5823a003d14777fcde7b",
      "package_uuid":"b20caf7301eb5d55b6a6cc25f78a1366",
      "package_name":"$25.00 - 1 CPU, 256MB RAM 10GB HDD",
      "zone_uuid":"36fe43313fe05e028d32c821be0ab75f",
      "state":"powered_on",
      "ip_addresses":["207.223.0.1"]
    }
  }`

var vm2 = `{
    "virtual_machine": {
      "template_name":"CentOS-5-32",
      "name":"TestVM2",
      "uuid":"00000000000000000000000000000001",
      "zone_name":"us1",
      "vm_tools_status":null,
      "template_uuid":"3e0fbcadc76351709ee654dea38c51a5",
      "package_uuid":"b20caf7301eb5d55b6a6cc25f78a1366",
      "package_name":"$25.00 - 1 CPU, 256MB RAM 10GB HDD",
      "zone_uuid":"36fe43313fe05e028d32c821be0ab75f",
      "state":"pending_deploy",
      "ip_addresses":["207.223.0.2","207.223.0.3"]
    }
  }`

var vmsList = fmt.Sprintf(`[
  %v,
  %v
]`, vm1, vm2)
