package controllers

import (
	"testapi/models"

	"fmt"
	"time"

	"strconv"

	"testapi/constants"

	"strings"

	"github.com/astaxie/beego"
)

type UserInfoController struct {
	beego.Controller
}

// @Title 个人信息
// @Description 个人信息接口
// @Param	userid		query 	string	true		"用户id"
// @Success 200
// @Failure 403
// @router /info [get]
func (this *UserInfoController) GetUserInfo() {
	id := this.GetString("userid")
	fmt.Println("id---->", id)
	nid, _ := strconv.Atoi(id)
	u, err := models.GetUserInfo(nid)
	r := new(models.Result)
	if err != nil {
		r.Code = 400
		r.Msg = "用户id不存在"
	} else {
		r.Code = 200
		r.Data = u
	}
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title 修改个人信息
// @Description 修改个人信息接口
// @Param	userid		query 	int		true		"用户id"
// @Param	head_url	query 	string	false		"头像地址"
// @Param	sex			query 	int		false		"性别"
// @Param	birth		query 	string	false		"生日"
// @Param	sign		query 	string	false		"签名"
// @Param	nick		query 	string	false		"昵称"
// @Success 200
// @Failure 403
// @router /update [post]
func (this *UserInfoController) UpdateUserInfo() {
	idstr := this.GetString("userid")
	fmt.Println("idstr----->", idstr)
	head := this.GetString("head_url")
	sexstr := this.GetString("sex")
	birth := this.GetString("birth")
	sign := this.GetString("sign")
	nick := this.GetString("nick")
	id, _ := strconv.Atoi(idstr)
	sex, _ := strconv.Atoi(sexstr)
	uinfo := &models.UserInfo{UserId: id, HeadUrl: head, Sex: sex, Birth: birth, Sign: sign, Nick: nick}
	err := models.UpdateUserInfo(uinfo)
	var res models.Result
	if err != nil {
		res.Code = 400
		res.Msg = "用户不存在"
	} else {
		res.Code = 200
		res.Msg = "更改成功"
	}
	this.Data["json"] = res
	this.ServeJSON()
}

// @Title 上传头像
// @Description 上传头像接口
// @Param	head_file	formData  	file	true		"图片"
// @Success 200
// @Failure 403
// @router /upload [post]
func (this *UserInfoController) UploadPic() {
	// header := this.Ctx.Request.Header
	// fmt.Println(header)

	r := new(models.Result)
	//判断文件是否是图片
	_, f, _ := this.GetFile("head_file")
	if !(strings.HasSuffix(f.Filename, ".jpg") || strings.HasSuffix(f.Filename, ".png")) {
		r.Code = 400
		r.Msg = "上传的文件不是图片"
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
		r.Code = 400
		r.Msg = "上传失败"
	} else {
		r.Code = 200
		r.Msg = "上传成功"
		r.Data = constants.PrefixImageURL + fileName
	}
	this.Data["json"] = r
	this.ServeJSON()
}
