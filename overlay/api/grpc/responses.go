package grpc

import "github.com/0x6flab/mpesaoverlay/pkg"

type getTokenResp struct {
	pkg.TokenResp
}

type expressQueryResp struct {
	pkg.ExpressQueryResp
}

type expressSimulateResp struct {
	pkg.ExpressSimulateResp
}

type b2cResp struct {
	pkg.B2CPaymentResp
}

type accountBalanceResp struct {
	pkg.AccountBalanceResp
}

type c2bRegisterURLResp struct {
	pkg.C2BRegisterURLResp
}

type c2bSimulateResp struct {
	pkg.C2BSimulateResp
}

type generateQRResp struct {
	pkg.GenerateQRResp
}

type reverseResp struct {
	pkg.ReverseResp
}

type transactionStatusResp struct {
	pkg.TransactionStatusResp
}

type remitTaxResp struct {
	pkg.RemitTaxResp
}
