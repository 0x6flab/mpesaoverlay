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

func RemitTax(sdk mpesa.SDK) error {
	var req = mpesa.RemitTaxReq{}

	var qs = []*survey.Question{
		{
			Name: "InitiatorName",
			Prompt: &survey.Input{
				Message: "InitiatorName",
				Help:    "Username of the API Initiator",
				Default: "testapi",
			},
			Validate: survey.Required,
		},
		{
			Name: "InitiatorPassword",
			Prompt: &survey.Password{
				Message: "InitiatorPassword",
				Help:    "Password of the API user",
			},
			Validate: survey.Required,
		},
		{
			Name: "Amount",
			Prompt: &survey.Input{
				Message: "Amount",
				Help:    "Amount to be transferred",
			},
			Validate: survey.Required,
		},
		{
			Name: "SenderIdentifierType",
			Prompt: &survey.Input{
				Message: "SenderIdentifierType",
				Help:    "The type of shortcode from which money is deducted.",
				Default: "4",
			},
			Validate: survey.Required,
		},
		{
			Name: "RecieverIdentifierType",
			Prompt: &survey.Input{
				Message: "RecieverIdentifierType",
				Help:    "The type of shortcode to which money is credited.",
				Default: "4",
			},
			Validate: survey.Required,
		},
		{
			Name: "PartyA",
			Prompt: &survey.Input{
				Message: "PartyA",
				Help:    "This is your own shortcode from which the money will be deducted",
				Default: "600978",
			},
			Validate: survey.Required,
		},
		{
			Name: "PartyB",
			Prompt: &survey.Input{
				Message: "PartyB",
				Help:    "The account to which money will be credited",
				Default: "572572",
			},
			Validate: survey.Required,
		},
		{
			Name: "AccountReference",
			Prompt: &survey.Input{
				Message: "AccountReference",
				Help:    "The payment registration number (PRN) issued by KRA",
				Default: "353353",
			},
			Validate: survey.Required,
		},
		{
			Name: "QueueTimeOutURL",
			Prompt: &survey.Input{
				Message: "QueueTimeOutURL",
				Help:    "URL to send notification incase the payment request is timed out",
				Default: "https://example.com/timeout",
			},
			Validate: survey.Required,
		},
		{
			Name: "ResultURL",
			Prompt: &survey.Input{
				Message: "ResultURL",
				Help:    "URL to send notification upon completion of the request",
				Default: "https://example.com/result",
			},
			Validate: survey.Required,
		},
		{
			Name: "Remarks",
			Prompt: &survey.Input{
				Message: "Remarks",
				Help:    "Comments that are sent along with the transaction.",
				Default: "test",
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	req.CommandID = "PayTaxToKRA"

	resp, err := sdk.RemitTax(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(resp)

	return nil
}
