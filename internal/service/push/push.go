package push

import (
	"fmt"
	"github/hanzhang2566/asp/global/constant"
	"github/hanzhang2566/asp/internal/model"
	"github/hanzhang2566/asp/internal/req/rPush"
	"github/hanzhang2566/asp/pkg/setting"
	"github/hanzhang2566/asp/pkg/util"
)

type Dto struct {
	Username          string
	RepoName          string
	TotalCommitsCount int
}

func NewDto(p rPush.Push) Dto {
	return Dto{
		Username:          p.UserName,
		RepoName:          p.Repository.Name,
		TotalCommitsCount: p.TotalCommitsCount,
	}
}

func (d Dto) GetJsonBody(robot *setting.ChatRobot) (model.Body, error) {
	robot.AttentionUsers = util.RemoveAttentionUsersDup(robot.AttentionUsers)
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

func (d Dto) getContent(robot *setting.ChatRobot) string {
	templates := robot.Templates
	tmpl, ok := templates[constant.Push]
	if !ok {
		return constant.PushTmpl
	}

	return fmt.Sprintf(tmpl, d.Username, d.RepoName, d.TotalCommitsCount, robot.AttentionUsers)
}
