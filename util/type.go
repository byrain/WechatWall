package util

type User struct {
	UserID    string
	AvatarUrl string
}

type Text struct {
	FromUserID string `json:"UserID"`
	AvatarUrl  string `json:"Avatar"`
	CreateTime int64  `json:"CreateTime"`
	MsgId      int64  `json:"MsgId"`   // 消息id, 64位整型
	Content    string `json:"Content"` // 文本消息内容
}
