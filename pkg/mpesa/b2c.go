package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var b2cEndpoint = "mpesa/b2c/v1/paymentrequest"

// B2CPayment Transact between an M-Pesa short code to a phone number registered on M-Pesa
func (sdk mSDK) B2CPayment(b2cReq B2Creq) (B2CResp, error) {
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
