package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var expressEndpoint = "mpesa/stkpush/v1"

// ExpressSimulate Initiates online payment on behalf of a customer.
func (sdk mSDK) ExpressSimulate(eReq ExpressSimulateReq) (ExpressSimulateResp, error) {
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

// ExpressQuery Check the status of a Lipa Na M-Pesa Online Payment.
func (sdk mSDK) ExpressQuery(eqReq ExpressQueryReq) (ExpressQueryResp, error) {
	data, err := json.Marshal(eqReq)
	if err != nil {
		return ExpressQueryResp{}, err
	}
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, expressEndpoint, "query")
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
