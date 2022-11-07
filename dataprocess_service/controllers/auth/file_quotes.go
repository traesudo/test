package auth

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//@router /file_quotation [post]
func (c *MainController) FileQuotation() {
	f, h, err := c.GetFile("filename")
	if err != nil {
		c.SendError(400100, "读取文件错误")
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	datastr := string(data)

	//var sliceHaiCoder = []string{"Hello", "3", "2", "1"}
	//sliceHaiCoder = append(sliceHaiCoder[:4], "4")
	//fmt.Println("sliceHaiCoder=", sliceHaiCoder)

	for i := 0; i < len(datastr); i++ {
		ch := datastr[i]
		//ctype := reflect.TypeOf(ch)
		ok := string(ch)

		if ok == "'" {
			data[i] = '"'
		}
		if ok == "\n" {
			lastnumber := strings.LastIndex(datastr, string(data[i]))
			if lastnumber == -1 {
				return
			}
			if (data[i+1]) != '"' {
				fmt.Println("check")
				//data = append(data[:i], '"')
				//data[i] = '"' + data[i+1] + '"'
				//strings.Replace(, string(datastr[i+1]),)
			}
		}

	}
	fmt.Println("check len", string(data))

	//for key, value := range datastr {
	//	fmt.Println(key)
	//	fmt.Println(value)
	//}
	//fmt.Println()

	//params := &dto.SeparateQuotation{}
	//err := json.Unmarshal(c.Body, params)
	//log.Info("check....", params)
	//if err != nil {
	//	c.SendError(100400, "no unmarshar error")
	//}
	//strsum, err := modules.ReadFile("./user")
	//fmt.Println(strsum)
	c.SaveToFile("filename", "static/upload/"+h.Filename)
	c.SuccessOutput(string(data), "")

}
