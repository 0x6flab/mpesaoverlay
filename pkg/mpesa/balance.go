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

func (sdk mSDK) AccountBalance(abReq AccountBalanceReq) (AccountBalanceResp, error) {
	if err := abReq.Validate(); err != nil {
		return AccountBalanceResp{}, err
	}

	var err error
	abReq.SecurityCredential, err = sdk.generateSecurityCredential(abReq.InitiatorPassword)
	if err != nil {
		return AccountBalanceResp{}, err
	}

	data, err := json.Marshal(abReq)
	if err != nil {
		return AccountBalanceResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, accbalanceEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return AccountBalanceResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return AccountBalanceResp{}, err
	}

	var abr AccountBalanceResp
	if err := json.Unmarshal(resp, &abr); err != nil {
		return AccountBalanceResp{}, err
	}

	return abr, nil
}
