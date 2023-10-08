// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package grpc_test

import (
	"testing"

	"github.com/0x6flab/mpesaoverlay/grpc"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func TestAccountBalanceReq(t *testing.T) {
	var req = grpc.AccountBalanceReq{
		InitiatorName:     "testapi",
		InitiatorPassword: "Safaricom999!*!",
		CommandID:         "AccountBalance",
		IdentifierType:    4,
		PartyA:            600772,
		QueueTimeOutURL:   "https://example.com/timeout",
		ResultURL:         "https://example.com/result",
		Remarks:           "test",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetInitiatorName()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorName, val2)

	val2 = req.GetInitiatorPassword()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorPassword, val2)

	val2 = req.GetCommandID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CommandID, val2)

	val3 := req.GetIdentifierType()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.IdentifierType, val3)

	val31 := req.GetPartyA()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.PartyA, val31)

	val2 = req.GetQueueTimeOutURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.QueueTimeOutURL, val2)

	val2 = req.GetResultURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ResultURL, val2)

	val2 = req.GetRemarks()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Remarks, val2)

	val4, val5 := req.Descriptor()
	assert.NotEmpty(t, val4)
	assert.NotEmpty(t, val5)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestB2CPaymentReq(t *testing.T) {
	var req = grpc.B2CPaymentReq{
		OriginatorConversationID: ulid.Make().String(),
		InitiatorName:            "testapi",
		InitiatorPassword:        "Safaricom999!*!",
		CommandID:                "BusinessPayment",
		Amount:                   10,
		PartyA:                   600986,
		PartyB:                   254712345678,
		QueueTimeOutURL:          "https://example.com/timeout",
		ResultURL:                "https://example.com/result",
		Remarks:                  "test",
		Occasion:                 "test",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetOriginatorConversationID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.OriginatorConversationID, val2)

	val2 = req.GetInitiatorName()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorName, val2)

	val2 = req.GetInitiatorPassword()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorPassword, val2)

	val2 = req.GetCommandID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CommandID, val2)

	val3 := req.GetAmount()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.Amount, val3)

	val31 := req.GetPartyA()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.PartyA, val31)

	val31 = req.GetPartyB()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.PartyB, val31)

	val2 = req.GetQueueTimeOutURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.QueueTimeOutURL, val2)

	val2 = req.GetResultURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ResultURL, val2)

	val2 = req.GetRemarks()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Remarks, val2)

	val2 = req.GetOccasion()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Occasion, val2)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestC2BRegisterURLReq(t *testing.T) {
	var req = grpc.C2BRegisterURLReq{
		ShortCode:       600981,
		ResponseType:    "Completed",
		ConfirmationURL: "https://example.com/confirmation",
		ValidationURL:   "https://example.com/validation",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val3 := req.GetShortCode()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.ShortCode, val3)

	val2 := req.GetResponseType()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ResponseType, val2)

	val2 = req.GetConfirmationURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ConfirmationURL, val2)

	val2 = req.GetValidationURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ValidationURL, val2)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestC2BSimulateReq(t *testing.T) {
	var req = grpc.C2BSimulateReq{
		CommandID:     "CustomerBuyGoodsOnline",
		Amount:        10,
		Msisdn:        "254712345678",
		BillRefNumber: "",
		ShortCode:     600986,
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetCommandID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CommandID, val2)

	val3 := req.GetAmount()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.Amount, val3)

	val31 := req.GetMsisdn()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.Msisdn, val31)

	val2 = req.GetBillRefNumber()
	assert.Empty(t, val2)
	assert.Equal(t, req.BillRefNumber, val2)

	val32 := req.GetShortCode()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.ShortCode, val32)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestExpressQueryReq(t *testing.T) {
	var req = grpc.ExpressQueryReq{
		PassKey:           "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
		BusinessShortCode: 174379,
		CheckoutRequestID: "ws_CO_07092023195244460720136609",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetPassKey()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.PassKey, val2)

	val2 = req.GetCheckoutRequestID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CheckoutRequestID, val2)

	val3 := req.GetBusinessShortCode()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.BusinessShortCode, val3)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestExpressSimulateReq(t *testing.T) {
	var req = grpc.ExpressSimulateReq{
		PassKey:           "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
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

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetPassKey()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.PassKey, val2)

	val2 = req.GetTransactionType()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.TransactionType, val2)

	val2 = req.GetCallBackURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CallBackURL, val2)

	val2 = req.GetAccountReference()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.AccountReference, val2)

	val2 = req.GetTransactionDesc()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.TransactionDesc, val2)

	val3 := req.GetBusinessShortCode()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.BusinessShortCode, val3)

	val31 := req.GetPhoneNumber()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.PhoneNumber, val31)

	val31 = req.GetAmount()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.Amount, val31)

	val31 = req.GetPartyA()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.PartyA, val31)

	val31 = req.GetPartyB()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.PartyB, val31)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestGenerateQRReq(t *testing.T) {
	var req = grpc.GenerateQRReq{
		MerchantName: "Test Supermarket",
		RefNo:        "Invoice No",
		Amount:       2000,
		TrxCode:      "BG",
		CPI:          "174379",
		Size:         "300",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetMerchantName()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.MerchantName, val2)

	val2 = req.GetRefNo()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.RefNo, val2)

	val3 := req.GetAmount()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.Amount, val3)

	val2 = req.GetTrxCode()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.TrxCode, val2)

	val2 = req.GetCPI()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CPI, val2)

	val2 = req.GetSize()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Size, val2)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestRemitTaxReq(t *testing.T) {
	var req = grpc.RemitTaxReq{
		InitiatorName:          "testapi",
		InitiatorPassword:      "Safaricom999!*!",
		CommandID:              "PayTaxToKRA",
		SenderIdentifierType:   4,
		RecieverIdentifierType: 4,
		Amount:                 239,
		PartyA:                 600978,
		PartyB:                 572572,
		AccountReference:       "353353",
		QueueTimeOutURL:        "https://example.com/timeout",
		ResultURL:              "https://example.com/result",
		Remarks:                "test",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetInitiatorName()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorName, val2)

	val2 = req.GetInitiatorPassword()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorPassword, val2)

	val2 = req.GetCommandID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CommandID, val2)

	val3 := req.GetSenderIdentifierType()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.SenderIdentifierType, val3)

	val31 := req.GetRecieverIdentifierType()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.RecieverIdentifierType, val31)

	val32 := req.GetAmount()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.Amount, val32)

	val32 = req.GetPartyA()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.PartyA, val32)

	val32 = req.GetPartyB()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.PartyB, val32)

	val2 = req.GetAccountReference()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.AccountReference, val2)

	val2 = req.GetQueueTimeOutURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.QueueTimeOutURL, val2)

	val2 = req.GetResultURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ResultURL, val2)

	val2 = req.GetRemarks()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Remarks, val2)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestReverseReq(t *testing.T) {
	var req = grpc.ReverseReq{
		InitiatorName:          "testapi",
		InitiatorPassword:      "Safaricom999!*!",
		CommandID:              "TransactionReversal",
		TransactionID:          "RI704KI9RW",
		Amount:                 10,
		ReceiverParty:          600992,
		RecieverIdentifierType: 11,
		QueueTimeOutURL:        "https://example.com/timeout",
		ResultURL:              "https://example.com/result",
		Remarks:                "test",
		Occasion:               "test",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetInitiatorName()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorName, val2)

	val2 = req.GetInitiatorPassword()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorPassword, val2)

	val2 = req.GetCommandID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CommandID, val2)

	val2 = req.GetTransactionID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.TransactionID, val2)

	val3 := req.GetAmount()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.Amount, val3)

	val31 := req.GetRecieverIdentifierType()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.RecieverIdentifierType, val31)

	val32 := req.GetAmount()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.Amount, val32)

	val32 = req.GetReceiverParty()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.ReceiverParty, val32)

	val2 = req.GetQueueTimeOutURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.QueueTimeOutURL, val2)

	val2 = req.GetResultURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ResultURL, val2)

	val2 = req.GetRemarks()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Remarks, val2)

	val2 = req.GetOccasion()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Occasion, val2)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestEmpty(t *testing.T) {
	var req = grpc.Empty{}

	val := req.String()
	assert.Empty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestTransactionStatusReq(t *testing.T) {
	var req = grpc.TransactionStatusReq{
		InitiatorName:     "testapi",
		InitiatorPassword: "Safaricom999!*!",
		CommandID:         "TransactionStatusQuery",
		IdentifierType:    1,
		TransactionID:     "RI704KI9RW",
		PartyA:            254759764065,
		QueueTimeOutURL:   "https://example.com/timeout",
		ResultURL:         "https://example.com/result",
		Remarks:           "test",
		Occasion:          "test",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetInitiatorName()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorName, val2)

	val2 = req.GetInitiatorPassword()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorPassword, val2)

	val2 = req.GetCommandID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CommandID, val2)

	val3 := req.GetIdentifierType()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.IdentifierType, val3)

	val31 := req.GetTransactionID()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.TransactionID, val31)

	val32 := req.GetPartyA()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.PartyA, val32)

	val2 = req.GetQueueTimeOutURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.QueueTimeOutURL, val2)

	val2 = req.GetResultURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ResultURL, val2)

	val2 = req.GetRemarks()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Remarks, val2)

	val2 = req.GetOccasion()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Occasion, val2)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}

