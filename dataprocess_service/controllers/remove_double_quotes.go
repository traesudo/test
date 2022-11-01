package controllers

import (
	"dataprocess_service/dto"
	"fmt"
)

//@router /removeDQ [post]
func (c *MainController) removeDoubleQuotes() {
	params := &dto.DataQuotes{}
	c.UserRequestBody(&params)
	fmt.Println("check......params", params)

	c.SuccessOutput(params, "")

}
