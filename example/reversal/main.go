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
		BaseURL:   "https://sandbox.safaricom.co.ke",
		AppKey:    cKey,
		AppSecret: cSecret,
	}

	mp, err := mpesa.NewSDK(conf)
	if err != nil {
		log.Fatal(err)
	}

	reverseReq := mpesa.ReverseReq{
		InitiatorName:          "testapi",
		InitiatorPassword:      "Safaricom999!*!",
		CommandID:              "TransactionReversal",
		TransactionID:          "RI704KI9RW",
		Amount:                 10,
		ReceiverParty:          600992,
		ReceiverIdentifierType: 11,
		QueueTimeOutURL:        "https://example.com/timeout",
		ResultURL:              "https://example.com/result",
		Remarks:                "test",
		Occasion:               "test",
	}

	resp, err := mp.Reverse(reverseReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
