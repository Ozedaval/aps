package rPush

import (
	"time"
)

type Push struct {
	ObjectKind        string     `json:"object_kind"`
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	UserID            int        `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserEmail         string     `json:"user_email"`
	ProjectID         int        `json:"project_id"`
	Repository        Repository `json:"repository"`
	Commits           []Commits  `json:"commits"`
	TotalCommitsCount int        `json:"total_commits_count"`
}
type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Commits struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Author    Author    `json:"author"`
}
