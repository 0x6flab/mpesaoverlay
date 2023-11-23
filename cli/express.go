// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/AlecAivazis/survey/v2"
)

func STKPush(sdk mpesa.SDK) error {
	req := mpesa.ExpressSimulateReq{}

	qs := []*survey.Question{
		{
			Name: "PassKey",
			Prompt: &survey.Password{
				Message: "PassKey",
				Help:    "Lipa Na Mpesa Online PassKey",
			},
			Validate: survey.Required,
		},
		{
			Name: "BusinessShortCode",
			Prompt: &survey.Input{
				Message: "BusinessShortCode",
				Help:    "Lipa Na Mpesa Online ShortCode",
				Default: "174379",
			},
			Validate: survey.Required,
		},
		{
			Name: "TransactionType",
			Prompt: &survey.Select{
				Message: "TransactionType",
				Options: []string{
					"CustomerPayBillOnline",
					"CustomerBuyGoodsOnline",
				},
				Help:    "Lipa Na Mpesa Online TransactionType",
				Default: "CustomerPayBillOnline",
			},
			Validate: survey.Required,
		},
		{
			Name: "PhoneNumber",
			Prompt: &survey.Input{
				Message: "PhoneNumber",
				Help:    "PhoneNumber to receive the STK Pin Prompt (format: 2547XXXXXXXX)",
			},
			Validate: survey.Required,
		},
		{
			Name: "Amount",
			Prompt: &survey.Input{
				Message: "Amount",
				Help:    "Amount to be charged",
				Default: "1",
			},
			Validate: survey.Required,
		},
		{
			Name: "PartyA",
			Prompt: &survey.Input{
				Message: "PartyA",
				Help:    "PhoneNumber phone number sending money (format: 2547XXXXXXXX)",
			},
			Validate: survey.Required,
		},
		{
			Name: "PartyB",
			Prompt: &survey.Input{
				Message: "PartyB",
				Help:    "Organization that receives the funds",
				Default: "174379",
			},
			Validate: survey.Required,
		},
		{
			Name: "CallBackURL",
			Prompt: &survey.Input{
				Message: "CallBackURL",
				Help:    "Callback URL used to receive notifications from M-Pesa API",
				Default: "https://example.com/callback",
			},
			Validate: survey.Required,
		},
		{
			Name: "AccountReference",
			Prompt: &survey.Input{
				Message: "AccountReference",
				Help:    "Alpha-Numeric parameter that is defined by your system as an Identifier of the transaction",
				Default: "MpesaOverlay",
			},
			Validate: survey.Required,
		},
		{
			Name: "TransactionDesc",
			Prompt: &survey.Input{
				Message: "TransactionDesc",
				Help:    "Additional information that can be sent along with the request",
				Default: "Payment of X",
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	resp, err := sdk.ExpressSimulate(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(resp)

	return nil
}

func STKPushQuery(sdk mpesa.SDK) error {
	req := mpesa.ExpressQueryReq{}

	qs := []*survey.Question{
		{
			Name: "PassKey",
			Prompt: &survey.Password{
				Message: "PassKey",
				Help:    "Lipa Na Mpesa Online PassKey",
			},
			Validate: survey.Required,
		},
		{
			Name: "BusinessShortCode",
			Prompt: &survey.Input{
				Message: "BusinessShortCode",
				Help:    "Lipa Na Mpesa Online ShortCode",
				Default: "174379",
			},
			Validate: survey.Required,
		},
		{
			Name: "CheckoutRequestID",
			Prompt: &survey.Input{
				Message: "CheckoutRequestID",
				Help:    "CheckoutRequestID",
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	resp, err := sdk.ExpressQuery(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(resp)

	return nil
}
