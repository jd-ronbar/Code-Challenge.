package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/drivers/edit/:id:int", &controllers.DriversController{}, "get,put,post:Edit")
	beego.Router("/drivers/add", &controllers.DriversController{}, "get,post:Add")
	beego.Router("/drivers/", &controllers.DriversController{}, "get:Get")
	beego.Router("/drivers/delete/:id:int", &controllers.DriversController{}, "*:Delete")
	
	beego.Router("/metrics/edit/:id:int", &controllers.MetricsController{}, "get,put,post:Edit")
	beego.Router("/metrics/add", &controllers.MetricsController{}, "get,post:Add")
	beego.Router("/metrics/", &controllers.MetricsController{}, "get:Get")
        beego.Router("/metrics/delete/:id:int", &controllers.MetricsController{}, "*:Delete")
	
	//beego.Router("/drivers/delete/:id([0-9]+)", &controllers.DriversController{}, "*:Delete")
	//beego.Router("/drivers/update/:id([0-9]+)", &controllers.DriversController{}, "*:Update")
}
