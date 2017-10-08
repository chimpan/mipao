package controllers

import (
	"github.com/astaxie/beego"
	"mipao/models"
	"mipao/constant"
	"mipao/util"
	"strings"
	"time"
	"strconv"
	"fmt"
)

//公共
type CommonController struct {
	beego.Controller
}

// @Title
// @Description 获取验证码
// @Param	phone		query 	string	true		"手机号码"
// @Success 200
// @Failure 403
// @router /getcode [get]
func (c *CommonController) GetVerifyCode() {
	p := c.GetString("phone")
	r := new(models.Result)
	if p == "" {
		r.Code = constant.Err_Phone_nil
		r.Msg = "手机号码是空"
		c.Data["json"] = r
		c.ServeJSON()
		return

	}
	if !util.IsPhoneNumber(p) {
		r.Code = constant.Err_Phone_Invaild
		r.Msg = "手机号码格式不对"
		c.Data["json"] = r
		c.ServeJSON()
		return
	}
	code := models.GetVerifyCode(p)
	r.Code = 200
	r.Data = code
	c.Data["json"] = r
	c.ServeJSON()
}

// @Title
// @Description 判断手机号是否注册
// @Param	phone		query 	string	true		"手机号码"
// @Success 200
// @Failure 403
// @router /is_registed [get]
func (c *CommonController) IsPhoneRegisted() {
	p := c.GetString("phone")
	r := new(models.Result)
	if p == "" {
		r.Code = constant.Err_Phone_nil
		r.Msg = "手机号码是空"
		c.Data["json"] = r
		c.ServeJSON()
		return

	}
	if !util.IsPhoneNumber(p) {
		r.Code = constant.Err_Phone_Invaild
		r.Msg = "手机号码格式不对"
		c.Data["json"] = r
		c.ServeJSON()
		return
	}
	co, str := models.IsPhoneRegisted(p)
	r.Code = co
	r.Msg = str
	c.Data["json"] = r
	c.ServeJSON()
}

// @Title 上传头像
// @Description 上传头像接口
// @Param	head_file	formData  	file	true		"图片"
// @Success 200
// @Failure 403
// @router /upload [post]
func (this *CommonController) UploadPic() {
	r := new(models.Result)
	//判断文件是否是图片
	_, f, _ := this.GetFile("head_file")
	if !(strings.HasSuffix(f.Filename, ".jpg") || strings.HasSuffix(f.Filename, ".png")) {
		r.Code = constant.Err_Not_Pic_File
		r.Msg = "不是图片文件"
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	cur := time.Now()
	timestamp := cur.Unix() //返回时间戳,自从1970年1月1号到现在
	fileName := "static/file/" + strconv.FormatInt(timestamp, 10) + ".png"
	err := this.SaveToFile("head_file", fileName)

	if err != nil {
		fmt.Println(err)
		r.Code = 500
		r.Msg = "上传失败"
	} else {
		r.Code = 200
		r.Msg = "上传成功"
		r.Data = fileName
	}
	this.Data["json"] = r
	this.ServeJSON()
}
