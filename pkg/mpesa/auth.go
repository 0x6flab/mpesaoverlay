package mpesa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (sdk mSDK) GetToken() (TokenResp, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, authEndpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return TokenResp{}, err
	}

	if sdk.context != nil {
		req = req.WithContext(sdk.context)
	}

	req.SetBasicAuth(sdk.appKey, sdk.appSecret)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	resp, err := sdk.client.Do(req)
	if err != nil {
		return TokenResp{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TokenResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return TokenResp{}, fmt.Errorf("failed to get token: %s", string(body))
	}

	var tr TokenResp
	if err := json.Unmarshal(body, &tr); err != nil {
		return TokenResp{}, err
	}

	return tr, nil
}
