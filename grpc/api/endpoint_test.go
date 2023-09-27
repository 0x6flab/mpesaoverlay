// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package api_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	grpcadapter "github.com/0x6flab/mpesaoverlay/grpc"
	grpcapi "github.com/0x6flab/mpesaoverlay/grpc/api"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa/mocks"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	sdk       = new(mocks.SDK)
	svc       grpcadapter.Service
	port      = 8080
	errMock   = errors.New("mock error")
	validResp = mpesa.ValidResp{
		OriginatorConversationID: "AG_20230907_2010325b025970fde878",
		ConversationID:           "AG_20230907_2010325b025970fde878",
		ResponseDescription:      "Accept the service request successfully.",
		ResponseCode:             "0",
	}
)

func TestMain(m *testing.M) {
	svc = grpcadapter.NewService(sdk)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	require.Nil(&testing.T{}, err, fmt.Sprintf("unexpected error: %s\n", err))

	server := grpc.NewServer()
	grpcadapter.RegisterServiceServer(server, grpcapi.NewServer(svc))
	go func() {
		err := server.Serve(listener)
		require.Nil(&testing.T{}, err, fmt.Sprintf("unexpected error: %s\n", err))
	}()

	code := m.Run()
	os.Exit(code)
}

func TestToken(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		sdkResponse mpesa.TokenResp
		sdkError    error
	}{
		"get token success": {
			code: codes.OK,
			sdkResponse: mpesa.TokenResp{
				AccessToken: "access_token",
				Expiry:      "3599",
			},
			sdkError: nil,
		},
		"get token failure": {
			code:        codes.Internal,
			sdkResponse: mpesa.TokenResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("Token").Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.Token(ctx, &grpcadapter.Empty{})
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestAccountBalance(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.AccountBalanceReq
		sdkResponse mpesa.AccountBalanceResp
		sdkError    error
	}{
		"get account balance success": {
			code: codes.OK,
			req: &grpcadapter.AccountBalanceReq{
				InitiatorName:     "testapi",
				InitiatorPassword: "Safaricom999!*!",
				CommandID:         "AccountBalance",
				IdentifierType:    4,
				PartyA:            600772,
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
			},
			sdkResponse: mpesa.AccountBalanceResp{
				ValidResp: validResp,
			},
			sdkError: nil,
		},
		"get account balance failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.AccountBalanceResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("AccountBalance", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.AccountBalance(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestC2BRegisterURL(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.C2BRegisterURLReq
		sdkResponse mpesa.C2BRegisterURLResp
		sdkError    error
	}{
		"register c2b url success": {
			code: codes.OK,
			req: &grpcadapter.C2BRegisterURLReq{
				ShortCode:       600772,
				ResponseType:    "Completed",
				ConfirmationURL: "https://example.com/confirmation",
				ValidationURL:   "https://example.com/validation",
			},
			sdkResponse: mpesa.C2BRegisterURLResp{
				ValidResp: validResp,
			},
			sdkError: nil,
		},
		"register c2b url failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.C2BRegisterURLResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("C2BRegisterURL", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.C2BRegisterURL(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestC2BSimulate(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.C2BSimulateReq
		sdkResponse mpesa.C2BSimulateResp
		sdkError    error
	}{
		"simulate c2b success": {
			code: codes.OK,
			req: &grpcadapter.C2BSimulateReq{
				ShortCode:     600772,
				CommandID:     "CustomerPayBillOnline",
				Amount:        100,
				Msisdn:        "254708374149",
				BillRefNumber: "test",
			},
			sdkResponse: mpesa.C2BSimulateResp{
				ValidResp: validResp,
			},
			sdkError: nil,
		},
		"simulate c2b failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.C2BSimulateResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("C2BSimulate", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.C2BSimulate(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestGenerateQR(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.GenerateQRReq
		sdkResponse mpesa.GenerateQRResp
		sdkError    error
	}{
		"generate qr success": {
			code: codes.OK,
			req: &grpcadapter.GenerateQRReq{
				MerchantName: "Test Supermarket",
				RefNo:        "Invoice No",
				Amount:       2000,
				TrxCode:      "BG",
				CPI:          "174379",
				Size:         "300",
			},
			sdkResponse: mpesa.GenerateQRResp{
				ResponseDescription: "The service request is processed successfully.",
				ResponseCode:        "00",
				RequestID:           "QRCode:...",
				QRCode:              "qr_code",
			},
			sdkError: nil,
		},
		"generate qr failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.GenerateQRResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("GenerateQR", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.GenerateQR(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestExpressQuery(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.ExpressQueryReq
		sdkResponse mpesa.ExpressQueryResp
		sdkError    error
	}{
		"express query success": {
			code: codes.OK,
			req: &grpcadapter.ExpressQueryReq{
				PassKey:           "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
				BusinessShortCode: 174379,
				CheckoutRequestID: "ws_CO_07092023195244460720136609",
			},
			sdkResponse: mpesa.ExpressQueryResp{
				ResponseDescription: "The service request has been accepted successsfully",
				ResponseCode:        "0",
				MerchantRequestID:   "92643-47073138-2",
				CheckoutRequestID:   "ws_CO_07092023195244460712345678",
				CustomerMessage:     "",
				ResultCode:          "1032",
				ResultDesc:          "Request cancelled by user",
			},
			sdkError: nil,
		},
		"express query failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.ExpressQueryResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("ExpressQuery", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.ExpressQuery(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestReverse(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.ReverseReq
		sdkResponse mpesa.ReverseResp
		sdkError    error
	}{
		"reverse success": {
			code: codes.OK,
			req: &grpcadapter.ReverseReq{
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
			},
			sdkResponse: mpesa.ReverseResp{
				ValidResp: validResp,
			},
			sdkError: nil,
		},
		"reverse failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.ReverseResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("Reverse", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.Reverse(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestExpressSimulate(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.ExpressSimulateReq
		sdkResponse mpesa.ExpressSimulateResp
		sdkError    error
	}{
		"express simulate success": {
			code: codes.OK,
			req: &grpcadapter.ExpressSimulateReq{
				PassKey:           "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
				BusinessShortCode: 174379,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       254712345678,
				Amount:            10,
				PartyA:            254712345678,
				PartyB:            174379,
				CallBackURL:       "https://69a2-105-163-2-116.ngrok.io",
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   "Payment of X",
			},
			sdkResponse: mpesa.ExpressSimulateResp{
				ResponseDescription: "Success. Request accepted for processing",
				ResponseCode:        "0",
				MerchantRequestID:   "27260-79456854-2",
				CheckoutRequestID:   "ws_CO_07092023004130971712345678",
				CustomerMessage:     "Success. Request accepted for processing",
			},
			sdkError: nil,
		},
		"express simulate failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.ExpressSimulateResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("ExpressSimulate", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.ExpressSimulate(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestRemitTax(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.RemitTaxReq
		sdkResponse mpesa.RemitTaxResp
		sdkError    error
	}{
		"remit tax success": {
			code: codes.OK,
			req: &grpcadapter.RemitTaxReq{
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
			},
			sdkResponse: mpesa.RemitTaxResp{
				ValidResp: validResp,
			},
			sdkError: nil,
		},
		"remit tax failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.RemitTaxResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("RemitTax", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.RemitTax(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestTransactionStatus(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.TransactionStatusReq
		sdkResponse mpesa.TransactionStatusResp
		sdkError    error
	}{
		"transaction status success": {
			code: codes.OK,
			req: &grpcadapter.TransactionStatusReq{
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
			},
			sdkResponse: mpesa.TransactionStatusResp{
				ValidResp: validResp,
			},
			sdkError: nil,
		},
		"transaction status failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.TransactionStatusResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("TransactionStatus", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.TransactionStatus(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s", desc, tc.code, e.Code()))
		call.Unset()
	}
}

