package main

import (
	"github.com/vitreuz/hackday/api/resource"
)

type API interface {
	ListApps() ([]resource.App, error)
}

func Run(api API) ([]string, error) {
	apps, err := api.ListApps()
	if err != nil {
		return nil, err
	}

	var names []string
	for _, app := range apps {
		names = append(names, app.Name)
	}

	return names, nil
}
