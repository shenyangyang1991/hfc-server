package controllers

import (
	"h/util"
	"github.com/astaxie/beego"
	"h/models"
)

type ApiErrorController struct {
	beego.Controller
}

func (this *ApiErrorController) Error404() {
	this.HttpError(util.NOT_FOUND)
}

func (this *ApiErrorController) Error500() {
	this.HttpError(util.ERROR)
}

func (this *ApiErrorController) HttpError(code int) {
	responseJson := models.ResponseJson{
		Code:    code,
		Success: false,
		Message: util.GetMsg(code),
		Data:    nil,
	}

	this.Data["json"] = &responseJson
	this.ServeJSON()
}
