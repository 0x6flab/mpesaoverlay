package pkg

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

	var tr TokenResp
	if err := json.Unmarshal(body, &tr); err != nil {
		return TokenResp{}, err
	}

	return tr, nil
}
