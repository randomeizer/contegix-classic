package contegixclassic

func (c *Client) SayHello() (*string, error) {
	var body *string

	_, err := c.DoRequest("GET", "/virtual_machines", nil, &body)

	return body, err
}
