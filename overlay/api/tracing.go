package api

import (
	"context"
	"fmt"

	"github.com/0x6flab/mpesaoverlay/overlay"
	"github.com/0x6flab/mpesaoverlay/pkg"
	"go.opentelemetry.io/otel/trace"
)

type tracingMiddleware struct {
	tracer trace.Tracer
	svc    overlay.Service
}

func TracingMiddleware(svc overlay.Service, tracer trace.Tracer) overlay.Service {
	return &tracingMiddleware{tracer, svc}
}

func (tm *tracingMiddleware) GetToken(ctx context.Context) (resp pkg.TokenResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "get_token")
	defer span.End()

	fmt.Println("GetToken")

	return tm.svc.GetToken(ctx)
}

func (tm *tracingMiddleware) ExpressQuery(ctx context.Context, eqReq pkg.ExpressQueryReq) (resp pkg.ExpressQueryResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "ExpressQuery")
	defer span.End()

	return tm.svc.ExpressQuery(ctx, eqReq)
}

func (tm *tracingMiddleware) ExpressSimulate(ctx context.Context, eReq pkg.ExpressSimulateReq) (resp pkg.ExpressSimulateResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "ExpressSimulate")
	defer span.End()

	return tm.svc.ExpressSimulate(ctx, eReq)
}

func (tm *tracingMiddleware) B2CPayment(ctx context.Context, b2cReq pkg.B2CPaymentReq) (resp pkg.B2CPaymentResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "B2CPayment")
	defer span.End()

	return tm.svc.B2CPayment(ctx, b2cReq)
}

func (tm *tracingMiddleware) AccountBalance(ctx context.Context, abReq pkg.AccountBalanceReq) (resp pkg.AccountBalanceResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "AccountBalance")
	defer span.End()

	return tm.svc.AccountBalance(ctx, abReq)
}

func (tm *tracingMiddleware) C2BRegisterURL(ctx context.Context, c2bReq pkg.C2BRegisterURLReq) (resp pkg.C2BRegisterURLResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "C2BRegisterURL")
	defer span.End()

	return tm.svc.C2BRegisterURL(ctx, c2bReq)
}

func (tm *tracingMiddleware) C2BSimulate(ctx context.Context, c2bReq pkg.C2BSimulateReq) (resp pkg.C2BSimulateResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "C2BSimulate")
	defer span.End()

	return tm.svc.C2BSimulate(ctx, c2bReq)
}
func (tm *tracingMiddleware) GenerateQR(ctx context.Context, qReq pkg.GenerateQRReq) (resp pkg.GenerateQRResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "GenerateQR")
	defer span.End()

	return tm.svc.GenerateQR(ctx, qReq)
}

func (tm *tracingMiddleware) Reverse(ctx context.Context, rReq pkg.ReverseReq) (resp pkg.ReverseResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "Reverse")
	defer span.End()

	return tm.svc.Reverse(ctx, rReq)
}

func (tm *tracingMiddleware) TransactionStatus(ctx context.Context, tsReq pkg.TransactionStatusReq) (resp pkg.TransactionStatusResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "TransactionStatus")
	defer span.End()

	return tm.svc.TransactionStatus(ctx, tsReq)
}

func (tm *tracingMiddleware) RemitTax(ctx context.Context, abqReq pkg.RemitTaxReq) (resp pkg.RemitTaxResp, err error) {
	ctx, span := tm.tracer.Start(ctx, "RemitTax")
	defer span.End()

	return tm.svc.RemitTax(ctx, abqReq)
}
