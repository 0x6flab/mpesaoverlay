package main

import (
	"log"
	"os"

	mpesa "github.com/0x6flab/mpesaoverlay/pkg"
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
		MerchantName: "Test Supermarket",
		RefNo:        "Invoice No",
		Amount:       2000,
		TrxCode:      "BG",
		CPI:          "174379",
		Size:         "300",
	}

	qrcode, err := mp.GenerateQR(qrReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("QR Code: %+v\n", qrcode)
}
