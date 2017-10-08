package controllers

import (
	"testapi/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 用户模块接口
type UserController struct {
	beego.Controller
}

type userInfo struct {
	Id int
}

// @Title 登录
// @Description 登录接口
// @Param	username		query 	string	true		"用户名"
// @Param	password		query 	string	true		"密码"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	b, i := models.Login(username, password)
	r := new(models.Result)
	if b {
		r.Code = 200
		r.Msg = "登录成功"
		uu := &userInfo{Id: i}
		r.Data = uu
	} else {
		switch i {
		case -1:
			r.Code = 400
			r.Msg = "用户不存在"
		case -2:
			r.Code = 400
			r.Msg = "密码错误"
		}
	}
	u.Data["json"] = r
	u.ServeJSON()
}

// @Title 注册
// @Description 注册接口
// @Param	username		query 	string	true		"用户名"
// @Param	password		query 	string	true		"密码"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /regist [post]
func (u *UserController) Regist() {
	username := u.GetString("username")
	password := u.GetString("password")

	r := new(models.Result)
	if username == "" {
		r.Code = 400
		r.Msg = "用户名为空"
	} else if password == "" {
		r.Code = 400
		r.Msg = "密码为空"
	} else if models.Regist(username, password) {
		r.Code = 200
		r.Msg = "注册成功"
		//插入一条个人信息
		insertUserInfo(username)
	} else {
		r.Code = 400
		r.Msg = "用户已存在"
	}
	u.Data["json"] = r
	u.ServeJSON()
}

func insertUserInfo(username string) {
	o := orm.NewOrm()
	var u models.User
	o.QueryTable("user").Filter("username", username).One(&u)
	uinfo := new(models.UserInfo)
	uinfo.UserId = u.Id
	models.SetUserInfo(uinfo)
}
