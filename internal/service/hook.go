package service

import (
	"github.com/gin-gonic/gin"
	"github/hanzhang2566/asp/internal/req/rCmt"
	"github/hanzhang2566/asp/internal/req/rPr"
	"github/hanzhang2566/asp/internal/req/rPush"
	"github/hanzhang2566/asp/internal/service/comment"
	"github/hanzhang2566/asp/internal/service/pr"
	"github/hanzhang2566/asp/internal/service/push"
	"github/hanzhang2566/asp/pkg/util"
	"net/http"
)

func Push(c *gin.Context) {
	p := rPush.Push{}
	err := c.ShouldBind(&p)
	if err != nil {
		panic(err)
	}
	util.CallbackPrint(p)
	chatRobot, err := util.GetChatRobot(c.Param("repo"))
	if err != nil {
		panic(err)
	}
	dto := push.NewDto(p)
	body, err := dto.GetJsonBody(chatRobot)
	if err != nil {
		panic(err)
	}
	err = util.NotifyDingTalk(chatRobot.Token, body)
	if err != nil {
		panic(err)
	}
}

func Comment(c *gin.Context) {
	cmt := rCmt.Comment{}
	err := c.ShouldBind(&cmt)
	if err != nil {
		panic(err)
	}
	util.CallbackPrint(cmt)
	chatRobot, err := util.GetChatRobot(c.Param("repo"))
	if err != nil {
		panic(err)
	}
	dto := comment.NewDto(cmt)
	body, err := dto.GetJsonBody(chatRobot)
	if err != nil {
		panic(err)
	}
	err = util.NotifyDingTalk(chatRobot.Token, body)
	if err != nil {
		panic(err)
	}
}

func Pr(c *gin.Context) {
	p := rPr.Pr{}
	err := c.ShouldBind(&p)
	if err != nil {
		panic(err)
	}
	util.CallbackPrint(p)
	chatRobot, err := util.GetChatRobot(c.Param("repo"))
	if err != nil {
		panic(err)
	}
	dto := pr.NewDto(p)
	body, err := dto.GetJsonBody(chatRobot)
	if err != nil {
		panic(err)
	}
	err = util.NotifyDingTalk(chatRobot.Token, body)
	if err != nil {
		panic(err)
	}
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
