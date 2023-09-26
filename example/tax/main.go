// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Package main provides an example of how to use RemitTax method.
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

	taxReq := mpesa.RemitTaxReq{
		InitiatorName:          "testapi",
		InitiatorPassword:      "Safaricom999!*!",
		CommandID:              "PayTaxToKRA",
		SenderIdentifierType:   4,
		RecieverIdentifierType: 4,
		Amount:                 239,
		PartyA:                 600978,
		PartyB:                 572572,
		AccountReference:       "353353",
		QueueTimeOutURL:        "https://example.com/timeout",
		ResultURL:              "https://example.com/result",
		Remarks:                "test",
	}

	resp, err := mp.RemitTax(taxReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Resp: %+v\n", resp)
}
