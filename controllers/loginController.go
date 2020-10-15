package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"shujvrenzhengzhichonglai/models"
)

type LoginController struct{
	beego.Controller
}

func (l *LoginController)Get(){
	l.TplName="login.html"
}

func (l *LoginController)Post(){
	var user models.User
	err:=l.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户信息解析失败，请重试")
		return
	}
	u,err:=user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登入失败，请重试")
		return
	}
	l.Data["Phone"]=u.Phone
	l.TplName="home.html"
}
