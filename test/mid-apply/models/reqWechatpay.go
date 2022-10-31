package models

type Reqwechat struct {
	SpAppid     string `json:"sp_appid"`
	SpMchid     string `json:"sp_mchid"`
	Name        string `json:"name"`
	Shortname   string `json:"shortname"`
	OfficePhone string `json:"office_phone"`
	Contact     struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"contact"`
	BusinessCategory              int    `json:"business_category"`
	MerchantCountryCode           string `json:"merchant_country_code"`
	MerchantType                  string `json:"merchant_type"`
	RegistrationCertificateNumber string `json:"registration_certificate_number"`
	RegistrationCertificateDate   string `json:"registration_certificate_date"`
	SettlementBankNumber          string `json:"settlement_bank_number"`
	Business                      struct {
		BusinessType string `json:"business_type"`
		Mcc          string `json:"mcc"`
		MiniProgram  string `json:"mini_program"`
		StoreAddress string `json:"store_address"`
	} `json:"business"`
	Director struct {
		Name   string `json:"name"`
		Number string `json:"number"`
	} `json:"director"`
}
