package contegixclassic

type Template struct {
	Name        string  `json:"name"`
	UUID        string  `json:"uuid"`
	Description *string `json:"description"`
}

func (c *Client) ListTemplates() ([]Template, error) {
	templateResponses := []templateResponse{}

	_, err := c.DoRequest("GET", "/templates", nil, &templateResponses)
	if err != nil {
		return nil, err
	}

	templates := make([]Template, len(templateResponses))
	for i, templateResponse := range templateResponses {
		templates[i] = templateResponse.Template
	}
	return templates, nil
}

type templateResponse struct {
	Template Template `json:"template"`
}
