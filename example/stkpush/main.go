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

	qrReq := mpesa.ExpressSimulateReq{
		PassKey:           "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919", // Get this from the developer portal under the test credentials section
		BusinessShortCode: 174379,
		TransactionType:   "CustomerPayBillOnline",
		PhoneNumber:       254712345678, // You can use your own phone number here
		Amount:            1,
		PartyA:            254712345678,
		PartyB:            174379,
		CallBackURL:       "https://69a2-105-163-2-116.ngrok.io",
		AccountReference:  "CompanyXLTD",
		TransactionDesc:   "Payment of X",
	}

	resp, err := mp.ExpressSimulate(qrReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
