package teamcity

type ServerLicensingData struct {
	MaxAgents         int
	AgentsLeft        int
	MaxBuildTypes     int
	BuildTypesLeft    int
	ServerLicenseType string
}

func (c client) GetServerLicensingData() (*ServerLicensingData, error) {
	response := ServerLicensingData{}
	err := c.httpGet("/server/licensingData", nil, &response)
	if err != nil {
		errorf("GetServerLicensingData failed with %s", err)
		return nil, err
	}

	return &response, nil
}
