package routers

import (
	"mybeegoapp/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
}
