package postgres

import (
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"gorm.io/gorm"
)

type expressQueryReq struct {
	gorm.Model
	mpesa.ExpressQueryReq
	id string
}

type expressSimulateReq struct {
	gorm.Model
	mpesa.ExpressSimulateReq
	id string
}

type b2cPaymentReq struct {
	gorm.Model
	mpesa.B2CPaymentReq
	id string
}

type accountBalanceReq struct {
	gorm.Model
	mpesa.AccountBalanceReq
	id string
}

type c2bRegisterURLReq struct {
	gorm.Model
	mpesa.C2BRegisterURLReq
	id string
}

type c2bSimulateReq struct {
	gorm.Model
	mpesa.C2BSimulateReq
	id string
}

type generateQRReq struct {
	gorm.Model
	mpesa.GenerateQRReq
	id string
}

type reverseReq struct {
	gorm.Model
	mpesa.ReverseReq
	id string
}

type transactionStatusReq struct {
	gorm.Model
	mpesa.TransactionStatusReq
	id string
}

type remitTaxReq struct {
	gorm.Model
	mpesa.RemitTaxReq
	id string
}
