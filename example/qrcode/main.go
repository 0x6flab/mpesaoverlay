package main

import (
	"fmt"
	"os"

	"github.com/mpesaoverlay/pkg/mpesa"
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
	mp := mpesa.NewSDK(conf)
	qrReq := mpesa.QRReq{
		MerchantName: "Safaricom LTD",
		RefNo:        "rf38f04",
		Amount:       "20000",
		TrxCode:      "BG",
		CPI:          "17408",
	}
	qrcode, err := mp.GenerateQR(qrReq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(qrcode)
}
