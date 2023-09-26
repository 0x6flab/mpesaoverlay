// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mpesa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var errFailedToGetToken = errors.New("failed to get token")

func (sdk mSDK) Token() (TokenResp, error) {
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

	if resp.StatusCode != http.StatusOK {
		return TokenResp{}, errFailedToGetToken
	}

	var tr TokenResp
	if err := json.Unmarshal(body, &tr); err != nil {
		return TokenResp{}, err
	}

	return tr, nil
}
