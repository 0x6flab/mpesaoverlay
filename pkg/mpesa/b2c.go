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

	"github.com/oklog/ulid/v2"
)

func (sdk mSDK) B2CPayment(b2cReq B2CPaymentReq) (B2CPaymentResp, error) {
	if err := b2cReq.Validate(); err != nil {
		return B2CPaymentResp{}, err
	}

	var err error
	b2cReq.SecurityCredential, err = sdk.generateSecurityCredential(b2cReq.InitiatorPassword)
	if err != nil {
		return B2CPaymentResp{}, err
	}

	if b2cReq.OriginatorConversationID == "" {
		b2cReq.OriginatorConversationID = ulid.Make().String()
	}

	data, err := json.Marshal(b2cReq)
	if err != nil {
		return B2CPaymentResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, b2cEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return B2CPaymentResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return B2CPaymentResp{}, err
	}

	var b2cr B2CPaymentResp
	if err := json.Unmarshal(resp, &b2cr); err != nil {
		return B2CPaymentResp{}, err
	}

	return b2cr, nil
}
