package grpc

import "github.com/0x6flab/mpesaoverlay/pkg"

type getTokenReq struct {
}

func (req getTokenReq) validate() error {
	return nil
}

type expressQueryReq struct {
	pkg.ExpressQueryReq
}

func (req expressQueryReq) validate() error {
	return req.ExpressQueryReq.Validate()
}

type expressSimulateReq struct {
	pkg.ExpressSimulateReq
}

func (req expressSimulateReq) validate() error {
	return req.ExpressSimulateReq.Validate()
}

type b2cReq struct {
	pkg.B2CPaymentReq
}

func (req b2cReq) validate() error {
	return req.B2CPaymentReq.Validate()
}

type accountBalanceReq struct {
	pkg.AccountBalanceReq
}

func (req accountBalanceReq) validate() error {
	return req.AccountBalanceReq.Validate()
}

type c2bRegisterURLReq struct {
	pkg.C2BRegisterURLReq
}

func (req c2bRegisterURLReq) validate() error {
	return req.C2BRegisterURLReq.Validate()
}

type c2bSimulateReq struct {
	pkg.C2BSimulateReq
}

func (req c2bSimulateReq) validate() error {
	return req.C2BSimulateReq.Validate()
}

type generateQRReq struct {
	pkg.GenerateQRReq
}

func (req generateQRReq) validate() error {
	return req.GenerateQRReq.Validate()
}

type reversalReq struct {
	pkg.ReverseReq
}

func (req reversalReq) validate() error {
	return req.ReverseReq.Validate()
}

type transactionReq struct {
	pkg.TransactionStatusReq
}

func (req transactionReq) validate() error {
	return req.TransactionStatusReq.Validate()
}

type remitTaxReq struct {
	pkg.RemitTaxReq
}

func (req remitTaxReq) validate() error {
	return req.RemitTaxReq.Validate()
}
