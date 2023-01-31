package models

type Organisation struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Environments []Environment `json:"environments"`
}
