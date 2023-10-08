// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Package main provides an example of how to use B2C payment method.
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

	b2cReq := mpesa.BusinessPayBillReq{
		Initiator:              "testapi",
		InitiatorPassword:      "Safaricom999!*!",
		CommandID:              "BusinessPayBill",
		SenderIdentifierType:   4,
		RecieverIdentifierType: 4,
		Amount:                 10,
		PartyA:                 600992,
		PartyB:                 600992,
		AccountReference:       "353353",
		Requester:              254700000000,
		QueueTimeOutURL:        "https://example.com/timeout",
		ResultURL:              "https://example.com/result",
		Remarks:                "test",
	}

	resp, err := mp.BusinessPayBill(b2cReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
