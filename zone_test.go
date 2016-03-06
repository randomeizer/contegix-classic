package contegixclassic

import (
	"fmt"
	. "gopkg.in/check.v1"
)

func (s *S) Test_ListZones(c *C) {
	testServer.Response(200, nil, zoneList)

	zones, err := s.client.ListZones()
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(zones, DeepEquals, []Zone{
		Zone{
			Name: "us1",
			UUID: "36fe43313fe05e028d32c821be0ab75f",
		},
	})
}

func (s *S) Test_GetZone(c *C) {
	testServer.Response(200, nil, zone1)

	zone, err := s.client.GetZone("3e0fbcadc76351709ee654dea38c51a5")
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(zone, DeepEquals,
		&Zone{
			Name: "us1",
			UUID: "36fe43313fe05e028d32c821be0ab75f",
		},
	)
}

var zone1 = `{"zone":{"name":"us1","uuid":"36fe43313fe05e028d32c821be0ab75f"}}`

var zoneList = fmt.Sprintf("[%v]", zone1)
