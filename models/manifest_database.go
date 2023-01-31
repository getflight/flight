package models

type ManifestDatabase struct {
	Name   string `json:"name" validate:"required,max=256"`
	Driver string `json:"driver" validate:"required,oneof=mysql postgresql"`
}
