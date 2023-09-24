package main

import (
	"log"
	"os"

	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/oklog/ulid/v2"
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

	b2cReq := mpesa.B2CPaymentReq{
		OriginatorConversationID: ulid.Make().String(),
		InitiatorName:            "testapi",
		InitiatorPassword:        "Safaricom999!*!",
		CommandID:                "BusinessPayment",
		Amount:                   10,
		PartyA:                   600986,
		PartyB:                   254712345678,
		QueueTimeOutURL:          "https://example.com/timeout",
		ResultURL:                "https://example.com/result",
		Remarks:                  "test",
		Occasion:                 "test",
	}

	resp, err := mp.B2CPayment(b2cReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
