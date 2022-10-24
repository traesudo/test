package service

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	_ "github.com/kylelemons/go-gypsy/yaml"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test/mid-apply/models"
	"time"
)

func GetApplicationData(key string) models.Application {

	url := "https://wonder_registry.bindolabs.com/items/business_applications/" + key
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header("authorization", "Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T")
	req.Header = map[string][]string{
		"authorization": {"Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T"},
	}
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	var application models.Application
	if err == nil {
		// 解析Response

		err = json.Unmarshal(body, &application)

	}

	return application

}

func GetCompanyData(key int) models.Compnay {

	url := fmt.Sprintf("https://wonder_registry.bindolabs.com/items/companies/%d", key)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header("authorization", "Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T")
	req.Header = map[string][]string{
		"authorization": {"Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T"},
	}
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("body---->", body)
	var company models.Compnay
	if err == nil {
		// 解析Response

		err = json.Unmarshal(body, &company)

	}

	return company

}
func GetBusinessDate(key string) models.Business {

	url := "https://wonder_registry.bindolabs.com/items/businesses/" + key
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header("authorization", "Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T")
	req.Header = map[string][]string{
		"authorization": {"Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T"},
	}
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	var business models.Business
	if err == nil {
		// 解析Response

		err = json.Unmarshal(body, &business)

	}

	return business
}
func GetBrData(keys []string) []models.Br {
	var brs []models.Br
	for key, value := range keys {

		keys[key] = value

		url := "https://wonder_registry.bindolabs.com/items/hk_brs/" + value
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		//req.Header("authorization", "Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T")
		req.Header = map[string][]string{
			"authorization": {"Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T"},
		}
		resp, err := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)

		var br models.Br
		if err == nil {
			// 解析Response

			err = json.Unmarshal(body, &br)

		}
		fmt.Println("br......", br)
		brs := append(brs, br)
		fmt.Println("brs......", brs)
		return brs
	}
	return brs

}

func SelectBr(brs []models.Br) []models.Br {
	for j := 0; j < len(brs); j++ {
		for i := 0; i < len(brs); i++ {
			fmt.Println(brs[j].Data.DateCreated)
			typeof := fmt.Sprintf("typeof %T", brs[j].Data.DateCreated)
			fmt.Println(typeof)
			if brs[j].Data.DateCreated.After(brs[j+1].Data.DateCreated) {
				temp := brs[j+1]
				brs[j+1] = brs[j]
				brs[j] = temp

			}

		}
	}
	return brs

}

func NewApply(brs models.Br, application models.Application, company models.Compnay, applymodel models.Applymodel) models.Applymodel {
	var merchant_typeVal = ""
	var business_typeVul = ""
	if company.Data.CompanyType == "Private unlimited company" {
		merchant_typeVal = "INDIVIDUAL"
	}
	if company.Data.CompanyType == "Private company limited by shares" || company.Data.CompanyType == "Company limited by guarantee" || company.Data.CompanyType == "Public unlimited company" || company.Data.CompanyType == "Public company limited by shares" {
		merchant_typeVal = "ENTERPRISE"
	}
	if application.Data.WebsiteUrl != "" {
		business_typeVul = "ONLINE"
	} else {
		business_typeVul = "OFFLINE"
	}
	applymodel.SpAppid = "wx31331a78e8a40eea"
	applymodel.SpMchid = "1501977491"
	applymodel.Name = company.Data.NameOfBusinessCorporationChinese
	applymodel.Shortname = company.Data.NameOfBusinessCorporationChinese
	applymodel.OfficePhone = application.Data.NaturalPersonPhone
	applymodel.BusinessCategory = 656
	applymodel.MerchantCountryCode = "344"
	applymodel.MerchantType = merchant_typeVal
	applymodel.RegistrationCertificateNumber = brs.Data.BrNumberWithoutCheckDigit
	applymodel.RegistrationCertificateDate = brs.Data.ExpirationDate
	applymodel.SettlementBankNumber = ""
	applymodel.Business.BusinessType = business_typeVul
	applymodel.Business.Mcc = "5812"
	applymodel.Business.StoreAddress = brs.Data.CompanyAddress
	return applymodel
}

