package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser") // 数据库用户
	mysqlpass := beego.AppConfig.String("mysqlpass") // 数据库密码
	mysqlurls := beego.AppConfig.String("mysqlurls") // 数据库链接
	mysqldb := beego.AppConfig.String("mysqldb")     // 数据库名
	// 设置数据库
	orm.RegisterModel(new(Subject), new(Topic), new(Comment), new(SubjectRelation), new(TopicRelation), new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(
		"default",
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", mysqluser, mysqlpass, mysqlurls, mysqldb),
	)
	// 开启构建数据表
	orm.RunSyncdb("default", false, true)
	// 开启日志
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
