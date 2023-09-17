package api

import (
	"context"
	"time"

	"github.com/0x6flab/mpesaoverlay/overlay"
	"github.com/0x6flab/mpesaoverlay/pkg"
	"go.uber.org/zap"
)

var _ overlay.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *zap.Logger
	svc    overlay.Service
}

func LoggingMiddleware(svc overlay.Service, logger *zap.Logger) overlay.Service {
	return &loggingMiddleware{logger, svc}
}

func (lm *loggingMiddleware) GetToken(ctx context.Context) (resp pkg.TokenResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"GetToken",
			zap.Error(err),
		)
	}(time.Now())

	return lm.svc.GetToken(ctx)
}

func (lm *loggingMiddleware) ExpressQuery(ctx context.Context, eqReq pkg.ExpressQueryReq) (resp pkg.ExpressQueryResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"ExpressQuery",
			zap.Error(err),
			zap.Uint64("BusinessShortCode", eqReq.BusinessShortCode),
			zap.String("CheckoutRequestID", eqReq.CheckoutRequestID),
		)
	}(time.Now())

	return lm.svc.ExpressQuery(ctx, eqReq)
}

func (lm *loggingMiddleware) ExpressSimulate(ctx context.Context, eReq pkg.ExpressSimulateReq) (resp pkg.ExpressSimulateResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"ExpressSimulate",
			zap.Error(err),
			zap.Uint64("BusinessShortCode", eReq.BusinessShortCode),
			zap.String("TransactionType", eReq.TransactionType),
			zap.Uint64("Amount", eReq.Amount),
			zap.Uint64("PartyA", eReq.PartyA),
			zap.Uint64("PartyB", eReq.PartyB),
			zap.Uint64("PhoneNumber", eReq.PhoneNumber),
			zap.String("AccountReference", eReq.AccountReference),
		)
	}(time.Now())

	return lm.svc.ExpressSimulate(ctx, eReq)
}

func (lm *loggingMiddleware) B2CPayment(ctx context.Context, b2cReq pkg.B2CPaymentReq) (resp pkg.B2CPaymentResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"B2CPayment",
			zap.Error(err),
			zap.String("InitiatorName", b2cReq.InitiatorName),
			zap.String("OriginatorConversationID", b2cReq.OriginatorConversationID),
			zap.String("CommandID", b2cReq.CommandID),
			zap.Uint64("Amount", b2cReq.Amount),
			zap.Uint64("PartyA", b2cReq.PartyA),
			zap.Uint64("PartyB", b2cReq.PartyB),
			zap.String("TransactionID", b2cReq.TransactionID),
		)
	}(time.Now())

	return lm.svc.B2CPayment(ctx, b2cReq)
}

func (lm *loggingMiddleware) AccountBalance(ctx context.Context, abReq pkg.AccountBalanceReq) (resp pkg.AccountBalanceResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"AccountBalance",
			zap.Error(err),
			zap.String("CommandID", abReq.CommandID),
			zap.Uint64("PartyA", abReq.PartyA),
			zap.Uint8("IdentifierType", abReq.IdentifierType),
			zap.String("InitiatorName", abReq.InitiatorName),
		)
	}(time.Now())

	return lm.svc.AccountBalance(ctx, abReq)
}

func (lm *loggingMiddleware) C2BRegisterURL(ctx context.Context, c2bReq pkg.C2BRegisterURLReq) (resp pkg.C2BRegisterURLResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"C2BRegisterURL",
			zap.Error(err),
			zap.String("ResponseType", c2bReq.ResponseType),
			zap.Uint64("ShortCode", c2bReq.ShortCode),
		)
	}(time.Now())

	return lm.svc.C2BRegisterURL(ctx, c2bReq)
}

func (lm *loggingMiddleware) C2BSimulate(ctx context.Context, c2bReq pkg.C2BSimulateReq) (resp pkg.C2BSimulateResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"C2BSimulate",
			zap.Error(err),
			zap.String("CommandID", c2bReq.CommandID),
			zap.Uint64("Amount", c2bReq.Amount),
			zap.String("Msisdn", c2bReq.Msisdn),
			zap.String("BillRefNumber", c2bReq.BillRefNumber),
			zap.Uint64("ShortCode", c2bReq.ShortCode),
		)
	}(time.Now())

	return lm.svc.C2BSimulate(ctx, c2bReq)
}

func (lm *loggingMiddleware) GenerateQR(ctx context.Context, qReq pkg.GenerateQRReq) (resp pkg.GenerateQRResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"GenerateQR",
			zap.Error(err),
			zap.String("MerchantName", qReq.MerchantName),
			zap.String("RefNo", qReq.RefNo),
			zap.Uint64("Amount", qReq.Amount),
			zap.String("TrxCode", qReq.TrxCode),
			zap.String("CPI", qReq.CPI),
			zap.String("Size", qReq.Size),
		)
	}(time.Now())

	return lm.svc.GenerateQR(ctx, qReq)
}

func (lm *loggingMiddleware) Reverse(ctx context.Context, rReq pkg.ReverseReq) (resp pkg.ReverseResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"Reverse",
			zap.Error(err),
			zap.String("CommandID", rReq.CommandID),
			zap.String("InitiatorName", rReq.InitiatorName),
			zap.String("TransactionID", rReq.TransactionID),
			zap.Uint64("Amount", rReq.Amount),
			zap.Uint64("ReceiverParty", rReq.ReceiverParty),
			zap.Uint8("ReceiverIdentifierType", rReq.RecieverIdentifierType),
		)
	}(time.Now())

	return lm.svc.Reverse(ctx, rReq)
}

func (lm *loggingMiddleware) TransactionStatus(ctx context.Context, tReq pkg.TransactionStatusReq) (resp pkg.TransactionStatusResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"TransactionStatus",
			zap.Error(err),
			zap.String("CommandID", tReq.CommandID),
			zap.String("Initiator", tReq.InitiatorName),
			zap.String("TransactionID", tReq.TransactionID),
			zap.Uint64("PartyA", tReq.PartyA),
			zap.Uint8("IdentifierType", tReq.IdentifierType),
		)
	}(time.Now())

	return lm.svc.TransactionStatus(ctx, tReq)
}

func (lm *loggingMiddleware) RemitTax(ctx context.Context, rReq pkg.RemitTaxReq) (resp pkg.RemitTaxResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"RemitTax",
			zap.Error(err),
			zap.String("CommandID", rReq.CommandID),
			zap.String("InitiatorName", rReq.InitiatorName),
			zap.String("CommandID", rReq.CommandID),
			zap.Uint8("SenderIdentifierType", rReq.SenderIdentifierType),
			zap.Uint8("ReceiverIdentifierType", rReq.RecieverIdentifierType),
			zap.Uint64("Amount", rReq.Amount),
			zap.Uint64("PartyA", rReq.PartyA),
			zap.Uint64("PartyB", rReq.PartyB),
			zap.String("AccountReference", rReq.AccountReference),
		)
	}(time.Now())

	return lm.svc.RemitTax(ctx, rReq)
}
