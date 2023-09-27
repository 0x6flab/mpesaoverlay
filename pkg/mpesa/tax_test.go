// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mpesa

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemitTax(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          RemitTaxReq
		expectedResponse RemitTaxResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: RemitTaxReq{
				InitiatorName:          initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "PayTaxToKRA",
				Amount:                 10,
				QueueTimeOutURL:        "https://example.com/timeout",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
				RecieverIdentifierType: 11,
				SenderIdentifierType:   4,
				PartyA:                 600978,
				PartyB:                 572572,
				AccountReference:       "353353",
			},
			expectedResponse: RemitTaxResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid CommandID",
			statusCode: http.StatusInternalServerError,
			request: RemitTaxReq{
				InitiatorName:          initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              invalidString,
				Amount:                 10,
				QueueTimeOutURL:        "https://example.com/timeout",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
				RecieverIdentifierType: 11,
				SenderIdentifierType:   4,
				PartyA:                 600978,
				PartyB:                 572572,
				AccountReference:       "353353",
			},
			expectedResponse: RemitTaxResp{},
			expectedErr:      errInvalidCommandID,
		},
		{
			name:       "failure with invalid QueueTimeOutURL",
			statusCode: http.StatusInternalServerError,
			request: RemitTaxReq{
				InitiatorName:          initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "PayTaxToKRA",
				Amount:                 10,
				QueueTimeOutURL:        invalidURL,
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
				RecieverIdentifierType: 11,
				SenderIdentifierType:   4,
				PartyA:                 600978,
				PartyB:                 572572,
				AccountReference:       "353353",
			},
			expectedResponse: RemitTaxResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid ResultURL",
			statusCode: http.StatusInternalServerError,
			request: RemitTaxReq{
				InitiatorName:          initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "PayTaxToKRA",
				Amount:                 10,
				QueueTimeOutURL:        "https://example.com/result",
				ResultURL:              invalidURL,
				Remarks:                "test",
				RecieverIdentifierType: 11,
				SenderIdentifierType:   4,
				PartyA:                 600978,
				PartyB:                 572572,
				AccountReference:       "353353",
			},
			expectedResponse: RemitTaxResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid Remarks",
			statusCode: http.StatusInternalServerError,
			request: RemitTaxReq{
				InitiatorName:          initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "PayTaxToKRA",
				Amount:                 10,
				QueueTimeOutURL:        "https://example.com/result",
				ResultURL:              "https://example.com/result",
				Remarks:                invalidString,
				RecieverIdentifierType: 11,
				SenderIdentifierType:   4,
				PartyA:                 600978,
				PartyB:                 572572,
				AccountReference:       "353353",
			},
			expectedResponse: RemitTaxResp{},
			expectedErr:      errInvalidRemarks,
		},
		{
			name:       "failure with invalid AccountReference",
			statusCode: http.StatusInternalServerError,
			request: RemitTaxReq{
				InitiatorName:          initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "PayTaxToKRA",
				Amount:                 10,
				QueueTimeOutURL:        "https://example.com/result",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
				RecieverIdentifierType: 11,
				SenderIdentifierType:   4,
				PartyA:                 600978,
				PartyB:                 572572,
				AccountReference:       invalidString,
			},
			expectedResponse: RemitTaxResp{},
			expectedErr:      errInvalidAccountReference,
		},
		{
			name:       "failure with invalid PartyA",
			statusCode: http.StatusInternalServerError,
			request: RemitTaxReq{
				InitiatorName:          initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "PayTaxToKRA",
				Amount:                 10,
				QueueTimeOutURL:        "https://example.com/result",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
				RecieverIdentifierType: 11,
				SenderIdentifierType:   4,
				PartyA:                 invalidShortCode,
				PartyB:                 572572,
				AccountReference:       "353353",
			},
			expectedResponse: RemitTaxResp{},
			expectedErr:      errInvalidShortCode,
		},
		{
			name:       "failure with invalid PartyB",
			statusCode: http.StatusInternalServerError,
			request: RemitTaxReq{
				InitiatorName:          initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "PayTaxToKRA",
				Amount:                 10,
				QueueTimeOutURL:        "https://example.com/result",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
				RecieverIdentifierType: 11,
				SenderIdentifierType:   4,
				PartyA:                 600978,
				PartyB:                 invalidShortCode,
				AccountReference:       "353353",
			},
			expectedResponse: RemitTaxResp{},
			expectedErr:      errInvalidShortCode,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet && r.URL.Path == "/"+strings.Split(authEndpoint, "?")[0] {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)

					if err := json.NewEncoder(w).Encode(validToken); err != nil {
						t.Errorf("Expected no error, got %v", err)
					}

					return
				}
				if r.Method != http.MethodPost {
					t.Errorf("Expected GET request, got %s", r.Method)
				}
				if r.URL.Path != "/"+taxEndpoint {
					t.Errorf("Expected URL path '%s', got %s", taxEndpoint, r.URL.Path)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.statusCode)

				if tc.statusCode == http.StatusOK {
					if err := json.NewEncoder(w).Encode(tc.expectedResponse); err != nil {
						t.Errorf("Expected no error, got %v", err)
					}
				}
			}))
			defer server.Close()

			sdk := mSDK{
				baseURL:   server.URL,
				appKey:    appKey,
				appSecret: appSecret,
				certFile:  sandboxCertificate,
				client:    server.Client(),
			}

			response, err := sdk.RemitTax(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}
