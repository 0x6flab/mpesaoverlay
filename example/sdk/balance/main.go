// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Package main provides an example of how to use AccountBalance method.
package main

import (
	"log"
	"os"

	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
)

var (
	cKey    = os.Getenv("MPESA_CONSUMER_KEY")
	cSecret = os.Getenv("MPESA_CONSUMER_SECRET")
)

func main() {
	conf := mpesa.Config{
		BaseURL:   "https://sandbox.safaricom.co.ke",
		AppKey:    cKey,
		AppSecret: cSecret,
	}

	mp, err := mpesa.NewSDK(conf)
	if err != nil {
		log.Fatal(err)
	}

	balReq := mpesa.AccountBalanceReq{
		InitiatorName:     "testapi",
		InitiatorPassword: "Safaricom999!*!",
		CommandID:         "AccountBalance",
		IdentifierType:    4,
		PartyA:            600772,
		QueueTimeOutURL:   "https://example.com/timeout",
		ResultURL:         "https://example.com/result",
		Remarks:           "test",
	}

	resp, err := mp.AccountBalance(balReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
