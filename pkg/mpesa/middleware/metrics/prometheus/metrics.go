// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package prometheus

import (
	"fmt"
	"strings"
	"time"

	"github.com/0x6flab/mpesaoverlay"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var _ mpesa.SDK = (*metricsMiddleware)(nil)

var funcNames = []string{"Token", "ExpressQuery", "ExpressSimulate", "B2CPayment", "AccountBalance", "C2BRegisterURL", "C2BSimulate", "GenerateQR", "Reverse", "TransactionStatus", "RemitTax"}

type metricsMiddleware struct {
	counters  map[string]prom.Counter
	latencies map[string]prom.Histogram
	svcName   string
	pusher    *push.Pusher
	sdk       mpesa.SDK
}

// WithMetrics returns a SDK middleware that instruments various metrics.
func WithMetrics(svcName, url string) mpesa.Option {
	return func(sdk mpesa.SDK) (mpesa.SDK, error) {
		var mm = &metricsMiddleware{
			svcName: fmt.Sprintf("%s_%s", mpesaoverlay.SVCName, strings.ReplaceAll(svcName, "-", "_")),
			sdk:     sdk,
		}

		var counters = make(map[string]prom.Counter)
		var latencies = make(map[string]prom.Histogram)

		var registry = prom.NewRegistry()

		for _, name := range funcNames {
			counters[name] = mm.counter(name)
			latencies[name] = mm.latency(name)
			registry.MustRegister(counters[name])
			registry.MustRegister(latencies[name])
		}

		mm.counters = counters
		mm.latencies = latencies

		mm.pusher = push.New(url, "mpesaoverlay").Gatherer(registry)

		return mm, nil
	}
}

func (mm *metricsMiddleware) Token() (resp mpesa.TokenResp, err error) {
	defer func(begin time.Time) {
		mm.counters["Token"].Inc()
		mm.latencies["Token"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.Token()
}

func (mm *metricsMiddleware) ExpressQuery(eqReq mpesa.ExpressQueryReq) (resp mpesa.ExpressQueryResp, err error) {
	defer func(begin time.Time) {
		mm.counters["ExpressQuery"].Inc()
		mm.latencies["ExpressQuery"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.ExpressQuery(eqReq)
}

func (mm *metricsMiddleware) ExpressSimulate(eReq mpesa.ExpressSimulateReq) (resp mpesa.ExpressSimulateResp, err error) {
	defer func(begin time.Time) {
		mm.counters["ExpressSimulate"].Inc()
		mm.latencies["ExpressSimulate"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.ExpressSimulate(eReq)
}

func (mm *metricsMiddleware) B2CPayment(b2cReq mpesa.B2CPaymentReq) (resp mpesa.B2CPaymentResp, err error) {
	defer func(begin time.Time) {
		mm.counters["B2CPayment"].Inc()
		mm.latencies["B2CPayment"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.B2CPayment(b2cReq)
}

func (mm *metricsMiddleware) AccountBalance(abReq mpesa.AccountBalanceReq) (resp mpesa.AccountBalanceResp, err error) {
	defer func(begin time.Time) {
		mm.counters["AccountBalance"].Inc()
		mm.latencies["AccountBalance"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.AccountBalance(abReq)
}

func (mm *metricsMiddleware) C2BRegisterURL(c2bReq mpesa.C2BRegisterURLReq) (resp mpesa.C2BRegisterURLResp, err error) {
	defer func(begin time.Time) {
		mm.counters["C2BRegisterURL"].Inc()
		mm.latencies["C2BRegisterURL"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.C2BRegisterURL(c2bReq)
}

func (mm *metricsMiddleware) C2BSimulate(c2bReq mpesa.C2BSimulateReq) (resp mpesa.C2BSimulateResp, err error) {
	defer func(begin time.Time) {
		mm.counters["C2BSimulate"].Inc()
		mm.latencies["C2BSimulate"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.C2BSimulate(c2bReq)
}

func (mm *metricsMiddleware) GenerateQR(qReq mpesa.GenerateQRReq) (resp mpesa.GenerateQRResp, err error) {
	defer func(begin time.Time) {
		mm.counters["GenerateQR"].Inc()
		mm.latencies["GenerateQR"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.GenerateQR(qReq)
}

func (mm *metricsMiddleware) Reverse(rReq mpesa.ReverseReq) (resp mpesa.ReverseResp, err error) {
	defer func(begin time.Time) {
		mm.counters["Reverse"].Inc()
		mm.latencies["Reverse"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.Reverse(rReq)
}

func (mm *metricsMiddleware) TransactionStatus(tReq mpesa.TransactionStatusReq) (resp mpesa.TransactionStatusResp, err error) {
	defer func(begin time.Time) {
		mm.counters["TransactionStatus"].Inc()
		mm.latencies["TransactionStatus"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.TransactionStatus(tReq)
}

func (mm *metricsMiddleware) RemitTax(rReq mpesa.RemitTaxReq) (resp mpesa.RemitTaxResp, err error) {
	defer func(begin time.Time) {
		mm.counters["RemitTax"].Inc()
		mm.latencies["RemitTax"].Observe(time.Since(begin).Seconds())
		if err1 := mm.pusher.Add(); err1 != nil {
			err = fmt.Errorf("%w: %w", err, err1)
		}
	}(time.Now())

	return mm.sdk.RemitTax(rReq)
}

func (mm *metricsMiddleware) counter(name string) prom.Counter {
	name = strings.ToLower(name)

	return prom.NewCounter(prom.CounterOpts{
		Namespace: mm.svcName,
		Subsystem: name,
		Name:      "request_count",
		Help:      fmt.Sprintf("Number of requests %s received.", name),
	})
}

func (mm *metricsMiddleware) latency(name string) prom.Histogram {
	name = strings.ToLower(name)

	return prom.NewHistogram(prom.HistogramOpts{
		Namespace: mm.svcName,
		Subsystem: name,
		Name:      "request_latency_microseconds",
		Help:      fmt.Sprintf("Total duration of %s requests in microseconds.", name),
	})
}
