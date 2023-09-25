package mqtt

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/0x6flab/mpesaoverlay"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"go.uber.org/zap"
)

type Hook struct {
	mqtt.HookBase
	serve  *mqtt.Server
	logger *zap.Logger
	Service
}

func NewHook(logger *zap.Logger, svc Service) *Hook {
	return &Hook{
		logger:  logger,
		Service: svc,
	}
}

func (h *Hook) SetServer(s *mqtt.Server) {
	h.serve = s
}

// ID returns the ID of the hook.
func (h *Hook) ID() string {
	return mpesaoverlay.SVCName + "-mqtt"
}

func (h *Hook) Provides(b byte) bool {
	return bytes.Contains(
		[]byte{b},
		[]byte{
			mqtt.OnConnect,
			mqtt.OnDisconnect,
			mqtt.OnPublished,
			mqtt.OnSubscribed,
			mqtt.OnUnsubscribed,
		},
	)
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
	h.logger.Info(
		"mqtt client disconnected",
		zap.String("client_id", cl.ID),
		zap.String("username", string(cl.Properties.Username)),
		zap.Bool("expired", expire),
		zap.Error(err),
	)
}

func (h *Hook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	h.logger.Info(
		"mqtt client published",
		zap.String("client_id", cl.ID),
		zap.String("username", string(cl.Properties.Username)),
		zap.String("topic", pk.TopicName),
	)

	h.handleMessages(cl, pk)
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

func (h *Hook) handleMessages(cl *mqtt.Client, pk packets.Packet) {
	// if pk.Payload == nil {
	// 	h.logger.Info("empty payload")
	// 	return
	// }
	switch pk.TopicName {
	case "mpesa/token":
		h.logger.Info("handling token")
		resp, _ := h.GetToken(pk)
		h.logger.Info("token", zap.Any("resp", resp))
		// if err != nil {
		// 	h.logger.Error("failed to handle token", zap.Error(err))
		// 	return
		// }
		h.publish(cl, "mpesa/token", resp)

	case "mpesa/express/query":
		h.logger.Info("handling express query")
		resp, err := h.ExpressQuery(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/express/query", resp)

	case "mpesa/express/simulate":
		h.logger.Info("handling express simulate")
		resp, err := h.ExpressSimulate(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/express/simulate", resp)

	case "mpesa/b2c/payment":
		h.logger.Info("handling b2c payment")
		resp, err := h.B2CPayment(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/b2c/payment", resp)

	case "mpesa/account/balance":
		h.logger.Info("handling account balance")
		resp, err := h.AccountBalance(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/account/balance", resp)

	case "mpesa/c2b/register":
		h.logger.Info("handling c2b register")
		resp, err := h.C2BRegisterURL(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/c2b/register", resp)

	case "mpesa/c2b/simulate":
		h.logger.Info("handling c2b simulate")
		resp, err := h.C2BSimulate(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/c2b/simulate", resp)

	case "mpesa/generate/qr":
		h.logger.Info("handling generate qr")
		resp, err := h.GenerateQR(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/generate/qr", resp)

	case "mpesa/reverse":
		h.logger.Info("handling reverse")
		resp, err := h.Reverse(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/reverse", resp)

	case "mpesa/transaction/status":
		h.logger.Info("handling transaction status")
		resp, err := h.TransactionStatus(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/transaction/status", resp)

	case "mpesa/remit/tax":
		h.logger.Info("handling remit tax")
		resp, err := h.RemitTax(pk)
		if err != nil {
			h.logger.Error("failed to handle express query", zap.Error(err))

			return
		}
		h.publish(cl, "mpesa/remit/tax", resp)

	default:
		switch strings.HasSuffix(pk.TopicName, "/response") {
		case true:
			h.logger.Info("handling response")
		case false:
			h.logger.Info("unknown topic")
		}
	}
}

func (h *Hook) publish(_ *mqtt.Client, topic string, payload interface{}) {
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
