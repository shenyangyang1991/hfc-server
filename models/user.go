package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int
	Username      string    `orm:"size(11);index"`              // 帐号
	Password      string    `orm:"size(6)"`                     // 密码
	Nickname      string    `orm:"size(16);index"`              // 昵称
	AvatarUrl     string    `orm:"size(255)"`                   // 头像
	LastLoginTime time.Time `orm:"type(datetime)"`              // 最后登录日期
	LastLoginIp   string    `orm:"size(16)"`                    // 最后登录IP
	Role          int                                           // 角色
	Status        bool                                          // 状态: 0为删除 1为正常
	Created       time.Time `orm:"auto_now_add;type(datetime)"` // 创建日期
}

func (u *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}

func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}

	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}

	return nil
}

// status=0 即为删除
func (u *User) Delete() error {
	u.Status = false
	return u.Update("Status")
}
