package models

import "time"

type DeploymentStep struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	State        string    `json:"state"`
	Result       string    `json:"result"`
	BeforeStepID string    `json:"after_step_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
