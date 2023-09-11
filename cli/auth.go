package cli

import "github.com/0x6flab/mpesaoverlay/pkg"

func GetToken(sdk pkg.SDK) error {
	token, err := sdk.GetToken()
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(token)

	return nil
}
