package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id           int
	CommentTitle string    `orm:"size(140)"`                   // 评论标题
	Author       int                                           // 作者
	Reply        int                                           // 回复
	Status       bool                                          // 状态: 0为删除 1为正常
	TopicId      int                                           // 帖子
	Created      time.Time `orm:"auto_now_add;type(datetime)"` // 创建日期
}

func (c *Comment) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}

	return nil
}

func (c *Comment) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}

	return nil
}

// status=0 即为删除
func (c *Comment) Delete() error {
	c.Status = false
	return c.Update("Status")
}
