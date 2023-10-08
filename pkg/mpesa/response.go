// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mpesa

// TokenResp is the response from the token endpoint.
type TokenResp struct {
	AccessToken string `json:"access_token,omitempty"` // Access token to access other APIs
	Expiry      string `json:"expires_in,omitempty"`   // Token expiry time in seconds
}

// RespError is a common response for all endpoints.
type RespError struct {
	RequestID string `json:"requestId,omitempty"`
	Code      string `json:"errorCode,omitempty"`
	Message   string `json:"errorMessage,omitempty"`
}

// ValidResp is a common response for all endpoints.
type ValidResp struct {
	OriginatorConversationID string `json:"OriginatorCoversationID,omitempty"` // The unique request ID for tracking a transaction
	ConversationID           string `json:"ConversationID,omitempty"`          // The unique request ID returned by mpesa for each request made
	ResponseDescription      string `json:"ResponseDescription,omitempty"`     // Response Description message
	ResponseCode             string `json:"ResponseCode,omitempty"`            // It indicates whether Mobile Money accepts the request or not.
}

// ExpressSimulateResp is the response from the ExpressSimulate endpoint.
type ExpressSimulateResp struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"` // Response description is an acknowledgment message from the API that gives the status of the request submission.
	ResponseCode        string `json:"ResponseCode,omitempty"`        // This is a Numeric status code that indicates the status of the transaction submission. 0 means successful submission and any other code means an error occurred.
	MerchantRequestID   string `json:"MerchantRequestID,omitempty"`   // This is a global unique Identifier for any submitted payment request.
	CheckoutRequestID   string `json:"CheckoutRequestID,omitempty"`   // This is a global unique identifier of the processed checkout transaction request.
	CustomerMessage     string `json:"CustomerMessage,omitempty"`     // This is a message that your system can display to the customer as an acknowledgment of the payment request submission.
}

// ExpressQueryResp is the response from the ExpressQuery endpoint.
type ExpressQueryResp struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"` // Response Description message. It can be a Success submission message or an error description.
	ResponseCode        string `json:"ResponseCode,omitempty"`        // This is a numeric status code that indicates the status of the transaction submission. 0 means successful submission and any other code means an error occurred.
	MerchantRequestID   string `json:"MerchantRequestID,omitempty"`   // This is a global unique Identifier for any submitted payment request.
	CheckoutRequestID   string `json:"CheckoutRequestID,omitempty"`   // This is a global unique identifier of the processed checkout transaction request.
	CustomerMessage     string `json:"CustomerMessage,omitempty"`     // This is a message that your system can display to the Customer as an acknowledgement of the payment request submission.
	ResultCode          string `json:"ResultCode,omitempty"`          // This is a numeric status code that indicates the status of the transaction processing. 0 means successful processing and any other code means an error occurred or the transaction failed.
	ResultDesc          string `json:"ResultDesc,omitempty"`          // Response description is an acknowledgment message from the API that gives the status of the request submission usually maps to a specific ResponseCode value. It can be a "Success" submission message or an error description.
}

// GenerateQRResp is the response from the GenerateQR endpoint.
type GenerateQRResp struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"` // This is a response describing the status of the transaction.
	ResponseCode        string `json:"ResponseCode,omitempty"`        // Used to return the Transaction Type.
	RequestID           string `json:"RequestID,omitempty"`
	QRCode              string `json:"QRCode,omitempty"` // QR Code Image Data/String/Image.
}

// C2BRegisterURLResp is the response from the C2BRegisterURL endpoint.
type C2BRegisterURLResp struct {
	ValidResp
}

// C2BSimulateResp is the response from the C2BSimulate endpoint.
type C2BSimulateResp struct {
	ValidResp
}

// B2CPaymentResp is the response from the B2CPayment endpoint.
type B2CPaymentResp struct {
	ValidResp
}

// TransactionStatusResp is the response from the TransactionStatus endpoint.
type TransactionStatusResp struct {
	ValidResp
}

// AccountBalanceResp is the response from the AccountBalance endpoint.
type AccountBalanceResp struct {
	ValidResp
}

// ReverseResp is the response from the Reverse endpoint.
type ReverseResp struct {
	ValidResp
}

// RemitTaxResp is the response from the RemitTax endpoint.
type RemitTaxResp struct {
	ValidResp
}

// BusinessPayBillResp is the response from the BusinessPayBill endpoint.
type BusinessPayBillResp struct {
	ValidResp
}
