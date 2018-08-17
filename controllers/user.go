package controllers

import (
	"h/util"
	"github.com/astaxie/beego"
	"h/models"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type UserController struct {
	baseController
}

func (this *UserController) URLMapping() {
	this.Mapping("ValidNickname", this.ValidNickname)
	this.Mapping("UpdateUser", this.UpdateUser)
}

func (this *UserController) NestPrepare() {
	if !this.isLogin {
		this.HttpError(util.TOKEN_INVALID)
	}
}

// @router / [post]
func (this *UserController) UpdateUser() {
	// 获取请求参数
	userdata := models.UserParam{}
	if err := this.ParseForm(&userdata); err != nil {
		beego.Error(err)
		this.HttpError(util.BAD_PARAM)
	}

	// 验证参数
	valid := validation.Validation{}
	pass, _ := valid.Valid(&userdata)
	if !pass {
		beego.Error(valid.Errors)

		for _, err := range valid.Errors {
			if err.Key == "Name" {
				this.HttpError(util.USER_INVALID_NAME)
				break
			} else {
				this.HttpError(util.USER_BAD_NAME)
				break
			}
		}
	}

	// 更新用户
	var (
		user         *models.User = new(models.User)
		responseJson              = models.UserJson{}
	)

	user.Nickname = userdata.Nickname
	if err := user.Update("Nickname"); err != nil {
		this.HttpError(util.USER_ERROR_UPDATE)
	}

	responseJson.Nickname = user.Nickname
	responseJson.AvatarUrl = user.AvatarUrl

	this.Result(responseJson)
}

// @router /exist [post]
func (this *UserController) ValidNickname() {
	// 获取请求参数
	userdata := models.UserParam{}
	if err := this.ParseForm(&userdata); err != nil {
		beego.Error(err)
		this.HttpError(util.BAD_PARAM)
	}

	// 验证参数
	valid := validation.Validation{}
	pass, _ := valid.Valid(&userdata)
	if !pass {
		beego.Error(valid.Errors)

		for _, err := range valid.Errors {
			if err.Key == "Name" {
				this.HttpError(util.USER_INVALID_NAME)
				break
			} else {
				this.HttpError(util.BAD_PARAM)
				break
			}
		}
	}

	// 验证是否存在
	var (
		user         *models.User = new(models.User)
		responseJson              = models.NicknameStateJson{}
	)

	responseJson.NicknameState = user.Query().Filter("Nickname", userdata.Nickname).Exist()

	this.Result(responseJson)
}

// @router /topics [get]
func (this *UserController) UserTopic() {
	var (
		topicPage    *models.TopicPage = new(models.TopicPage)
		responseJson                   = models.TopicListJson{}
		offset       int
	)

	pageParam := this.Input().Get("offset")
	page, _ := strconv.Atoi(pageParam)

	valid := validation.Validation{}
	valid.Min(page, 1, "offset")

	if valid.HasErrors() {
		beego.Error(valid.Errors)
		this.HttpError(util.BAD_PARAM)
	}

	offset = (page - 1) * 10

	topicPage.Author = this.user.Id
	list := topicPage.GetMyTopics(offset)

	size := len(list)
	if size > 0 {
		responseJson.List = list
		responseJson.Next = size == 10
	}

	responseJson.Page = page

	this.Result(responseJson)
}
