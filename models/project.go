package models

import "time"

type Project struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	Variables []Variable `json:"variables"`
	Artifact  Artifact   `json:"artifact"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
