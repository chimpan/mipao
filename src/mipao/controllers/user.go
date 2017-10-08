package controllers

import (
	"github.com/astaxie/beego"
	_"mipao/models"
	"mipao/models"
	"mipao/constant"
	"mipao/util"
)

// 用户
type UserController struct {
	beego.Controller
}

// @Title 注册
// @Description 注册
// @Param	phone		query 	string	true		"手机号码"
// @Param	password		query 	string	true		"密码"
// @Param	code		query 	string	true	"验证码"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /regist [post]
func (u *UserController) Regist() {
	phone := u.GetString("phone")
	pwd := u.GetString("password")
	code := u.GetString("code")
	r := models.Result{}
	if phone == "" {
		r.Code = constant.Err_Phone_nil
		r.Msg = "手机号是空"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	if !util.IsPhoneNumber(phone) {
		r.Code = constant.Err_Phone_Invaild
		r.Msg = "手机号码格式不对"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	if pwd == "" {
		r.Code = constant.Err_Password_nil
		r.Msg = "密码是空"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	if code == "" {
		r.Code = constant.Err_Code_nil
		r.Msg = "验证码是空"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	c, s := models.AddUser(phone, pwd, code)
	r.Code = c
	r.Msg = s
	u.Data["json"] = r
	u.ServeJSON()
}

// @Title
// @Description 登录
// @Param	phone		query 	string	true		"手机号码"
// @Param	password		query 	string	true		"密码"
// @Success 200
// @Failure 403
// @router /login [post]
func (u *UserController) Login() {
	phone := u.GetString("phone")
	pwd := u.GetString("password")
	r := models.Result{}
	if phone == "" {
		r.Code = constant.Err_Phone_nil
		r.Msg = "手机号是空"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	if !util.IsPhoneNumber(phone) {
		r.Code = constant.Err_Phone_Invaild
		r.Msg = "手机号码格式不对"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	if pwd == "" {
		r.Code = constant.Err_Password_nil
		r.Msg = "密码是空"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	c, s, user := models.Login(phone, pwd)
	r.Code = c
	r.Msg = s
	r.Data = user
	u.Data["json"] = r
	u.ServeJSON()
}

// @Title
// @Description  验证码登录
// @Param	phone		query 	string	true		"手机号码"
// @Param	code		query 	string	true	"验证码"
// @Success 200
// @Failure 403
// @router /code_login [post]
func (u *UserController) CodeLogin() {
	phone := u.GetString("phone")
	code := u.GetString("code")
	r := models.Result{}
	if phone == "" {
		r.Code = constant.Err_Phone_nil
		r.Msg = "手机号是空"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	if !util.IsPhoneNumber(phone) {
		r.Code = constant.Err_Phone_Invaild
		r.Msg = "手机号码格式不对"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	if code == "" {
		r.Code = constant.Err_Code_nil
		r.Msg = "验证码是空"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	c, s := models.IsPhoneRegisted(phone)
	if c != 200 {
		r.Code = c
		r.Msg = s
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	c1, s1, user := models.CodeLogin(phone, code)
	r.Code = c1
	r.Msg = s1
	r.Data = user
	u.Data["json"] = r
	u.ServeJSON()
}

// @Title
// @Description  获取周变跑友
// @Param	id		query 	int	 true		"用户id"
// @Success 200
// @Failure 403
// @router /getNear [get]
func (u *UserController) GetNear() {
	id, _ := u.GetInt("user_id", 0)
	lat, _ := u.GetFloat("lat", 0)
	lon, _ := u.GetFloat("lon", 0)

	r := new(models.Result)
	if id == 0 {
		r.Code = constant.Err_Params
		r.Msg = "参数错误"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	c, s, us := models.GetNear(id, lat, lon)
	r.Code = c
	r.Msg = s
	r.Data = us
	u.Data["json"] = r
	u.ServeJSON()
}

// @Title
// @Description  获取个人信息
// @Param	user_id		query 	int	 true		"用户id"
// @Success 200
// @Failure 403
// @router /info [get]
func (u *UserController) GetInfo() {
	id, _ := u.GetInt("user_id", 0)
	r := new(models.Result)
	if id == 0 {
		r.Code = constant.Err_Params
		r.Msg = "参数错误"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	uinfo := models.GetUserInfo(id)
	if uinfo.Id == 0 {
		r.Code = constant.Err_User_Id_NO_FOUND
		r.Msg = "用户未找到"
	} else {
		r.Code = 200
		r.Data = uinfo
	}
	u.Data["json"] = r
	u.ServeJSON()
}

// @Title
// @Description  获取个人信息
// @Param user_id query int true "用户id"
// @Param	sex			query 	int	 	true		"性别 0-未设置 1-男 2-女"
// @Param	sign		query 	string	 true		"个性签名"
// @Param	home_town	query 	string	 true		"故乡"
// @Param	now_live	query 	string	 true		"现居地"
// @Param	school		query 	string	 true		"毕业学校"
// @Param	birth		query 	int	 	true		"生日（秒）"
// @Param	profession	query 	string	 true		"职业"
// @Param	email		query 	string	 true		"email"
// @Success 200
// @Failure 403
// @router /update_info [post]
func (u *UserController) UpdateInfo() {
	user_id, _ := u.GetInt("user_id", 0)
	r := new(models.Result)
	if user_id == 0 {
		r.Code = constant.Err_Params
		r.Msg = "参数错误"
		u.Data["json"] = r
		u.ServeJSON()
		return
	}
	sex, _ := u.GetInt("sex", 0)
	birth, _ := u.GetInt64("birth", 0)
	sign := u.GetString("sign", "")
	home_town := u.GetString("home_town", "")
	now_live := u.GetString("now_live", "")
	school := u.GetString("school", "")
	profession := u.GetString("profession", "")
	email := u.GetString("email", "")
	code, msg := models.UpdateUserInfo(user_id, sex, birth, sign, home_town, now_live, school, profession, email)
	r.Code = code
	r.Msg = msg
	u.Data["json"] = r
	u.ServeJSON()
}
