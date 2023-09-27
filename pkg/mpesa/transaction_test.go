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

func TestTransactionStatus(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          TransactionStatusReq
		expectedResponse TransactionStatusResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: TransactionStatusReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "TransactionStatusQuery",
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
				Occasion:          "test",
				TransactionID:     "RI704KI9RW",
				IdentifierType:    1,
				PartyA:            254759764065,
			},
			expectedResponse: TransactionStatusResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid CommandID",
			statusCode: http.StatusInternalServerError,
			request: TransactionStatusReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         invalidString,
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
				Occasion:          "test",
				TransactionID:     "RI704KI9RW",
				IdentifierType:    1,
				PartyA:            254759764065,
			},
			expectedResponse: TransactionStatusResp{},
			expectedErr:      errInvalidCommandID,
		},
		{
			name:       "failure with invalid QueueTimeOutURL",
			statusCode: http.StatusInternalServerError,
			request: TransactionStatusReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "TransactionStatusQuery",
				QueueTimeOutURL:   invalidURL,
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
				Occasion:          "test",
				TransactionID:     "RI704KI9RW",
				IdentifierType:    1,
				PartyA:            254759764065,
			},
			expectedResponse: TransactionStatusResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid ResultURL",
			statusCode: http.StatusInternalServerError,
			request: TransactionStatusReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "TransactionStatusQuery",
				QueueTimeOutURL:   "https://example.com/result",
				ResultURL:         invalidURL,
				Remarks:           "test",
				Occasion:          "test",
				TransactionID:     "RI704KI9RW",
				IdentifierType:    1,
				PartyA:            254759764065,
			},
			expectedResponse: TransactionStatusResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid Occasion",
			statusCode: http.StatusInternalServerError,
			request: TransactionStatusReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "TransactionStatusQuery",
				QueueTimeOutURL:   "https://example.com/result",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
				Occasion:          invalidString,
				TransactionID:     "RI704KI9RW",
				IdentifierType:    1,
				PartyA:            254759764065,
			},
			expectedResponse: TransactionStatusResp{},
			expectedErr:      errInvalidOccasion,
		},
		{
			name:       "failure with invalid Remarks",
			statusCode: http.StatusInternalServerError,
			request: TransactionStatusReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "TransactionStatusQuery",
				QueueTimeOutURL:   "https://example.com/result",
				ResultURL:         "https://example.com/result",
				Remarks:           invalidString,
				Occasion:          "test",
				TransactionID:     "RI704KI9RW",
				IdentifierType:    1,
				PartyA:            254759764065,
			},
			expectedResponse: TransactionStatusResp{},
			expectedErr:      errInvalidRemarks,
		},
		{
			name:       "failure with invalid IdentifierType",
			statusCode: http.StatusInternalServerError,
			request: TransactionStatusReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "TransactionStatusQuery",
				QueueTimeOutURL:   "https://example.com/result",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
				Occasion:          "test",
				TransactionID:     "RI704KI9RW",
				IdentifierType:    10,
				PartyA:            254759764065,
			},
			expectedResponse: TransactionStatusResp{},
			expectedErr:      errInvalidIdentifierType,
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
					t.Errorf("Expected POST request, got %s", r.Method)
				}
				if r.URL.Path != "/"+transactionEndpoint {
					t.Errorf("Expected URL path '%s', got %s", transactionEndpoint, r.URL.Path)
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

			response, err := sdk.TransactionStatus(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}
