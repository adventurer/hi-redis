package routers

import (
	"hi-redis/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/echo", &controllers.MainController{}, "*:Echo")
	beego.Router("/welcome", &controllers.MainController{}, "*:Welcome")
	beego.Router("/data", &controllers.MainController{}, "*:Data")
	beego.Router("/get", &controllers.MainController{}, "*:Key")
	beego.Router("/cmd", &controllers.MainController{}, "*:Command")

}
