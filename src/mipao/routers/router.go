// @APIVersion 1.0.0
// @Title  一起约跑吧 API
// @Description 全民健身 一起约跑
package routers

import (
	"mipao/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/common",
			beego.NSInclude(
				&controllers.CommonController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
