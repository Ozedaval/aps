package comment

import (
	"errors"
	"fmt"
	"github/hanzhang2566/asp/global/constant"
	"github/hanzhang2566/asp/internal/model"
	"github/hanzhang2566/asp/internal/req/rCmt"
	"github/hanzhang2566/asp/pkg/setting"
	"github/hanzhang2566/asp/pkg/util"
	"strings"
)

type Dto struct {
	Username           string
	Note               string
	Url                string
	LastCommitUsername string
}

func NewDto(c rCmt.Comment) Dto {
	return Dto{
		Username:           c.User.Name,
		Note:               c.ObjectAttributes.Note,
		Url:                c.ObjectAttributes.URL,
		LastCommitUsername: c.MergeRequest.LastCommit.Author.Name,
	}
}

func (d Dto) GetJsonBody(robot *setting.ChatRobot) (model.Body, error) {
	attentionUsers, err := d.getAttentionUser()
	if err != nil {
		return model.Body{}, err
	}
	robot.AttentionUsers = util.RemoveAttentionUsersDup(attentionUsers)
	return model.Body{
		At: model.At{
			AtMobiles: util.GetMobiles(robot.AttentionUsers),
		},
		Text: model.Text{
			Content: d.getContent(robot),
		},
		Msgtype: constant.DingMsgType,
	}, nil
}

func (d Dto) getAttentionUser() (users []string, err error) {
	criticUsername := d.Username
	commitUsername := d.LastCommitUsername
	note := d.Note
	if !strings.EqualFold(criticUsername, commitUsername) && !strings.HasPrefix(note, "@") {
		users = append(users, commitUsername)
		return users, nil
	} else if strings.HasPrefix(note, "@") {
		users = append(users, util.GetAttentionUsername(note))
		return users, nil
	} else {
		return nil, errors.New(fmt.Sprintf("不处理: %s", d.Url))
	}
}

func (d Dto) getContent(robot *setting.ChatRobot) string {
	templates := robot.Templates
	tmpl, ok := templates[constant.Comment]
	if !ok {
		return constant.CommentTmpl
	}
	return fmt.Sprintf(tmpl, d.Username, d.Note, robot.AttentionUsers, d.Url)
}
