package main

import (
	"fmt"
	"os"

	"github.com/mpesaoverlay/pkg/mpesa"
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
	mp := mpesa.NewSDK(conf)
	token, err := mp.GetToken()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token.AccessToken)
}
