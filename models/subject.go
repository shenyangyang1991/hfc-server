package models

import "github.com/astaxie/beego/orm"

// 话题模型
type Subject struct {
	Id            int
	SubjectTitle string `orm:"size(14)"`   // 标题
	SubjectPoster string `orm:"size(255)"`  // 海报
	SubjectWeight int    `orm:"index"`      // 权重
	FollowerCnt   int                       // 关注量
	TopicCnt      int                       // 发帖量
	TopicList     string `orm:"type(text)"` // 前3条帖子
	Status        bool                      // 状态: 0为删除 1为正常
}

func (s *Subject) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(s)
}

func (s *Subject) Insert() error {
	if _, err := orm.NewOrm().Insert(s); err != nil {
		return err
	}

	return nil
}

func (s *Subject) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(s, fields...); err != nil {
		return err
	}

	return nil
}

// status=0 即为删除
func (s *Subject) Delete() error {
	s.Status = false
	return s.Update("Status")
}
