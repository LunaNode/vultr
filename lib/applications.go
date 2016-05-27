package lib

// application image on Vultr
type Application struct {
	ID         int    `json:"APPID,string"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	DeployName string `json:"deploy_name"`
	Surcharge  int    `json:"surcharge"`
}

func (c *Client) GetApplications() ([]Application, error) {
	var appMap map[string]Application
	if err := c.get(`app/list`, &appMap); err != nil {
		return nil, err
	}

	var appList []Application
	for _, app := range appMap {
		appList = append(appList, app)
	}
	return appList, nil
}
