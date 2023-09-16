package cli

import (
	"github.com/0x6flab/mpesaoverlay/pkg"
	"github.com/AlecAivazis/survey/v2"
)

func QRCode(sdk pkg.SDK) error {
	var req = pkg.GenerateQRReq{}

	var qs = []*survey.Question{
		{
			Name: "MerchantName",
			Prompt: &survey.Input{
				Message: "Mpesa Overlay",
				Help:    "Name of the Company/M-Pesa Merchant Name",
				Default: "600981",
			},
			Validate: survey.Required,
		},
		{
			Name: "RefNo",
			Prompt: &survey.Input{
				Message: "Reference Number",
				Help:    "Transaction Reference",
				Default: "Invoice No",
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
			Name: "TrxCode",
			Prompt: &survey.Select{
				Message: "Transaction Code",
				Help:    "Transaction Type",
				Options: []string{
					"SB",
					"WA",
					"PB",
					"SM",
					"BG",
				},
				Default: "BG",
			},
			Validate: survey.Required,
		},
		{
			Name: "CPI",
			Prompt: &survey.Input{
				Message: "Credit Party Identifier",
				Help:    "Credit Party Identifier",
				Default: "174379",
			},
			Validate: survey.Required,
		},
		{
			Name: "Size",
			Prompt: &survey.Input{
				Message: "Size",
				Help:    "Size of the QR code image in pixels",
				Default: "300",
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(qs, &req, survey.WithHideCharacter('*'), survey.WithShowCursor(true)); err != nil {
		logError(err)

		return nil
	}

	resp, err := sdk.GenerateQR(req)
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(resp)

	return nil
}
