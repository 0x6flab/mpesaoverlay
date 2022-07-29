package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var reversalEndpoint = "mpesa/reversal/v1/request"

// Reverse Reverses an M-Pesa transaction.
func (sdk mSDK) Reverse(rReq ReversalReq) (ReversalResp, error) {
	if err := rReq.Validate(); err != nil {
		return ReversalResp{}, err
	}
	data, err := json.Marshal(rReq)
	if err != nil {
		return ReversalResp{}, err
	}
	url := fmt.Sprintf("%s/%s", sdk.baseURL, reversalEndpoint)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return ReversalResp{}, err
	}
	resp, err := sdk.sendRequest(req)
	if err != nil {
		return ReversalResp{}, err
	}

	var rr ReversalResp
	if err := json.Unmarshal(resp, &rr); err != nil {
		return ReversalResp{}, err
	}
	return rr, nil
}
