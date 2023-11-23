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

var passKey = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"

func TestExpressSimulate(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          ExpressSimulateReq
		expectedResponse ExpressSimulateResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       254712345678,
				Amount:            1,
				PartyA:            254712345678,
				PartyB:            174379,
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   "Payment of X",
			},
			expectedResponse: ExpressSimulateResp{
				ResponseDescription: "Success. Request accepted for processing",
				ResponseCode:        "0",
				MerchantRequestID:   "27260-79456854-2",
				CheckoutRequestID:   "ws_CO_07092023004130971712345678",
				CustomerMessage:     "Success. Request accepted for processing",
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid TransactionType",
			statusCode: http.StatusInternalServerError,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				TransactionType:   invalidString,
				PhoneNumber:       254712345678,
				Amount:            1,
				PartyA:            254712345678,
				PartyB:            174379,
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   "Payment of X",
			},
			expectedResponse: ExpressSimulateResp{},
			expectedErr:      errInvalidTransactionType,
		},
		{
			name:       "failure with invalid BusinessShortCode",
			statusCode: http.StatusInternalServerError,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: invalidShortCode,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       254712345678,
				Amount:            1,
				PartyA:            254712345678,
				PartyB:            174379,
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   "Payment of X",
			},
			expectedResponse: ExpressSimulateResp{},
			expectedErr:      errInvalidShortCode,
		},
		{
			name:       "failure with invalid PhoneNumber",
			statusCode: http.StatusInternalServerError,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       invalidPhoneNumber,
				Amount:            1,
				PartyA:            254712345678,
				PartyB:            174379,
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   "Payment of X",
			},
			expectedResponse: ExpressSimulateResp{},
			expectedErr:      errInvalidPhoneNumber,
		},
		{
			name:       "failure with invalid PartyA",
			statusCode: http.StatusInternalServerError,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       254712345678,
				Amount:            1,
				PartyA:            invalidPhoneNumber,
				PartyB:            174379,
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   "Payment of X",
			},
			expectedResponse: ExpressSimulateResp{},
			expectedErr:      errInvalidPhoneNumber,
		},
		{
			name:       "failure with invalid PartyB",
			statusCode: http.StatusInternalServerError,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       254712345678,
				Amount:            1,
				PartyA:            254712345678,
				PartyB:            invalidShortCode,
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   "Payment of X",
			},
			expectedResponse: ExpressSimulateResp{},
			expectedErr:      errInvalidShortCode,
		},
		{
			name:       "failure with invalid CallBackURL",
			statusCode: http.StatusInternalServerError,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       254712345678,
				Amount:            1,
				PartyA:            254712345678,
				PartyB:            174379,
				CallBackURL:       invalidURL,
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   "Payment of X",
			},
			expectedResponse: ExpressSimulateResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid AccountReference",
			statusCode: http.StatusInternalServerError,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       254712345678,
				Amount:            1,
				PartyA:            254712345678,
				PartyB:            174379,
				CallBackURL:       "https://example.com/callback",
				AccountReference:  invalidString,
				TransactionDesc:   "Payment of X",
			},
			expectedResponse: ExpressSimulateResp{},
			expectedErr:      errInvalidAccountReference,
		},
		{
			name:       "failure with invalid TransactionDesc",
			statusCode: http.StatusInternalServerError,
			request: ExpressSimulateReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				TransactionType:   "CustomerPayBillOnline",
				PhoneNumber:       254712345678,
				Amount:            1,
				PartyA:            254712345678,
				PartyB:            174379,
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "CompanyXLTD",
				TransactionDesc:   invalidString,
			},
			expectedResponse: ExpressSimulateResp{},
			expectedErr:      errInvalidTransactionDesc,
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
				if r.URL.Path != "/"+expressSimulateEndpoint {
					t.Errorf("Expected URL path '%s', got %s", expressSimulateEndpoint, r.URL.Path)
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

			response, err := sdk.ExpressSimulate(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}

func TestExpressQuery(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          ExpressQueryReq
		expectedResponse ExpressQueryResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: ExpressQueryReq{
				PassKey:           passKey,
				BusinessShortCode: 174379,
				CheckoutRequestID: "ws_CO_07092023195244460720136609",
			},
			expectedResponse: ExpressQueryResp{
				ResponseDescription: "The service request has been accepted successsfully",
				ResponseCode:        "0",
				MerchantRequestID:   "92643-47073138-2",
				CheckoutRequestID:   "ws_CO_07092023195244460712345678",
				CustomerMessage:     "",
				ResultCode:          "1032",
				ResultDesc:          "Request cancelled by user",
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid BusinessShortCode",
			statusCode: http.StatusInternalServerError,
			request: ExpressQueryReq{
				PassKey:           passKey,
				BusinessShortCode: invalidShortCode,
				CheckoutRequestID: "ws_CO_07092023195244460720136609",
			},
			expectedResponse: ExpressQueryResp{},
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
						http.Error(w, err.Error(), http.StatusInternalServerError)

						return
					}

					return
				}
				if r.Method != http.MethodPost {
					t.Errorf("Expected POST request, got %s", r.Method)
				}
				if r.URL.Path != "/"+queryEndpoint {
					t.Errorf("Expected URL path '%s', got %s", queryEndpoint, r.URL.Path)
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

			response, err := sdk.ExpressQuery(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}
