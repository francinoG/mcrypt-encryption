package main

import (
	"encoding/json"
	"fmt"
	"mcrypt-encryption/mcrypt"
)

var (
	ApiKey = "fKQGN3snxNJryjCM"
)

type Data struct {
	Voucher        string `json:"voucher"`
	CustomerNumber string `json:"customer_number"`
	ProductType    string `json:"product_type,omitempty"`
}

type DataSign struct {
	SignData  string `json:"sign_data"`
	PartnerId string `json:"partner_id"`
}

func main() {
	voucher_code := "hOQ9LJ"
	data := Data{
		Voucher:        voucher_code,
		CustomerNumber: "081234000001",
	}
	str, _ := json.Marshal(data)
	key := []byte(ApiKey)
	signData := mcrypt.Encrypt(key, []byte(str))

	encryptedData := DataSign{
		SignData:  signData,
		PartnerId: "50911",
	}

	fmt.Println("encryptedData >>>", encryptedData)

	// signData = "QsZ5XIucj1GV5c1828eVg3VqFZKjI3b7aawmtC99xR+CQAabu0TgS0SDpp6kl5gIh0hmLVJU3id1RVDo1SxXSQ=="
	decryptData := mcrypt.Decrypt(key, signData)
	fmt.Println("decryptedData >>>", decryptData)
}
