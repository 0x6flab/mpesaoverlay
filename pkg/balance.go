package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sdk mSDK) AccountBalance(abReq AccBalanceReq) (AccBalanceResp, error) {
	if err := abReq.validate(); err != nil {
		return AccBalanceResp{}, err
	}

	var err error
	abReq.SecurityCredential, err = sdk.generateSecurityCredential(abReq.InitiatorPassword)
	if err != nil {
		return AccBalanceResp{}, err
	}

	data, err := json.Marshal(abReq)
	if err != nil {
		return AccBalanceResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, accbalanceEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return AccBalanceResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return AccBalanceResp{}, err
	}

	var abr AccBalanceResp
	if err := json.Unmarshal(resp, &abr); err != nil {
		return AccBalanceResp{}, err
	}

	return abr, nil
}