func TestB2CPayment(t *testing.T) {
	mpesaAddr := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(mpesaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	require.Nil(t, err, fmt.Sprintf("unexpected error: %s\n", err))

	cli := grpcapi.NewClient(conn, time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	cases := map[string]struct {
		code        codes.Code
		req         *grpcadapter.B2CPaymentReq
		sdkResponse mpesa.B2CPaymentResp
		sdkError    error
	}{
		"b2c payment success": {
			code: codes.OK,
			req: &grpcadapter.B2CPaymentReq{
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
			},
			sdkResponse: mpesa.B2CPaymentResp{
				ValidResp: validResp,
			},
			sdkError: nil,
		},
		"b2c payment failure": {
			code:        codes.InvalidArgument,
			sdkResponse: mpesa.B2CPaymentResp{},
			sdkError:    errMock,
		},
	}

	for desc, tc := range cases {
		call := sdk.On("B2CPayment", mock.Anything).Return(tc.sdkResponse, tc.sdkError)
		_, err := cli.B2CPayment(ctx, tc.req)
		e, ok := status.FromError(err)
		assert.True(t, ok, "OK expected to be true")
		assert.Equal(t, tc.code, e.Code(), fmt.Sprintf("%s: expected %s got %s\n", desc, tc.code, e.Code()))
		call.Unset()
	}
}
