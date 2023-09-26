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

func (sdk mSDK) TransactionStatus(tReq TransactionStatusReq) (TransactionStatusResp, error) {
	if err := tReq.Validate(); err != nil {
		return TransactionStatusResp{}, err
	}

	var err error
	tReq.SecurityCredential, err = sdk.generateSecurityCredential(tReq.InitiatorPassword)
	if err != nil {
		return TransactionStatusResp{}, err
	}

	data, err := json.Marshal(tReq)
	if err != nil {
		return TransactionStatusResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, transactionEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return TransactionStatusResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return TransactionStatusResp{}, err
	}

	var tr TransactionStatusResp
	if err := json.Unmarshal(resp, &tr); err != nil {
		return TransactionStatusResp{}, err
	}

	return tr, nil
}
