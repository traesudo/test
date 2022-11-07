package auth

import (
	"dataprocess_service/controllers/formbind/json"
	"dataprocess_service/dto"
	"dataprocess_service/modules"
	"fmt"
	log "github.com/sirupsen/logrus"
	//""
)

//@router /removedq [post]
func (c *MainController) RemoveDoubleQuotes() {
	params := &dto.DataQuotes{}
	err := json.Unmarshal(c.Body, params)
	log.Info(params.Type)
	if err != nil {
		c.SendError(100400, "no unmarshar error")
	}
	if params.QuotesInfo == nil {
		c.SendError(100400, "params QuotesInfo is null")
	}
	if params.Type == "" {
		c.SendError(100400, "params Type is null")
	}
	//csv
	log.Info("[check params ]", params)
	strSum := ""
	if params.Type == "数组转CSV" {
		strSum = modules.ToCsv(params.QuotesInfo, "transverse")
	}
	fmt.Println(strSum)
	c.SuccessOutput(strSum, "")

}
