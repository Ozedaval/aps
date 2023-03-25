package model

type Body struct {
	At      At     `json:"at"`
	Text    Text   `json:"text"`
	Msgtype string `json:"msgtype"`
}
type At struct {
	AtMobiles []string `json:"atMobiles"`
}
type Text struct {
	Content string `json:"content"`
}
