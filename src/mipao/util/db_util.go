package util

import "github.com/astaxie/beego/orm"

func GetOrm() orm.Ormer {
	o:=orm.NewOrm()
	return o
}