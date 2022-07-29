package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var c2bEndpoint = "mpesa/c2b/v1"

// C2BRegisterURL Register validation and confirmation URLs on M-Pesa
func (sdk mSDK) C2BRegisterURL(c2bReq C2BRegisterURLReq) (C2BRegisterURLResp, error) {
	data, err := json.Marshal(c2bReq)
	if err != nil {
		return C2BRegisterURLResp{}, err
	}
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, c2bEndpoint, "registerurl")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return C2BRegisterURLResp{}, err
	}
	resp, err := sdk.sendRequest(req)
	if err != nil {
		return C2BRegisterURLResp{}, err
	}

	var c2br C2BRegisterURLResp
	if err := json.Unmarshal(resp, &c2br); err != nil {
		return C2BRegisterURLResp{}, err
	}
	return c2br, nil

}

// C2BSimulate Make payment requests from Client to Business (C2B)
func (sdk mSDK) C2BSimulate(c2bReq C2BSimulateReq) (C2BSimulateResp, error) {
	data, err := json.Marshal(c2bReq)
	if err != nil {
		return C2BSimulateResp{}, err
	}
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, c2bEndpoint, "simulate")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return C2BSimulateResp{}, err
	}
	resp, err := sdk.sendRequest(req)
	if err != nil {
		return C2BSimulateResp{}, err
	}

	var c2bsr C2BSimulateResp
	if err := json.Unmarshal(resp, &c2bsr); err != nil {
		return C2BSimulateResp{}, err
	}
	return c2bsr, nil

}
