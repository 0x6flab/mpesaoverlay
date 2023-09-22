package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sdk mSDK) RemitTax(rReq RemitTaxReq) (RemitTaxResp, error) {
	if err := rReq.Validate(); err != nil {
		return RemitTaxResp{}, err
	}

	var err error
	rReq.SecurityCredential, err = sdk.generateSecurityCredential(rReq.InitiatorPassword)
	if err != nil {
		return RemitTaxResp{}, err
	}

	data, err := json.Marshal(rReq)
	if err != nil {
		return RemitTaxResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, taxEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return RemitTaxResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return RemitTaxResp{}, err
	}

	var tr RemitTaxResp
	if err := json.Unmarshal(resp, &tr); err != nil {
		return RemitTaxResp{}, err
	}

	return tr, nil
}