func TestBusinessPayBillReq(t *testing.T) {
	var req = grpc.BusinessPayBillReq{
		Initiator:              "testapi",
		InitiatorPassword:      "Safaricom999!*!",
		CommandID:              "BusinessPayBill",
		SenderIdentifierType:   4,
		RecieverIdentifierType: 4,
		Amount:                 239,
		PartyA:                 600978,
		PartyB:                 572572,
		AccountReference:       "353353",
		Requester:              254700000000,
		QueueTimeOutURL:        "https://example.com/timeout",
		ResultURL:              "https://example.com/result",
		Remarks:                "test",
	}

	val := req.String()
	assert.NotEmpty(t, val)

	req.ProtoMessage()

	val1 := req.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := req.GetInitiator()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Initiator, val2)

	val2 = req.GetInitiatorPassword()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.InitiatorPassword, val2)

	val2 = req.GetCommandID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.CommandID, val2)

	val3 := req.GetSenderIdentifierType()
	assert.NotEmpty(t, val3)
	assert.Equal(t, req.SenderIdentifierType, val3)

	val31 := req.GetRecieverIdentifierType()
	assert.NotEmpty(t, val31)
	assert.Equal(t, req.RecieverIdentifierType, val31)

	val32 := req.GetAmount()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.Amount, val32)

	val32 = req.GetPartyA()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.PartyA, val32)

	val32 = req.GetPartyB()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.PartyB, val32)

	val32 = req.GetRequester()
	assert.NotEmpty(t, val32)
	assert.Equal(t, req.Requester, val32)

	val2 = req.GetAccountReference()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.AccountReference, val2)

	val2 = req.GetQueueTimeOutURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.QueueTimeOutURL, val2)

	val2 = req.GetResultURL()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.ResultURL, val2)

	val2 = req.GetRemarks()
	assert.NotEmpty(t, val2)
	assert.Equal(t, req.Remarks, val2)

	val9, val10 := req.Descriptor()
	assert.NotEmpty(t, val9)
	assert.NotEmpty(t, val10)

	req.Reset()

	val = req.String()
	assert.Empty(t, val)
}
