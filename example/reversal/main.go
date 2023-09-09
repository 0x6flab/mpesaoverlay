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

	reverseReq := mpesa.ReversalReq{
		InitiatorName:          "testapi",
		InitiatorPassword:      "Safaricom999!*!",
		CommandID:              "TransactionReversal",
		TransactionID:          "RI704KI9RW",
		Amount:                 10,
		ReceiverParty:          600992,
		RecieverIdentifierType: 11,
		QueueTimeOutURL:        "https://8e76-105-163-2-116.ngrok.io",
		ResultURL:              "https://8e76-105-163-2-116.ngrok.io",
		Remarks:                "test",
		Occasion:               "test",
	}

	resp, err := mp.Reverse(reverseReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
