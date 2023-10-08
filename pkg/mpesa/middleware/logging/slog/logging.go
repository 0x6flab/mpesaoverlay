// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package slog

import (
	log "log/slog"
	"time"

	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
)

var _ mpesa.SDK = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *log.Logger
	sdk    mpesa.SDK
}

// WithLogger returns a logging middleware using slog.
func WithLogger(logger *log.Logger) mpesa.Option {
	return func(sdk mpesa.SDK) (mpesa.SDK, error) {
		return &loggingMiddleware{logger, sdk}, nil
	}
}

func (lm *loggingMiddleware) Token() (resp mpesa.TokenResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"Token",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
		)
	}(time.Now())

	return lm.sdk.Token()
}

func (lm *loggingMiddleware) ExpressQuery(eqReq mpesa.ExpressQueryReq) (resp mpesa.ExpressQueryResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"ExpressQuery",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.Uint64("BusinessShortCode", eqReq.BusinessShortCode),
			log.String("CheckoutRequestID", eqReq.CheckoutRequestID),
		)
	}(time.Now())

	return lm.sdk.ExpressQuery(eqReq)
}

func (lm *loggingMiddleware) ExpressSimulate(eReq mpesa.ExpressSimulateReq) (resp mpesa.ExpressSimulateResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"ExpressSimulate",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.Uint64("BusinessShortCode", eReq.BusinessShortCode),
			log.String("TransactionType", eReq.TransactionType),
			log.Uint64("Amount", eReq.Amount),
			log.Uint64("PartyA", eReq.PartyA),
			log.Uint64("PartyB", eReq.PartyB),
			log.Uint64("PhoneNumber", eReq.PhoneNumber),
			log.String("AccountReference", eReq.AccountReference),
		)
	}(time.Now())

	return lm.sdk.ExpressSimulate(eReq)
}

func (lm *loggingMiddleware) B2CPayment(b2cReq mpesa.B2CPaymentReq) (resp mpesa.B2CPaymentResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"B2CPayment",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("InitiatorName", b2cReq.InitiatorName),
			log.String("OriginatorConversationID", b2cReq.OriginatorConversationID),
			log.String("CommandID", b2cReq.CommandID),
			log.Uint64("Amount", b2cReq.Amount),
			log.Uint64("PartyA", b2cReq.PartyA),
			log.Uint64("PartyB", b2cReq.PartyB),
			log.String("TransactionID", b2cReq.TransactionID),
		)
	}(time.Now())

	return lm.sdk.B2CPayment(b2cReq)
}

func (lm *loggingMiddleware) AccountBalance(abReq mpesa.AccountBalanceReq) (resp mpesa.AccountBalanceResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"AccountBalance",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("CommandID", abReq.CommandID),
			log.Uint64("PartyA", abReq.PartyA),
			log.Int("IdentifierType", int(abReq.IdentifierType)),
			log.String("InitiatorName", abReq.InitiatorName),
		)
	}(time.Now())

	return lm.sdk.AccountBalance(abReq)
}

func (lm *loggingMiddleware) C2BRegisterURL(c2bReq mpesa.C2BRegisterURLReq) (resp mpesa.C2BRegisterURLResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"C2BRegisterURL",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("ResponseType", c2bReq.ResponseType),
			log.Uint64("ShortCode", c2bReq.ShortCode),
		)
	}(time.Now())

	return lm.sdk.C2BRegisterURL(c2bReq)
}

func (lm *loggingMiddleware) C2BSimulate(c2bReq mpesa.C2BSimulateReq) (resp mpesa.C2BSimulateResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"C2BSimulate",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("CommandID", c2bReq.CommandID),
			log.Uint64("Amount", c2bReq.Amount),
			log.String("Msisdn", c2bReq.Msisdn),
			log.String("BillRefNumber", c2bReq.BillRefNumber),
			log.Uint64("ShortCode", c2bReq.ShortCode),
		)
	}(time.Now())

	return lm.sdk.C2BSimulate(c2bReq)
}

func (lm *loggingMiddleware) GenerateQR(qReq mpesa.GenerateQRReq) (resp mpesa.GenerateQRResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"GenerateQR",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("MerchantName", qReq.MerchantName),
			log.String("RefNo", qReq.RefNo),
			log.Uint64("Amount", qReq.Amount),
			log.String("TrxCode", qReq.TrxCode),
			log.String("CPI", qReq.CPI),
			log.String("Size", qReq.Size),
		)
	}(time.Now())

	return lm.sdk.GenerateQR(qReq)
}

func (lm *loggingMiddleware) Reverse(rReq mpesa.ReverseReq) (resp mpesa.ReverseResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"Reverse",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("CommandID", rReq.CommandID),
			log.String("InitiatorName", rReq.InitiatorName),
			log.String("TransactionID", rReq.TransactionID),
			log.Uint64("Amount", rReq.Amount),
			log.Uint64("ReceiverParty", rReq.ReceiverParty),
			log.Int("RecieverIdentifierType", int(rReq.RecieverIdentifierType)),
		)
	}(time.Now())

	return lm.sdk.Reverse(rReq)
}

func (lm *loggingMiddleware) TransactionStatus(tReq mpesa.TransactionStatusReq) (resp mpesa.TransactionStatusResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"TransactionStatus",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("CommandID", tReq.CommandID),
			log.String("Initiator", tReq.InitiatorName),
			log.String("TransactionID", tReq.TransactionID),
			log.Uint64("PartyA", tReq.PartyA),
			log.Int("IdentifierType", int(tReq.IdentifierType)),
		)
	}(time.Now())

	return lm.sdk.TransactionStatus(tReq)
}

func (lm *loggingMiddleware) RemitTax(rReq mpesa.RemitTaxReq) (resp mpesa.RemitTaxResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"RemitTax",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("CommandID", rReq.CommandID),
			log.String("InitiatorName", rReq.InitiatorName),
			log.String("CommandID", rReq.CommandID),
			log.Int("SenderIdentifierType", int(rReq.SenderIdentifierType)),
			log.Int("RecieverIdentifierType", int(rReq.RecieverIdentifierType)),
			log.Uint64("Amount", rReq.Amount),
			log.Uint64("PartyA", rReq.PartyA),
			log.Uint64("PartyB", rReq.PartyB),
			log.String("AccountReference", rReq.AccountReference),
		)
	}(time.Now())

	return lm.sdk.RemitTax(rReq)
}

func (lm *loggingMiddleware) BusinessPayBill(bReq mpesa.BusinessPayBillReq) (resp mpesa.BusinessPayBillResp, err error) {
	defer func(begin time.Time) {
		lm.logger.Info(
			"BusinessPayBill",
			log.Any("error", err),
			log.String("duration", time.Since(begin).String()),
			log.String("CommandID", bReq.CommandID),
			log.String("Initiator", bReq.Initiator),
			log.Int("SenderIdentifierType", int(bReq.SenderIdentifierType)),
			log.Int("RecieverIdentifierType", int(bReq.RecieverIdentifierType)),
			log.Uint64("Amount", bReq.Amount),
			log.Uint64("PartyA", bReq.PartyA),
			log.Uint64("PartyB", bReq.PartyB),
			log.String("AccountReference", bReq.AccountReference),
			log.Uint64("Requester", bReq.Requester),
		)
	}(time.Now())

	return lm.sdk.BusinessPayBill(bReq)
}
