package models

type Deployment struct {
	ID          string           `json:"id"`
	State       string           `json:"state"`
	Artifact    string           `json:"artifact"`
	Environment string           `json:"environment"`
	Count       string           `json:"count"`
	Manifest    Manifest         `json:"manifest"`
	Steps       []DeploymentStep `json:"steps"`
}
