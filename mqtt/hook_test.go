// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mqtt

import (
	"errors"
	"fmt"
	"testing"

	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa/mocks"
	mochimqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

var (
	invalidPayload = []byte(`123`)
	errMock        = errors.New("mock error")
	errInvalidJSON = errors.New("json: cannot unmarshal")
	validResp      = mpesa.ValidResp{
		OriginatorConversationID: "AG_20230907_2010325b025970fde878",
		ConversationID:           "AG_20230907_2010325b025970fde878",
		ResponseDescription:      "Accept the service request successfully.",
		ResponseCode:             "0",
	}
)

func TestSetServer(_ *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name   string
		server *mochimqtt.Server
	}{
		{
			name:   "set server",
			server: &mochimqtt.Server{},
		},
	}

	for _, c := range cases {
		hook.SetServer(c.server)
	}
}

func TestID(t *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name string
		want string
	}{
		{
			name: "get hook id",
			want: "mpesaoverlay-mqtt",
		},
	}

	for _, c := range cases {
		got := hook.ID()
		assert.Equal(t, c.want, got, "ID() = %v, want %v", got, c.want)
	}
}

func TestProvides(t *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name string
		b    byte
		want bool
	}{
		{
			name: "provides on connect",
			b:    mochimqtt.OnConnect,
			want: true,
		},
		{
			name: "provides on disconnect",
			b:    mochimqtt.OnDisconnect,
			want: true,
		},
		{
			name: "provides on published",
			b:    mochimqtt.OnPublished,
			want: true,
		},
		{
			name: "provides on subscribed",
			b:    mochimqtt.OnSubscribed,
			want: true,
		},
		{
			name: "provides on unsubscribed",
			b:    mochimqtt.OnUnsubscribed,
			want: true,
		},
		{
			name: "does not provide",
			b:    0x00,
			want: false,
		},
	}

	for _, c := range cases {
		ok := hook.Provides(c.b)
		if ok != c.want {
			assert.False(t, ok, fmt.Sprintf("%s %s", c.name, "failed"))
		}
	}
}

func TestInit(t *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name string
	}{
		{
			name: "init hook",
		},
	}

	for _, c := range cases {
		err := hook.Init(nil)
		assert.NoError(t, err, fmt.Sprintf("%s %s", c.name, "failed"))
	}
}

func TestOnConnect(t *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name string
	}{
		{
			name: "on connect",
		},
	}

	for _, c := range cases {
		err := hook.OnConnect(&mochimqtt.Client{}, packets.Packet{})
		assert.NoError(t, err, fmt.Sprintf("%s %s", c.name, "failed"))
	}
}

func TestOnDisconnect(t *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name string
	}{
		{
			name: "on disconnect",
		},
	}

	for _, c := range cases {
		hook.OnDisconnect(&mochimqtt.Client{}, nil, false)
		assert.NoError(t, nil, fmt.Sprintf("%s %s", c.name, "failed"))
	}
}

func TestOnPublished(t *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name string
	}{
		{
			name: "on published",
		},
	}

	for _, c := range cases {
		hook.OnPublished(&mochimqtt.Client{}, packets.Packet{})
		assert.NoError(t, nil, fmt.Sprintf("%s %s", c.name, "failed"))
	}
}

func TestOnSubscribed(t *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name string
	}{
		{
			name: "on subscribed",
		},
	}

	for _, c := range cases {
		hook.OnSubscribed(&mochimqtt.Client{}, packets.Packet{}, []byte{})
		assert.NoError(t, nil, fmt.Sprintf("%s %s", c.name, "failed"))
	}
}

func TestOnUnsubscribed(t *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)

	cases := []struct {
		name string
	}{
		{
			name: "on unsubscribed",
		},
	}

	for _, c := range cases {
		hook.OnUnsubscribed(&mochimqtt.Client{}, packets.Packet{})
		assert.NoError(t, nil, fmt.Sprintf("%s %s", c.name, "failed"))
	}
}

