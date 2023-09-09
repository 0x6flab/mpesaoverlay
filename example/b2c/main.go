package main

import (
	"log"
	"os"

	"github.com/google/uuid"
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

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	b2cReq := mpesa.B2Creq{
		OriginatorConversationID: uuid.String(),
		InitiatorName:            "testapi",
		InitiatorPassword:        "Safaricom999!*!",
		CommandID:                "BusinessPayment",
		Amount:                   10,
		PartyA:                   600986,
		PartyB:                   254720136609,
		QueueTimeOutURL:          "https://8e76-105-163-2-116.ngrok.io",
		ResultURL:                "https://8e76-105-163-2-116.ngrok.io",
		Remarks:                  "test",
		Occasion:                 "test",
	}

	resp, err := mp.B2CPayment(b2cReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
