// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mpesa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defaultTimeout = 1 * time.Minute

	authEndpoint            = "oauth/v1/generate?grant_type=client_credentials"
	b2cEndpoint             = "mpesa/b2c/v1/paymentrequest"
	accbalanceEndpoint      = "mpesa/accountbalance/v1/query"
	c2bRegisterURLEndpoint  = "mpesa/c2b/v1/registerurl"
	c2bSimulateEndpoint     = "mpesa/c2b/v1/simulate"
	qrCodeEndpoint          = "mpesa/qrcode/v1/generate"
	expressSimulateEndpoint = "mpesa/stkpush/v1/processrequest"
	queryEndpoint           = "mpesa/stkpushquery/v1/query"
	reversalEndpoint        = "mpesa/reversal/v1/request"
	transactionEndpoint     = "mpesa/transactionstatus/v1/query"
	taxEndpoint             = "mpesa/b2b/v1/remittax"
	prodCertificate         = "https://developer.safaricom.co.ke/api/v1/GenerateSecurityCredential/ProductionCertificate.cer"
	sandboxCertificate      = "https://developer.safaricom.co.ke/api/v1/GenerateSecurityCredential/SandboxCertificate.cer"
)

var (
	errInvalidBaseURL   = errors.New("invalid base url, must be either https://api.safaricom.co.ke or https://sandbox.safaricom.co.ke")
	errMissingBaseURL   = errors.New("missing base url")
	errMissingAppKey    = errors.New("missing app key")
	errMissingAppSecret = errors.New("missing app secret")
	errFailedToSendReq  = errors.New("failed to send request")
)

var _ SDK = (*mSDK)(nil)

