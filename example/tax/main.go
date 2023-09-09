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

	taxReq := mpesa.RemitTax{
		InitiatorName:          "testapi",
		InitiatorPassword:      "Safaricom999!*!",
		CommandID:              "PayTaxToKRA",
		SenderIdentifierType:   4,
		RecieverIdentifierType: 4,
		Amount:                 239,
		PartyA:                 600978,
		PartyB:                 572572,
		AccountReference:       "353353",
		QueueTimeOutURL:        "https://8e76-105-163-2-116.ngrok.io",
		ResultURL:              "https://8e76-105-163-2-116.ngrok.io",
		Remarks:                "test",
	}

	resp, err := mp.RemitTax(taxReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
