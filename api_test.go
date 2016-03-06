package contegixclassic

import . "gopkg.in/check.v1"

func makeClient(c *C) *Client {
	client, err := NewClient("foobartoken")

	c.Assert(err, IsNil)
	c.Assert(client.Token, Equals, "foobartoken")

	return client
}

func makeCustomClient(baseUrl string, c *C) *Client {
	client, err := NewCustomClient("foobartoken", baseUrl)

	c.Assert(err, IsNil)
	c.Assert(client.Token, Equals, "foobartoken")
	c.Assert(client.URL, Equals, baseUrl)

	return client
}

func (s *S) Test_NewGetRequest(c *C) {
	client := makeClient(c)

	body := map[string]interface{}{
		"foo": "bar",
		"baz": "bar",
	}

	req, err := client.NewRequest("GET", "/bar", body)

	c.Assert(err, IsNil)
	c.Assert(req.URL.Host, Equals, "classic.contegix.com")
	c.Assert(req.URL.Path, Equals, "/api/v1/bar")

	c.Assert(req.FormValue("auth_token"), Equals, "foobartoken")
	c.Assert(req.FormValue("foo"), Equals, "bar")
	c.Assert(req.FormValue("baz"), Equals, "bar")

	c.Assert(req.Method, Equals, "GET")
}

func (s *S) TestClient_NewPostRequest(c *C) {
	client := makeClient(c)

	body := map[string]interface{}{
		"foo": "bar",
		"baz": "bar",
	}
	req, err := client.NewRequest("POST", "/bar", body)

	c.Assert(err, IsNil)

	c.Assert(req.URL.Host, Equals, "classic.contegix.com")
	c.Assert(req.URL.Path, Equals, "/api/v1/bar")

	c.Assert(req.PostFormValue("auth_token"), Equals, "foobartoken")
	c.Assert(req.PostFormValue("foo"), Equals, "bar")
	c.Assert(req.PostFormValue("baz"), Equals, "bar")

	c.Assert(req.Method, Equals, "POST")
}

func (s *S) TestClient_NewCustomPostRequest(c *C) {
	client := makeCustomClient("http://foobar", c)

	body := map[string]interface{}{
		"foo": "bar",
		"baz": "bar",
	}

	req, err := client.NewRequest("POST", "/bar", body)

	c.Assert(err, IsNil)

	c.Assert(req.URL.Host, Equals, "foobar")
	c.Assert(req.URL.Path, Equals, "/api/v1/bar")

	c.Assert(req.PostFormValue("auth_token"), Equals, "foobartoken")
	c.Assert(req.PostFormValue("foo"), Equals, "bar")
	c.Assert(req.PostFormValue("baz"), Equals, "bar")

	c.Assert(req.Method, Equals, "POST")
}
