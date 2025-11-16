package model

type CloneGitRequest struct {
	URL string `json:"url" validate:"required,url"`
}
