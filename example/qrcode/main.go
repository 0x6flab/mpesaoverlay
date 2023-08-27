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
	
	qrReq := mpesa.QRReq{
		MerchantName: "Safaricom LTD",
		RefNo:        "rf38f04",
		Amount:       "20000",
		TrxCode:      "BG",
		CPI:          "17408",
	}
	
	qrcode, err := mp.GenerateQR(qrReq)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println(qrcode.QRCode)
}
