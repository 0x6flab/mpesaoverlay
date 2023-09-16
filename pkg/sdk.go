package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defaultTimeout = 1 * time.Minute

	authEndpoint        = "oauth/v1/generate?grant_type=client_credentials"
	b2cEndpoint         = "mpesa/b2c/v1/paymentrequest"
	accbalanceEndpoint  = "mpesa/accountbalance/v1/query"
	c2bEndpoint         = "mpesa/c2b/v1"
	qrCodeEndpoint      = "mpesa/qrcode/v1/generate"
	expressEndpoint     = "mpesa/stkpush/v1"
	queryEndpoint       = "mpesa/stkpushquery/v1/query"
	reversalEndpoint    = "mpesa/reversal/v1/request"
	transactionEndpoint = "mpesa/transactionstatus/v1/query"
	taxEndpoint         = "mpesa/b2b/v1/remittax"
	prodCertificate     = "https://developer.safaricom.co.ke/api/v1/GenerateSecurityCredential/ProductionCertificate.cer"
	sandboxCertificate  = "https://developer.safaricom.co.ke/api/v1/GenerateSecurityCredential/SandboxCertificate.cer"
)

var _ SDK = (*mSDK)(nil)

// SDK contains MpesaOverlay interface API.
type SDK interface {
	// GetToken Gives you a time bound access token to call allowed APIs.
	//
	// The token is valid for the specified time duration, which is usually an hour.
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
	GetToken() (TokenResp, error)

	// ExpressQuery Check the status of a Lipa Na M-Pesa Online Payment.
	//
	// Query the payment status of a Lipa Na M-Pesa Online Payment using the M-Pesa transaction ID.
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

	// ExpressSimulate Initiates online payment on behalf of a customer.
	//
	// Sends a USSD push to the customer’s phone to prompt them to enter their PIN to authorize the payment.
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
	// B2C API is an API used to make payments from a Business to Customers (Pay Outs), also known as Bulk Disbursements.
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

	// C2BRegisterURL Register validation and confirmation URLs on M-Pesa
	//
	// Register URL API works hand in hand with Customer to Business (C2B) APIs and allows receiving payment notifications to your paybill.
	//
	// This API enables you to register the callback URLs via which you shall receive notifications for payments to your pay bill/till number.
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

	// GenerateQR Generates a dynamic M-PESA QR Code.
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

	// TransactionStatus Check the status of a transaction
	//
	// Check the status of a transaction.
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

type mSDK struct {
	baseURL   string
	appKey    string
	appSecret string
	certFile  string
	client    *http.Client
}

// Config contains sdk configuration parameters.
type Config struct {
	BaseURL      string
	AppKey       string
	AppSecret    string
	MaxIdleConns int
}

func (cfg Config) validate() error {
	if cfg.BaseURL == "" {
		return fmt.Errorf("base url is required")
	}

	if cfg.BaseURL != "https://api.safaricom.co.ke" && cfg.BaseURL != "https://sandbox.safaricom.co.ke" {
		return fmt.Errorf("base url must be either https://api.safaricom.co.ke or https://sandbox.safaricom.co.ke")
	}

	if cfg.AppKey == "" {
		return fmt.Errorf("app key is required")
	}

	if cfg.AppSecret == "" {
		return fmt.Errorf("app secret is required")
	}

	return nil
}

// NewSDK returns new mpesa SDK instance.
func NewSDK(conf Config, opts ...Options) (SDK, error) {
	if err := conf.validate(); err != nil {
		return nil, err
	}

	fileName := prodCertificate
	if strings.Contains(conf.BaseURL, "sandbox") {
		fileName = sandboxCertificate
	}

	sdk := &mSDK{
		baseURL:   conf.BaseURL,
		appKey:    conf.AppKey,
		appSecret: conf.AppSecret,
		certFile:  fileName,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
				},
				MaxIdleConns: conf.MaxIdleConns,
			},
			Timeout: defaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(&conf)
	}

	return sdk, nil
}

func (sdk mSDK) sendRequest(req *http.Request) ([]byte, error) {
	token, err := sdk.GetToken()
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
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResp
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response: %w", err)
		}

		return nil, fmt.Errorf("failed to send request: requestID %s, errorCode %s, errorMessage %s", errResp.RequestID, errResp.Code, errResp.Message)
	}

	return body, nil
}

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
