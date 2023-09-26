// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Package main provides the entrypoint for the mpesaoverlay cli.
//
// The main package is responsible for parsing the command line arguments and
// passing them to the appropriate function.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/0x6flab/mpesaoverlay"
	"github.com/0x6flab/mpesaoverlay/cli"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/caarlos0/env/v9"
	"github.com/choria-io/fisk"
)

type Config struct {
	ConsumerKey    string `env:"MPESA_CONSUMER_KEY"`
	ConsumerSecret string `env:"MPESA_CONSUMER_SECRET"`
	BaseURL        string `env:"MPESA_BASE_URL"         envDefault:"https://sandbox.safaricom.co.ke"`
}

var help = `Mpesa Daraja CLI

	See 'mpesa cheat' for a quick tutorial.
	`

func main() {
	var cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf(fmt.Sprintf("failed to parse env: %v", err))
	}

	mpesaCfg := mpesa.Config{
		BaseURL:   cfg.BaseURL,
		AppKey:    cfg.ConsumerKey,
		AppSecret: cfg.ConsumerSecret,
	}
	sdk, err := mpesa.NewSDK(mpesaCfg)
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to create mpesa sdk: %v", err))
	}

	mpesaCLI := fisk.New("mpesa", help)
	mpesaCLI.Author("MpesaOverlay <socials@rodneyosodo.com>")
	mpesaCLI.UsageWriter(os.Stdout)
	mpesaCLI.ErrorWriter(os.Stderr)
	mpesaCLI.Version(mpesaoverlay.Version)
	mpesaCLI.WithCheats("cheat")

	mpesaCLI.Flag("consumer-key", "Mpesa Consumer Key").Short('k').Envar("MPESA_CONSUMER_KEY").StringVar(&cfg.ConsumerKey)
	mpesaCLI.Flag("consumer-secret", "Mpesa Consumer Secret").Short('s').Envar("MPESA_CONSUMER_SECRET").StringVar(&cfg.ConsumerSecret)
	mpesaCLI.Flag("base-url", "Mpesa Base URL").Short('b').Envar("MPESA_BASE_URL").StringVar(&cfg.BaseURL)

	cli.AddCommands(mpesaCLI, sdk)

	log.SetFlags(log.Ltime)
	log.SetPrefix("[mpesa] ")

	mpesaCLI.MustParseWithUsage(os.Args[1:])
}
