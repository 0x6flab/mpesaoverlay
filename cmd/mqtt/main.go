// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Package main provides the entrypoint for the mqtt service.
//
// The mqtt service is responsible for listening to mqtt messages and
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqttadapter "github.com/0x6flab/mpesaoverlay/mqtt"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	zapm "github.com/0x6flab/mpesaoverlay/pkg/mpesa/middleware/logging/zap"
	prometheusm "github.com/0x6flab/mpesaoverlay/pkg/mpesa/middleware/metrics/prometheus"
	"github.com/caarlos0/env/v9"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	svcName      = "mqtt-adapter"
	stopWaitTime = 5 * time.Second
)

type config struct {
	LogLevel       string `env:"MO_LOG_LEVEL"          envDefault:"info"`
	ConsumerKey    string `env:"MPESA_CONSUMER_KEY"`
	ConsumerSecret string `env:"MPESA_CONSUMER_SECRET"`
	BaseURL        string `env:"MPESA_BASE_URL"        envDefault:"https://sandbox.safaricom.co.ke"`
	MQTTURL        string `env:"MO_MQTT_URL"           envDefault:"localhost:1883"`
	MQTTServerCert string `env:"MO_MQTT_SERVER_CERT"`
	MQTTServerKey  string `env:"MO_MQTT_SERVER_KEY"`
	PrometheusURL  string `env:"MO_PROMETHEUS_URL"     envDefault:""`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to load configuration : %s", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to init logger: %s", err)
	}

	server := mqtt.New(&mqtt.Options{InlineClient: true})

	if err = server.AddHook(new(auth.AllowHook), nil); err != nil {
		logger.Error(fmt.Sprintf("failed to add auth hook: %s", err))
	}

	hook, err := newService(cfg, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create %s hook: %s", svcName, err))
	}
	hook.SetServer(server)

	if err = server.AddHook(hook, nil); err != nil {
		logger.Error(fmt.Sprintf("failed to add %s hook: %s", svcName, err))
	}

	g.Go(func() error {
		return startMQTTServer(cfg, server)
	})

	g.Go(func() error {
		return StopSignalHandler(ctx, cancel, logger, svcName, server)
	})

	if err := g.Wait(); err != nil {
		logger.Error(fmt.Sprintf("%s service terminated: %s", svcName, err))
	}
}

func newService(cfg config, logger *zap.Logger) (*mqttadapter.Hook, error) {
	mpesaCfg := mpesa.Config{
		BaseURL:   cfg.BaseURL,
		AppKey:    cfg.ConsumerKey,
		AppSecret: cfg.ConsumerSecret,
	}

	opts := []mpesa.Option{zapm.WithLogger(logger)}
	if cfg.PrometheusURL != "" {
		opts = append(opts, prometheusm.WithMetrics(svcName, cfg.PrometheusURL))
	}

	sdk, err := mpesa.NewSDK(mpesaCfg, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create mpesa sdk: %w", err)
	}

	svc := mqttadapter.NewService(sdk)

	hook := mqttadapter.NewHook(logger, svc)

	return hook, nil
}

func startMQTTServer(cfg config, server *mqtt.Server) error {
	mqtt := listeners.NewTCP(fmt.Sprintf("%s-mqtt", svcName), cfg.MQTTURL, nil)

	if err := server.AddListener(mqtt); err != nil {
		return fmt.Errorf("failed to add mqtt listener: %w", err)
	}

	return server.Serve()
}

func StopSignalHandler(ctx context.Context, cancel context.CancelFunc, logger *zap.Logger, svcName string, server *mqtt.Server) error {
	var err error
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	select {
	case sig := <-c:
		defer cancel()
		c := make(chan bool)
		go func() {
			defer close(c)
			if err = server.Close(); err != nil {
				logger.Error(fmt.Sprintf("failed to close %s server: %s", svcName, err))
			}
		}()
		select {
		case <-c:
		case <-time.After(stopWaitTime):
		}

		logger.Info(fmt.Sprintf("%s gRPC service shutdown by signal: %s", svcName, sig))

		return err
	case <-ctx.Done():
		return nil
	}
}
