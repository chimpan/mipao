package models

import (
	"fmt"

	"time"

	"github.com/astaxie/beego/orm"
)

type Commend struct {
	Id             int
	UserInfo       *UserInfo `orm:"rel(fk)"`
	CommendContent string
	CommendTime    int64
}

type News struct {
	Id             int
	Userid         int
	UserHeadUrl    string
	UserNick       string
	PublishImage   string `orm:"null"`
	PublishContent string `orm:"null"`
	PublishTime    int64
	GoodNum        int
	// Commend        *Commend `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(News), new(Commend))
}

func InsertPublishData(news *News) error {

	o := orm.NewOrm()
	u := new(UserInfo)
	o.QueryTable("user_info").Filter("user_id", news.Userid).One(u)
	fmt.Println(2, u)
	news.UserHeadUrl = u.HeadUrl
	news.UserNick = u.Nick
	news.PublishTime = time.Now().Unix()
	_, err := o.Insert(news)
	return err
}

func FindAllNews() ([]News, error) {
	// n := &News{1, 1, "", "", "", "", 1}
	// news := []*News{n, n, n}
	var newss []News
	o := orm.NewOrm()
	_, err := o.QueryTable("news").OrderBy("-publish_time").All(&newss)

	return newss, err
}

// func Commend1(content string, userid, newsid int) error {
// 	var userinfo UserInfo
// 	var new News
// 	o := orm.NewOrm()
// 	err := o.QueryTable("user_info").Filter("user_id", userid).One(&userinfo)
// 	err1 := o.QueryTable("news").Filter("id", newsid).One(&new)
// 	if err != nil || err1 != nil {
// 		return err
// 	}
// 	return err
// }
