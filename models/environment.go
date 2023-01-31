package models

import "time"

type Environment struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Projects  []Project  `json:"projects"`
	Databases []Database `json:"databases"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
