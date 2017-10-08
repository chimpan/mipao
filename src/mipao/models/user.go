package models

import (
	"github.com/astaxie/beego/orm"
	"mipao/util"
	"mipao/constant"
	"fmt"
)

func init() {
	orm.RegisterModel(new(User), new(UserInfo))
}

type User struct {
	Id       int `json:"id"`
	Phone    string `orm:"unique" json:"phone"`
	Password string    `json:"-"`
}

type UserInfo struct {
	Id         int `json:"-"`
	UserId     int `orm:"unique" json:"user_id"`
	HeaderUrl  string    `orm:"null" json:"header_url"`
	Sex        int    `orm:"null" json:"sex"`
	Sign       string    `orm:"null" json:"sign"`
	Lat        float64    `orm:"null" json:"lat"`
	Lon        float64`orm:"null" json:"lon"`
	HomeTown   string`orm:"null" json:"home_town"`
	NowLive    string`orm:"null" json:"now_live"`
	Phone      string `orm:"unique" json:"phone"`
	School     string `orm:"null" json:"school"`
	Age        int `json:"age"`
	Birth      int64 `json:"birth"`
	Profession string `orm:"null" json:"profession"`
	Email      string `orm:"null json:"email""`
	NiceCount  int `json:"nice_count"`
}

//注册
func AddUser(phone, pwd, code string) (int, string) {
	var a []User
	util.GetOrm().QueryTable("user").Filter("phone", phone).All(&a)
	if len(a) > 0 {
		return constant.Err_Phone_Exit, "手机号码已存在"
	}
	c, s := VerifyCode(phone, code)
	if c != 0 {
		return c, s
	}
	u := new(User)
	u.Password = pwd
	u.Phone = phone
	util.GetOrm().Insert(u)
	uinfo := new(UserInfo)
	uinfo.UserId = u.Id
	uinfo.Phone = phone
	util.GetOrm().Insert(uinfo)
	return 200, "注册成功"
}

//登录
func Login(phone, pwd string) (int, string, *UserInfo) {
	u := new(User)
	util.GetOrm().QueryTable("user").Filter("phone", phone).One(u)
	if u.Id == 0 || pwd != u.Password {
		return constant.Err_Login_Info, "账号或密码有误", nil
	}
	uinfo := new(UserInfo)
	util.GetOrm().QueryTable("user_info").Filter("user_id", u.Id).One(uinfo)
	return 200, "登录成功", uinfo
}

//验证码登录
func CodeLogin(phone, code string) (int, string, *UserInfo) {
	c, s := VerifyCode(phone, code)
	if c != 0 {
		return c, s, nil
	}
	uinfo := new(UserInfo)
	util.GetOrm().QueryTable("user_info").Filter("phone", phone).One(uinfo)
	return 200, "登录成功", uinfo
}

//获取周边用户
func GetNear(user_id int, lat, lon float64) (int, string, []UserInfo) {
	code, msg, _ := CommitLocation(user_id, lat, lon)
	if code != 200 {
		return code, msg, nil
	}
	var us []UserInfo
	cond := orm.NewCondition()
	cond1 := cond.And("lat__lte", 0.1+lat).And("lat__gte", lat-0.1)
	cond2 := cond.And("lon__lte", lon+0.32).And("lon__gte", lon-0.32)
	qs := util.GetOrm().QueryTable("user_info")
	cond3 := cond.AndCond(cond1).AndCond(cond2).AndNot("user_id", user_id)
	qs.SetCond(cond3).All(&us)
	fmt.Println(us)
	return 200, "", us
}

//提交经纬度
func CommitLocation(user_id int, lat, lon float64) (int, string, []UserInfo) {
	u := new(UserInfo)
	util.GetOrm().QueryTable("user_info").Filter("user_id", user_id).One(u)
	if u.Id == 0 {
		return constant.Err_User_Id_NO_FOUND, "用户未找到", nil
	}
	u.Lat = lat
	u.Lon = lon
	util.GetOrm().Update(u)
	return 200, "", nil
}

//获取用户信息
func GetUserInfo(userId int) (*UserInfo) {
	uinfo := new(UserInfo)
	util.GetOrm().QueryTable("user_info").Filter("user_id", userId).One(uinfo)
	return uinfo
}

//更新用户信息
func UpdateUserInfo(userId int,sex int,birth int64,sign ,home_town ,now_live,school,profession,email string)(int, string)  {
	uinfo := new(UserInfo)
	util.GetOrm().QueryTable("user_info").Filter("user_id", userId).One(uinfo)
	if uinfo.Id==0 {
		 return constant.Err_User_Id_NO_FOUND, "用户未找到"
	}
	uinfo.Sex = sex
	uinfo.Sign = sign
	uinfo.HomeTown = home_town
	uinfo.NowLive = now_live
	uinfo.School = school
	uinfo.Profession = profession
	uinfo.Email = email
	util.GetOrm().Update(uinfo)
	return 200,""
}
