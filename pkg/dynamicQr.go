package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sdk mSDK) GenerateQR(qReq QRReq) (QRResp, error) {
	if err := qReq.validate(); err != nil {
		return QRResp{}, err
	}

	data, err := json.Marshal(qReq)
	if err != nil {
		return QRResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, qrCodeEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return QRResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return QRResp{}, err
	}

	var qrr QRResp
	if err := json.Unmarshal(resp, &qrr); err != nil {
		return QRResp{}, err
	}

	return qrr, nil
}
