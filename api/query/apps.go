package query

import "github.com/vitreuz/hackday/api/resource"

type Apps struct {
	GUIDs             []resource.UUID
	Name              []string
	SpaceGUIDS        []resource.UUID
	OrganizationGUIDS []resource.UUID
	Page              int
	PerPage           int
	OrderBy           string
	Include           string
}

func ByNames(names ...string) func(*Apps) {
	return func(a *Apps) { a.Name = names }
}
