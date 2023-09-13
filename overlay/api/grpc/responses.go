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
	pkg.B2CResp
}

type accountBalanceResp struct {
	pkg.AccBalanceResp
}

type c2bRegisterURLResp struct {
	pkg.C2BRegisterURLResp
}

type c2bSimulateResp struct {
	pkg.C2BSimulateResp
}

type generateQRResp struct {
	pkg.QRResp
}

type reverseResp struct {
	pkg.ReversalResp
}

type transactionStatusResp struct {
	pkg.TransactionResp
}

type remitTaxResp struct {
	pkg.RemitTaxResp
}
