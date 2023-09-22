package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sdk mSDK) GenerateQR(qReq GenerateQRReq) (GenerateQRResp, error) {
	if err := qReq.Validate(); err != nil {
		return GenerateQRResp{}, err
	}

	data, err := json.Marshal(qReq)
	if err != nil {
		return GenerateQRResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, qrCodeEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return GenerateQRResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return GenerateQRResp{}, err
	}

	var qrr GenerateQRResp
	if err := json.Unmarshal(resp, &qrr); err != nil {
		return GenerateQRResp{}, err
	}

	return qrr, nil
}