// SDK contains MpesaOverlay interface API.
type SDK interface {
	// GetToken gives you a time bound access token to call allowed APIs.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/Authorization
	//
	// Example:
	// 	token, err := mp.GetToken()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("Token: %+v\n", token)
	// Output:
	// 	2023/09/06 22:43:21 Token: {AccessToken:unU9joKpPqIsZ1jFiDmQoNJ1cIvK Expiry:3599}
	Token() (TokenResp, error)

	// ExpressQuery check the status of a Lipa Na M-Pesa Online Payment.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/MpesaExpressQuery
	//
	// Example:
	// 	eqReq := mpesa.ExpressQueryReq{
	// 		PassKey:           "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919", // Get this from the developer portal under the test credentials section
	// 		BusinessShortCode: 174379,
	// 		CheckoutRequestID: "ws_CO_07092023195244460712345678",
	// 	}
	//
	// 	resp, err := mp.ExpressQuery(eqReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 20:08:05 Resp: {ResponseDescription:The service request has been accepted successsfully ResponseCode:0 MerchantRequestID:92643-47073138-2 CheckoutRequestID:ws_CO_07092023195244460712345678 CustomerMessage: ResultCode:1032 ResultDesc:Request cancelled by user}
	ExpressQuery(eqReq ExpressQueryReq) (ExpressQueryResp, error)

	// ExpressSimulate initiates online payment on behalf of a customer.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/MpesaExpressSimulate
	//
	// Example:
	// 	qrReq := mpesa.ExpressSimulateReq{
	// 		PassKey:           "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919", // Get this from the developer portal under the test credentials section
	// 		BusinessShortCode: 174379,
	// 		TransactionType:   "CustomerPayBillOnline",
	// 		PhoneNumber:       254712345678, // You can use your own phone number here
	// 		Amount:            10,
	// 		PartyA:            254712345678,
	// 		PartyB:            174379,
	// 		CallBackURL:       "https://69a2-105-163-2-116.ngrok.io",
	// 		AccountReference:  "CompanyXLTD",
	// 		TransactionDesc:   "Payment of X",
	// 	}
	//
	// 	resp, err := mp.ExpressSimulate(qrReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 00:39:08 Resp: {ResponseDescription:Success. Request accepted for processing ResponseCode:0 MerchantRequestID:27260-79456854-2 CheckoutRequestID:ws_CO_07092023004130971712345678 CustomerMessage:Success. Request accepted for processing}
	ExpressSimulate(eReq ExpressSimulateReq) (ExpressSimulateResp, error)

	// B2CPayment Transact between an M-Pesa short code to a phone number registered on M-Pesa
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/BusinessToCustomer
	//
	// Example:
	// 	b2cReq := mpesa.B2CPaymentReq{
	//      OriginatorConversationID: uuid.String(),
	// 		InitiatorName:            "testapi",
	// 		SecurityCredential:       "Safaricom111!",
	// 		CommandID:                "BusinessPayment",
	// 		Amount:                   10,
	// 		PartyA:                   174379,
	// 		PartyB:                   254712345678,
	// 		Remarks:                  "Test",
	// 		QueueTimeOutURL:          "https://69a2-105-163-2-116.ngrok.io",
	// 		ResultURL:                "https://69a2-105-163-2-116.ngrok.io",
	// 		Occasion:                 "Test",
	// 	}
	//
	// 	resp, err := mp.B2CPayment(b2cReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 21:12:26 Resp: {ValidResp:{OriginatorConversationID: ConversationID:AG_20230907_2010325b025970fde878 ResponseDescription:Accept the service request successfully. ResponseCode:0}}
	B2CPayment(b2cReq B2CPaymentReq) (B2CPaymentResp, error)

	// AccountBalance Enquire the balance on an M-Pesa BuyGoods (Till Number)
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/AccountBalance
	//
	// Example:
	// 	balReq := mpesa.AccountBalanceReq{
	// 		InitiatorName:     "testapi",
	// 		InitiatorPassword: "Safaricom999!*!",
	// 		CommandID:         "AccountBalance",
	// 		IdentifierType:    4,
	// 		PartyA:            600772,
	// 		QueueTimeOutURL:   "https://example.com/timeout",
	// 		ResultURL:         "https://example.com/result",
	// 		Remarks:           "test",
	// 	}
	//
	// 	resp, err := mp.AccountBalance(balReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 22:05:44 Resp: {ValidResp:{OriginatorConversationID: ConversationID:AG_20230907_201045e9b4e4f9bcb4d6 ResponseDescription:Accept the service request successfully. ResponseCode:0}}
	AccountBalance(abReq AccountBalanceReq) (AccountBalanceResp, error)

	// C2BRegisterURL register validation and confirmation URLs on M-Pesa
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/CustomerToBusinessRegisterURL
	//
	// Example:
	// 	c2bReq := mpesa.C2BRegisterURLReq{
	// 		ShortCode: 174379,
	// 		ResponseType: "Completed",
	// 		ConfirmationURL: "https://69a2-105-163-2-116.ngrok.io",
	// 		ValidationURL: "https://69a2-105-163-2-116.ngrok.io",
	// 	}
	//
	// 	resp, err := mp.C2BRegisterURL(c2bReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 20:23:39 Resp: {ValidResp:{OriginatorConversationID:29607-261203248-2 ConversationID: ResponseDescription:Success ResponseCode:0}}
	C2BRegisterURL(c2bReq C2BRegisterURLReq) (C2BRegisterURLResp, error)

	// C2BSimulate Make payment requests from Client to Business (C2B)
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/CustomerToBusinessRegisterURL
	//
	// Example:
	// 	c2bReq := mpesa.C2BSimulateReq{
	// 		ShortCode: 174379,
	// 		CommandID: "CustomerPayBillOnline",
	// 		Amount: 10,
	// 		Msisdn: 254712345678,
	// 		BillRefNumber: "",
	// 	}
	//
	// 	resp, err := mp.C2BSimulate(c2bReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 20:33:56 Resp: {ValidResp:{OriginatorConversationID:92647-47234949-2 ConversationID: ResponseDescription:Accept the service request successfully. ResponseCode:0}}
	C2BSimulate(c2bReq C2BSimulateReq) (C2BSimulateResp, error)

	// GenerateQR generates a dynamic M-PESA QR Code.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/DynamicQRCode
	//
	// Example:
	// 	qrReq := mpesa.GenerateQRReq{
	// 		MerchantName: "Test Supermarket",
	// 		RefNo:        "Invoice No",
	// 		Amount:       "2000",
	// 		TrxCode:      "BG",
	// 		CPI:          "174379",
	// 		Size:         "300",
	// 	}
	// 	qrcode, err := mp.GenerateQR(qrReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("QR Code: %+v\n", qrcode)
	// Output:
	//  2023/09/06 23:22:51 QR Code: {ResponseDescription:The service request is processed successfully. ResponseCode:00 RequestID: QRCode:...}
	GenerateQR(qReq GenerateQRReq) (GenerateQRResp, error)

	// Reverse Reverses an M-Pesa transaction.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/Reversal
	//
	// Example:
	// 	rReq := mpesa.ReverseReq{
	// 		InitiatorName:          "testapi",
	// 		InitiatorPassword:      "Safaricom999!*!",
	// 		CommandID:              "TransactionReversal",
	// 		TransactionID:          "RI704KI9RW",
	// 		Amount:                 10,
	// 		ReceiverParty:          600992,
	// 		RecieverIdentifierType: 11,
	// 		QueueTimeOutURL:        "https://example.com/timeout",
	// 		ResultURL:              "https://example.com/result",
	// 		Remarks:                "test",
	// 		Occasion:               "test",
	// 	}
	//
	// 	resp, err := mp.Reverse(rReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 22:11:13 Resp: {ValidResp:{OriginatorConversationID: ConversationID:AG_20230907_20106204c62f8f1a3f21 ResponseDescription:Accept the service request successfully. ResponseCode:0}}
	Reverse(rReq ReverseReq) (ReverseResp, error)

	// TransactionStatus check the status of a transaction
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/TransactionStatus
	//
	// Example:
	// 	tReq := mpesa.TransactionStatusReq{
	// 		InitiatorName:     "testapi",
	// 		InitiatorPassword: "Safaricom999!*!",
	// 		CommandID:         "TransactionStatusQuery",
	// 		IdentifierType:    1,
	// 		TransactionID:     "RI704KI9RW",
	// 		PartyA:            254759764065,
	// 		QueueTimeOutURL:   "https://example.com/timeout",
	// 		ResultURL:         "https://example.com/result",
	// 		Remarks:           "test",
	// 		Occasion:          "test",
	// 	}
	//
	// 	resp, err := mp.TransactionStatus(tReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 21:56:50 Resp: {ValidResp:{OriginatorConversationID: ConversationID:AG_20230907_20102e33b7103b4f7b0e ResponseDescription:Accept the service request successfully. ResponseCode:0}}
	TransactionStatus(tReq TransactionStatusReq) (TransactionStatusResp, error)

	// RemitTax enables businesses to remit tax to Kenya Revenue Authority (KRA).
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/TaxRemittance
	//
	// Example:
	// 	taxReq := mpesa.RemitTax{
	// 		InitiatorName:          "testapi",
	// 		InitiatorPassword:      "Safaricom999!*!",
	// 		CommandID:              "PayTaxToKRA",
	// 		SenderIdentifierType:   4,
	// 		RecieverIdentifierType: 4,
	// 		Amount:                 239,
	// 		PartyA:                 600978,
	// 		PartyB:                 572572,
	// 		AccountReference:       "353353",
	// 		QueueTimeOutURL:        "https://example.com/timeout",
	// 		ResultURL:              "https://example.com/result",
	// 		Remarks:                "test",
	// 	}
	//
	// 	resp, err := mp.RemitTax(taxReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	log.Printf("Resp: %+v\n", resp)
	// Output:
	//  2023/09/07 22:30:00 Resp: {ValidResp:{OriginatorConversationID: ConversationID:AG_20230907_201001484b176c67b3fb ResponseDescription:Accept the service request successfully. ResponseCode:0}}
	RemitTax(rReq RemitTaxReq) (RemitTaxResp, error)
}