func TestHandleMessages(_ *testing.T) {
	mockSDK := new(mocks.SDK)
	svc := NewService(mockSDK)
	hook := NewHook(zap.NewNop(), svc)
	server := mochimqtt.New(&mochimqtt.Options{InlineClient: true})
	hook.SetServer(server)

	cases := []struct {
		name    string
		topic   string
		payload []byte
		mockErr error
	}{
		{
			name:    "handle mpesa/token success",
			topic:   "mpesa/token",
			payload: []byte(``),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/token failure",
			topic:   "mpesa/token",
			payload: []byte(``),
			mockErr: errMock,
		},
		{
			name:  "handle mpesa/express/query success",
			topic: "mpesa/express/query",
			payload: []byte(`{
				"PassKey": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/express/query failure",
			topic:   "mpesa/express/query",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/express/simulate success",
			topic: "mpesa/express/simulate",
			payload: []byte(`{
				"CommandID": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/express/simulate failure",
			topic:   "mpesa/express/simulate",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/b2c/payment success",
			topic: "mpesa/b2c/payment",
			payload: []byte(`{
				"Occasion": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/b2c/payment failure",
			topic:   "mpesa/b2c/payment",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/account/balance success",
			topic: "mpesa/account/balance",
			payload: []byte(`{
				"Remarks": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/account/balance failure",
			topic:   "mpesa/account/balance",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/c2b/register success",
			topic: "mpesa/c2b/register",
			payload: []byte(`{
				"ResponseType": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/c2b/register failure",
			topic:   "mpesa/c2b/register",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/c2b/simulate success",
			topic: "mpesa/c2b/simulate",
			payload: []byte(`{
				"CommandID": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/c2b/simulate failure",
			topic:   "mpesa/c2b/simulate",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/generate/qr success",
			topic: "mpesa/generate/qr",
			payload: []byte(`{
				"MerchantName": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/generate/qr failure",
			topic:   "mpesa/generate/qr",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/reverse success",
			topic: "mpesa/reverse",
			payload: []byte(`{
				"Occasion": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/reverse failure",
			topic:   "mpesa/reverse",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/transaction/status success",
			topic: "mpesa/transaction/status",
			payload: []byte(`{
				"Occasion": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/transaction/status failure",
			topic:   "mpesa/transaction/status",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/remit/tax success",
			topic: "mpesa/remit/tax",
			payload: []byte(`{
				"CommandID": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/remit/tax failure",
			topic:   "mpesa/remit/tax",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
		{
			name:  "handle mpesa/b2b/payment success",
			topic: "mpesa/b2b/payment",
			payload: []byte(`{
				"CommandID": "test"
			}`),
			mockErr: nil,
		},
		{
			name:    "handle mpesa/b2b/payment failure",
			topic:   "mpesa/b2b/payment",
			payload: invalidPayload,
			mockErr: errInvalidJSON,
		},
	}

	for _, c := range cases {
		call1 := mockSDK.On("Token").Return(mpesa.TokenResp{}, c.mockErr)
		call2 := mockSDK.On("ExpressQuery", mock.Anything).Return(mpesa.ExpressQueryResp{
			ResponseDescription: "The service request has been accepted successsfully",
			ResponseCode:        "0",
			MerchantRequestID:   "92643-47073138-2",
			CheckoutRequestID:   "ws_CO_07092023195244460712345678",
			CustomerMessage:     "",
			ResultCode:          "1032",
			ResultDesc:          "Request cancelled by user",
		}, c.mockErr)
		call3 := mockSDK.On("ExpressSimulate", mock.Anything).Return(mpesa.ExpressSimulateResp{
			ResponseDescription: "Success. Request accepted for processing",
			ResponseCode:        "0",
			MerchantRequestID:   "27260-79456854-2",
			CheckoutRequestID:   "ws_CO_07092023004130971712345678",
			CustomerMessage:     "Success. Request accepted for processing",
		}, c.mockErr)
		call4 := mockSDK.On("B2CPayment", mock.Anything).Return(mpesa.B2CPaymentResp{
			ValidResp: validResp,
		}, c.mockErr)
		call5 := mockSDK.On("AccountBalance", mock.Anything).Return(mpesa.AccountBalanceResp{
			ValidResp: validResp,
		}, c.mockErr)
		call6 := mockSDK.On("C2BRegisterURL", mock.Anything).Return(mpesa.C2BRegisterURLResp{
			ValidResp: validResp,
		}, c.mockErr)
		call7 := mockSDK.On("C2BSimulate", mock.Anything).Return(mpesa.C2BSimulateResp{
			ValidResp: validResp,
		}, c.mockErr)
		call8 := mockSDK.On("GenerateQR", mock.Anything).Return(mpesa.GenerateQRResp{
			ResponseDescription: "The service request is processed successfully.",
			ResponseCode:        "00",
			RequestID:           "QRCode:...",
			QRCode:              "qr_code",
		}, c.mockErr)
		call9 := mockSDK.On("Reverse", mock.Anything).Return(mpesa.ReverseResp{
			ValidResp: validResp,
		}, c.mockErr)
		call10 := mockSDK.On("TransactionStatus", mock.Anything).Return(mpesa.TransactionStatusResp{
			ValidResp: validResp,
		}, c.mockErr)
		call11 := mockSDK.On("RemitTax", mock.Anything).Return(mpesa.RemitTaxResp{
			ValidResp: validResp,
		}, c.mockErr)
		call12 := mockSDK.On("BusinessPayBill", mock.Anything).Return(mpesa.BusinessPayBillResp{
			ValidResp: validResp,
		}, c.mockErr)

		hook.handleMessages(packets.Packet{
			TopicName: c.topic,
			Payload:   c.payload,
		})

		call1.Unset()
		call2.Unset()
		call3.Unset()
		call4.Unset()
		call5.Unset()
		call6.Unset()
		call7.Unset()
		call8.Unset()
		call9.Unset()
		call10.Unset()
		call11.Unset()
		call12.Unset()
	}
}
