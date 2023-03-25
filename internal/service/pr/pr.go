package pr

import (
	"errors"
	"fmt"
	"github/hanzhang2566/asp/global/constant"
	"github/hanzhang2566/asp/internal/model"
	"github/hanzhang2566/asp/internal/req/rPr"
	"github/hanzhang2566/asp/pkg/setting"
	"github/hanzhang2566/asp/pkg/util"
	"strings"
)

type Dto struct {
	Username           string
	Url                string
	LastCommitUsername string
	MergeRequestStatus string
	Action             string
	SourceBranch       string
	TargetBranch       string
	Title              string
}

func NewDto(c rPr.Pr) Dto {
	return Dto{
		Username:           c.User.Name,
		Url:                c.ObjectAttributes.URL,
		LastCommitUsername: c.ObjectAttributes.LastCommit.Author.Name,
		MergeRequestStatus: c.ObjectAttributes.MergeStatus,
		Action:             c.ObjectAttributes.Action,
		SourceBranch:       c.ObjectAttributes.SourceBranch,
		TargetBranch:       c.ObjectAttributes.TargetBranch,
		Title:              c.ObjectAttributes.Title,
	}
}

func (p Dto) GetJsonBody(robot *setting.ChatRobot) (model.Body, error) {
	merge := strings.EqualFold(p.MergeRequestStatus, "can_be_merged")
	attentionUsers, err := p.getAttentionUser(robot, merge)
	if err != nil {
		return model.Body{}, nil
	}
	robot.AttentionUsers = util.RemoveAttentionUsersDup(attentionUsers)
	return model.Body{
		At: model.At{
			AtMobiles: util.GetMobiles(robot.AttentionUsers),
		},
		Text: model.Text{
			Content: p.getContent(robot, merge),
		},
		Msgtype: constant.DingMsgType,
	}, nil
}

func (p Dto) getAttentionUser(robot *setting.ChatRobot, merge bool) (users []string, err error) {
	if !merge {
		users = append(users, p.Username, p.LastCommitUsername)
		return users, nil
	}
	switch p.Action {
	case "open", "reopen", "close", "merge":
		users = robot.AttentionUsers
		users = append(users, p.LastCommitUsername)
		return users, nil
	default:
		return nil, errors.New(fmt.Sprintf("不处理: %s", p.Url))
	}
}

func (p Dto) getContent(robot *setting.ChatRobot, merge bool) string {
	templates := robot.Templates
	if merge {
		tmpl, ok := templates[constant.CanMerge]
		if !ok {
			return constant.CanMergeTmpl
		}
		return fmt.Sprintf(tmpl, p.Username, p.Action,
			p.SourceBranch, p.TargetBranch,
			robot.AttentionUsers, p.Url)
	} else {
		tmpl, ok := templates[constant.CannotMerge]
		if !ok {
			return constant.CannotMergeTmpl
		}
		return fmt.Sprintf(tmpl, p.Title, robot.AttentionUsers, p.Url)
	}
}
