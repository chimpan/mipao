package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["mipao/controllers:CommonController"] = append(beego.GlobalControllerRouter["mipao/controllers:CommonController"],
		beego.ControllerComments{
			Method: "GetVerifyCode",
			Router: `/getcode`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mipao/controllers:CommonController"] = append(beego.GlobalControllerRouter["mipao/controllers:CommonController"],
		beego.ControllerComments{
			Method: "IsPhoneRegisted",
			Router: `/is_registed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mipao/controllers:CommonController"] = append(beego.GlobalControllerRouter["mipao/controllers:CommonController"],
		beego.ControllerComments{
			Method: "UploadPic",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mipao/controllers:UserController"] = append(beego.GlobalControllerRouter["mipao/controllers:UserController"],
		beego.ControllerComments{
			Method: "CodeLogin",
			Router: `/code_login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mipao/controllers:UserController"] = append(beego.GlobalControllerRouter["mipao/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetNear",
			Router: `/getNear`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mipao/controllers:UserController"] = append(beego.GlobalControllerRouter["mipao/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetInfo",
			Router: `/info`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mipao/controllers:UserController"] = append(beego.GlobalControllerRouter["mipao/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mipao/controllers:UserController"] = append(beego.GlobalControllerRouter["mipao/controllers:UserController"],
		beego.ControllerComments{
			Method: "Regist",
			Router: `/regist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mipao/controllers:UserController"] = append(beego.GlobalControllerRouter["mipao/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateInfo",
			Router: `/update_info`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
