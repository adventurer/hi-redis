package main

import (
	_ "hi-redis/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/assets", "static")
	beego.SetStaticPath("/", "dist")
	beego.SetStaticPath("/js", "dist/js")
	beego.SetStaticPath("/css", "dist/css")

	beego.Run()
}
