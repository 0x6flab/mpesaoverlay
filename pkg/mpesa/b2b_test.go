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

func TestBusinessPayBill(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		request          BusinessPayBillReq
		expectedResponse BusinessPayBillResp
		expectedErr      error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   4,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 600986,
				QueueTimeOutURL:        "https://example.com/timeout",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
			},
			expectedResponse: BusinessPayBillResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:       "failure with invalid commanID",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              invalidString,
				SenderIdentifierType:   4,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 600986,
				QueueTimeOutURL:        "https://example.com/timeout",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidCommandID,
		},
		{
			name:       "failure with invalid SenderIdentifierType",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   5,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 600986,
				QueueTimeOutURL:        "https://example.com/timeout",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidIdentifierType,
		},
		{
			name:       "failure with invalid RecieverIdentifierType",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   4,
				RecieverIdentifierType: 5,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 600986,
				QueueTimeOutURL:        "https://example.com/timeout",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidIdentifierType,
		},
		{
			name:       "failure with invalid party A",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   4,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 invalidShortCode,
				PartyB:                 600986,
				QueueTimeOutURL:        "https://example.com/timeout",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidShortCode,
		},
		{
			name:       "failure with invalid party B",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   4,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 invalidShortCode,
				QueueTimeOutURL:        "https://example.com/timeout",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidShortCode,
		},
		{
			name:       "failure with invalid QueueTimeOutURL",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   4,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 600986,
				QueueTimeOutURL:        invalidURL,
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid ResultURL",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   4,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 600986,
				QueueTimeOutURL:        "https://example.com/result",
				ResultURL:              invalidURL,
				Remarks:                "test",
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidURL,
		},
		{
			name:       "failure with invalid Remarks",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   4,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 600986,
				QueueTimeOutURL:        "https://example.com/result",
				ResultURL:              "https://example.com/result",
				Remarks:                invalidString,
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidRemarks,
		},
		{
			name:       "failure with invalid Account Reference",
			statusCode: http.StatusInternalServerError,
			request: BusinessPayBillReq{
				Initiator:              initiatorName,
				InitiatorPassword:      initiatorPassword,
				CommandID:              "BusinessPayBill",
				SenderIdentifierType:   4,
				RecieverIdentifierType: 4,
				Amount:                 10,
				PartyA:                 600986,
				PartyB:                 600986,
				QueueTimeOutURL:        "https://example.com/result",
				ResultURL:              "https://example.com/result",
				Remarks:                "test",
				AccountReference:       invalidString,
			},
			expectedResponse: BusinessPayBillResp{},
			expectedErr:      errInvalidAccountReference,
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
				if r.URL.Path != "/"+b2bEndpoint {
					t.Errorf("Expected URL path '%s', got %s", b2bEndpoint, r.URL.Path)
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

			response, err := sdk.BusinessPayBill(tc.request)
			assert.ErrorIs(t, err, tc.expectedErr, "%s: Expected error '%v', got '%v'", tc.name, tc.expectedErr, err)

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response '%v', got '%v'", tc.expectedResponse, response)
			}
		})
	}
}
