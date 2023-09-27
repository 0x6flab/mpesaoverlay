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

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

var (
	validResp = ValidResp{
		OriginatorConversationID: "AG_20230907_2010325b025970fde878",
		ConversationID:           "AG_20230907_2010325b025970fde878",
		ResponseDescription:      "Accept the service request successfully.",
		ResponseCode:             "0",
	}
	invalidURL                = "ws://invalid"
	invalidPhoneNumber uint64 = 1
	invalidShortCode   uint64 = 1
	invalidString             = strings.Repeat("a", 256)
	initiatorName             = "testapi"
	initiatorPassword         = "Safaricom999!*!"
)

func TestB2CPayment(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          B2CPaymentReq
		expectedResponse B2CPaymentResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: B2CPaymentReq{
				OriginatorConversationID: ulid.Make().String(),
				InitiatorName:            initiatorName,
				InitiatorPassword:        initiatorPassword,
				CommandID:                "BusinessPayment",
				Amount:                   10,
				PartyA:                   600986,
				PartyB:                   254712345678,
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "test",
				Occasion:                 "test",
			},
			expectedResponse: B2CPaymentResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid commanID",
			statusCode: http.StatusInternalServerError,
			request: B2CPaymentReq{
				OriginatorConversationID: ulid.Make().String(),
				InitiatorName:            initiatorName,
				InitiatorPassword:        initiatorPassword,
				CommandID:                invalidString,
				Amount:                   10,
				PartyA:                   invalidShortCode,
				PartyB:                   254712345678,
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "test",
				Occasion:                 "test",
			},
			expectedResponse: B2CPaymentResp{},
			expectedErr:      errInvalidCommandID,
		},
		{
			name:       "failure with invalid party A",
			statusCode: http.StatusInternalServerError,
			request: B2CPaymentReq{
				OriginatorConversationID: ulid.Make().String(),
				InitiatorName:            initiatorName,
				InitiatorPassword:        initiatorPassword,
				CommandID:                "BusinessPayment",
				Amount:                   10,
				PartyA:                   invalidShortCode,
				PartyB:                   254712345678,
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "test",
				Occasion:                 "test",
			},
			expectedResponse: B2CPaymentResp{},
			expectedErr:      errInvalidShortCode,
		},
		{
			name:       "failure with invalid party B",
			statusCode: http.StatusInternalServerError,
			request: B2CPaymentReq{
				OriginatorConversationID: ulid.Make().String(),
				InitiatorName:            initiatorName,
				InitiatorPassword:        initiatorPassword,
				CommandID:                "BusinessPayment",
				Amount:                   10,
				PartyA:                   600986,
				PartyB:                   invalidPhoneNumber,
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "test",
				Occasion:                 "test",
			},
			expectedResponse: B2CPaymentResp{},
			expectedErr:      errInvalidPhoneNumber,
		},
		{
			name:       "failure with invalid QueueTimeOutURL",
			statusCode: http.StatusInternalServerError,
			request: B2CPaymentReq{
				OriginatorConversationID: ulid.Make().String(),
				InitiatorName:            initiatorName,
				InitiatorPassword:        initiatorPassword,
				CommandID:                "BusinessPayment",
				Amount:                   10,
				PartyA:                   600986,
				PartyB:                   254712345678,
				QueueTimeOutURL:          invalidURL,
				ResultURL:                "https://example.com/result",
				Remarks:                  "test",
				Occasion:                 "test",
			},
			expectedResponse: B2CPaymentResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid ResultURL",
			statusCode: http.StatusInternalServerError,
			request: B2CPaymentReq{
				OriginatorConversationID: ulid.Make().String(),
				InitiatorName:            initiatorName,
				InitiatorPassword:        initiatorPassword,
				CommandID:                "BusinessPayment",
				Amount:                   10,
				PartyA:                   600986,
				PartyB:                   254712345678,
				QueueTimeOutURL:          "https://example.com/result",
				ResultURL:                invalidURL,
				Remarks:                  "test",
				Occasion:                 "test",
			},
			expectedResponse: B2CPaymentResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid Occasion",
			statusCode: http.StatusInternalServerError,
			request: B2CPaymentReq{
				OriginatorConversationID: ulid.Make().String(),
				InitiatorName:            initiatorName,
				InitiatorPassword:        initiatorPassword,
				CommandID:                "BusinessPayment",
				Amount:                   10,
				PartyA:                   600986,
				PartyB:                   254712345678,
				QueueTimeOutURL:          "https://example.com/result",
				ResultURL:                "https://example.com/result",
				Remarks:                  "test",
				Occasion:                 invalidString,
			},
			expectedResponse: B2CPaymentResp{},
			expectedErr:      errInvalidOccasion,
		},
		{
			name:       "failure with invalid Remarks",
			statusCode: http.StatusInternalServerError,
			request: B2CPaymentReq{
				OriginatorConversationID: ulid.Make().String(),
				InitiatorName:            initiatorName,
				InitiatorPassword:        initiatorPassword,
				CommandID:                "BusinessPayment",
				Amount:                   10,
				PartyA:                   600986,
				PartyB:                   254712345678,
				QueueTimeOutURL:          "https://example.com/result",
				ResultURL:                "https://example.com/result",
				Remarks:                  invalidString,
				Occasion:                 "test",
			},
			expectedResponse: B2CPaymentResp{},
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
				if r.URL.Path != "/"+b2cEndpoint {
					t.Errorf("Expected URL path '%s', got %s", b2cEndpoint, r.URL.Path)
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

			response, err := sdk.B2CPayment(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}
