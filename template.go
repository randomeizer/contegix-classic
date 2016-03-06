package contegixclassic

type Template struct {
	Name        string  `json:"name"`
	UUID        string  `json:"uuid"`
	Description *string `json:"description"`
}

// ListTemplates returns an array of the available VM templates.
func (c *Client) ListTemplates() ([]Template, error) {
	responses := []templateResponse{}

	_, err := c.DoRequest("GET", "/templates", nil, &responses)
	if err != nil {
		return nil, err
	}

	templates := make([]Template, len(responses))
	for i, response := range responses {
		templates[i] = response.Template
	}
	return templates, nil
}

// GetTemplate returns the details for the specified VM Template
func (c *Client) GetTemplate(uuid string) (*Template, error) {
	response := templateResponse{}

	_, err := c.DoRequest("GET", "/templates/"+uuid, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response.Template, nil
}

type templateResponse struct {
	Template Template `json:"template"`
}
