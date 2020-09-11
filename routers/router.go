package routers

import (
	"github.com/yxydde/go-server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ServerController{})
}
