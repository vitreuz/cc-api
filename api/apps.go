package api

import (
	"encoding/json"

	"github.com/vitreuz/hackday/api/resource"
)

const appsEndpoint = "/v3/apps"

type Apps struct {
	*conn
}

func (a Apps) ListApps() ([]resource.App, error) {
	resp, err := a.get(appsEndpoint)
	if err != nil {
		return nil, err
	}

	var apps struct {
		Resources []resource.App
	}
	if err := json.NewDecoder(resp.Body).Decode(&apps); err != nil {
		return nil, err
	}

	return apps.Resources, nil
}

func (Apps) CreateApp() (resource.App, error) {
	// var app struct {
	// 	Name                 string                     `json:"name"`
	// 	Space                resource.ToOneRelationship `json:"relationship"`
	// 	EnvironmentVarialbes resource.Object            `json:"environment_varialbes,omitempty"`
	// 	Lifecycle            resource.Lifecycle         `json:"lifecycle,omitempty"`
	// }
	return resource.App{}, nil
}

func (Apps) GetApp() (resource.App, error) {
	return resource.App{}, nil
}
