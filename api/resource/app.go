package resource

import "time"

type App struct {
	Name      string            `json:"name"`
	State     string            `json:"state"`
	Lifecycle Lifecycle         `json:"lifecycle"`
	GUID      UUID              `json:"guid"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Space     ToOneRelationship `json:"space"`
}
