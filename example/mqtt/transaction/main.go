// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Package main provides an example of how to use TransactionStatus method.
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
	respTopic = "mpesa/transaction/status/response"
	topic     = "mpesa/transaction/status"
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

	trxReq := mpesa.TransactionStatusReq{
		InitiatorName:     "testapi",
		InitiatorPassword: "Safaricom999!*!",
		CommandID:         "TransactionStatusQuery",
		IdentifierType:    1,
		TransactionID:     "RI704KI9RW",
		PartyA:            254759764065,
		QueueTimeOutURL:   "https://example.com/timeout",
		ResultURL:         "https://example.com/result",
		Remarks:           "test",
		Occasion:          "test",
	}
	message, err := json.Marshal(trxReq)
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
