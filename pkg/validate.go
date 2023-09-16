// This file is separated from request.go make it easy
// to generate protobuf files using https://github.com/anjmao/go2proto

package pkg

import (
	"errors"
	"net/url"
)

var (
	// errInvalidCommandID indicates the CommandID is invalid.
	errInvalidCommandID = errors.New("invalid command id")

	// errInvalidTransactionType indicates invalid transaction type.
	errInvalidTransactionType = errors.New("invalid transaction type")

	// errInvalidPhoneNumber indicates invalid phone number.
	errInvalidPhoneNumber = errors.New("invalid phone number")

	// errInvalidShortCode indicates invalid short code.
	errInvalidShortCode = errors.New("invalid short code")

	// errInvalidAccountReference indicates invalid account reference.
	errInvalidAccountReference = errors.New("invalid account reference")

	// errInvalidTransactionDesc indicates invalid transaction description.
	errInvalidTransactionDesc = errors.New("invalid transaction description")

	// errInvalidRemarks indicates invalid remarks.
	errInvalidRemarks = errors.New("invalid remarks")

	// errInvalidOccasion indicates invalid occasion.
	errInvalidOccasion = errors.New("invalid occasion")

	// errInvalidResponseType indicates invalid response type.
	errInvalidResponseType = errors.New("invalid response type")

	// errInvalidIdentifierType indicates invalid identifier type.
	errInvalidIdentifierType = errors.New("invalid identifier type")

	// errInvalidURL indicates invalid url.
	errInvalidURL = errors.New("invalid url")
)

// Validate validate the ExpressSimulateReq Request.
func (esr ExpressSimulateReq) Validate() error {
	if ok := isShortCode(esr.BusinessShortCode); !ok {
		return errInvalidShortCode
	}

	if esr.TransactionType != "CustomerPayBillOnline" && esr.TransactionType != "CustomerBuyGoodsOnline" {
		return errInvalidTransactionType
	}

	if ok := isPhoneNumber(esr.PartyA); !ok {
		return errInvalidPhoneNumber
	}

	if ok := isShortCode(esr.PartyB); !ok {
		return errInvalidShortCode
	}

	if ok := isPhoneNumber(esr.PhoneNumber); !ok {
		return errInvalidPhoneNumber
	}

	if len(esr.AccountReference) > 12 {
		return errInvalidAccountReference
	}

	if len(esr.TransactionDesc) > 13 {
		return errInvalidTransactionDesc
	}
	if ok := isValidURL(esr.CallBackURL); !ok {
		return errInvalidURL
	}

	return nil
}

// Validate validate the ExpressQueryReq Request.
func (eqr ExpressQueryReq) Validate() error {
	if ok := isShortCode(eqr.BusinessShortCode); !ok {
		return errInvalidShortCode
	}

	return nil
}

// Validate validate the GenerateQRReq Request.
func (qr GenerateQRReq) Validate() error {
	if qr.TrxCode != "SB" && qr.TrxCode != "SM" && qr.TrxCode != "PB" && qr.TrxCode != "WA" && qr.TrxCode != "BG" {
		return errInvalidTransactionType
	}

	return nil
}

func (c2b C2BRegisterURLReq) Validate() error {
	if ok := isShortCode(c2b.ShortCode); !ok {
		return errInvalidShortCode
	}
	if c2b.ResponseType != "Completed" && c2b.ResponseType != "Cancelled" {
		return errInvalidResponseType
	}
	if ok := isValidURL(c2b.ValidationURL); !ok {
		return errInvalidURL
	}
	if ok := isValidURL(c2b.ConfirmationURL); !ok {
		return errInvalidURL
	}

	return nil
}

func (c2b C2BSimulateReq) Validate() error {
	if c2b.CommandID != "CustomerPayBillOnline" && c2b.CommandID != "CustomerBuyGoodsOnline" {
		return errInvalidCommandID
	}

	return nil
}

// Validate validate the struct.
func (r B2CPaymentReq) Validate() error {
	if r.CommandID != "BusinessPayment" && r.CommandID != "SalaryPayment" && r.CommandID != "PromotionPayment" {
		return errInvalidCommandID
	}

	if ok := isShortCode(r.PartyA); !ok {
		return errInvalidShortCode
	}

	if ok := isPhoneNumber(r.PartyB); !ok {
		return errInvalidPhoneNumber
	}
	if ok := isValidURL(r.QueueTimeOutURL); !ok {
		return errInvalidURL
	}
	if ok := isValidURL(r.ResultURL); !ok {
		return errInvalidURL
	}

	return nil
}

// Validate validate the struct.
func (r TransactionStatusReq) Validate() error {
	if r.CommandID != "TransactionStatusQuery" {
		return errInvalidCommandID
	}

	if len(r.Remarks) > 100 {
		return errInvalidRemarks
	}

	if r.Occasion != "" && len(r.Occasion) > 100 {
		return errInvalidOccasion
	}

	if r.IdentifierType != 1 && r.IdentifierType != 2 && r.IdentifierType != 4 {
		return errInvalidIdentifierType
	}

	if ok := isValidURL(r.QueueTimeOutURL); !ok {
		return errInvalidURL
	}

	if ok := isValidURL(r.ResultURL); !ok {
		return errInvalidURL
	}

	return nil
}

// Validate validate the struct.
func (r AccountBalanceReq) Validate() error {
	if r.CommandID != "AccountBalance" {
		return errInvalidCommandID
	}

	if r.IdentifierType != 1 && r.IdentifierType != 2 && r.IdentifierType != 4 {
		return errInvalidIdentifierType
	}

	if ok := isValidURL(r.QueueTimeOutURL); !ok {
		return errInvalidURL
	}

	if ok := isValidURL(r.ResultURL); !ok {
		return errInvalidURL
	}

	return nil
}

// Validate validate the struct.
func (r ReverseReq) Validate() error {
	if r.CommandID != "TransactionReversal" {
		return errInvalidCommandID
	}

	if ok := isValidURL(r.QueueTimeOutURL); !ok {
		return errInvalidURL
	}

	if ok := isValidURL(r.ResultURL); !ok {
		return errInvalidURL
	}

	return nil
}

// Validate validate the struct.
func (r RemitTaxReq) Validate() error {
	if r.CommandID != "PayTaxToKRA" {
		return errInvalidCommandID
	}

	if len(r.Remarks) > 100 {
		return errInvalidRemarks
	}

	if ok := isValidURL(r.QueueTimeOutURL); !ok {
		return errInvalidURL
	}

	if ok := isValidURL(r.ResultURL); !ok {
		return errInvalidURL
	}

	return nil
}

// isPhoneNumber checks if the number is a valid phone number.
// MSISDN (12 digits Mobile Number) e.g. 2547XXXXXXXX.
func isPhoneNumber(number uint64) bool {
	if number < 100000000000 || number > 999999999999 {
		return false
	}

	return true
}

// isShortCode checks if the number is a valid short code.
// Shortcode (5 to 7 digits) e.g. 654321.
func isShortCode(number uint64) bool {
	if number < 10000 || number > 9999999 {
		return false
	}

	return true
}

func isValidURL(inputURL string) bool {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	if parsedURL.Host == "" {
		return false
	}

	return true
}
