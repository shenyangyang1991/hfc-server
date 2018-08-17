package main

import (
	_ "h/routers"
	"github.com/astaxie/beego"
	"h/controllers"
)

func main() {
	// 设置日志文件
	beego.SetLogger(
		"multifile",
		`{"filename":"hfc.log","separate":["error","info","warning","notice"]}`,
	)
	// 开启行号输出
	beego.SetLogFuncCall(true)
	// 删除控制输出
	beego.BeeLogger.DelLogger("console")
	// 错误处理
	beego.ErrorController(&controllers.ApiErrorController{})
	// 开启服务
	beego.Run()
}