func (c *Conf) DownCertificate() {
	timeUnix := time.Now().Unix()
	time1 := strconv.Itoa(int(timeUnix))
	var number1 = "adssw223"
	var mchid1 = "1501977491"
	var serialvalue = "57309C2E21D7D22688A6FC8A04245526E005A553"
	var data = ""
	var message = "GET" + "\n" + "/v3/certificates" + "\n" + time1 + "\n" + number1 + "\n" + data + "\n"
	conf := c.getConf()
	signkeys, _ := Sha256WithRsa(conf, message)
	fmt.Println("check key-------chekc key ", signkeys)
	finstr := fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",serial_no="%s",nonce_str="%s",timestamp="%s",signature="%s"`, mchid1, serialvalue, number1, time1, signkeys)
	url := "https://api.mch.weixin.qq.com/v3/certificates"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header("authorization", "Bearer 3op6y9nMdIEHJG2-KPH88uHaU2xuQk8T")
	req.Header = map[string][]string{
		"content-type":  {"application/json"},
		"Accept":        {"application/json"},
		"authorization": {finstr},
		"User-Agent":    {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"},
	}
	resp, err := client.Do(req)

	fmt.Printf("%T", resp.Body)
	if err != nil {
		log.Fatal("request error", err)

	}
	var Serfile models.Downfile
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &Serfile)
		err = json.Unmarshal(body, &Serfile)

	}
	fmt.Println("check ........11", Serfile)

	fmt.Println("check ........11", Serfile.Data[0].SerialNo)
	fmt.Println("check ........11", Serfile.Data[0].EncryptCertificate)

}

func (c *Conf) getConf() string {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("check.....conf ", yamlFile)
	viper.SetConfigName("config")
	viper.AddConfigPath("/Users/developer/develop/src/test/mid-apply")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
	fmt.Println(viper.Get("semiprivate"))
	key := viper.Get("semiprivate")

	if err != nil {
		fmt.Println(err.Error())
	}
	return key.(string)
}

type Conf struct {
	lineprivatekey string `yaml:"Lineprivatekey"`
}

// let body={
//            method: "get",
//            url: 'https://api.mch.weixin.qq.com/v3/certificates',
//            headers: {
//                "content-type":"application/json",
//                "Accept":"application/json",
//                "authorization": finstr,
//                "User-Agent":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
//            },
//        }

//`WECHATPAY2-SHA256-RSA2048 mchid="${mchid1}",serial_no="${serialvalue}",nonce_str="${number1}",timestamp="${time2}",signature="${keysign}"`
////step1
//async  downCertificate($){
////时间戳
//const time1 =Math.floor(Date.now() / 1000);
//const time2= time1.toString()
////随机串
//const number1="adssw223"
////加密商户id

//const serialvalue="57309C2E21D7D22688A6FC8A04245526E005A553"
////message串
//var data = "";

//const keysign = await this.sha256withRSA($,message)
//console.log("打印签名加密keysign",keysign)
//var finstr =`WECHATPAY2-SHA256-RSA2048 mchid="${mchid1}",serial_no="${serialvalue}",nonce_str="${number1}",timestamp="${time2}",signature="${keysign}"`
//console.log("get finstr info",finstr)
//let body={
//method: "get",
//url: 'https://api.mch.weixin.qq.com/v3/certificates',
//headers: {
//"content-type":"application/json",
//"Accept":"application/json",
//"authorization": finstr,
//"User-Agent":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
//},
//}
//try{
//let res = await axios($,body);
////console.log(res.data)
//return res
//}catch(err){
//console.log(err)
//return err
//}
//
//
//
//}
func Sha256WithRsa(privateRaw string, msg string) (string, error) {
	privateRaw = strings.Trim(privateRaw, "\n")
	if !strings.HasPrefix(privateRaw, "-----BEGIN RSA PRIVATE KEY-----") {
		privateRaw = fmt.Sprintf("%s\n%s\n%s", "-----BEGIN RSA PRIVATE KEY-----", privateRaw, "-----END RSA PRIVATE KEY-----")
	}

	blockPri, _ := pem.Decode([]byte(privateRaw))
	if blockPri == nil {
		return "", fmt.Errorf("blockPri is nil")
	}

	rsaPri, e := genPriKey(blockPri.Bytes, PKCS8)
	if e != nil {
		panic(e)
	}

	h := sha256.New()
	h.Write([]byte(msg))
	d := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPri, crypto.SHA256, d)
	if err != nil {
		log.Fatal("change ..... SignPKCS1v15", err)

	}
	encodedSig := base64.StdEncoding.EncodeToString(signature)
	return encodedSig, nil
}

const (
	PKCS1 int64 = iota
	PKCS8
)

func genPriKey(privateKey []byte, privateKeyType int64) (*rsa.PrivateKey, error) {
	var priKey *rsa.PrivateKey
	var err error
	switch privateKeyType {
	case PKCS1:
		{
			priKey, err = x509.ParsePKCS1PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
		}
	case PKCS8:
		{
			prkI, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
			priKey = prkI.(*rsa.PrivateKey)
		}
	default:
		{
			return nil, fmt.Errorf("unsupport private key type")
		}
	}
	return priKey, nil
}
