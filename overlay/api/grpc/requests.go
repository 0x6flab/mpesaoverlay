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
	pkg.B2Creq
}

func (req b2cReq) validate() error {
	return req.B2Creq.Validate()
}

type accountBalanceReq struct {
	pkg.AccBalanceReq
}

func (req accountBalanceReq) validate() error {
	return req.AccBalanceReq.Validate()
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
	pkg.QRReq
}

func (req generateQRReq) validate() error {
	return req.QRReq.Validate()
}

type reversalReq struct {
	pkg.ReversalReq
}

func (req reversalReq) validate() error {
	return req.ReversalReq.Validate()
}

type transactionReq struct {
	pkg.TransactionReq
}

func (req transactionReq) validate() error {
	return req.TransactionReq.Validate()
}

type remitTaxReq struct {
	pkg.RemitTax
}

func (req remitTaxReq) validate() error {
	return req.RemitTax.Validate()
}
