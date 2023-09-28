// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Package main provides an example of how to use ExpressSimulate method.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	respTopic = "mpesa/express/simulate/response"
	topic     = "mpesa/express/simulate"
	closeChan = make(chan struct{})
)

func onMessageReceived(_ mqtt.Client, message mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", string(message.Payload()), message.Topic())
	close(closeChan)
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(fmt.Errorf("failed to connect to MQTT broker: %w", token.Error()))
	}

	if token := client.Subscribe(respTopic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		log.Fatal(fmt.Errorf("error subscribing to topic: %w", token.Error()))
	}

	qrReq := mpesa.ExpressSimulateReq{
		PassKey:           "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919", // Get this from the developer portal under the test credentials section
		BusinessShortCode: 174379,
		TransactionType:   "CustomerPayBillOnline",
		PhoneNumber:       254712345678, // You can use your own phone number here
		Amount:            1,
		PartyA:            254712345678,
		PartyB:            174379,
		CallBackURL:       "https://69a2-105-163-2-116.ngrok.io",
		AccountReference:  "CompanyXLTD",
		TransactionDesc:   "Payment of X",
	}
	message, err := json.Marshal(qrReq)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to marshal b2c request: %w", err))
	}

	token := client.Publish(topic, 0, false, message)
	token.Wait()

	select {
	case <-closeChan:
		log.Println("received message")
	case <-time.After(5 * time.Second):
		log.Println("timed out")
	}

	client.Unsubscribe(topic)
	client.Disconnect(250)
}
