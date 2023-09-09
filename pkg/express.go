package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sdk mSDK) ExpressSimulate(eReq ExpressSimulateReq) (ExpressSimulateResp, error) {
	if err := eReq.validate(); err != nil {
		return ExpressSimulateResp{}, err
	}

	eReq.Timestamp, eReq.Password = sdk.generateTimestampAndPassword(eReq.BusinessShortCode, eReq.PassKey)

	data, err := json.Marshal(eReq)
	if err != nil {
		return ExpressSimulateResp{}, err
	}

	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, expressEndpoint, "processrequest")

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return ExpressSimulateResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return ExpressSimulateResp{}, err
	}

	var esr ExpressSimulateResp
	if err := json.Unmarshal(resp, &esr); err != nil {
		return ExpressSimulateResp{}, err
	}

	return esr, nil
}

func (sdk mSDK) ExpressQuery(eqReq ExpressQueryReq) (ExpressQueryResp, error) {
	if err := eqReq.validate(); err != nil {
		return ExpressQueryResp{}, err
	}

	eqReq.Timestamp, eqReq.Password = sdk.generateTimestampAndPassword(eqReq.BusinessShortCode, eqReq.PassKey)

	data, err := json.Marshal(eqReq)
	if err != nil {
		return ExpressQueryResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, queryEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return ExpressQueryResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return ExpressQueryResp{}, err
	}

	var eqr ExpressQueryResp
	if err := json.Unmarshal(resp, &eqr); err != nil {
		return ExpressQueryResp{}, err
	}

	return eqr, nil
}
