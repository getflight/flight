package models

type Manifest struct {
	Name         string                `json:"name" validate:"required,max=256"`
	Files        *[]string             `json:"files"`
	Trigger      string                `json:"trigger" validate:"required,oneof=gateway queue"`
	Environments []ManifestEnvironment `json:"environments" validate:"required,gt=0,dive"`
}
