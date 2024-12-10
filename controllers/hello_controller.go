// controllers/hello_controller.go
package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	web.Controller
}

func (c *HelloController) Get() {
	c.Data["json"] = map[string]string{"message": "Hello, Beego!"}
	c.ServeJSON()
}
