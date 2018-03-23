package routers

import (
	"scrape/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/api/live.do", &controllers.MainController{})
}
