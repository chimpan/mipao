package controllers

import (
	"testapi/models"

	"strings"

	"github.com/astaxie/beego"
)

//动态
type NewsController struct {
	beego.Controller
}

// @Title 发布动态
// @Description 发布动态接口
// @Param	userid	query  	int	true		"用户id"
// @Param	image	query  	string	false		"图片"
// @Param	content	query  	string	true		"发布内容"
// @Success 200
// @Failure 403
// @router /publish [post]
func (this *NewsController) Publish() {
	id, _ := this.GetInt("userid", -1)
	image := this.GetString("image")
	content := this.GetString("content")
	news := &models.News{Userid: id, PublishImage: image, PublishContent: content}
	err := models.InsertPublishData(news)
	r := &models.Result{}
	if err != nil {
		r.Code = 400
		r.Msg = "发表失败"
	} else {
		r.Code = 200
		r.Msg = "发表成功"
	}
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title 所有动态
// @Description 所有动态接口
// @Success 200
// @Failure 403
// @router /all [get]
func (this *NewsController) All() {
	news, err := models.FindAllNews()
	r := &models.Result{}
	if err != nil {
		r.Code = 400
		r.Msg = "获取动态失败"
	} else {
		r.Code = 200
		r.Data = news
	}
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title 评论
// @Description 评论接口
// @Param	userid	query  	int	true		"用户id"
//@Param	newsid	query  	int	true		"动态id"
//@Param	content	query  	string	true		"评论内容"
// @Success 200
// @Failure 403
// @router /commend [post]
func (this *NewsController) Commend() {
	// _, err := this.GetInt("useid", 0)
	// _, err := this.GetInt("newsid", 0)
	content := this.GetString("content")
	r := &models.Result{}
	if strings.EqualFold("", content) {
		r.Code = 400
		r.Msg = "评论内容不能为空"
	} else {

	}
	this.Data["json"] = r
	this.ServeJSON()
}
