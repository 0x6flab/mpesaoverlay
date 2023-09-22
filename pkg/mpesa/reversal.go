package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sdk mSDK) Reverse(rReq ReverseReq) (ReverseResp, error) {
	if err := rReq.Validate(); err != nil {
		return ReverseResp{}, err
	}

	var err error
	rReq.SecurityCredential, err = sdk.generateSecurityCredential(rReq.InitiatorPassword)
	if err != nil {
		return ReverseResp{}, err
	}

	data, err := json.Marshal(rReq)
	if err != nil {
		return ReverseResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, reversalEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return ReverseResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return ReverseResp{}, err
	}

	var rr ReverseResp
	if err := json.Unmarshal(resp, &rr); err != nil {
		return ReverseResp{}, err
	}

	return rr, nil
}
