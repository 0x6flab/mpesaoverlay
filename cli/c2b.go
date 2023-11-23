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

// C2BRegisterURL registers the confirmation and validation urls.
func C2BRegisterURL(sdk mpesa.SDK) error {
	req := mpesa.C2BRegisterURLReq{}

	qs := []*survey.Question{
		{
			Name: "ShortCode",
			Prompt: &survey.Input{
				Message: "ShortCode",
				Help:    "Organization receiving the transaction",
				Default: "600981",
			},
			Validate: survey.Required,
		},
		{
			Name: "ResponseType",
			Prompt: &survey.Select{
				Message: "ResponseType",
				Options: []string{
					"Completed",
					"Cancelled",
				},
				Help:    "Confirmation response type",
				Default: "Completed",
			},
			Validate: survey.Required,
		},
		{
			Name: "ConfirmationURL",
			Prompt: &survey.Input{
				Message: "ConfirmationURL",
				Help:    "URL that receives the confirmation request",
				Default: "https://example.com/confirmation",
			},
			Validate: survey.Required,
		},
		{
			Name: "ValidationURL",
			Prompt: &survey.Input{
				Message: "ValidationURL",
				Help:    "URL that receives the validation request",
				Default: "https://example.com/validation",
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	resp, err := sdk.C2BRegisterURL(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(resp)

	return nil
}

func C2BSimulate(sdk mpesa.SDK) error {
	req := mpesa.C2BSimulateReq{}

	qs := []*survey.Question{
		{
			Name: "CommandID",
			Prompt: &survey.Select{
				Message: "CommandID",
				Options: []string{
					"CustomerPayBillOnline",
					"CustomerBuyGoodsOnline",
				},
				Help:    "Unique command for each transaction type",
				Default: "CustomerPayBillOnline",
			},
			Validate: survey.Required,
		},
		{
			Name: "Amount",
			Prompt: &survey.Input{
				Message: "Amount",
				Default: "10",
				Help:    "Amount to be charged",
			},
			Validate: survey.Required,
		},
		{
			Name: "Msisdn",
			Prompt: &survey.Input{
				Message: "Msisdn",
				Help:    "Phone number initiating the C2B transaction",
				Default: "254708374149",
			},
			Validate: survey.Required,
		},
		{
			Name: "BillRefNumber",
			Prompt: &survey.Input{
				Message: "BillRefNumber",
				Help:    "Used on CustomerPayBillOnline option only. Unique bill identifier, e.g. an Account Number.",
			},
		},
		{
			Name: "ShortCode",
			Prompt: &survey.Input{
				Message: "ShortCode",
				Help:    "Organization receiving the transaction",
				Default: "600986",
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	resp, err := sdk.C2BSimulate(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(resp)

	return nil
}
