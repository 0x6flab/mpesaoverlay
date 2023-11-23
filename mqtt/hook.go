// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package mqtt

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/0x6flab/mpesaoverlay"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Hook struct {
	mqtt.HookBase
	serve  *mqtt.Server
	logger *zap.Logger
	Service
}

// NewHook returns a new MQTT hook.
func NewHook(logger *zap.Logger, svc Service) *Hook {
	return &Hook{
		logger:  logger,
		Service: svc,
	}
}

func (h *Hook) SetServer(s *mqtt.Server) {
	h.serve = s
}

func (h *Hook) ID() string {
	return mpesaoverlay.SVCName + "-mqtt"
}

func (h *Hook) Provides(b byte) bool {
	return bytes.Contains( //nolint: gocritic
		[]byte{
			mqtt.OnConnect,
			mqtt.OnDisconnect,
			mqtt.OnPublished,
			mqtt.OnSubscribed,
			mqtt.OnUnsubscribed,
		}, []byte{b})
}

func (h *Hook) Init(_ interface{}) error {
	h.logger.Info("initializing mqtt hook")

	return nil
}

func (h *Hook) OnConnect(cl *mqtt.Client, _ packets.Packet) error {
	h.logger.Info(
		"mqtt client connected",
		zap.String("client_id", cl.ID),
		zap.String("username", string(cl.Properties.Username)),
	)

	return nil
}

func (h *Hook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	fields := []zapcore.Field{
		zap.String("client_id", cl.ID),
		zap.String("username", string(cl.Properties.Username)),
		zap.Bool("expired", expire),
	}
	switch err {
	case nil:
		h.logger.Info(
			"mqtt client disconnected",
			fields...,
		)
	default:
		fields = append(fields, zap.Error(err))
		h.logger.Error(
			"mqtt client disconnected",
			fields...,
		)
	}
}

func (h *Hook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	h.logger.Info(
		"mqtt client published",
		zap.String("client_id", cl.ID),
		zap.String("username", string(cl.Properties.Username)),
		zap.String("topic", pk.TopicName),
	)

	h.handleMessages(pk)
}

func (h *Hook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, _ []byte) {
	h.logger.Info(
		"mqtt client subscribed",
		zap.String("client_id", cl.ID),
		zap.String("username", string(cl.Properties.Username)),
		zap.String("topic", pk.TopicName),
	)
}

func (h *Hook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
	h.logger.Info(
		"mqtt client unsubscribed",
		zap.String("client_id", cl.ID),
		zap.String("username", string(cl.Properties.Username)),
		zap.String("topic", pk.TopicName),
	)
}

// handleMessages handles the inbound MQTT messages.
func (h *Hook) handleMessages(pk packets.Packet) {
	switch pk.TopicName {
	case "mpesa/token":
		h.logger.Info("handling token")
		resp, _ := h.Token(pk)
		h.logger.Info("token", zap.Any("resp", resp))

		h.publish("mpesa/token", resp)

	case "mpesa/express/query":
		h.logger.Info("handling express query")
		resp, err := h.ExpressQuery(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish("mpesa/express/query", resp)

	case "mpesa/express/simulate":
		h.logger.Info("handling express simulate")
		resp, err := h.ExpressSimulate(pk)
		if err != nil {
			h.logger.Error("failed to handle express simulate", zap.Error(err))

			return
		}
		h.publish("mpesa/express/simulate", resp)

	case "mpesa/b2c/payment":
		h.logger.Info("handling b2c payment")
		resp, err := h.B2CPayment(pk)
		if err != nil {
			h.logger.Error("failed to handle b2c payment", zap.Error(err))

			return
		}
		h.publish("mpesa/b2c/payment", resp)

	case "mpesa/account/balance":
		h.logger.Info("handling account balance")
		resp, err := h.AccountBalance(pk)
		if err != nil {
			h.logger.Error("failed to handle account balance", zap.Error(err))

			return
		}
		h.publish("mpesa/account/balance", resp)

	case "mpesa/c2b/register":
		h.logger.Info("handling c2b register")
		resp, err := h.C2BRegisterURL(pk)
		if err != nil {
			h.logger.Error("failed to handle c2b register", zap.Error(err))

			return
		}
		h.publish("mpesa/c2b/register", resp)

	case "mpesa/c2b/simulate":
		h.logger.Info("handling c2b simulate")
		resp, err := h.C2BSimulate(pk)
		if err != nil {
			h.logger.Error("failed to handle c2b simulate", zap.Error(err))

			return
		}
		h.publish("mpesa/c2b/simulate", resp)

	case "mpesa/generate/qr":
		h.logger.Info("handling generate qr")
		resp, err := h.GenerateQR(pk)
		if err != nil {
			h.logger.Error("failed to handle generate qr", zap.Error(err))

			return
		}
		h.publish("mpesa/generate/qr", resp)

	case "mpesa/reverse":
		h.logger.Info("handling reverse")
		resp, err := h.Reverse(pk)
		if err != nil {
			h.logger.Error("failed to handle reverse", zap.Error(err))

			return
		}
		h.publish("mpesa/reverse", resp)

	case "mpesa/transaction/status":
		h.logger.Info("handling transaction status")
		resp, err := h.TransactionStatus(pk)
		if err != nil {
			h.logger.Error("failed to handle transaction status", zap.Error(err))

			return
		}
		h.publish("mpesa/transaction/status", resp)

	case "mpesa/remit/tax":
		h.logger.Info("handling remit tax")
		resp, err := h.RemitTax(pk)
		if err != nil {
			h.logger.Error("failed to handle remit tax", zap.Error(err))

			return
		}
		h.publish("mpesa/remit/tax", resp)

	case "mpesa/b2b/payment":
		h.logger.Info("handling b2b payment")
		resp, err := h.BusinessPayBill(pk)
		if err != nil {
			h.logger.Error("failed to handle b2b payment", zap.Error(err))

			return
		}
		h.publish("mpesa/b2b/payment", resp)

	default:
		switch strings.HasSuffix(pk.TopicName, "/response") {
		case true:
			h.logger.Info("handling response")
		case false:
			h.logger.Info("unknown topic")
		}
	}
}

// publish publishes the response to the MQTT broker.
func (h *Hook) publish(topic string, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		h.logger.Error("failed to marshal payload", zap.Error(err))

		return
	}

	topic += "/response"

	if err = h.serve.Publish(topic, data, false, 0); err != nil {
		h.logger.Error("failed to publish", zap.Error(err))

		return
	}
}
