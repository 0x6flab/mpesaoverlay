package pkg

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
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
	reversalEndpoint    = "mpesa/reversal/v1/request"
	transactionEndpoint = "mpesa/transactionstatus/v1/query"
)

var _ SDK = (*mSDK)(nil)

// SDK contains MpesaOverlay interface API.
type SDK interface {
	// GetToken Gives you a time bound access token to call allowed APIs.
	GetToken() (TokenResp, error)

	// ExpressQuery Check the status of a Lipa Na M-Pesa Online Payment.
	ExpressQuery(eqReq ExpressQueryReq) (ExpressQueryResp, error)

	// ExpressSimulate Initiates online payment on behalf of a customer.
	ExpressSimulate(eReq ExpressSimulateReq) (ExpressSimulateResp, error)

	// B2CPayment Transact between an M-Pesa short code to a phone number registered on M-Pesa
	B2CPayment(b2cReq B2Creq) (B2CResp, error)

	// AccountBalance Enquire the balance on an M-Pesa BuyGoods (Till Number)
	AccountBalance(abReq AccBalanceReq) (AccBalanceResp, error)

	// C2BRegisterURL Register validation and confirmation URLs on M-Pesa
	C2BRegisterURL(c2bReq C2BRegisterURLReq) (C2BRegisterURLResp, error)

	// C2BSimulate Make payment requests from Client to Business (C2B)
	C2BSimulate(c2bReq C2BSimulateReq) (C2BSimulateResp, error)

	// GenerateQR Generates a dynamic M-PESA QR Code.
	GenerateQR(qReq QRReq) (QRResp, error)

	// Reverse Reverses an M-Pesa transaction.
	Reverse(rReq ReversalReq) (ReversalResp, error)

	// TransactionStatus Check the status of a transaction
	TransactionStatus(tReq TransactionReq) (TransactionResp, error)
}

type mSDK struct {
	ctx       context.Context
	baseURL   string
	appKey    string
	appSecret string
	client    *http.Client
}

// Config contains sdk configuration parameters.
type Config struct {
	CTX          context.Context
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
func NewSDK(conf Config, opts ...SDKOption) (SDK, error) {
	if err := conf.validate(); err != nil {
		return nil, err
	}

	sdk := &mSDK{
		baseURL:   conf.BaseURL,
		appKey:    conf.AppKey,
		appSecret: conf.AppSecret,
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

	req = req.WithContext(sdk.ctx)

	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	return body, nil
}

func (sdk mSDK) generateTimestampAndPassword(shortcode, passkey string) (string, string) {
	timestamp := time.Now().Local().Format("20060102150405")
	password := fmt.Sprintf("%s%s%s", shortcode, passkey, timestamp)

	return timestamp, base64.StdEncoding.EncodeToString([]byte(password))
}

// GetSecurityCredential generates a security credential.
func (sdk mSDK) GetSecurityCredential(password string) (string, error) {
	fileName := "https://developer.safaricom.co.ke/sites/default/files/cert/cert_prod/cert.cer"
	if strings.Contains(sdk.baseURL, "sandbox") {
		fileName = "https://developer.safaricom.co.ke/sites/default/files/cert/cert_sandbox/cert.cer"
	}

	resp, err := http.Get(fileName)
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