// mSDK implements SDK interface.
type mSDK struct {
	baseURL           string
	appKey            string
	appSecret         string
	certFile          string
	client            *http.Client
	initiatorName     string
	initiatorPassword string
}

// Config contains sdk configuration parameters.
type Config struct {
	BaseURL           string
	AppKey            string
	AppSecret         string
	CertFile          string
	HTTPClient        *http.Client
	InitiatorName     string
	InitiatorPassword string
}

// validate validates the configuration parameters.
func (cfg Config) validate() error {
	if cfg.BaseURL == "" {
		return errMissingBaseURL
	}

	if cfg.BaseURL != "https://api.safaricom.co.ke" && cfg.BaseURL != "https://sandbox.safaricom.co.ke" {
		return errInvalidBaseURL
	}

	if cfg.AppKey == "" {
		return errMissingAppKey
	}

	if cfg.AppSecret == "" {
		return errMissingAppSecret
	}

	return nil
}

// newSDK returns new mpesa SDK instance.
func newSDK(conf Config) (SDK, error) {
	conf.CertFile = prodCertificate
	if strings.Contains(conf.BaseURL, "sandbox") {
		conf.CertFile = sandboxCertificate
	}

	conf.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
		Timeout: defaultTimeout,
	}

	if err := conf.validate(); err != nil {
		return nil, err
	}

	sdk := &mSDK{
		baseURL:           conf.BaseURL,
		appKey:            conf.AppKey,
		appSecret:         conf.AppSecret,
		certFile:          conf.CertFile,
		client:            conf.HTTPClient,
		initiatorName:     conf.InitiatorName,
		initiatorPassword: conf.InitiatorPassword,
	}

	return sdk, nil
}

