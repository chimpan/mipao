package models

import (
	"github.com/astaxie/beego/orm"
	"mipao/util"
	"mipao/constant"
	"time"
	"math/rand"
	"strconv"
)

type Msg struct {
	Id      int
	Phone   string
	Code    string // 验证码
	State   int    //0-未使用  1-已使用
	AddTime int64  //获取验证码时间
}

func init() {
	orm.RegisterModel(new(Msg))
}

//判断手机号是否注册
func IsPhoneRegisted(phone string) (int,string) {
	var u []User
	util.GetOrm().QueryTable("user").Filter("phone", phone).All(&u)
	if len(u) == 0 {
		return constant.Err_Phone_NO_FOUND,"手机号码未注册"
	}

	return 200,"手机号码已注册"
}

//获取验证码
func GetVerifyCode(phone string) (string) {
	msg := new(Msg)
	msg.Phone = phone
	msg.State = 0
	msg.AddTime = time.Now().Unix()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var s string
	for i := 0; i < 4; i++ {
		s += strconv.Itoa(r.Intn(10))
	}
	msg.Code = s
	util.GetOrm().Insert(msg)
	return msg.Code
}

/**
验证验证码
 */
func VerifyCode(phone, code string) (int, string) {
	msg := new(Msg)
	util.GetOrm().QueryTable("msg").Filter("phone", phone).OrderBy("-add_time").One(msg)
	if msg.State == 1 {
		return constant.Err_Code_Invaild, "验证码已使用"
	}
	if msg.Code != code {
		return constant.Err_Code_Unmatch, "验证码有误"
	}
	msg.State = 1
	util.GetOrm().Update(msg, "state")
	return 0, ""
}

