package routers

import (
	"shujvrenzhengzhichonglai/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
