package cli

import (
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/AlecAivazis/survey/v2"
	"github.com/oklog/ulid/v2"
)

func B2CPayment(sdk mpesa.SDK) error {
	var req = mpesa.B2CPaymentReq{}

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
			Name: "CommandID",
			Prompt: &survey.Select{
				Message: "CommandID",
				Options: []string{
					"BusinessPayment",
					"SalaryPayment",
					"PromotionPayment",
				},
				Help:    "Command ID that specifies B2C transaction type",
				Default: "BusinessPayment",
			},
			Validate: survey.Required,
		},
		{
			Name: "OriginatorConversationID",
			Prompt: &survey.Input{
				Message: "OriginatorConversationID",
				Help:    "Unique unique string you specify for a transaction",
				Suggest: func(_ string) []string {
					return []string{ulid.Make().String()}
				},
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
				Help:    "Customer mobile number to receive the amount",
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
		{
			Name: "Occasion",
			Prompt: &survey.Input{
				Message: "Occasion",
				Help:    "Additional information to be associated with the transaction",
				Default: "test",
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	b2cResp, err := sdk.B2CPayment(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(b2cResp)

	return nil
}
