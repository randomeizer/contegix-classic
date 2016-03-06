package contegixclassic

type Zone struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// ListZones returns an array of the available hosting zones.
func (c *Client) ListZones() ([]Zone, error) {
	responses := []zoneResponse{}

	_, err := c.DoRequest("GET", "/zones", nil, &responses)
	if err != nil {
		return nil, err
	}

	zones := make([]Zone, len(responses))
	for i, response := range responses {
		zones[i] = response.Zone
	}
	return zones, nil
}

// GetZone returns the details for the specified hosting zone
func (c *Client) GetZone(uuid string) (*Zone, error) {
	response := zoneResponse{}

	_, err := c.DoRequest("GET", "/zones/"+uuid, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response.Zone, nil
}

type zoneResponse struct {
	Zone Zone `json:"zone"`
}
