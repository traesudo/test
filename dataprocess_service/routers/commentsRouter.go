package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["dataprocess_service/controllers/auth:MainController"] = append(beego.GlobalControllerRouter["dataprocess_service/controllers/auth:MainController"],
		beego.ControllerComments{
			Method:           "RemoveDoubleQuotes",
			Router:           `/removedq`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
