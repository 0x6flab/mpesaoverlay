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

func (sdk mSDK) BusinessPayBill(bpbReq BusinessPayBillReq) (BusinessPayBillResp, error) {
	if err := bpbReq.Validate(); err != nil {
		return BusinessPayBillResp{}, err
	}

	var err error
	bpbReq.SecurityCredential, err = sdk.generateSecurityCredential(bpbReq.InitiatorPassword)
	if err != nil {
		return BusinessPayBillResp{}, err
	}

	data, err := json.Marshal(bpbReq)
	if err != nil {
		return BusinessPayBillResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, b2bEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return BusinessPayBillResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return BusinessPayBillResp{}, err
	}

	var b2cr BusinessPayBillResp
	if err := json.Unmarshal(resp, &b2cr); err != nil {
		return BusinessPayBillResp{}, err
	}

	return b2cr, nil
}
