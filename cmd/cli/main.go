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
	CONSUMER_KEY    string `env:"MPESA_CONSUMER_KEY"`
	CONSUMER_SECRET string `env:"MPESA_CONSUMER_SECRET"`
	BASE_URL        string `env:"MPESA_BASE_URL" envDefault:"https://sandbox.safaricom.co.ke"`
	MAX_IDLE_CONNS  int    `env:"MPESA_MAX_IDLE_CONNS" envDefault:"10"`
}

var (
	help = `Mpesa Daraja CLI

	See 'mpesa cheat' for a quick tutorial.
	`
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf(fmt.Sprintf("failed to parse env: %v", err))
	}

	mpesaCfg := mpesa.Config{
		CTX:          ctx,
		BaseURL:      cfg.BASE_URL,
		AppKey:       cfg.CONSUMER_KEY,
		AppSecret:    cfg.CONSUMER_SECRET,
		MaxIdleConns: cfg.MAX_IDLE_CONNS,
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

	mpesaCLI.Flag("consumer-key", "Mpesa Consumer Key").Short('k').Envar("MPESA_CONSUMER_KEY").StringVar(&cfg.CONSUMER_KEY)
	mpesaCLI.Flag("consumer-secret", "Mpesa Consumer Secret").Short('s').Envar("MPESA_CONSUMER_SECRET").StringVar(&cfg.CONSUMER_SECRET)
	mpesaCLI.Flag("base-url", "Mpesa Base URL").Short('b').Envar("MPESA_BASE_URL").StringVar(&cfg.BASE_URL)
	mpesaCLI.Flag("max-idle-conns", "Mpesa Max Idle Connections").Short('m').Envar("MPESA_MAX_IDLE_CONNS").IntVar(&cfg.MAX_IDLE_CONNS)

	cli.AddCommands(mpesaCLI, sdk)

	log.SetFlags(log.Ltime)
	log.SetPrefix("[mpesa] ")

	mpesaCLI.MustParseWithUsage(os.Args[1:])
}
