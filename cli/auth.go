package cli

import "github.com/0x6flab/mpesaoverlay/pkg/mpesa"

func GetToken(sdk mpesa.SDK) error {
	token, err := sdk.GetToken()
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(token)

	return nil
}
