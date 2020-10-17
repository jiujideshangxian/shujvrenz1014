package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"shujvrenzhengzhichonglai/db_mysql"

)

func main() {
	//连接数据库
	db_mysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run() //阻塞
	fmt.Println("nan")
	//http.ListenAndServe(":8080")
}

