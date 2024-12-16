package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Set up static directories
	web.SetStaticPath("/css", "static/css")
	web.SetStaticPath("/js", "static/js")
	web.SetStaticPath("/images", "static/images")

	// Serve the HTML file
	web.Router("/", &MainController{})

	web.Run() // Start the server
}

// MainController handles the root route
type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html" // Template file to serve
}