// NewSDK returns new mpesa SDK instance.
func NewSDK(conf Config, opts ...Option) (SDK, error) {
	sdk, err := newSDK(conf)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		sdk, err = opt(sdk)
		if err != nil {
			return nil, err
		}
	}

	return sdk, nil
}

// sendRequest sends a request to the Mpesa API.
func (sdk mSDK) sendRequest(req *http.Request) ([]byte, error) {
	token, err := sdk.Token()
	if err != nil {
		return nil, err
	}

	if token.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, errors.Join(errFailedToSendReq, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp RespError
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response: %w", err)
		}

		return nil, errors.Join(errFailedToSendReq, errResp)
	}

	return body, nil
}

// generateTimestampAndPassword generates a timestamp and password.
func (sdk mSDK) generateTimestampAndPassword(shortcode uint64, passkey string) (string, string) {
	timestamp := time.Now().Local().Format("20060102150405")
	password := fmt.Sprintf("%d%s%s", shortcode, passkey, timestamp)

	return timestamp, base64.StdEncoding.EncodeToString([]byte(password))
}

// generateSecurityCredential generates a security credential.
func (sdk mSDK) generateSecurityCredential(password string) (string, error) {
	resp, err := http.Get(sdk.certFile)
	if err != nil {
		return "", fmt.Errorf("failed to get certificate: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read certificate: %w", err)
	}

	cBlock, _ := pem.Decode(data)

	cert, err := x509.ParseCertificate(cBlock.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse certificate: %w", err)
	}

	pubKey := cert.PublicKey.(*rsa.PublicKey)

	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(password))
	if err != nil {
		return "", fmt.Errorf("failed to encrypt password: %w", err)
	}

	return base64.StdEncoding.EncodeToString(cipher), nil
}
