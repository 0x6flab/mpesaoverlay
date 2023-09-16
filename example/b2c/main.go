package main

import (
	"log"
	"os"

	mpesa "github.com/0x6flab/mpesaoverlay/pkg"
	"github.com/google/uuid"
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

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	b2cReq := mpesa.B2CPaymentReq{
		OriginatorConversationID: uuid.String(),
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
