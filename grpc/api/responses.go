// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package api

import "github.com/0x6flab/mpesaoverlay/pkg/mpesa"

type tokenResp struct {
	mpesa.TokenResp
}

type expressQueryResp struct {
	mpesa.ExpressQueryResp
}

type expressSimulateResp struct {
	mpesa.ExpressSimulateResp
}

type b2cResp struct {
	mpesa.B2CPaymentResp
}

type accountBalanceResp struct {
	mpesa.AccountBalanceResp
}

type c2bRegisterURLResp struct {
	mpesa.C2BRegisterURLResp
}

type c2bSimulateResp struct {
	mpesa.C2BSimulateResp
}

type generateQRResp struct {
	mpesa.GenerateQRResp
}

type reverseResp struct {
	mpesa.ReverseResp
}

type transactionStatusResp struct {
	mpesa.TransactionStatusResp
}

type remitTaxResp struct {
	mpesa.RemitTaxResp
}

type businessPayBillResp struct {
	mpesa.BusinessPayBillResp
}
