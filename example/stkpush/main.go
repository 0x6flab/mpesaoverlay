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
	qrReq := mpesa.ExpressSimulateReq{
		BusinessShortCode: "174379",
		TransactionType:   "CustomerPayBillOnline",
		PhoneNumber:       "",
		Amount:            "10",
		PartyA:            "",
		PartyB:            "174379",
		CallBackURL:       "rodneyosodo.com",
		AccountReference:  "Test",
		TransactionDesc:   "Test",
	}
	resp, err := mp.ExpressSimulate(qrReq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
