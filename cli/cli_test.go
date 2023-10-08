// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"testing"

	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa/mocks"
	"github.com/choria-io/fisk"
)

func TestToken(t *testing.T) {
	sdk := new(mocks.SDK)

	call := sdk.On("Token").Return(mpesa.TokenResp{}, nil)

	if err := Token(sdk); err != nil {
		t.Errorf("Token() error = %v", err)
	}

	call.Unset()
}

func TestSTKPush(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := STKPush(sdk); err != nil {
		t.Errorf("STKPush() error = %v", err)
	}
}

func TestB2CPayment(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := B2CPayment(sdk); err != nil {
		t.Errorf("B2CPayment() error = %v", err)
	}
}

func TestAccountBalance(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := AccountBalance(sdk); err != nil {
		t.Errorf("AccountBalance() error = %v", err)
	}
}

func TestC2BRegisterURL(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := C2BRegisterURL(sdk); err != nil {
		t.Errorf("C2BRegisterURL() error = %v", err)
	}
}

func TestSTKPushQuery(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := STKPushQuery(sdk); err != nil {
		t.Errorf("STKPushQuery() error = %v", err)
	}
}

func TestC2BSimulate(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := C2BSimulate(sdk); err != nil {
		t.Errorf("C2BSimulate() error = %v", err)
	}
}

func TestQRGenerate(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := QRCode(sdk); err != nil {
		t.Errorf("QRGenerate() error = %v", err)
	}
}

func TestReversal(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := Reversal(sdk); err != nil {
		t.Errorf("Reversal() error = %v", err)
	}
}

func TestRemitTax(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := RemitTax(sdk); err != nil {
		t.Errorf("RemitTax() error = %v", err)
	}
}

func TestTransactionStatus(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := TransactionStatus(sdk); err != nil {
		t.Errorf("TransactionStatus() error = %v", err)
	}
}

func TestBusinessPayBill(t *testing.T) {
	sdk := new(mocks.SDK)

	if err := BusinessPayBill(sdk); err != nil {
		t.Errorf("BusinessPayBill() error = %v", err)
	}
}

func TestAddCommands(_ *testing.T) {
	sdk := new(mocks.SDK)
	app := fisk.New("mpesa-cli", "0.0.1")
	AddCommands(app, sdk)
}
