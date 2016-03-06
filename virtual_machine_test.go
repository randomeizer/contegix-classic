package contegixclassic

import (
	. "gopkg.in/check.v1"
)

func (s *S) Test_ListVirtualMachines(c *C) {
	testServer.Response(202, nil, vmsExample)

	vms, err := s.client.ListVirtualMachines()
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(vms, DeepEquals, []VirtualMachine{
		VirtualMachine{
			"TestVM",
			"00000000000000000000000000000000",
			"powered_on",
			[]string{"207.223.0.1"},
			"Ubuntu-810-32",
			"a78624bd5e2d5823a003d14777fcde7b",
			"$25.00 - 1 CPU, 256MB RAM 10GB HDD",
			"b20caf7301eb5d55b6a6cc25f78a1366",
			"us1",
			"36fe43313fe05e028d32c821be0ab75f",
			nil,
		},
		VirtualMachine{
			"TestVM2",
			"00000000000000000000000000000001",
			"powered_on",
			[]string{"207.223.0.2", "207.223.0.3"},
			"CentOS-5-32",
			"3e0fbcadc76351709ee654dea38c51a5",
			"$25.00 - 1 CPU, 256MB RAM 10GB HDD",
			"b20caf7301eb5d55b6a6cc25f78a1366",
			"us1",
			"36fe43313fe05e028d32c821be0ab75f",
			nil,
		},
	})
}

var vmsExample = `[
  {
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
  },
  {
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
      "state":"powered_on",
      "ip_addresses":["207.223.0.2","207.223.0.3"]
    }
  }
]`
