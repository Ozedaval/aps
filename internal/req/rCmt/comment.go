package rCmt

type Comment struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	ProjectID        int              `json:"project_id"`
	Repository       Repository       `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	MergeRequest     MergeRequest     `json:"merge_request"`
}
type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}
type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}
type ObjectAttributes struct {
	ID           int         `json:"id"`
	Note         string      `json:"note"`
	NoteableType string      `json:"noteable_type"`
	AuthorID     int         `json:"author_id"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	ProjectID    int         `json:"project_id"`
	Attachment   interface{} `json:"attachment"`
	LineCode     interface{} `json:"line_code"`
	CommitID     string      `json:"commit_id"`
	NoteableID   int         `json:"noteable_id"`
	System       bool        `json:"system"`
	StDiff       interface{} `json:"st_diff"`
	URL          string      `json:"url"`
}
type Source struct {
	Name            string `json:"name"`
	SSHURL          string `json:"ssh_url"`
	HTTPURL         string `json:"http_url"`
	WebURL          string `json:"web_url"`
	Namespace       string `json:"namespace"`
	VisibilityLevel int    `json:"visibility_level"`
}
type Target struct {
	Name            string `json:"name"`
	SSHURL          string `json:"ssh_url"`
	HTTPURL         string `json:"http_url"`
	WebURL          string `json:"web_url"`
	Namespace       string `json:"namespace"`
	VisibilityLevel int    `json:"visibility_level"`
}
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type LastCommit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}
type MergeRequest struct {
	ID              int         `json:"id"`
	TargetBranch    string      `json:"target_branch"`
	SourceBranch    string      `json:"source_branch"`
	SourceProjectID int         `json:"source_project_id"`
	AuthorID        int         `json:"author_id"`
	AssigneeID      int         `json:"assignee_id"`
	Title           string      `json:"title"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	MilestoneID     int         `json:"milestone_id"`
	State           string      `json:"state"`
	MergeStatus     string      `json:"merge_status"`
	TargetProjectID int         `json:"target_project_id"`
	Iid             int         `json:"iid"`
	Description     string      `json:"description"`
	Position        int         `json:"position"`
	LockedAt        interface{} `json:"locked_at"`
	Source          Source      `json:"source"`
	Target          Target      `json:"target"`
	LastCommit      LastCommit  `json:"last_commit"`
	WorkInProgress  bool        `json:"work_in_progress"`
}
