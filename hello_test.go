package contegixclassic

import (
	. "gopkg.in/check.v1"
)

func (s *S) Test_SayHello(c *C) {
	testServer.Response(202, nil, "Hello from Contegix Cloud")

	hello, err := s.client.SayHello()
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(*hello, Equals, "Hello from Contegix Cloud")
}
