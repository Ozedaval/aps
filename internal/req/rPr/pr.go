package rPr

type Pr struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
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
type ObjectAttributes struct {
	ID              int         `json:"id"`
	TargetBranch    string      `json:"target_branch"`
	SourceBranch    string      `json:"source_branch"`
	SourceProjectID int         `json:"source_project_id"`
	AuthorID        int         `json:"author_id"`
	AssigneeID      int         `json:"assignee_id"`
	Title           string      `json:"title"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	StCommits       interface{} `json:"st_commits"`
	StDiffs         interface{} `json:"st_diffs"`
	MilestoneID     interface{} `json:"milestone_id"`
	State           string      `json:"state"`
	MergeStatus     string      `json:"merge_status"`
	TargetProjectID int         `json:"target_project_id"`
	Iid             int         `json:"iid"`
	Description     string      `json:"description"`
	Source          Source      `json:"source"`
	Target          Target      `json:"target"`
	LastCommit      LastCommit  `json:"last_commit"`
	WorkInProgress  bool        `json:"work_in_progress"`
	URL             string      `json:"url"`
	Action          string      `json:"action"`
}
