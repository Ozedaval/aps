package constant

const (
	DingMethod      = "POST"
	DingRobotPath   = "https://oapi.dingtalk.com/robot/send?access_token=%s"
	DingContentType = "application/json"
	DingMsgType     = "text"
)

const (
	Push         = "push"
	PushTmpl     = "%s在%s仓库中做了【分支推送】，共有%d次提交，请%s关注"
	Comment      = "comment"
	CommentTmpl  = "%s发表了【代码评论】：%s，请%s关注。详情：%s"
	CanMerge     = "canmerge"
	CanMergeTmpl = "%s已经%s从%s分支到%s分支的【合并请求】，请%s关注。详情：%s"

	CannotMerge     = "cannotmerge"
	CannotMergeTmpl = "%s这次【合并请求】可能存在冲突，不可以被合并，请%s关注。详情：%s"
)

const (
	PORT            = "7201"
	ContextPath     = "/api/v1"
	ApplicationName = "automation"
	ApiV1           = "/api/v1"
)
