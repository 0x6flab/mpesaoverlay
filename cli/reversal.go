package cli

import (
	"github.com/0x6flab/mpesaoverlay/pkg"
	"github.com/AlecAivazis/survey/v2"
)

func Reversal(sdk pkg.SDK) error {
	var req = pkg.ReversalReq{}

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
			Name: "TransactionID",
			Prompt: &survey.Input{
				Message: "TransactionID",
				Help:    "Organization Receiving the funds",
			},
			Validate: survey.Required,
		},
		{
			Name: "Amount",
			Prompt: &survey.Input{
				Message: "Amount",
				Help:    "Amount to be reversed",
			},
			Validate: survey.Required,
		},
		{
			Name: "ReceiverParty",
			Prompt: &survey.Input{
				Message: "ReceiverParty",
				Help:    "Organization that receives the transaction",
				Default: "600992",
			},
			Validate: survey.Required,
		},
		{
			Name: "RecieverIdentifierType",
			Prompt: &survey.Input{
				Message: "RecieverIdentifierType",
				Help:    "Type of organization receiving the transaction",
				Default: "11",
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
			Name: "QueueTimeOutURL",
			Prompt: &survey.Input{

				Message: "QueueTimeOutURL",
				Help:    "URL to send notification incase the payment request is timed out",
				Default: "https://example.com/timeout",
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
		{
			Name: "Occasion",
			Prompt: &survey.Input{
				Message: "Occasion",
				Help:    "Optional Parameter",
				Default: "test",
			},
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	req.CommandID = "TransactionReversal"

	resp, err := sdk.Reverse(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(resp)

	return nil
}
