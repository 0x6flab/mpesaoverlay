// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// This file is separated from request.go make it easy
// to generate protobuf files using https://github.com/anjmao/go2proto

package mpesa

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	maxAccountReferenceLen = 12
	maxTransactionDescLen  = 13
	maxOccasionLen         = 100
	maxRemarksLen          = 100
	customerPayBillOnline  = "CustomerPayBillOnline"
	customerBuyGoodsOnline = "CustomerBuyGoodsOnline"
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
	if !isShortCode(esr.BusinessShortCode) {
		return errInvalidShortCode
	}
	if esr.TransactionType != customerPayBillOnline && esr.TransactionType != customerBuyGoodsOnline {
		return errInvalidTransactionType
	}
	if !isPhoneNumber(esr.PartyA) || !isPhoneNumber(esr.PhoneNumber) {
		return errInvalidPhoneNumber
	}
	if !isShortCode(esr.PartyB) {
		return errInvalidShortCode
	}
	if len(esr.AccountReference) > maxAccountReferenceLen {
		return errInvalidAccountReference
	}
	if len(esr.TransactionDesc) > maxTransactionDescLen {
		return errInvalidTransactionDesc
	}
	if !isValidURL(esr.CallBackURL) {
		return errInvalidURL
	}

	return nil
}

// Validate validate the ExpressQueryReq Request.
func (eqr ExpressQueryReq) Validate() error {
	if !isShortCode(eqr.BusinessShortCode) {
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

// Validate validate the C2BRegisterURLReq Request.
func (c2b C2BRegisterURLReq) Validate() error {
	if !isShortCode(c2b.ShortCode) {
		return errInvalidShortCode
	}
	if c2b.ResponseType != "Completed" && c2b.ResponseType != "Cancelled" {
		return errInvalidResponseType
	}
	if !isValidURL(c2b.ValidationURL) || !isValidURL(c2b.ConfirmationURL) {
		return errInvalidURL
	}

	return nil
}

// Validate validate the C2BSimulateReq Request.
func (c2b C2BSimulateReq) Validate() error {
	if c2b.CommandID != customerPayBillOnline && c2b.CommandID != customerBuyGoodsOnline {
		return errInvalidCommandID
	}
	if !isShortCode(c2b.ShortCode) {
		return errInvalidShortCode
	}

	return nil
}

// Validate validate the struct.
func (r B2CPaymentReq) Validate() error {
	if r.CommandID != "BusinessPayment" && r.CommandID != "SalaryPayment" && r.CommandID != "PromotionPayment" {
		return errInvalidCommandID
	}
	if !isShortCode(r.PartyA) {
		return errInvalidShortCode
	}
	if !isPhoneNumber(r.PartyB) {
		return errInvalidPhoneNumber
	}
	if !isValidURL(r.QueueTimeOutURL) || !isValidURL(r.ResultURL) {
		return errInvalidURL
	}
	if r.Remarks != "" && len(r.Remarks) > maxRemarksLen {
		return errInvalidRemarks
	}
	if r.Occasion != "" && len(r.Occasion) > maxOccasionLen {
		return errInvalidOccasion
	}

	return nil
}

// Validate validate the struct.
func (r TransactionStatusReq) Validate() error {
	if r.CommandID != "TransactionStatusQuery" {
		return errInvalidCommandID
	}
	if r.Remarks != "" && len(r.Remarks) > maxRemarksLen {
		return errInvalidRemarks
	}
	if r.Occasion != "" && len(r.Occasion) > maxOccasionLen {
		return errInvalidOccasion
	}
	if r.IdentifierType != 1 && r.IdentifierType != 2 && r.IdentifierType != 4 {
		return errInvalidIdentifierType
	}
	if !isValidURL(r.QueueTimeOutURL) || !isValidURL(r.ResultURL) {
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
	if !isValidURL(r.QueueTimeOutURL) || !isValidURL(r.ResultURL) {
		return errInvalidURL
	}
	if r.Remarks != "" && len(r.Remarks) > maxRemarksLen {
		return errInvalidRemarks
	}
	if !isShortCode(r.PartyA) {
		return errInvalidShortCode
	}

	return nil
}

// Validate validate the struct.
func (r ReverseReq) Validate() error {
	if r.CommandID != "TransactionReversal" {
		return errInvalidCommandID
	}
	if !isValidURL(r.QueueTimeOutURL) || !isValidURL(r.ResultURL) {
		return errInvalidURL
	}
	if r.Remarks != "" && len(r.Remarks) > maxRemarksLen {
		return errInvalidRemarks
	}
	if r.Occasion != "" && len(r.Occasion) > maxOccasionLen {
		return errInvalidOccasion
	}

	return nil
}

// Validate validate the struct.
func (r RemitTaxReq) Validate() error {
	if r.CommandID != "PayTaxToKRA" {
		return errInvalidCommandID
	}
	if r.Remarks != "" && len(r.Remarks) > maxRemarksLen {
		return errInvalidRemarks
	}
	if !isValidURL(r.QueueTimeOutURL) || !isValidURL(r.ResultURL) {
		return errInvalidURL
	}
	if !isShortCode(r.PartyA) || !isShortCode(r.PartyB) {
		return errInvalidShortCode
	}
	if len(r.AccountReference) > maxAccountReferenceLen {
		return errInvalidAccountReference
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

// isValidURL checks if the url is valid.
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

// Error is the error returned by the Mpesa API.
func (e RespError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
