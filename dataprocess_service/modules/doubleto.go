package modules

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func ToCsv(str []string, Type string) string {
	log.Info("[check params ]", str)
	strSum := "start"
	for _, value := range str {
		//fmt.Println(params.QuotesInfo[key])
		//fmt.Println(value)
		//保留 换行符号
		strSum = strSum + fmt.Sprintln(",", value)
	}

	//不保留换行
	if Type == "transverse" {
		strSum = strings.Replace(strSum, "\n", "", -1)
		fmt.Println("[sum]", strSum)
	}
	return strSum

}
