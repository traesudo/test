package auth

import (
	"dataprocess_service/dto"
	"dataprocess_service/modules"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

//@router /separate_quotation [post]
func (c *MainController) SeparateQuotation() {

	params := &dto.SeparateQuotation{}
	err := json.Unmarshal(c.Body, params)
	log.Info("check....", params)
	if err != nil {
		c.SendError(100400, "no unmarshar error")
	}
	strsum, err := modules.ReadFile("./user")
	fmt.Println(strsum)
	c.SuccessOutput(strsum, "")

}
