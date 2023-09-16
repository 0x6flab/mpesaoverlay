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

	trxReq := mpesa.TransactionStatusReq{
		InitiatorName:     "testapi",
		InitiatorPassword: "Safaricom999!*!",
		CommandID:         "TransactionStatusQuery",
		IdentifierType:    1,
		TransactionID:     "RI704KI9RW",
		PartyA:            254759764065,
		QueueTimeOutURL:   "https://example.com/timeout",
		ResultURL:         "https://example.com/result",
		Remarks:           "test",
		Occasion:          "test",
	}

	resp, err := mp.TransactionStatus(trxReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
