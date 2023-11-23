// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mpesa

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	accessToken = "token"
	appKey      = "your-app-key"
	appSecret   = "your-app-secret"
	validToken  = TokenResp{
		AccessToken: accessToken,
		Expiry:      "3599",
	}
)

func TestToken(t *testing.T) {
	testCases := []struct {
		name           string
		statusCode     int
		expectedToken  string
		expectedExpiry string
		expectedErr    error
	}{
		{
			name:           "success",
			statusCode:     http.StatusOK,
			expectedToken:  accessToken,
			expectedExpiry: "3599",
			expectedErr:    nil,
		},
		{
			name:           "failure",
			statusCode:     http.StatusInternalServerError,
			expectedToken:  "",
			expectedExpiry: "",
			expectedErr:    errFailedToGetToken,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("Expected GET request, got %s", r.Method)
				}
				if r.URL.Path != "/"+strings.Split(authEndpoint, "?")[0] {
					t.Errorf("Expected URL path '%s', got %s", strings.Split(authEndpoint, "?")[0], r.URL.Path)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.statusCode)

				if tc.statusCode == http.StatusOK {
					tr := TokenResp{
						AccessToken: tc.expectedToken,
						Expiry:      tc.expectedExpiry,
					}

					if err := json.NewEncoder(w).Encode(tr); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)

						return
					}
				}
			}))
			defer server.Close()

			sdk := mSDK{
				baseURL:   server.URL,
				appKey:    appKey,
				appSecret: appSecret,
				client:    server.Client(),
			}

			tokenResp, err := sdk.Token()
			if errors.Is(err, tc.expectedErr) == false {
				t.Errorf("Expected error '%v', got '%v'", tc.expectedErr, err)
			}
			if tokenResp.AccessToken != tc.expectedToken {
				t.Errorf("Expected token value '%s', got '%s'", tc.expectedToken, tokenResp.AccessToken)
			}
			if tokenResp.Expiry != tc.expectedExpiry {
				t.Errorf("Expected expiry value '%s', got '%s'", tc.expectedExpiry, tokenResp.Expiry)
			}
		})
	}
}
