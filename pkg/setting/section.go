package setting

type ChatRobot struct {
	RepoName       string
	Token          string
	AttentionUsers []string
	Templates      map[string]string
}
