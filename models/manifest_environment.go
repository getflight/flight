package models

type ManifestEnvironment struct {
	Name      string             `json:"name" validate:"required,max=256"`
	Databases []ManifestDatabase `json:"databases" validate:"dive"`
	Variables []ManifestVariable `json:"variables" validate:"dive"`
}
