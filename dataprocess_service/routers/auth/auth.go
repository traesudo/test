package auth

import (
	"dataprocess_service/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func Init() {
	ns := beego.NewNamespace("/data",
		beego.NSNamespace("/service",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
