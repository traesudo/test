package auth

import (
	"dataprocess_service/controllers/auth"
	beego "github.com/beego/beego/v2/server/web"
)

func Init() {
	//beego.Include(&auth.AuthController{})
	beego.Include(&auth.MainController{})

	//beego.AddNamespace(ns)
}
