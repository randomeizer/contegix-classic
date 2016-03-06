package contegixclassic

import (
	"fmt"
	. "gopkg.in/check.v1"
)

func (s *S) Test_ListTemplates(c *C) {
	testServer.Response(200, nil, templateList)

	templates, err := s.client.ListTemplates()
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(templates, DeepEquals, []Template{
		Template{
			Name:        "CentOS-5-32",
			UUID:        "3e0fbcadc76351709ee654dea38c51a5",
			Description: nil,
		},
		Template{
			Name:        "CentOS-5-64",
			UUID:        "c422177956d45f9fa957b82f8c9394f9",
			Description: nil,
		},
		Template{
			Name:        "Ubuntu-810-32",
			UUID:        "a78624bd5e2d5823a003d14777fcde7b",
			Description: nil,
		},
		Template{
			Name:        "Ubuntu-810-64",
			UUID:        "51e9899f0ed656eb8ce0e1b5207217d1",
			Description: nil,
		},
	})
}

func (s *S) Test_GetTemplate(c *C) {
	testServer.Response(200, nil, template1)

	vm, err := s.client.GetTemplate("3e0fbcadc76351709ee654dea38c51a5")
	c.Assert(err, IsNil)

	_ = testServer.WaitRequest()

	c.Assert(vm, DeepEquals,
		&Template{
			Name:        "CentOS-5-32",
			UUID:        "3e0fbcadc76351709ee654dea38c51a5",
			Description: nil,
		},
	)
}

var template1 = `{"template":{
    "name":"CentOS-5-32",
    "uuid":"3e0fbcadc76351709ee654dea38c51a5",
    "description":null
  }}`

var template2 = `{"template":{
    "name":"CentOS-5-64",
    "uuid":"c422177956d45f9fa957b82f8c9394f9",
    "description":null
  }}`

var template3 = `{"template":{
    "name":"Ubuntu-810-32",
    "uuid":"a78624bd5e2d5823a003d14777fcde7b",
    "description":null
  }}`

var template4 = `{"template":{
    "name":"Ubuntu-810-64",
    "uuid":"51e9899f0ed656eb8ce0e1b5207217d1",
    "description":null
  }}`

var templateList = fmt.Sprintf("[%v, %v, %v, %v]", template1, template2, template3, template4)
