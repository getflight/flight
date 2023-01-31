package models

import "time"

type User struct {
	ID            string         `json:"id"`
	Email         string         `json:"email"`
	Organisations []Organisation `json:"organisations,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}
