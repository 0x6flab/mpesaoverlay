package api

import (
	"context"
	"time"

	"github.com/0x6flab/mpesaoverlay/overlay"
	"github.com/0x6flab/mpesaoverlay/pkg"
	"github.com/go-kit/kit/metrics"
)

var _ overlay.Service = (*metricsMiddleware)(nil)

type metricsMiddleware struct {
	counter metrics.Counter
	latency metrics.Histogram
	svc     overlay.Service
}

// MetricsMiddleware instruments policies service by tracking request count and latency.
func MetricsMiddleware(svc overlay.Service, counter metrics.Counter, latency metrics.Histogram) overlay.Service {
	return &metricsMiddleware{
		counter: counter,
		latency: latency,
		svc:     svc,
	}
}

func (mm *metricsMiddleware) GetToken(ctx context.Context) (resp pkg.TokenResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "GetToken").Add(1)
		mm.latency.With("method", "GetToken").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.GetToken(ctx)
}

func (mm *metricsMiddleware) ExpressQuery(ctx context.Context, eqReq pkg.ExpressQueryReq) (resp pkg.ExpressQueryResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "ExpressQuery").Add(1)
		mm.latency.With("method", "ExpressQuery").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.ExpressQuery(ctx, eqReq)
}

func (mm *metricsMiddleware) ExpressSimulate(ctx context.Context, eReq pkg.ExpressSimulateReq) (resp pkg.ExpressSimulateResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "ExpressSimulate").Add(1)
		mm.latency.With("method", "ExpressSimulate").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.ExpressSimulate(ctx, eReq)
}

func (mm *metricsMiddleware) B2CPayment(ctx context.Context, b2cReq pkg.B2Creq) (resp pkg.B2CResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "B2CPayment").Add(1)
		mm.latency.With("method", "B2CPayment").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.B2CPayment(ctx, b2cReq)
}

func (mm *metricsMiddleware) AccountBalance(ctx context.Context, abReq pkg.AccBalanceReq) (resp pkg.AccBalanceResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "AccountBalance").Add(1)
		mm.latency.With("method", "AccountBalance").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.AccountBalance(ctx, abReq)
}

func (mm *metricsMiddleware) C2BRegisterURL(ctx context.Context, c2bReq pkg.C2BRegisterURLReq) (resp pkg.C2BRegisterURLResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "C2BRegisterURL").Add(1)
		mm.latency.With("method", "C2BRegisterURL").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.C2BRegisterURL(ctx, c2bReq)
}

func (mm *metricsMiddleware) C2BSimulate(ctx context.Context, c2bReq pkg.C2BSimulateReq) (resp pkg.C2BSimulateResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "C2BSimulate").Add(1)
		mm.latency.With("method", "C2BSimulate").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.C2BSimulate(ctx, c2bReq)
}

func (mm *metricsMiddleware) GenerateQR(ctx context.Context, qReq pkg.QRReq) (resp pkg.QRResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "GenerateQR").Add(1)
		mm.latency.With("method", "GenerateQR").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.GenerateQR(ctx, qReq)
}

func (mm *metricsMiddleware) Reverse(ctx context.Context, rReq pkg.ReversalReq) (resp pkg.ReversalResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "Reverse").Add(1)
		mm.latency.With("method", "Reverse").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.Reverse(ctx, rReq)
}

func (mm *metricsMiddleware) TransactionStatus(ctx context.Context, tReq pkg.TransactionReq) (resp pkg.TransactionResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "TransactionStatus").Add(1)
		mm.latency.With("method", "TransactionStatus").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.TransactionStatus(ctx, tReq)
}

func (mm *metricsMiddleware) RemitTax(ctx context.Context, rReq pkg.RemitTax) (resp pkg.RemitTaxResp, err error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "RemitTax").Add(1)
		mm.latency.With("method", "RemitTax").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.RemitTax(ctx, rReq)
}
