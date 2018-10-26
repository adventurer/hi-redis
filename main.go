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

	// go func() {
	// 	var handel *exec.Cmd
	// 	cmd := "serve -s dist"
	// 	handel = exec.Command("/bin/sh", "-c", cmd)
	// 	output, err := handel.Output()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println(string(output))
	// }()
	beego.SetLevel(beego.LevelDebug)
	// beego.BeeLogger.DelLogger("console")

	beego.Run()
}
