package contegixclassic

import (
	"fmt"
	. "gopkg.in/check.v1"
)

func (s *S) Test_ListPackages(c *C) {
	testServer.Response(200, nil, packageList)

	packages, err := s.client.ListPackages()
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(packages, DeepEquals, []Package{
		Package{
			Name:       "$75.00 - 1 CPU, 1024MB RAM 40GB HDD",
			UUID:       "d92570500fad55bdbf7d14b558332e41",
			DiskSizeGB: 40,
			NumCPUs:    1,
			MemoryMB:   1024,
		},
		Package{
			Name:       "$135.00 - 2 CPUs, 2048MB RAM 80GB HDD",
			UUID:       "59dca474f32a54a384316516de190aab",
			DiskSizeGB: 80,
			NumCPUs:    2,
			MemoryMB:   2048,
		},
		Package{
			Name:       "$40.00 - 1 CPU, 512MB RAM 20GB HDD",
			UUID:       "bdcb5ebdb7c753ea90156b110d0195ca",
			DiskSizeGB: 20,
			NumCPUs:    1,
			MemoryMB:   512,
		},
	})
}

func (s *S) Test_GetPackage(c *C) {
	testServer.Response(200, nil, package1)

	pkg, err := s.client.GetPackage("d92570500fad55bdbf7d14b558332e41")
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(pkg, DeepEquals,
		&Package{
			Name:       "$75.00 - 1 CPU, 1024MB RAM 40GB HDD",
			UUID:       "d92570500fad55bdbf7d14b558332e41",
			DiskSizeGB: 40,
			NumCPUs:    1,
			MemoryMB:   1024,
		},
	)
}

var package1 = `{"package":{
  "disk_size_in_gigabytes":40,
  "name":"$75.00 - 1 CPU, 1024MB RAM 40GB HDD",
  "uuid":"d92570500fad55bdbf7d14b558332e41",
  "num_cpus":1,
  "memory_mb":1024
}}`

var package2 = `{"package":{
  "disk_size_in_gigabytes":80,
  "name":"$135.00 - 2 CPUs, 2048MB RAM 80GB HDD",
  "uuid":"59dca474f32a54a384316516de190aab",
  "num_cpus":2,
  "memory_mb":2048
}}`

var package3 = `{"package":{
  "disk_size_in_gigabytes":20,
  "name":"$40.00 - 1 CPU, 512MB RAM 20GB HDD",
  "uuid":"bdcb5ebdb7c753ea90156b110d0195ca",
  "num_cpus":1,
  "memory_mb":512
}}`

var packageList = fmt.Sprintf("[%v, %v, %v]", package1, package2, package3)
