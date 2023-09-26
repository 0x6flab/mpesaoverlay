// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/choria-io/fisk"
)

// AddCommands adds the mpesa commands to the application.
func AddCommands(app *fisk.Application, sdk mpesa.SDK) {
	token := app.Command("token", "Get a token")
	token.Action(func(ctx *fisk.ParseContext) error {
		return Token(sdk)
	})
	token.Alias("auth")
	token.Cheat("token", `Get an access token
For example: mpesa-cli token`)

	stkpush := app.Command("stkpush", "Simulate STK Push")
	stkpush.Action(func(_ *fisk.ParseContext) error {
		return STKPush(sdk)
	})
	stkpush.Cheat("stkpush", `Simulate STK Push
For example: mpesa-cli stkpush`)
	stkpush.Alias("express")

	stkpushquery := app.Command("stkpushquery", "Query STK Push")
	stkpushquery.Action(func(_ *fisk.ParseContext) error {
		return STKPushQuery(sdk)
	})
	stkpushquery.Cheat("stkpushquery", `Query STK Push
For example: mpesa-cli stkpushquery`)
	stkpushquery.Alias("expressquery")

	b2c := app.Command("b2c", "Simulate B2C Payment")
	b2c.Action(func(_ *fisk.ParseContext) error {
		return B2CPayment(sdk)
	})
	b2c.Cheat("b2c", `Simulate B2C Payment
For example: mpesa-cli b2c`)
	b2c.Alias("pay")

	balance := app.Command("balance", "Check Account Balance")
	balance.Action(func(_ *fisk.ParseContext) error {
		return AccountBalance(sdk)
	})
	balance.Cheat("balance", `Check Account Balance
For example: mpesa-cli balance`)
	balance.Alias("accbalance")
	balance.Alias("bal")

	c2bregisterurl := app.Command("c2bregisterurl", "Register C2B URL")
	c2bregisterurl.Action(func(_ *fisk.ParseContext) error {
		return C2BRegisterURL(sdk)
	})
	c2bregisterurl.Cheat("c2bregisterurl", `Register C2B URL
For example: mpesa-cli c2bregisterurl`)
	c2bregisterurl.Alias("registerurl")
	c2bregisterurl.Alias("regurl")

	c2bsimulate := app.Command("c2bsimulate", "Simulate C2B Payment")
	c2bsimulate.Action(func(_ *fisk.ParseContext) error {
		return C2BSimulate(sdk)
	})
	c2bsimulate.Cheat("c2bsimulate", `Simulate C2B Payment
For example: mpesa-cli c2bsimulate`)
	c2bsimulate.Alias("simulate")
	c2bsimulate.Alias("c2b")

	qrcode := app.Command("qrcode", "Generate QR Code")
	qrcode.Action(func(_ *fisk.ParseContext) error {
		return QRCode(sdk)
	})
	qrcode.Cheat("qrcode", `Generate QR Code
For example: mpesa-cli qrcode`)
	qrcode.Alias("qr")
	qrcode.Alias("code")

	reversal := app.Command("reversal", "Simulate Reversal")
	reversal.Action(func(_ *fisk.ParseContext) error {
		return Reversal(sdk)
	})
	reversal.Cheat("reversal", `Simulate Reversal
For example: mpesa-cli reversal`)
	reversal.Alias("reverse")
	reversal.Alias("rev")

	remittax := app.Command("remittax", "Simulate Remittance Tax")
	remittax.Action(func(_ *fisk.ParseContext) error {
		return RemitTax(sdk)
	})
	remittax.Cheat("remittax", `Simulate Remittance Tax
For example: mpesa-cli remittax`)
	remittax.Alias("tax")
	reversal.Alias("remit")

	transactionstatus := app.Command("transactionstatus", "Simulate Transaction Status")
	transactionstatus.Action(func(_ *fisk.ParseContext) error {
		return TransactionStatus(sdk)
	})
	transactionstatus.Cheat("transactionstatus", `Simulate Transaction Status
For example: mpesa-cli transactionstatus`)
	transactionstatus.Alias("status")
	transactionstatus.Alias("transstatus")
	transactionstatus.Alias("transstat")
}
