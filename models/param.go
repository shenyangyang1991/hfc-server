package models

import (
	"github.com/astaxie/beego/validation"
	"strings"
)

// 登录参数
type LoginParam struct {
	Username string `valid:"Mobile" form:"username"`
	Password string `valid:"Length(4)" form:"password"`
}

//用户参数
type UserParam struct {
	Nickname  string `valid:"MinSize(1);MaxSize(16);" form:"nickname"`
	AvatarUrl string `form:"avatar_url"`
}

func (u *UserParam) Valid(v *validation.Validation) {
	if strings.Index(u.Nickname, "admin") != -1 {
		v.SetError("Name", "名称里不能含有 admin")
	}

	if strings.Index(u.Nickname, "hfc") != -1 {
		v.SetError("Name", "名称里不能含有 hfc")
	}

	if strings.Index(u.Nickname, "嗨范儿") != -1 {
		v.SetError("Name", "名称里不能含有 嗨范儿")
	}

	if strings.Index(u.Nickname, "社区") != -1 {
		v.SetError("Name", "名称里不能含有 社区")
	}

	if strings.Index(u.Nickname, "区块链") != -1 {
		v.SetError("Name", "名称里不能含有 区块链")
	}
}

type TopicImageInfo struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
