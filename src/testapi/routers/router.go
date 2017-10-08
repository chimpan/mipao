// @APIVersion 1.0.0
// @Title beego cyp的测试 API
// @Description 第一次写api接口&文档，有点小激动
package routers

import (
	"testapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/userinfo",
			beego.NSInclude(
				&controllers.UserInfoController{},
			),
		),
		beego.NSNamespace("/news",
			beego.NSInclude(
				&controllers.NewsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
