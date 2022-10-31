package main

import (
	_ "dataprocess_service/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
