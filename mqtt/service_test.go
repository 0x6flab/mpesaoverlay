// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mqtt_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/0x6flab/mpesaoverlay/mqtt"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa/mocks"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	invalidPacket = packets.Packet{
		Payload: []byte(`123`),
	}
	errMock        = errors.New("mock error")
	errInvalidJSON = errors.New("json: cannot unmarshal")
	validResp      = mpesa.ValidResp{
		OriginatorConversationID: "AG_20230907_2010325b025970fde878",
		ConversationID:           "AG_20230907_2010325b025970fde878",
		ResponseDescription:      "Accept the service request successfully.",
		ResponseCode:             "0",
	}
)

func TestToken(t *testing.T) {
	mockSDK := new(mocks.SDK)
	mockPacket := packets.Packet{}
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		expectedResp mpesa.TokenResp
		expectedErr  error
	}{
		{
			name: "Token success",
			expectedResp: mpesa.TokenResp{
				AccessToken: "mocked-token",
				Expiry:      "3559",
			},
			expectedErr: nil,
		},
		{
			name:         "Token error",
			expectedResp: mpesa.TokenResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("Token").Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.Token(mockPacket)

		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		assert.Equal(t, tc.expectedErr, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))

		call.Parent.AssertCalled(t, "Token")
		call.Unset()
	}
}

