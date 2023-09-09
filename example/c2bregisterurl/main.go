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

	c2bReq := mpesa.C2BRegisterURLReq{
		ShortCode:       600981,
		ResponseType:    "Completed",
		ConfirmationURL: "https://8e76-105-163-2-116.ngrok.io",
		ValidationURL:   "https://8e76-105-163-2-116.ngrok.io",
	}

	resp, err := mp.C2BRegisterURL(c2bReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
