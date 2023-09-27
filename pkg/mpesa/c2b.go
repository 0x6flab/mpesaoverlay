// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (sdk mSDK) C2BRegisterURL(c2bReq C2BRegisterURLReq) (C2BRegisterURLResp, error) {
	if err := c2bReq.Validate(); err != nil {
		return C2BRegisterURLResp{}, err
	}

	data, err := json.Marshal(c2bReq)
	if err != nil {
		return C2BRegisterURLResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, c2bRegisterURLEndpoint)

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

func (sdk mSDK) C2BSimulate(c2bReq C2BSimulateReq) (C2BSimulateResp, error) {
	if err := c2bReq.Validate(); err != nil {
		return C2BSimulateResp{}, err
	}

	data, err := json.Marshal(c2bReq)
	if err != nil {
		return C2BSimulateResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, c2bSimulateEndpoint)

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
