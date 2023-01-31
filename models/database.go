package models

type Database struct {
	Name        string `json:"name"`
	Driver      string `json:"driver"`
	MinCapacity *int64 `json:"min_capacity"`
	MaxCapacity *int64 `json:"max_capacity"`
}
