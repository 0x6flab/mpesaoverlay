// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package grpc_test

import (
	"testing"

	"github.com/0x6flab/mpesaoverlay/grpc"
	"github.com/stretchr/testify/assert"
)

var validGRPCResp = &grpc.ValidResp{
	OriginatorConversationID: "AG_20230907_2010325b025970fde878",
	ConversationID:           "AG_20230907_2010325b025970fde878",
	ResponseDescription:      "Accept the service request successfully.",
	ResponseCode:             "0",
}

func TestAccountBalanceResp(t *testing.T) {
	resp := grpc.AccountBalanceResp{
		ValidResp: validGRPCResp,
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetValidResp()
	assert.NotEmpty(t, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestB2CPaymentResp(t *testing.T) {
	resp := grpc.B2CPaymentResp{
		ValidResp: validGRPCResp,
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetValidResp()
	assert.NotEmpty(t, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestC2BRegisterURLResp(t *testing.T) {
	resp := grpc.C2BRegisterURLResp{
		ValidResp: validGRPCResp,
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetValidResp()
	assert.NotEmpty(t, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestC2BSimulateResp(t *testing.T) {
	resp := grpc.C2BSimulateResp{
		ValidResp: validGRPCResp,
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetValidResp()
	assert.NotEmpty(t, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestExpressQueryResp(t *testing.T) {
	resp := grpc.ExpressQueryResp{
		ResponseDescription: "The service request has been accepted successsfully",
		ResponseCode:        "0",
		MerchantRequestID:   "92643-47073138-2",
		CheckoutRequestID:   "ws_CO_07092023195244460712345678",
		CustomerMessage:     "",
		ResultCode:          "1032",
		ResultDesc:          "Request cancelled by user",
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetResponseDescription()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResponseDescription, val2)

	val2 = resp.GetResponseCode()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResponseCode, val2)

	val2 = resp.GetMerchantRequestID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.MerchantRequestID, val2)

	val2 = resp.GetCheckoutRequestID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.CheckoutRequestID, val2)

	val2 = resp.GetCustomerMessage()
	assert.Empty(t, val2)
	assert.Equal(t, resp.CustomerMessage, val2)

	val2 = resp.GetResultCode()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResultCode, val2)

	val2 = resp.GetResultDesc()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResultDesc, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestExpressSimulateResp(t *testing.T) {
	resp := grpc.ExpressSimulateResp{
		ResponseDescription: "Success. Request accepted for processing",
		ResponseCode:        "0",
		MerchantRequestID:   "27260-79456854-2",
		CheckoutRequestID:   "ws_CO_07092023004130971712345678",
		CustomerMessage:     "Success. Request accepted for processing",
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetResponseDescription()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResponseDescription, val2)

	val2 = resp.GetResponseCode()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResponseCode, val2)

	val2 = resp.GetMerchantRequestID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.MerchantRequestID, val2)

	val2 = resp.GetCheckoutRequestID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.CheckoutRequestID, val2)

	val2 = resp.GetCustomerMessage()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.CustomerMessage, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestGenerateQRResp(t *testing.T) {
	resp := grpc.GenerateQRResp{
		ResponseDescription: "The service request is processed successfully.",
		ResponseCode:        "00",
		RequestID:           "QRCode:...",
		QRCode:              "qr_code",
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetResponseDescription()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResponseDescription, val2)

	val2 = resp.GetResponseCode()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResponseCode, val2)

	val2 = resp.GetRequestID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.RequestID, val2)

	val2 = resp.GetQRCode()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.QRCode, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestRemitTaxResp(t *testing.T) {
	resp := grpc.RemitTaxResp{
		ValidResp: validGRPCResp,
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetValidResp()
	assert.NotEmpty(t, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestRespError(t *testing.T) {
	resp := grpc.RespError{
		RequestID: "AG_20230907_2010325b025970fde878",
		Code:      "500",
		Message:   "Internal Server Error",
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetRequestID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.RequestID, val2)

	val2 = resp.GetCode()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.Code, val2)

	val2 = resp.GetMessage()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.Message, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestReverseResp(t *testing.T) {
	resp := grpc.ReverseResp{
		ValidResp: validGRPCResp,
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetValidResp()
	assert.NotEmpty(t, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestTokenResp(t *testing.T) {
	resp := grpc.TokenResp{
		AccessToken: "access_token",
		Expiry:      "3600",
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetAccessToken()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.AccessToken, val2)

	val2 = resp.GetExpiry()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.Expiry, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestTransactionStatusResp(t *testing.T) {
	resp := grpc.TransactionStatusResp{
		ValidResp: validGRPCResp,
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetValidResp()
	assert.NotEmpty(t, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestValidResp(t *testing.T) {
	resp := validGRPCResp

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetOriginatorConversationID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.OriginatorConversationID, val2)

	val2 = resp.GetConversationID()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ConversationID, val2)

	val2 = resp.GetResponseDescription()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResponseDescription, val2)

	val2 = resp.GetResponseCode()
	assert.NotEmpty(t, val2)
	assert.Equal(t, resp.ResponseCode, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}

func TestBusinessPayBillResp(t *testing.T) {
	resp := grpc.BusinessPayBillResp{
		ValidResp: validGRPCResp,
	}

	val := resp.String()
	assert.NotEmpty(t, val)

	resp.ProtoMessage()

	val1 := resp.ProtoReflect()
	assert.NotEmpty(t, val1)

	val2 := resp.GetValidResp()
	assert.NotEmpty(t, val2)

	val3, val4 := resp.Descriptor()
	assert.NotEmpty(t, val3)
	assert.NotEmpty(t, val4)

	resp.Reset()

	val = resp.String()
	assert.Empty(t, val)
}
