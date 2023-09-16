package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/0x6flab/mpesaoverlay"
	"github.com/0x6flab/mpesaoverlay/cli"
	mpesa "github.com/0x6flab/mpesaoverlay/pkg"
	"github.com/caarlos0/env/v9"
	"github.com/choria-io/fisk"
)

type Config struct {
	ConsumerKey    string `env:"MPESA_CONSUMER_KEY"`
	ConsumerSecret string `env:"MPESA_CONSUMER_SECRET"`
	BaseURL        string `env:"MPESA_BASE_URL"         envDefault:"https://sandbox.safaricom.co.ke"`
}

var (
	help = `Mpesa Daraja CLI

	See 'mpesa cheat' for a quick tutorial.
	`
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		cancel()
		log.Fatalf(fmt.Sprintf("failed to parse env: %v", err))
	}

	mpesaCfg := mpesa.Config{
		BaseURL:   cfg.BaseURL,
		AppKey:    cfg.ConsumerKey,
		AppSecret: cfg.ConsumerSecret,
		Context:   ctx,
	}
	sdk, err := mpesa.NewSDK(mpesaCfg)
	if err != nil {
		cancel()
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
