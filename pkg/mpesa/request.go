package mpesa

import "github.com/mpesaoverlay/pkg/errors"

// ExpressSimulateReq struct
type ExpressSimulateReq struct {
	BusinessShortCode string `json:"BusinessShortCode,omitempty"`
	Password          string `json:"Password,omitempty"`
	Timestamp         string `json:"Timestamp,omitempty"`
	TransactionType   string `json:"TransactionType,omitempty"`
	PhoneNumber       string `json:"PhoneNumber,omitempty"`
	Amount            string `json:"Amount,omitempty"`
	PartyA            string `json:"PartyA,omitempty"`
	PartyB            string `json:"PartyB,omitempty"`
	CallBackURL       string `json:"CallBackURL,omitempty"`
	AccountReference  string `json:"AccountReference,omitempty"`
	TransactionDesc   string `json:"TransactionDesc,omitempty"`
}

// Validate validate the ExpressSimulateReq Request
func (esr ExpressSimulateReq) Validate() error {
	if esr.TransactionType != "CustomerPayBillOnline" && esr.TransactionType != "CustomerBuyGoodsOnline" {
		return errors.ErrInvalidTransactionType
	}
	return nil
}

// ExpressQueryReq struct
type ExpressQueryReq struct {
	BusinessShortCode string `json:"BusinessShortCode,omitempty"` // This is organizations shortcode (Paybill or Buygoods - A 5 to 7 digit account number) used to identify an organization and receive the transaction.
	Password          string `json:"Password,omitempty"`          // This is the password used for encrypting the request sent: A base64 encoded string.
	Timestamp         string `json:"Timestamp,omitempty"`         // This is the Timestamp of the transaction, normaly in the formart of YEAR+MONTH+DATE+HOUR+MINUTE+SECOND (YYYYMMDDHHMMSS)
	CheckoutRequestID string `json:"CheckoutRequestID,omitempty"` // This is a global unique identifier of the processed checkout transaction request.
}

// Validate validate the ExpressQueryReq Request
func (eqr ExpressQueryReq) Validate() error {
	return nil
}

// QRReq struct
type QRReq struct {
	MerchantName string `json:"MerchantName,omitempty"` //  Name of the Company/M-Pesa Merchant Name
	RefNo        string `json:"RefNo,omitempty"`        // Transaction Reference
	Amount       string `json:"Amount,omitempty"`       //  The total amount for the sale/transaction
	TrxCode      string `json:"TrxCode,omitempty"`      // Transaction Type
	CPI          string `json:"CPI,omitempty"`          // Credit Party Identifier. Can be a Mobile Number, Business Number, Agent Till, Paybill or Business number, Merchant Buy Goods.
}

// Validate validate the QRReq Request
func (qr QRReq) Validate() error {
	if qr.TrxCode == "SB" || qr.TrxCode == "SM" || qr.TrxCode == "PB" || qr.TrxCode == "WA" || qr.TrxCode == "BG" {
		return nil
	}
	return errors.ErrInvalidTrxCode
}

// C2BRegisterURLReq struct
type C2BRegisterURLReq struct {
	ValidationURL   string `json:"ValidationURL,omitempty"`   // This is the URL that receives the validation request from API upon payment submission. The validation URL is only called if external validation on the registered shortcode is enabled. (By default external validation is disabled)
	ConfirmationURL string `json:"ConfirmationURL,omitempty"` // Thie is the URL that receives the confirmation request from API upon payment completion
	ShortCode       string `json:"ShortCode,omitempty"`       // The shortcode of the organization
	ResponseType    string `json:"ResponseType,omitempty"`
}

// C2BSimulateReq struct
type C2BSimulateReq struct {
	CommandID     string `json:"CommandID,omitempty"`     // This is a unique identifier of the transaction type: There are two types of these Identifiers:
	Msisdn        uint16 `json:"Msisdn,omitempty"`        // This is the phone number initiating the C2B transaction.
	BillRefNumber string `json:"BillRefNumber,omitempty"` // This is used on CustomerPayBillOnline option only. This is where a customer is expected to enter a unique bill identifier, e.g. an Account Number.
	Amount        uint64 `json:"Amount,omitempty"`        // This is the amount being transacted.
	ShortCode     string `json:"ShortCode,omitempty"`     // This is the Short Code receiving the amount being transacted.
}

