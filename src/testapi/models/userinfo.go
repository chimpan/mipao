package models

import (
	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	Id      int
	UserId  int    `orm:"unique"`
	HeadUrl string `orm:"null"`
	Sex     int    `orm:"null"`
	Birth   string `orm:"null"`
	Sign    string `orm:"null"`
	Nick    string `orm:"null"`
}

func init() {
	orm.RegisterModel(new(UserInfo))
}

func GetUserInfo(userId int) (UserInfo, error) {
	var u UserInfo
	//id不存在
	err := isUserIdExist(userId)
	if err != nil {
		return u, err
	}
	o := orm.NewOrm()

	o.QueryTable("user_info").Filter("userid", userId).One(&u)
	return u, nil
}

/*
判断用户id是否存在
*/
func isUserIdExist(userId int) error {
	o := orm.NewOrm()
	var uu User
	err := o.QueryTable("user").Filter("id", userId).One(&uu)
	if err != nil { //id不存在
		return err
	}
	return nil
}

/*
插入一条个人信息
*/
func SetUserInfo(u *UserInfo) {
	o := orm.NewOrm()
	o.Insert(u)
}

/*
修改个人信息
*/
func UpdateUserInfo(uinfo *UserInfo) error {
	err := isUserIdExist(uinfo.UserId)
	if err != nil {
		return err
	}
	u := new(UserInfo)
	o := orm.NewOrm()
	o.QueryTable("user_info").Filter("user_id", uinfo.UserId).One(u)
	u.Birth = uinfo.Birth
	u.HeadUrl = uinfo.HeadUrl
	u.Sex = uinfo.Sex
	u.Sign = uinfo.Sign
	u.Nick = uinfo.Nick
	o.Update(u)

	var news []News
	o.QueryTable("news").Filter("userid", uinfo.UserId).All(&news)
	for _, n := range news {
		n.UserNick = uinfo.Nick
		n.UserHeadUrl = uinfo.HeadUrl
		o.Update(&n)
	}
	return nil
}
