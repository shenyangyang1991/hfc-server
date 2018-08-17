package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Topic struct {
	Id             int
	TopicTitle     string    `orm:"size(140)"`                   // 帖子标题
	Author         int       `orm:"index"`                       // 作者
	LikerCnt       int                                           // 点赞量
	CommentCnt     int                                           // 评论量
	SubjectId      int                                           // 话题
	TopicMedia     int                                           // 帖子媒体类型: 1为图片 2为语音
	TopicMediaUrl  string    `orm:"size(255)"`                   // 帖子媒体
	TopicMediaInfo string    `orm:"size(255)"`                   // 帖子媒体信息
	CommentList    string    `orm:"type(text)"`                  // 前3条评论
	Status         bool                                          // 状态: 0为删除 1为正常
	Created        time.Time `orm:"auto_now_add;type(datetime)"` // 创建日期
}

func (t *Topic) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(t)
}

func (t *Topic) Insert() error {
	if _, err := orm.NewOrm().Insert(t); err != nil {
		return err
	}

	return nil
}

func (t *Topic) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}

	return nil
}

// status=0 即为删除
func (t *Topic) Delete() error {
	t.Status = false
	return t.Update("Status")
}
