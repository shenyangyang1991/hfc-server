package controllers

import (
	"h/models"
	"h/util"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoginController struct {
	baseController
}

func (this *LoginController) URLMapping() {
	this.Mapping("Login", this.Login)
}

// @router / [post]
func (this *LoginController) Login() {
	// 获取请求参数
	login := models.LoginParam{}
	if err := this.ParseForm(&login); err != nil {
		beego.Error(err)
		this.HttpError(util.LOGIN_BAD_PARAM)
	}
	// 验证参数
	valid := validation.Validation{}
	pass, _ := valid.Valid(&login)
	if !pass {
		beego.Error(valid.Errors)
		this.HttpError(util.LOGIN_INVALID_PARAM)
	}

	// 验证短信验证码
	if login.Password != "1111" {
		this.HttpError(util.LOGIN_INVALID_SMS)
	}

	// 登录
	var (
		user         *models.User      = new(models.User)
		responseJson *models.LoginJson = new(models.LoginJson)
		state        bool
	)

	user.Username = login.Username
	err := user.Read("Username")
	if err != nil && err != orm.ErrNoRows {
		beego.Error(err)
		this.HttpError(util.LOGIN_ERROR)
	}

	state = user.Id > 0

	if user.Id > 0 {
		if !user.Status {
			this.HttpError(util.LOGIN_INVALID_USER)
		}

		user.LastLoginIp = this.getClientIp()
		user.LastLoginTime = time.Now()
		err := user.Update("LastLoginIp", "LastLoginTime")
		if err != nil {
			beego.Error(err)
			this.HttpError(util.LOGIN_ERROR)
		}
	} else {
		user.Password = ""
		user.Nickname = util.GenRandNickname()
		user.AvatarUrl = ""
		user.Status = true
		user.Role = 0
		user.LastLoginTime = time.Now()
		user.LastLoginIp = this.getClientIp()

		err := user.Insert()
		if err != nil {
			beego.Error(err)
			this.HttpError(util.REGISTER_ERROR)
		}
	}

	this.user.Id = user.Id
	this.user.Nickname = user.Nickname
	this.user.AvatarUrl = user.AvatarUrl

	responseJson.AccessToken = this.user.TokenGenerator()
	responseJson.UserState = state

	this.Result(responseJson)
}
