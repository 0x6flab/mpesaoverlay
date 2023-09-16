package cli

import (
	"github.com/0x6flab/mpesaoverlay/pkg"
	"github.com/AlecAivazis/survey/v2"
)

func Balance(sdk pkg.SDK) error {
	var req = pkg.AccountBalanceReq{}

	var qs = []*survey.Question{
		{
			Name: "InitiatorName",
			Prompt: &survey.Input{
				Message: "InitiatorName",
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
			Name: "PartyA",
			Prompt: &survey.Input{
				Message: "PartyA",
				Help:    "Organization receiving the transaction",
				Default: "600772",
			},
			Validate: survey.Required,
		},
		{
			Name: "IdentifierType",
			Prompt: &survey.Input{
				Message: "IdentifierType",
				Help:    "Type of organization querying for the account balance.",
				Default: "4",
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
				Help:    "Comments that are sent along with the transaction",
				Default: "test",
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	req.CommandID = "AccountBalance"

	balance, err := sdk.AccountBalance(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(balance)

	return nil
}
