package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github/hanzhang2566/asp/global"
	"github/hanzhang2566/asp/global/constant"
	"github/hanzhang2566/asp/internal/model"
	"github/hanzhang2566/asp/pkg/setting"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

func GetChatRobot(repoName string) (*setting.ChatRobot, error) {
	for _, chatRobot := range global.ChatRobots {
		if strings.EqualFold(repoName, chatRobot.RepoName) {
			return &chatRobot, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("chatRobot[%s] not in server.", repoName))
}

func NotifyDingTalk(token string, body model.Body) error {
	if strings.Contains(body.Text.Content, "%") {
		return errors.New(fmt.Sprintf("notify failure. body: %s", body))
	}
	marshal, err := json.Marshal(body)
	if err != nil {
		return err
	}
	dingHookPath := fmt.Sprintf(constant.DingRobotPath, token)
	log.Printf("\n[path]: %s\n[body]:%s\n", dingHookPath, string(marshal))
	request, err := http.NewRequest(constant.DingMethod, dingHookPath, bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", constant.DingContentType)
	response, err := http.DefaultClient.Do(request)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		log.Println(err)
	}(response.Body)
	if err != nil {
		return err
	}
	log.Printf("\n[response.status]: %s\n", response.Status)
	return nil
}

func GetAttentionUsername(content string) string {
	start := strings.Index(content, "@")
	end := strings.Index(content, " ")
	return content[start+1 : end]
}

func RemoveAttentionUsersDup(list []string) []string {
	sort.Strings(list)
	i := 0
	ret := []string{""}
	for j := 0; j < len(list); j++ {
		if strings.Compare(ret[i], list[j]) == -1 {
			ret = append(ret, list[j])
			i++
		}
	}

	return ret[1:]
}

func SliceString(slice []string) string {
	builder := strings.Builder{}
	for index, s := range slice {
		builder.WriteString(s)
		if index != len(slice)-1 {
			builder.WriteString("、")
		}
	}
	return builder.String()
}

func CallbackPrint(c interface{}) {
	marshal, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	log.Printf("callback: %s\n", string(marshal))
}

func GetMobiles(usersName []string) (mobiles []string) {
	for _, username := range usersName {
		if mobile, ok := global.UsersMobile[username]; ok {
			mobiles = append(mobiles, mobile)
		}
	}
	return mobiles
}
