// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
)

// Service is the interface that provides methods for the MpesaOverlay SDK.
type Service interface {
	Token() (mpesa.TokenResp, error)

	ExpressQuery(eqReq mpesa.ExpressQueryReq) (mpesa.ExpressQueryResp, error)

	ExpressSimulate(eReq mpesa.ExpressSimulateReq) (mpesa.ExpressSimulateResp, error)

	B2CPayment(b2cReq mpesa.B2CPaymentReq) (mpesa.B2CPaymentResp, error)

	AccountBalance(abReq mpesa.AccountBalanceReq) (mpesa.AccountBalanceResp, error)

	C2BRegisterURL(c2bReq mpesa.C2BRegisterURLReq) (mpesa.C2BRegisterURLResp, error)

	C2BSimulate(c2bReq mpesa.C2BSimulateReq) (mpesa.C2BSimulateResp, error)

	GenerateQR(qReq mpesa.GenerateQRReq) (mpesa.GenerateQRResp, error)

	Reverse(rReq mpesa.ReverseReq) (mpesa.ReverseResp, error)

	TransactionStatus(tReq mpesa.TransactionStatusReq) (mpesa.TransactionStatusResp, error)

	RemitTax(rReq mpesa.RemitTaxReq) (mpesa.RemitTaxResp, error)
}

// service implements the Service interface.
type service struct {
	sdk mpesa.SDK
}

var _ Service = (*service)(nil)

// NewService returns a new gRPC service.
func NewService(sdk mpesa.SDK) Service {
	return &service{sdk: sdk}
}

func (s *service) Token() (mpesa.TokenResp, error) {
	return s.sdk.Token()
}

func (s *service) ExpressQuery(eqReq mpesa.ExpressQueryReq) (mpesa.ExpressQueryResp, error) {
	return s.sdk.ExpressQuery(eqReq)
}

func (s *service) ExpressSimulate(eReq mpesa.ExpressSimulateReq) (mpesa.ExpressSimulateResp, error) {
	return s.sdk.ExpressSimulate(eReq)
}

func (s *service) B2CPayment(b2cReq mpesa.B2CPaymentReq) (mpesa.B2CPaymentResp, error) {
	return s.sdk.B2CPayment(b2cReq)
}

func (s *service) AccountBalance(abReq mpesa.AccountBalanceReq) (mpesa.AccountBalanceResp, error) {
	return s.sdk.AccountBalance(abReq)
}

func (s *service) C2BRegisterURL(c2bReq mpesa.C2BRegisterURLReq) (mpesa.C2BRegisterURLResp, error) {
	return s.sdk.C2BRegisterURL(c2bReq)
}

func (s *service) C2BSimulate(c2bReq mpesa.C2BSimulateReq) (mpesa.C2BSimulateResp, error) {
	return s.sdk.C2BSimulate(c2bReq)
}

func (s *service) GenerateQR(qReq mpesa.GenerateQRReq) (mpesa.GenerateQRResp, error) {
	return s.sdk.GenerateQR(qReq)
}

func (s *service) Reverse(rReq mpesa.ReverseReq) (mpesa.ReverseResp, error) {
	return s.sdk.Reverse(rReq)
}

func (s *service) TransactionStatus(tReq mpesa.TransactionStatusReq) (mpesa.TransactionStatusResp, error) {
	return s.sdk.TransactionStatus(tReq)
}

func (s *service) RemitTax(rReq mpesa.RemitTaxReq) (mpesa.RemitTaxResp, error) {
	return s.sdk.RemitTax(rReq)
}
