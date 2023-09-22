package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcadapter "github.com/0x6flab/mpesaoverlay/grpc"
	"github.com/0x6flab/mpesaoverlay/grpc/api"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	zapm "github.com/0x6flab/mpesaoverlay/pkg/mpesa/middleware/logging/zap"
	prometheusm "github.com/0x6flab/mpesaoverlay/pkg/mpesa/middleware/metrics/prometheus"
	"github.com/caarlos0/env/v9"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	svcName      = "grpc-adapter"
	stopWaitTime = 5 * time.Second
)

type config struct {
	LogLevel       string `env:"MO_LOG_LEVEL"            envDefault:"info"`
	ConsumerKey    string `env:"MPESA_CONSUMER_KEY"`
	ConsumerSecret string `env:"MPESA_CONSUMER_SECRET"`
	BaseURL        string `env:"MPESA_BASE_URL"          envDefault:"https://sandbox.safaricom.co.ke"`
	GRPCURL        string `env:"MO_GRPC_URL"             envDefault:"localhost:9000"`
	GRPCServerCert string `env:"MO_GRPC_SERVER_CERT"`
	GRPCServerKey  string `env:"MO_GRPC_SERVER_KEY"`
	PrometheusURL  string `env:"MO_PROMETHEUS_URL"       envDefault:""`
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

	svc, err := newService(ctx, cfg, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create %s service: %s", svcName, err))
	}

	grpcServer, err := initGRPCServer(svc, cfg, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to init %s gRPC server: %s", svcName, err))
	}

	g.Go(func() error {
		return startGRPCServer(cfg, grpcServer)
	})

	g.Go(func() error {
		return StopSignalHandler(ctx, cancel, logger, svcName, grpcServer)
	})

	if err := g.Wait(); err != nil {
		logger.Error(fmt.Sprintf("%s service terminated: %s", svcName, err))
	}
}

func newService(ctx context.Context, cfg config, logger *zap.Logger) (grpcadapter.Service, error) {
	mpesaCfg := mpesa.Config{
		BaseURL:   cfg.BaseURL,
		AppKey:    cfg.ConsumerKey,
		AppSecret: cfg.ConsumerSecret,
		Context:   ctx,
	}
	var opts = []mpesa.Option{zapm.WithLogger(logger)}
	if cfg.PrometheusURL != "" {
		opts = append(opts, prometheusm.WithMetrics(svcName, cfg.PrometheusURL))
	}

	sdk, err := mpesa.NewSDK(mpesaCfg, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create mpesa sdk: %w", err)
	}

	svc := grpcadapter.NewService(sdk)

	return svc, nil
}

func initGRPCServer(svc grpcadapter.Service, cfg config, logger *zap.Logger) (*grpc.Server, error) {
	grpcServerOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	}

	var server *grpc.Server
	if cfg.GRPCServerCert != "" || cfg.GRPCServerKey != "" {
		creds, err := credentials.NewServerTLSFromFile(cfg.GRPCServerCert, cfg.GRPCServerKey)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s certificates: %w", svcName, err)
		}

		grpcServerOptions = append(grpcServerOptions, grpc.Creds(creds))

		logger.Info(fmt.Sprintf("%s gRPC service started using https on url %s with cert %s key %s", svcName, cfg.GRPCURL, cfg.GRPCServerCert, cfg.GRPCServerKey))
	} else {
		grpcServerOptions = append(grpcServerOptions, grpc.Creds(insecure.NewCredentials()))

		logger.Info(fmt.Sprintf("%s gRPC service started using http on url %s", svcName, cfg.GRPCURL))
	}

	server = grpc.NewServer(grpcServerOptions...)

	reflection.Register(server)
	grpcadapter.RegisterServiceServer(server, api.NewServer(svc))

	return server, nil
}

func startGRPCServer(cfg config, server *grpc.Server) error {
	listener, err := net.Listen("tcp", cfg.GRPCURL)
	if err != nil {
		return fmt.Errorf("failed to start %s gRPC service: %w", svcName, err)
	}

	return server.Serve(listener)
}

func StopSignalHandler(ctx context.Context, cancel context.CancelFunc, logger *zap.Logger, svcName string, server *grpc.Server) error {
	var err error
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	select {
	case sig := <-c:
		defer cancel()
		c := make(chan bool)
		go func() {
			defer close(c)
			server.GracefulStop()
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
