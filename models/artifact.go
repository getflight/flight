package models

type Artifact struct {
	ID            string `json:"id"`
	CommitMessage string `json:"commit_message"`
	CommitHash    string `json:"commit_hash"`
	UploadURL     string `json:"upload_url"`
}
