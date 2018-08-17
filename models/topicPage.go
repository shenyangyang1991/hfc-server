package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type TopicPage struct {
	Id             int
	TopicTitle     string    // 帖子标题
	Author         int       // 作者
	AuthorName     string    // 作者名称
	AuthorAvatar   string    // 作者头像
	LikerId        int       // 是否点赞
	LikerCnt       int       // 点赞量
	CommentCnt     int       // 评论量
	SubjectId      int       // 话题
	SubjectTitle   string    // 话题标题
	TopicMedia     int       // 帖子媒体类型: 1为图片 2为语音
	TopicMediaUrl  string    // 帖子媒体
	TopicMediaInfo string    // 帖子媒体信息
	CommentList    string    // 前3条评论
	Created        time.Time // 创建日期
}

func (t *TopicPage) QueryBuilder() orm.QueryBuilder {
	builder, _ := orm.NewQueryBuilder("mysql")
	return builder.
		Select("topic.id",
		"topic.topic_title",
		"topic.author",
		"user.nickname as author_name",
		"user.avatar_url as author_avatar",
		"topic_relation.id as liker_id",
		"topic.liker_cnt",
		"topic.comment_cnt",
		"topic.subject_id",
		"subject.subject_title",
		"topic.topic_media",
		"topic.topic_media_url",
		"topic.topic_media_info",
		"topic.comment_list",
		"topic.created")
}

func (t *TopicPage) GetUserTopics(offset int) []*TopicPage {
	var (
		list = make([]*TopicPage, 0)
	)

	sql := t.QueryBuilder().
		From("topic").
		InnerJoin("user").On("user.id = topic.author").
		LeftJoin("subject").On("subject.id = topic.subject_id").
		LeftJoin("topic_relation").On("topic_relation.topic_id = topic.id").
		Where("topic.author = ?").
		OrderBy("topic.created").Desc().Limit(10).Offset(offset).String()

	o := orm.NewOrm()
	if _, err := o.Raw(sql, t.Author).QueryRows(&list); err != nil {
		beego.Error(err)
	}

	return list
}

func (t *TopicPage) GetMyTopics(offset int) []*TopicJson {
	var list = make([]*TopicJson, 0)

	data := t.GetUserTopics(offset)
	if len(data) > 0 {
		for _, item := range data {
			topicJson := TopicJson{}
			topicJson.Id = item.Id
			topicJson.TopicTitle = item.TopicTitle
			topicJson.Author = item.Author
			topicJson.AuthorAvatar = item.AuthorAvatar
			topicJson.AuthorName = item.AuthorName
			topicJson.SubjectId = item.SubjectId
			topicJson.SubjectTitle = item.SubjectTitle
			topicJson.Liker = item.LikerId > 0
			topicJson.CommentCnt = item.CommentCnt
			topicJson.LikerCnt = item.LikerCnt
			topicJson.Created = item.Created

			if item.TopicMedia == 1 {
				topicJson.TopicPoster = item.TopicMediaUrl
				if item.TopicMediaInfo != "" {
					imageInfo := TopicImageInfo{}
					json.Unmarshal([]byte(item.TopicMediaInfo), &imageInfo)
					topicJson.PosterWidth = imageInfo.Width
					topicJson.PosterHeight = imageInfo.Height
				}
			} else if item.TopicMedia == 2 {
				topicJson.TopicVoice = item.TopicMediaUrl
			}

			if item.CommentList != "" {
				var comments = make([]*CommentJson, 0)
				json.Unmarshal([]byte(item.CommentList), &comments)
				topicJson.CommentList = comments
			}

			list = append(list, &topicJson)
		}
	}

	return list
}
