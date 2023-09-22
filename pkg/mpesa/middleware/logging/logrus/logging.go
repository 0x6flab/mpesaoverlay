package logrus

import (
	"time"

	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	log "github.com/sirupsen/logrus"
)

var _ mpesa.SDK = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *log.Logger
	sdk    mpesa.SDK
}

func WithLogger(logger *log.Logger) mpesa.Option {
	return func(sdk mpesa.SDK) mpesa.SDK {
		logger.SetFormatter(&log.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
		logger.SetReportCaller(true)

		return &loggingMiddleware{logger, sdk}
	}
}

func (lm *loggingMiddleware) GetToken() (resp mpesa.TokenResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration": time.Since(begin).String(),
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("GetToken")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("GetToken")
		}
	}(time.Now())

	return lm.sdk.GetToken()
}

func (lm *loggingMiddleware) ExpressQuery(eqReq mpesa.ExpressQueryReq) (resp mpesa.ExpressQueryResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":          time.Since(begin).String(),
			"BusinessShortCode": eqReq.BusinessShortCode,
			"CheckoutRequestID": eqReq.CheckoutRequestID,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("ExpressQuery")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("ExpressQuery")
		}
	}(time.Now())

	return lm.sdk.ExpressQuery(eqReq)
}

func (lm *loggingMiddleware) ExpressSimulate(eReq mpesa.ExpressSimulateReq) (resp mpesa.ExpressSimulateResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":          time.Since(begin).String(),
			"BusinessShortCode": eReq.BusinessShortCode,
			"TransactionType":   eReq.TransactionType,
			"Amount":            eReq.Amount,
			"PartyA":            eReq.PartyA,
			"PartyB":            eReq.PartyB,
			"PhoneNumber":       eReq.PhoneNumber,
			"AccountReference":  eReq.AccountReference,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("ExpressSimulate")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("ExpressSimulate")
		}
	}(time.Now())

	return lm.sdk.ExpressSimulate(eReq)
}

func (lm *loggingMiddleware) B2CPayment(b2cReq mpesa.B2CPaymentReq) (resp mpesa.B2CPaymentResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":               time.Since(begin).String(),
			"InitiatorName":          b2cReq.InitiatorName,
			"OriginatorConversation": b2cReq.OriginatorConversationID,
			"CommandID":              b2cReq.CommandID,
			"Amount":                 b2cReq.Amount,
			"PartyA":                 b2cReq.PartyA,
			"PartyB":                 b2cReq.PartyB,
			"TransactionID":          b2cReq.TransactionID,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("B2CPayment")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("B2CPayment")
		}
	}(time.Now())

	return lm.sdk.B2CPayment(b2cReq)
}

func (lm *loggingMiddleware) AccountBalance(abReq mpesa.AccountBalanceReq) (resp mpesa.AccountBalanceResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":           time.Since(begin).String(),
			"CommandID":          abReq.CommandID,
			"PartyA":             abReq.PartyA,
			"IdentifierType":     abReq.IdentifierType,
			"InitiatorName":      abReq.InitiatorName,
			"SecurityCredential": abReq.SecurityCredential,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("AccountBalance")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("AccountBalance")
		}
	}(time.Now())

	return lm.sdk.AccountBalance(abReq)
}

func (lm *loggingMiddleware) C2BRegisterURL(c2bReq mpesa.C2BRegisterURLReq) (resp mpesa.C2BRegisterURLResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":     time.Since(begin).String(),
			"ResponseType": c2bReq.ResponseType,
			"ShortCode":    c2bReq.ShortCode,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("C2BRegisterURL")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("C2BRegisterURL")
		}
	}(time.Now())

	return lm.sdk.C2BRegisterURL(c2bReq)
}

func (lm *loggingMiddleware) C2BSimulate(c2bReq mpesa.C2BSimulateReq) (resp mpesa.C2BSimulateResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":      time.Since(begin).String(),
			"CommandID":     c2bReq.CommandID,
			"Amount":        c2bReq.Amount,
			"Msisdn":        c2bReq.Msisdn,
			"BillRefNumber": c2bReq.BillRefNumber,
			"ShortCode":     c2bReq.ShortCode,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("C2BSimulate")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("C2BSimulate")
		}
	}(time.Now())

	return lm.sdk.C2BSimulate(c2bReq)
}

func (lm *loggingMiddleware) GenerateQR(qReq mpesa.GenerateQRReq) (resp mpesa.GenerateQRResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":     time.Since(begin).String(),
			"MerchantName": qReq.MerchantName,
			"RefNo":        qReq.RefNo,
			"Amount":       qReq.Amount,
			"TrxCode":      qReq.TrxCode,
			"CPI":          qReq.CPI,
			"Size":         qReq.Size,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("GenerateQR")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("GenerateQR")
		}
	}(time.Now())

	return lm.sdk.GenerateQR(qReq)
}

func (lm *loggingMiddleware) Reverse(rReq mpesa.ReverseReq) (resp mpesa.ReverseResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":               time.Since(begin).String(),
			"CommandID":              rReq.CommandID,
			"InitiatorName":          rReq.InitiatorName,
			"TransactionID":          rReq.TransactionID,
			"Amount":                 rReq.Amount,
			"ReceiverParty":          rReq.ReceiverParty,
			"RecieverIdentifierType": rReq.RecieverIdentifierType,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("Reverse")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("Reverse")
		}
	}(time.Now())

	return lm.sdk.Reverse(rReq)
}

func (lm *loggingMiddleware) TransactionStatus(tReq mpesa.TransactionStatusReq) (resp mpesa.TransactionStatusResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":       time.Since(begin).String(),
			"CommandID":      tReq.CommandID,
			"Initiator":      tReq.InitiatorName,
			"TransactionID":  tReq.TransactionID,
			"PartyA":         tReq.PartyA,
			"IdentifierType": tReq.IdentifierType,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("TransactionStatus")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("TransactionStatus")
		}
	}(time.Now())

	return lm.sdk.TransactionStatus(tReq)
}

func (lm *loggingMiddleware) RemitTax(rReq mpesa.RemitTaxReq) (resp mpesa.RemitTaxResp, err error) {
	defer func(begin time.Time) {
		var fields = log.Fields{
			"duration":               time.Since(begin).String(),
			"CommandID":              rReq.CommandID,
			"InitiatorName":          rReq.InitiatorName,
			"SenderIdentifierType":   rReq.SenderIdentifierType,
			"ReceiverIdentifierType": rReq.RecieverIdentifierType,
			"Amount":                 rReq.Amount,
			"PartyA":                 rReq.PartyA,
			"PartyB":                 rReq.PartyB,
			"AccountReference":       rReq.AccountReference,
		}
		switch err {
		case nil:
			lm.logger.WithFields(fields).Info("RemitTax")
		default:
			fields["error"] = err
			lm.logger.WithFields(fields).Error("RemitTax")
		}
	}(time.Now())

	return lm.sdk.RemitTax(rReq)
}