// B2Creq struct
type B2Creq struct {
	CommandID          string `json:"CommandID,omitempty"`          // This is a unique command that specifies B2C transaction type.
	PartyA             uint8  `json:"PartyA,omitempty"`             // This is the B2C organization shortcode from which the money is sent from.
	PartyB             uint16 `json:"PartyB,omitempty"`             // This is the customer mobile number to receive the amount. - The number should have the country code (254) without the plus sign.
	Remarks            string `json:"Remarks,omitempty"`            // Any additional information to be associated with the transaction.
	InitiatorName      string `json:"InitiatorName,omitempty"`      // This is an API user created by the Business Administrator of the M-PESA Bulk disbursement account that is active and authorized to initiate B2C transactions via API.
	SecurityCredential string `json:"SecurityCredential,omitempty"` // This is the value obtained after encrypting the API initiator password. The password on Sandbox has been provisioned on the simulator. However, on production the password is created when the user is being created on the M-PESA organization portal.
	QueueTimeOutURL    string `json:"QueueTimeOutURL,omitempty"`    // This is the URL to be specified in your request that will be used by API Proxy to send notification incase the payment request is timed out while awaiting processing in the queue
	ResultURL          string `json:"ResultURL,omitempty"`          // This is the URL to be specified in your request that will be used by M-PESA to send notification upon processing of the payment request.
	TransactionID      string `json:"TransactionID,omitempty"`
	Occassion          string `json:"Occassion,omitempty"` // Any additional information to be associated with the transaction.
	Amount             string `json:"Amount,omitempty"`    // The amount of money being sent to the customer.
}

// Validate validate the struct
func (r B2Creq) Validate() error {
	if r.CommandID != "BusinessPayment" {
		return errors.ErrInvalidCommandID
	}
	return nil
}

// TransactionReq struct
type TransactionReq struct {
	CommandID          string `json:"CommandID,omitempty"`
	PartyA             string `json:"PartyA,omitempty"`             // Organization/MSISDN receiving the transaction
	IdentifierType     string `json:"IdentifierType,omitempty"`     // Type of organization receiving the transaction
	Remarks            string `json:"Remarks,omitempty"`            // Comments that are sent along with the transaction.
	Initiator          string `json:"Initiator,omitempty"`          // The name of Initiator to initiating  the request
	SecurityCredential string `json:"SecurityCredential,omitempty"` // Encrypted Credential of user getting transaction amoun
	QueueTimeOutURL    string `json:"QueueTimeOutURL,omitempty"`    // The path that stores information of time out transaction
	ResultURL          string `json:"ResultURL,omitempty"`          // The path that stores information of transaction
	TransactionID      string `json:"TransactionID,omitempty"`      // Unique identifier to identify a transaction on M-Pesa
	Occassion          string `json:"Occassion,omitempty"`
}

// Validate validate the struct
func (r TransactionReq) Validate() error {
	if r.CommandID != "TransactionStatusQuery" {
		return errors.ErrInvalidCommandID
	}
	return nil
}

// AccBalanceReq struct
type AccBalanceReq struct {
	CommandID          string `json:"CommandID,omitempty"`
	PartyA             string `json:"PartyA,omitempty"`             // Type of organization receiving the transaction
	IdentifierType     string `json:"IdentifierType,omitempty"`     // Type of organization receiving the transaction
	Remarks            string `json:"Remarks,omitempty"`            // Comments that are sent along with the transaction.
	Initiator          string `json:"Initiator,omitempty"`          // The name of Initiator to initiating  the request
	SecurityCredential string `json:"SecurityCredential,omitempty"` // Encrypted Credential of user getting transaction amoun
	QueueTimeOutURL    string `json:"QueueTimeOutURL,omitempty"`    // The path that stores information of time out transaction
	ResultURL          string `json:"ResultURL,omitempty"`          // The path that stores information of transaction
}

// Validate validate the struct
func (r AccBalanceReq) Validate() error {
	if r.CommandID != "AccountBalance" {
		return errors.ErrInvalidCommandID
	}
	return nil
}

// ReversalReq struct
type ReversalReq struct {
	CommandID              string `json:"CommandID,omitempty"`
	ReceiverParty          uint8  `json:"ReceiverParty,omitempty"`
	RecieverIdentifierType uint16 `json:"RecieverIdentifierType,omitempty"`
	Remarks                string `json:"Remarks,omitempty"`
	Initiator              string `json:"Initiator,omitempty"`
	SecurityCredential     string `json:"SecurityCredential,omitempty"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL,omitempty"`
	ResultURL              string `json:"ResultURL,omitempty"`
	TransactionID          string `json:"TransactionID,omitempty"`
	Occassion              string `json:"Occassion,omitempty"`
	Amount                 string `json:"Amount,omitempty"`
}

// Validate validate the struct
func (r ReversalReq) Validate() error {
	if r.CommandID != "TransactionReversal" {
		return errors.ErrInvalidCommandID
	}
	return nil
}