func TestAccountBalance(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.AccountBalanceResp
		expectedErr  error
	}{
		{
			name: "AccountBalance success",
			packet: packets.Packet{
				Payload: []byte(`{
					"Remarks": "test"
				}`),
			},
			expectedResp: mpesa.AccountBalanceResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name: "AccountBalance error",
			packet: packets.Packet{
				Payload: []byte(`{
					"Remarks": "test"
				}`),
			},
			expectedResp: mpesa.AccountBalanceResp{},
			expectedErr:  errMock,
		},
		{
			name:         "AccountBalance invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.AccountBalanceResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("AccountBalance", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.AccountBalance(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestC2BRegisterURL(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.C2BRegisterURLResp
		expectedErr  error
	}{
		{
			name: "C2BRegisterURL success",
			packet: packets.Packet{
				Payload: []byte(`{
					"ResponseType": "Completed"
				}`),
			},
			expectedResp: mpesa.C2BRegisterURLResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name: "C2BRegisterURL error",
			packet: packets.Packet{
				Payload: []byte(`{
					"ResponseType": "Completed"
				}`),
			},
			expectedResp: mpesa.C2BRegisterURLResp{},
			expectedErr:  errMock,
		},
		{
			name:         "C2BRegisterURL invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.C2BRegisterURLResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("C2BRegisterURL", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.C2BRegisterURL(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestC2BSimulate(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.C2BSimulateResp
		expectedErr  error
	}{
		{
			name: "C2BSimulate success",
			packet: packets.Packet{
				Payload: []byte(`{
					"CommandID": "CustomerBuyGoodsOnline"
				}`),
			},
			expectedResp: mpesa.C2BSimulateResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name: "C2BSimulate error",
			packet: packets.Packet{
				Payload: []byte(`{
					"CommandID": "CustomerBuyGoodsOnline"
				}`),
			},
			expectedResp: mpesa.C2BSimulateResp{},
			expectedErr:  errMock,
		},
		{
			name:         "C2BSimulate invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.C2BSimulateResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("C2BSimulate", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.C2BSimulate(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestGenerateQR(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.GenerateQRResp
		expectedErr  error
	}{
		{
			name: "GenerateQR success",
			packet: packets.Packet{
				Payload: []byte(`{
					"MerchantName": "test"
				}`),
			},
			expectedResp: mpesa.GenerateQRResp{
				ResponseDescription: "The service request is processed successfully.",
				ResponseCode:        "00",
				RequestID:           "QRCode:...",
				QRCode:              "qr_code",
			},
			expectedErr: nil,
		},
		{
			name: "GenerateQR error",
			packet: packets.Packet{
				Payload: []byte(`{
					"MerchantName": "test"
				}`),
			},
			expectedResp: mpesa.GenerateQRResp{},
			expectedErr:  errMock,
		},
		{
			name:         "GenerateQR invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.GenerateQRResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("GenerateQR", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.GenerateQR(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestExpressQuery(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.ExpressQueryResp
		expectedErr  error
	}{
		{
			name: "ExpressQuery success",
			packet: packets.Packet{
				Payload: []byte(`{
					"PassKey": "test"
				}`),
			},
			expectedResp: mpesa.ExpressQueryResp{
				ResponseDescription: "The service request has been accepted successsfully",
				ResponseCode:        "0",
				MerchantRequestID:   "92643-47073138-2",
				CheckoutRequestID:   "ws_CO_07092023195244460712345678",
				CustomerMessage:     "",
				ResultCode:          "1032",
				ResultDesc:          "Request cancelled by user",
			},
			expectedErr: nil,
		},
		{
			name: "ExpressQuery error",
			packet: packets.Packet{
				Payload: []byte(`{
					"PassKey": "test"
				}`),
			},
			expectedResp: mpesa.ExpressQueryResp{},
			expectedErr:  errMock,
		},
		{
			name:         "ExpressQuery invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.ExpressQueryResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("ExpressQuery", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.ExpressQuery(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestReverse(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.ReverseResp
		expectedErr  error
	}{
		{
			name: "Reverse success",
			packet: packets.Packet{
				Payload: []byte(`{
					"Occasion": "test"
				}`),
			},
			expectedResp: mpesa.ReverseResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name: "Reverse error",
			packet: packets.Packet{
				Payload: []byte(`{
					"Occasion": "test"
				}`),
			},
			expectedResp: mpesa.ReverseResp{},
			expectedErr:  errMock,
		},
		{
			name:         "Reverse invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.ReverseResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("Reverse", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.Reverse(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestExpressSimulate(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.ExpressSimulateResp
		expectedErr  error
	}{
		{
			name: "ExpressSimulate success",
			packet: packets.Packet{
				Payload: []byte(`{
					"PassKey": "test"
				}`),
			},
			expectedResp: mpesa.ExpressSimulateResp{
				ResponseDescription: "Success. Request accepted for processing",
				ResponseCode:        "0",
				MerchantRequestID:   "27260-79456854-2",
				CheckoutRequestID:   "ws_CO_07092023004130971712345678",
				CustomerMessage:     "Success. Request accepted for processing",
			},
			expectedErr: nil,
		},
		{
			name: "ExpressSimulate error",
			packet: packets.Packet{
				Payload: []byte(`{
					"PassKey": "test"
				}`),
			},
			expectedResp: mpesa.ExpressSimulateResp{},
			expectedErr:  errMock,
		},
		{
			name:         "ExpressSimulate invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.ExpressSimulateResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("ExpressSimulate", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.ExpressSimulate(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestRemitTax(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.RemitTaxResp
		expectedErr  error
	}{
		{
			name: "RemitTax success",
			packet: packets.Packet{
				Payload: []byte(`{
					"CommandID": "test"
				}`),
			},
			expectedResp: mpesa.RemitTaxResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name: "RemitTax error",
			packet: packets.Packet{
				Payload: []byte(`{
					"CommandID": "test"
				}`),
			},
			expectedResp: mpesa.RemitTaxResp{},
			expectedErr:  errMock,
		},
		{
			name:         "RemitTax invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.RemitTaxResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("RemitTax", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.RemitTax(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestTransactionStatus(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.TransactionStatusResp
		expectedErr  error
	}{
		{
			name: "TransactionStatus success",
			packet: packets.Packet{
				Payload: []byte(`{
					"Occasion": "test"
				}`),
			},
			expectedResp: mpesa.TransactionStatusResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name: "TransactionStatus error",
			packet: packets.Packet{
				Payload: []byte(`{
					"Occasion": "test"
				}`),
			},
			expectedResp: mpesa.TransactionStatusResp{},
			expectedErr:  errMock,
		},
		{
			name:         "TransactionStatus invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.TransactionStatusResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("TransactionStatus", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.TransactionStatus(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestB2CPayment(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := mqtt.NewService(mockSDK)

	cases := []struct {
		name         string
		packet       packets.Packet
		expectedResp mpesa.B2CPaymentResp
		expectedErr  error
	}{
		{
			name: "B2CPayment success",
			packet: packets.Packet{
				Payload: []byte(`{
					"Occasion": "test"
				}`),
			},
			expectedResp: mpesa.B2CPaymentResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name: "B2CPayment error",
			packet: packets.Packet{
				Payload: []byte(`{
					"Occasion": "test"
				}`),
			},
			expectedResp: mpesa.B2CPaymentResp{},
			expectedErr:  errMock,
		},
		{
			name:         "B2CPayment invalid payload",
			packet:       invalidPacket,
			expectedResp: mpesa.B2CPaymentResp{},
			expectedErr:  errInvalidJSON,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("B2CPayment", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.B2CPayment(tc.packet)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}
