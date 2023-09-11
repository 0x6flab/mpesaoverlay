package cli

import (
	"github.com/0x6flab/mpesaoverlay/pkg"
	"github.com/AlecAivazis/survey/v2"
)

func TransactionStatus(sdk pkg.SDK) error {
	var req = pkg.TransactionReq{}

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
			Name: "IdentifierType",
			Prompt: &survey.Input{
				Message: "IdentifierType",
				Help:    "Type of organization receiving the transaction",
				Default: "1",
			},
			Validate: survey.Required,
		},
		{
			Name: "PartyA",
			Prompt: &survey.Input{
				Message: "PartyA",
				Help:    "Organization/MSISDN receiving the transaction",
				Default: "254759764065",
			},
			Validate: survey.Required,
		},
		{
			Name: "TransactionID",
			Prompt: &survey.Input{
				Message: "TransactionID",
				Help:    "Unique identifier to identify a transaction on M-Pesa",
				Default: "RI704KI9RW",
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
		{
			Name: "Occasion",
			Prompt: &survey.Input{
				Message: "Occasion",
				Help:    "Optional parameter",
				Default: "test",
			},
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	req.CommandID = "TransactionStatusQuery"

	resp, err := sdk.TransactionStatus(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(resp)

	return nil
}
