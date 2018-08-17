package models

import "time"

// 响应结构
type ResponseJson struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 登录响应结构
type LoginJson struct {
	AccessToken string `json:"access_token"`
	UserState   bool   `json:"user_state"`
}

// 检查昵称可用
type NicknameStateJson struct {
	NicknameState bool `json:"nickname_state"`
}

// 用户响应结构
type UserJson struct {
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
}

// 帖子结构
type TopicJson struct {
	Id           int            `json:"topic_id"`
	TopicTitle   string         `json:"topic_context"`
	TopicVoice   string         `json:"topic_voice"`
	TopicPoster  string         `json:"topic_poster"`
	PosterWidth  int            `json:"poster_width"`
	PosterHeight int            `json:"poster_height"`
	Liker        bool           `json:"liker"`
	Author       int            `json:"user_id"`
	AuthorName   string         `json:"nickname"`
	AuthorAvatar string         `json:"avatar_url"`
	LikerCnt     int            `json:"like_cnt"`
	CommentCnt   int            `json:"comment_cnt"`
	SubjectId    int            `json:"subject_id"`
	SubjectTitle string         `json:"subject_title"`
	CommentList  []*CommentJson `json:"comment_list"`
	Created      time.Time      `json:"created"`
}

// 帖子列表
type TopicListJson struct {
	List []*TopicJson `json:"list"`
	Page int          `json:"page"`
	Next bool         `json:"next"`
}

type CommentJson struct {
	Id           int       `json:"comment_id"`
	CommentTitle string    `json:"comment_context"`
	Author       int       `json:"user_id"`
	AuthorName   string    `json:"nickname"`
	AuthorAvatar string    `json:"avatar_url"`
	Reply        int       `json:"reply_id"`
	ReplyName    string    `json:"reply_name"`
	ReplyAvatar  string    `json:"reply_avatar"`
	TopicId      int       `json:"topic_id"`
	Created      time.Time `json:"created"`
}
