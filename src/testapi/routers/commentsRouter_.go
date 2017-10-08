package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["testapi/controllers:NewsController"] = append(beego.GlobalControllerRouter["testapi/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Publish",
			Router: `/publish`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testapi/controllers:NewsController"] = append(beego.GlobalControllerRouter["testapi/controllers:NewsController"],
		beego.ControllerComments{
			Method: "All",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testapi/controllers:NewsController"] = append(beego.GlobalControllerRouter["testapi/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Commend",
			Router: `/commend`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testapi/controllers:UserController"] = append(beego.GlobalControllerRouter["testapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testapi/controllers:UserController"] = append(beego.GlobalControllerRouter["testapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Regist",
			Router: `/regist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testapi/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["testapi/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/info`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testapi/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["testapi/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "UpdateUserInfo",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["testapi/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["testapi/controllers:UserInfoController"],
		beego.ControllerComments{
			Method: "UploadPic",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
