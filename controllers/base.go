package controllers

import (
	"github.com/astaxie/beego"
	"h/util"
	"h/models"
)

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
	user    util.UserSession
	isLogin bool
}

func (this *baseController) Prepare() {
	// 认证
	err := this.user.TokenAuthenticator(this.Ctx.Input.Header("Authorization"))
	if err != nil {
		this.isLogin = false
	} else {
		this.isLogin = true
	}

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (this *baseController) HttpError(code int) {
	responseJson := models.ResponseJson{
		Code:    code,
		Success: false,
		Message: util.GetMsg(code),
		Data:    nil,
	}

	this.Data["json"] = &responseJson
	this.ServeJSON()
	this.StopRun()
}

func (this *baseController) Result(data interface{}) {
	responseJson := models.ResponseJson{
		Code:    util.SUCCESS,
		Success: true,
		Message: util.GetMsg(util.SUCCESS),
		Data:    data,
	}

	this.Data["json"] = &responseJson
	this.ServeJSON()
}

func (this *baseController) getClientIp() string {
	return this.Ctx.Input.IP()
}
