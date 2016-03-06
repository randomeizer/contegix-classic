package contegixclassic

func (c *Client) SayHello() (*string, error) {
	var body *string

	_, err := c.DoRequest("GET", "/hello", nil, &body)

	return body, err
}
