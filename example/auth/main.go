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
	token, err := mp.GetToken()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(token.AccessToken)
}
