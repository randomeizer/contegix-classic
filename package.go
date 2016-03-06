package contegixclassic

type Package struct {
	Name       string `json:"name"`
	UUID       string `json:"uuid"`
	DiskSizeGB int    `json:"disk_size_in_gigabytes"`
	NumCPUs    int    `json:"num_cpus"`
	MemoryMB   int    `json:"memory_mb"`
}

// ListPackages returns an array of the available VM Packages.
func (c *Client) ListPackages() ([]Package, error) {
	responses := []packageResponse{}

	_, err := c.DoRequest("GET", "/packages", nil, &responses)
	if err != nil {
		return nil, err
	}

	packages := make([]Package, len(responses))
	for i, response := range responses {
		packages[i] = response.Package
	}
	return packages, nil
}

// GetPackage returns the details for the specified VM Package
func (c *Client) GetPackage(uuid string) (*Package, error) {
	response := packageResponse{}

	_, err := c.DoRequest("GET", "/packages/"+uuid, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response.Package, nil
}

type packageResponse struct {
	Package Package `json:"package"`
}
