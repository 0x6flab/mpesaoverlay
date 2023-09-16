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

	balReq := mpesa.AccountBalanceReq{
		InitiatorName:     "testapi",
		InitiatorPassword: "Safaricom999!*!",
		CommandID:         "AccountBalance",
		IdentifierType:    4,
		PartyA:            600772,
		QueueTimeOutURL:   "https://example.com/timeout",
		ResultURL:         "https://example.com/result",
		Remarks:           "test",
	}

	resp, err := mp.AccountBalance(balReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
