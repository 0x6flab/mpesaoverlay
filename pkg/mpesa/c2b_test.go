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

func TestC2BRegisterURL(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          C2BRegisterURLReq
		expectedResponse C2BRegisterURLResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: C2BRegisterURLReq{
				ValidationURL:   "https://example.com/validation",
				ConfirmationURL: "https://example.com/confirmation",
				ShortCode:       600981,
				ResponseType:    "Completed",
			},
			expectedResponse: C2BRegisterURLResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid ResponseType",
			statusCode: http.StatusInternalServerError,
			request: C2BRegisterURLReq{
				ValidationURL:   "https://example.com/validation",
				ConfirmationURL: "https://example.com/confirmation",
				ShortCode:       600981,
				ResponseType:    invalidString,
			},
			expectedResponse: C2BRegisterURLResp{},
			expectedErr:      errInvalidResponseType,
		},
		{
			name:       "failure with invalid ValidationURL",
			statusCode: http.StatusInternalServerError,
			request: C2BRegisterURLReq{
				ValidationURL:   invalidURL,
				ConfirmationURL: "https://example.com/confirmation",
				ShortCode:       600981,
				ResponseType:    "Completed",
			},
			expectedResponse: C2BRegisterURLResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid ConfirmationURL",
			statusCode: http.StatusInternalServerError,
			request: C2BRegisterURLReq{
				ValidationURL:   "https://example.com/validation",
				ConfirmationURL: invalidURL,
				ShortCode:       600981,
				ResponseType:    "Completed",
			},
			expectedResponse: C2BRegisterURLResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid ShortCode",
			statusCode: http.StatusInternalServerError,
			request: C2BRegisterURLReq{
				ValidationURL:   "https://example.com/validation",
				ConfirmationURL: "https://example.com/confirmation",
				ShortCode:       invalidShortCode,
				ResponseType:    "Completed",
			},
			expectedResponse: C2BRegisterURLResp{},
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
					t.Errorf("Expected GET request, got %s", r.Method)
				}
				if r.URL.Path != "/"+c2bRegisterURLEndpoint {
					t.Errorf("Expected URL path '%s', got %s", c2bRegisterURLEndpoint, r.URL.Path)
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

			response, err := sdk.C2BRegisterURL(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}

func TestC2BSimulate(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          C2BSimulateReq
		expectedResponse C2BSimulateResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: C2BSimulateReq{
				CommandID:     "CustomerBuyGoodsOnline",
				Amount:        10,
				Msisdn:        "254712345678",
				BillRefNumber: "",
				ShortCode:     600986,
			},
			expectedResponse: C2BSimulateResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid CommandID",
			statusCode: http.StatusInternalServerError,
			request: C2BSimulateReq{
				CommandID:     invalidString,
				Amount:        10,
				Msisdn:        "254712345678",
				BillRefNumber: "",
				ShortCode:     600986,
			},
			expectedResponse: C2BSimulateResp{},
			expectedErr:      errInvalidCommandID,
		},
		{
			name:       "failure with invalid ShortCode",
			statusCode: http.StatusInternalServerError,
			request: C2BSimulateReq{
				CommandID:     "CustomerBuyGoodsOnline",
				Amount:        10,
				Msisdn:        "254712345678",
				BillRefNumber: "",
				ShortCode:     invalidShortCode,
			},
			expectedResponse: C2BSimulateResp{},
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
				if r.URL.Path != "/"+c2bSimulateEndpoint {
					t.Errorf("Expected URL path '%s', got %s", c2bSimulateEndpoint, r.URL.Path)
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

			response, err := sdk.C2BSimulate(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}
