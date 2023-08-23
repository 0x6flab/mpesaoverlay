package main

import (
	"log"
	"os"

	mpesa "github.com/mpesaoverlay/pkg"
)

var (
	cKey    = os.Getenv("MPESA_CONSUMER_KEY")
	cSecret = os.Getenv("MPESA_CONSUMER_SECRET")
)

func main() {
	conf := mpesa.Config{
		BaseURL:      "https://sandbox.safaricom.co.ke",
		AppKey:       cKey,
		AppSecret:    cSecret,
		MaxIdleConns: 10,
	}
	mp, err := mpesa.NewSDK(conf)
	if err != nil {
		log.Fatal(err)
	}
	qrReq := mpesa.ExpressSimulateReq{
		BusinessShortCode: "174379",
		TransactionType:   "CustomerPayBillOnline",
		PhoneNumber:       "",
		Amount:            "10",
		PartyA:            "",
		PartyB:            "174379",
		CallBackURL:       "rodneyosodo.com",
		AccountReference:  "Test",
		TransactionDesc:   "Test",
	}
	resp, err := mp.ExpressSimulate(qrReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
