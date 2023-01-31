package models

type ManifestVariable struct {
	Key   string `json:"key" validate:"required,max=256"`
	Value string `json:"value" validate:"required,max=256"`
}
