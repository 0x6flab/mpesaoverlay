package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sdk mSDK) B2CPayment(b2cReq B2Creq) (B2CResp, error) {
	if err := b2cReq.validate(); err != nil {
		return B2CResp{}, err
	}

	var err error
	b2cReq.SecurityCredential, err = sdk.generateSecurityCredential(b2cReq.InitiatorPassword)
	if err != nil {
		return B2CResp{}, err
	}

	data, err := json.Marshal(b2cReq)
	if err != nil {
		return B2CResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, b2cEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return B2CResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return B2CResp{}, err
	}

	var b2cr B2CResp
	if err := json.Unmarshal(resp, &b2cr); err != nil {
		return B2CResp{}, err
	}

	return b2cr, nil
}
