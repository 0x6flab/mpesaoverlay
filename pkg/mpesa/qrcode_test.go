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

func TestGenerateQR(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          GenerateQRReq
		expectedResponse GenerateQRResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: GenerateQRReq{
				MerchantName: "Test Supermarket",
				RefNo:        "Invoice No",
				Amount:       2000,
				TrxCode:      "BG",
				CPI:          "174379",
				Size:         "300",
			},
			expectedResponse: GenerateQRResp{
				ResponseDescription: "The service request is processed successfully.",
				ResponseCode:        "00",
				RequestID:           "QRCode:...",
				QRCode:              "qr_code",
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid TrxCode",
			statusCode: http.StatusInternalServerError,
			request: GenerateQRReq{
				MerchantName: "Test Supermarket",
				RefNo:        "Invoice No",
				Amount:       2000,
				TrxCode:      invalidString,
				CPI:          "174379",
				Size:         "300",
			},
			expectedResponse: GenerateQRResp{},
			expectedErr:      errInvalidTransactionType,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet && r.URL.Path == "/"+strings.Split(authEndpoint, "?")[0] {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)

					if err := json.NewEncoder(w).Encode(validToken); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)

						return
					}

					return
				}
				if r.Method != http.MethodPost {
					t.Errorf("Expected POST request, got %s", r.Method)
				}
				if r.URL.Path != "/"+qrCodeEndpoint {
					t.Errorf("Expected URL path '%s', got %s", qrCodeEndpoint, r.URL.Path)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.statusCode)

				if tc.statusCode == http.StatusOK {
					if err := json.NewEncoder(w).Encode(tc.expectedResponse); err != nil {
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
				certFile:  sandboxCertificate,
				client:    server.Client(),
			}

			response, err := sdk.GenerateQR(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}
