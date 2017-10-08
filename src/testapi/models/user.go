package models

import (
	"strings"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Username string `orm:"unique"`
	Password string
 
}

func init() {
	orm.RegisterModel(new(User))
}

/*
	注册
*/
func Regist(username, password  string) bool {
	o := orm.NewOrm()
	o.Using("default")

	u := new(User)
	u.Username = username
	 
	u.Password = password

	_, err := o.Insert(u)
	if err != nil {
		return false
	}

	return true
}

/*
登录
成功 返回true user id
失败 返回false -1用户不存在  -2密码错误
*/
func Login(user, pwd string) (bool, int) {
	var u User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("username", user).One(&u)
	if err != nil {
		return false, -1
	}
	if strings.EqualFold(user, u.Username) && strings.EqualFold(pwd, u.Password) {
		return true, u.Id
	} else {
		return false, -2
	}
}
