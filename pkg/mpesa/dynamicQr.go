package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var qrCodeEndpoint = "mpesa/qrcode/v1/generate"

// GenerateQR Generates a dynamic M-PESA QR Code.
func (sdk mSDK) GenerateQR(qReq QRReq) (QRResp, error) {
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
