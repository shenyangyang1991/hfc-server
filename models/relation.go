package models

import "time"

type SubjectRelation struct {
	Id        int
	SubjectId int                                           // 话题
	UserId    int       `orm:"index"`                       // 用户
	Created   time.Time `orm:"auto_now_add;type(datetime)"` // 创建日期
}

type TopicRelation struct {
	Id      int
	TopicId int                                           // 帖子
	UserId  int       `orm:"index"`                       // 用户
	Created time.Time `orm:"auto_now_add;type(datetime)"` // 创建日期
}
