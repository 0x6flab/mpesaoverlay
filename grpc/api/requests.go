// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package api

import "github.com/0x6flab/mpesaoverlay/pkg/mpesa"

type tokenReq struct{}

func (req tokenReq) validate() error {
	return nil
}

type expressQueryReq struct {
	mpesa.ExpressQueryReq
}

func (req expressQueryReq) validate() error {
	return req.ExpressQueryReq.Validate()
}

type expressSimulateReq struct {
	mpesa.ExpressSimulateReq
}

func (req expressSimulateReq) validate() error {
	return req.ExpressSimulateReq.Validate()
}

type b2cReq struct {
	mpesa.B2CPaymentReq
}

func (req b2cReq) validate() error {
	return req.B2CPaymentReq.Validate()
}

type accountBalanceReq struct {
	mpesa.AccountBalanceReq
}

func (req accountBalanceReq) validate() error {
	return req.AccountBalanceReq.Validate()
}

type c2bRegisterURLReq struct {
	mpesa.C2BRegisterURLReq
}

func (req c2bRegisterURLReq) validate() error {
	return req.C2BRegisterURLReq.Validate()
}

type c2bSimulateReq struct {
	mpesa.C2BSimulateReq
}

func (req c2bSimulateReq) validate() error {
	return req.C2BSimulateReq.Validate()
}

type generateQRReq struct {
	mpesa.GenerateQRReq
}

func (req generateQRReq) validate() error {
	return req.GenerateQRReq.Validate()
}

type reversalReq struct {
	mpesa.ReverseReq
}

func (req reversalReq) validate() error {
	return req.ReverseReq.Validate()
}

type transactionReq struct {
	mpesa.TransactionStatusReq
}

func (req transactionReq) validate() error {
	return req.TransactionStatusReq.Validate()
}

type remitTaxReq struct {
	mpesa.RemitTaxReq
}

func (req remitTaxReq) validate() error {
	return req.RemitTaxReq.Validate()
}

type businessPayBillReq struct {
	mpesa.BusinessPayBillReq
}

func (req businessPayBillReq) validate() error {
	return req.BusinessPayBillReq.Validate()
}
