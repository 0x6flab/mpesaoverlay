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

func TestAccountBalance(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          AccountBalanceReq
		expectedResponse AccountBalanceResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: AccountBalanceReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "AccountBalance",
				PartyA:            600986,
				IdentifierType:    4,
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
			},
			expectedResponse: AccountBalanceResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid commanID",
			statusCode: http.StatusInternalServerError,
			request: AccountBalanceReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         invalidString,
				PartyA:            600986,
				IdentifierType:    4,
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
			},
			expectedResponse: AccountBalanceResp{},
			expectedErr:      errInvalidCommandID,
		},
		{
			name:       "failure with invalid PartyA",
			statusCode: http.StatusInternalServerError,
			request: AccountBalanceReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "AccountBalance",
				PartyA:            invalidShortCode,
				IdentifierType:    4,
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
			},
			expectedResponse: AccountBalanceResp{},
			expectedErr:      errInvalidShortCode,
		},
		{
			name:       "failure with invalid identifier type",
			statusCode: http.StatusInternalServerError,
			request: AccountBalanceReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "AccountBalance",
				PartyA:            600986,
				IdentifierType:    10,
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
			},
			expectedResponse: AccountBalanceResp{},
			expectedErr:      errInvalidIdentifierType,
		},
		{
			name:       "failure with invalid QueueTimeOutURL",
			statusCode: http.StatusInternalServerError,
			request: AccountBalanceReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "AccountBalance",
				PartyA:            600986,
				IdentifierType:    4,
				QueueTimeOutURL:   invalidURL,
				ResultURL:         "https://example.com/result",
				Remarks:           "test",
			},
			expectedResponse: AccountBalanceResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid ResultURL",
			statusCode: http.StatusInternalServerError,
			request: AccountBalanceReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "AccountBalance",
				PartyA:            600986,
				IdentifierType:    4,
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         invalidURL,
				Remarks:           "test",
			},
			expectedResponse: AccountBalanceResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid Remarks",
			statusCode: http.StatusInternalServerError,
			request: AccountBalanceReq{
				InitiatorName:     initiatorName,
				InitiatorPassword: initiatorPassword,
				CommandID:         "AccountBalance",
				PartyA:            600986,
				IdentifierType:    4,
				QueueTimeOutURL:   "https://example.com/timeout",
				ResultURL:         "https://example.com/result",
				Remarks:           invalidString,
			},
			expectedResponse: AccountBalanceResp{},
			expectedErr:      errInvalidRemarks,
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
				if r.URL.Path != "/"+accbalanceEndpoint {
					t.Errorf("Expected URL path '%s', got %s", accbalanceEndpoint, r.URL.Path)
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

			response, err := sdk.AccountBalance(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}
