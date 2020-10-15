package controllers

import (
	"github.com/astaxie/beego"
	"shujvrenzhengzhichonglai/models"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController)Post(){
	var user models.User
	err:=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析失败，请重试")
		return
	}

	_,err=user.AddUser()
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试！")
		return
	}
	r.TplName="login.html"
}
