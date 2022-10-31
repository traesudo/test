package models

import "time"

type Downfile struct {
	Data []struct {
		EffectiveTime      time.Time `json:"effective_time"`
		EncryptCertificate EncryptCertificate
		ExpireTime         time.Time `json:"expire_time"`
		SerialNo           string    `json:"serial_no"`
	} `json:"data"`
}

type EncryptCertificate struct {
	Algorithm      string `json:"algorithm"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
	Nonce          string `json:"nonce"`
}

type Serfile struct {
	Data []struct {
		SerialNo           string    `json:"serial_no"`
		EffectiveTime      time.Time `json:"effective_time"`
		ExpireTime         time.Time `json:"expire_time"`
		EncryptCertificate struct {
			Algorithm      string `json:"algorithm"`
			Nonce          string `json:"nonce"`
			AssociatedData string `json:"associated_data"`
			Ciphertext     string `json:"ciphertext"`
		} `json:"encrypt_certificate"`
	} `json:"data"`
}
