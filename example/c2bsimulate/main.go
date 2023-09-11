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

	c2bReq := mpesa.C2BSimulateReq{
		CommandID:     "CustomerBuyGoodsOnline",
		Amount:        10,
		Msisdn:        "254712345678",
		BillRefNumber: "",
		ShortCode:     600986,
	}

	resp, err := mp.C2BSimulate(c2bReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
