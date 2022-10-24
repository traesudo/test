package main

import (
	_ "fmt"
	_ "github.com/gocolly/colly"
	"test/mid-apply/test"
)

func main() {
	//applicationInfo := service.GetApplicationData("1740b300-9dd1-41cd-9f06-0447818067dd")
	//companykeys := applicationInfo.Data.ApplicationCompany
	//businesskeys := applicationInfo.Data.ApplicationBusiness
	//companyInfo := service.GetCompanyData(companykeys)
	//fmt.Println(companyInfo.Data.Id)
	//if companyInfo.Data.Id == 0 {
	//	//写入log.......
	//	log.Fatal("getcompany failed")
	//}
	//businessInfo := service.GetBusinessDate(businesskeys)
	//fmt.Println("check.....business", businessInfo)
	//brkeys := companyInfo.Data.HkBrs
	//fmt.Println("check ..brInfo", brkeys)
	//
	////checkBR
	//brInfos := service.GetBrData(brkeys)
	//lenbr := len(brInfos)
	//if lenbr < 1 {
	//	log.Fatal("getBr failed")
	//
	//}
	//if lenbr > 1 {
	//	brInfos := service.SelectBr(brInfos)
	//	fmt.Println(brInfos)
	//}
	//
	//fmt.Println(lenbr)
	//var applymodes models.Applymodel
	//ok := service.NewApply(brInfos[0], applicationInfo, companyInfo, applymodes)
	//fmt.Println("ok", ok)
	//var c *service.Conf
	//c.DownCertificate()

	//var data = make(url.Values)
	//data.Set("grant_type", "client_credentials")
	//fmt.Println(strings.NewReader(data.Encode()))

	//	c := make(chan int)
	//	quit := make(chan int)
	//	go func() {
	//		for i := 0; i < 10; i++ {
	//			fmt.Println(<-c)
	//		}
	//		quit <- 0
	//	}()
	//	fibonacci(c, quit)
	//
	//}
	//
	//func fibonacci(c, quit chan int) {
	//	x, y := 0, 1
	//	for {
	//		select {
	//		case c <- x:
	//			x, y = y, x+y
	//		case <-quit:
	//			fmt.Println("quit")
	//			return
	//		}
	//	}
	var animal test.Animal
	animal.Write()
}
