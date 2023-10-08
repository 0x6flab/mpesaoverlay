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

// BusinessPayBill initiates a B2B payment request.
func BusinessPayBill(sdk mpesa.SDK) error {
	var req = mpesa.BusinessPayBillReq{}

	var qs = []*survey.Question{
		{
			Name: "Initiator",
			Prompt: &survey.Input{
				Message: "Initiator",
				Help:    "User authorized to initiate B2C transactions via API",
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
				Help:    "Amount to be charged",
				Default: "10",
			},
			Validate: survey.Required,
		},
		{
			Name: "PartyA",
			Prompt: &survey.Input{
				Message: "PartyA",
				Help:    "B2C organization shortcode from which the money is sent from",
				Default: "600986",
			},
			Validate: survey.Required,
		},
		{
			Name: "PartyB",
			Prompt: &survey.Input{
				Message: "PartyB",
				Help:    "B2C organization shortcode from which the money is sent to",
				Default: "600986",
			},
			Validate: survey.Required,
		},
		{
			Name: "AccountReference",
			Prompt: &survey.Input{
				Message: "AccountReference",
				Help:    "Account number associated with the party sending the transaction",
				Default: "353353",
			},
			Validate: survey.Required,
		},
		{
			Name: "Requester",
			Prompt: &survey.Input{
				Message: "Requester",
				Help:    "Customer's phone number who is sending the transaction",
				Default: "254700000000",
			},
			Validate: survey.Required,
		},
		{
			Name: "Remarks",
			Prompt: &survey.Input{
				Message: "Remarks",
				Help:    "Additional information to be associated with the transaction",
				Default: "test",
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
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	req.CommandID = "BusinessPayBill"
	req.SenderIdentifierType = 4
	req.RecieverIdentifierType = 4

	b2cResp, err := sdk.BusinessPayBill(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(b2cResp)

	return nil
}
