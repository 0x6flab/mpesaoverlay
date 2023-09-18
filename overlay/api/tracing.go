package api

import (
	"context"

	"github.com/0x6flab/mpesaoverlay/overlay"
	"github.com/0x6flab/mpesaoverlay/pkg"
	"go.opentelemetry.io/otel/attribute"
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
	ctx, span := tm.tracer.Start(ctx,
		"get_token",
		trace.WithAttributes(),
	)
	defer span.End()

	return tm.svc.GetToken(ctx)
}

func (tm *tracingMiddleware) ExpressQuery(ctx context.Context, eqReq pkg.ExpressQueryReq) (resp pkg.ExpressQueryResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"ExpressQuery",
		trace.WithAttributes(
			attribute.Int64("BusinessShortCode", int64(eqReq.BusinessShortCode)),
			attribute.String("CheckoutRequestID", eqReq.CheckoutRequestID),
		),
	)
	defer span.End()

	return tm.svc.ExpressQuery(ctx, eqReq)
}

func (tm *tracingMiddleware) ExpressSimulate(ctx context.Context, eReq pkg.ExpressSimulateReq) (resp pkg.ExpressSimulateResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"ExpressSimulate",
		trace.WithAttributes(
			attribute.Int64("BusinessShortCode", int64(eReq.BusinessShortCode)),
			attribute.String("TransactionType", eReq.TransactionType),
			attribute.Int64("Amount", int64(eReq.Amount)),
			attribute.Int64("PartyA", int64(eReq.PartyA)),
			attribute.Int64("PartyB", int64(eReq.PartyB)),
			attribute.Int64("PhoneNumber", int64(eReq.PhoneNumber)),
			attribute.String("AccountReference", eReq.AccountReference),
		),
	)
	defer span.End()

	return tm.svc.ExpressSimulate(ctx, eReq)
}

func (tm *tracingMiddleware) B2CPayment(ctx context.Context, b2cReq pkg.B2CPaymentReq) (resp pkg.B2CPaymentResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"B2CPayment",
		trace.WithAttributes(
			attribute.String("InitiatorName", b2cReq.InitiatorName),
			attribute.String("OriginatorConversationID", b2cReq.OriginatorConversationID),
			attribute.String("CommandID", b2cReq.CommandID),
			attribute.Int64("Amount", int64(b2cReq.Amount)),
			attribute.Int64("PartyA", int64(b2cReq.PartyA)),
			attribute.Int64("PartyB", int64(b2cReq.PartyB)),
			attribute.String("TransactionID", b2cReq.TransactionID),
		),
	)
	defer span.End()

	return tm.svc.B2CPayment(ctx, b2cReq)
}

func (tm *tracingMiddleware) AccountBalance(ctx context.Context, abReq pkg.AccountBalanceReq) (resp pkg.AccountBalanceResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"AccountBalance",
		trace.WithAttributes(
			attribute.String("CommandID", abReq.CommandID),
			attribute.Int64("PartyA", int64(abReq.PartyA)),
			attribute.Int64("IdentifierType", int64(abReq.IdentifierType)),
			attribute.String("InitiatorName", abReq.InitiatorName),
		),
	)
	defer span.End()

	return tm.svc.AccountBalance(ctx, abReq)
}

func (tm *tracingMiddleware) C2BRegisterURL(ctx context.Context, c2bReq pkg.C2BRegisterURLReq) (resp pkg.C2BRegisterURLResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"C2BRegisterURL",
		trace.WithAttributes(
			attribute.String("ResponseType", c2bReq.ResponseType),
			attribute.Int64("ShortCode", int64(c2bReq.ShortCode)),
		),
	)
	defer span.End()

	return tm.svc.C2BRegisterURL(ctx, c2bReq)
}

func (tm *tracingMiddleware) C2BSimulate(ctx context.Context, c2bReq pkg.C2BSimulateReq) (resp pkg.C2BSimulateResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"C2BSimulate",
		trace.WithAttributes(
			attribute.String("CommandID", c2bReq.CommandID),
			attribute.Int64("Amount", int64(c2bReq.Amount)),
			attribute.String("Msisdn", c2bReq.Msisdn),
			attribute.String("BillRefNumber", c2bReq.BillRefNumber),
			attribute.Int64("ShortCode", int64(c2bReq.ShortCode)),
		),
	)
	defer span.End()

	return tm.svc.C2BSimulate(ctx, c2bReq)
}
func (tm *tracingMiddleware) GenerateQR(ctx context.Context, qReq pkg.GenerateQRReq) (resp pkg.GenerateQRResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"GenerateQR",
		trace.WithAttributes(
			attribute.String("MerchantName", qReq.MerchantName),
			attribute.String("RefNo", qReq.RefNo),
			attribute.Int64("Amount", int64(qReq.Amount)),
			attribute.String("TrxCode", qReq.TrxCode),
			attribute.String("CPI", qReq.CPI),
			attribute.String("Size", qReq.Size),
		),
	)
	defer span.End()

	return tm.svc.GenerateQR(ctx, qReq)
}

func (tm *tracingMiddleware) Reverse(ctx context.Context, rReq pkg.ReverseReq) (resp pkg.ReverseResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"Reverse",
		trace.WithAttributes(
			attribute.String("CommandID", rReq.CommandID),
			attribute.String("InitiatorName", rReq.InitiatorName),
			attribute.String("TransactionID", rReq.TransactionID),
			attribute.Int64("Amount", int64(rReq.Amount)),
			attribute.Int64("ReceiverParty", int64(rReq.ReceiverParty)),
			attribute.Int64("ReceiverIdentifierType", int64(rReq.RecieverIdentifierType)),
		),
	)
	defer span.End()

	return tm.svc.Reverse(ctx, rReq)
}

func (tm *tracingMiddleware) TransactionStatus(ctx context.Context, tsReq pkg.TransactionStatusReq) (resp pkg.TransactionStatusResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"TransactionStatus",
		trace.WithAttributes(
			attribute.String("CommandID", tsReq.CommandID),
			attribute.String("Initiator", tsReq.InitiatorName),
			attribute.String("TransactionID", tsReq.TransactionID),
			attribute.Int64("PartyA", int64(tsReq.PartyA)),
			attribute.Int64("IdentifierType", int64(tsReq.IdentifierType)),
		),
	)
	defer span.End()

	return tm.svc.TransactionStatus(ctx, tsReq)
}

func (tm *tracingMiddleware) RemitTax(ctx context.Context, rReq pkg.RemitTaxReq) (resp pkg.RemitTaxResp, err error) {
	ctx, span := tm.tracer.Start(ctx,
		"RemitTax",
		trace.WithAttributes(
			attribute.String("CommandID", rReq.CommandID),
			attribute.String("InitiatorName", rReq.InitiatorName),
			attribute.String("CommandID", rReq.CommandID),
			attribute.Int64("SenderIdentifierType", int64(rReq.SenderIdentifierType)),
			attribute.Int64("ReceiverIdentifierType", int64(rReq.RecieverIdentifierType)),
			attribute.Int64("Amount", int64(rReq.Amount)),
			attribute.Int64("PartyA", int64(rReq.PartyA)),
			attribute.Int64("PartyB", int64(rReq.PartyB)),
			attribute.String("AccountReference", rReq.AccountReference),
		),
	)
	defer span.End()

	return tm.svc.RemitTax(ctx, rReq)
}
