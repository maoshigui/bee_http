package routers

import (
	"bee_http/goms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/cell", &controllers.CellController{})
	// beego.AutoRouter(&controllers.CellController{})
	beego.Router("/testcase", &controllers.TestcaseController{})
	beego.Router("/statistic", &controllers.StatisticController{})
}
