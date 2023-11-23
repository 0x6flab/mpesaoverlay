// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package postgres

import (
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/oklog/ulid/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ mpesa.SDK = (*postgresMiddleware)(nil)

type postgresMiddleware struct {
	db  *gorm.DB
	sdk mpesa.SDK
}

// WithDatabase returns a database middleware using postgres.
func WithDatabase(url string) mpesa.Option {
	return func(sdk mpesa.SDK) (mpesa.SDK, error) {
		db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			return sdk, err
		}
		tables := []interface{}{
			&expressQueryReq{},
			&expressSimulateReq{},
			&b2cPaymentReq{},
			&accountBalanceReq{},
			&c2bRegisterURLReq{},
			&c2bSimulateReq{},
			&generateQRReq{},
			&reverseReq{},
			&transactionStatusReq{},
			&remitTaxReq{},
			&businessPayBillReq{},
		}

		if err := db.AutoMigrate(tables...); err != nil {
			return sdk, err
		}

		return &postgresMiddleware{db, sdk}, nil
	}
}

func (pm *postgresMiddleware) Token() (resp mpesa.TokenResp, err error) {
	return pm.sdk.Token()
}

func (pm *postgresMiddleware) ExpressQuery(eqReq mpesa.ExpressQueryReq) (resp mpesa.ExpressQueryResp, err error) {
	defer func() {
		req := expressQueryReq{
			ExpressQueryReq: eqReq,
			id:              ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.ExpressQuery(eqReq)
}

func (pm *postgresMiddleware) ExpressSimulate(eReq mpesa.ExpressSimulateReq) (resp mpesa.ExpressSimulateResp, err error) {
	defer func() {
		req := expressSimulateReq{
			ExpressSimulateReq: eReq,
			id:                 ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.ExpressSimulate(eReq)
}

func (pm *postgresMiddleware) B2CPayment(b2cReq mpesa.B2CPaymentReq) (resp mpesa.B2CPaymentResp, err error) {
	defer func() {
		req := b2cPaymentReq{
			B2CPaymentReq: b2cReq,
			id:            ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.B2CPayment(b2cReq)
}

func (pm *postgresMiddleware) AccountBalance(abReq mpesa.AccountBalanceReq) (resp mpesa.AccountBalanceResp, err error) {
	defer func() {
		req := accountBalanceReq{
			AccountBalanceReq: abReq,
			id:                ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.AccountBalance(abReq)
}

func (pm *postgresMiddleware) C2BRegisterURL(c2bReq mpesa.C2BRegisterURLReq) (resp mpesa.C2BRegisterURLResp, err error) {
	defer func() {
		req := c2bRegisterURLReq{
			C2BRegisterURLReq: c2bReq,
			id:                ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.C2BRegisterURL(c2bReq)
}

func (pm *postgresMiddleware) C2BSimulate(c2bReq mpesa.C2BSimulateReq) (resp mpesa.C2BSimulateResp, err error) {
	defer func() {
		req := c2bSimulateReq{
			C2BSimulateReq: c2bReq,
			id:             ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.C2BSimulate(c2bReq)
}

func (pm *postgresMiddleware) GenerateQR(gqrReq mpesa.GenerateQRReq) (resp mpesa.GenerateQRResp, err error) {
	defer func() {
		req := generateQRReq{
			GenerateQRReq: gqrReq,
			id:            ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.GenerateQR(gqrReq)
}

func (pm *postgresMiddleware) Reverse(rReq mpesa.ReverseReq) (resp mpesa.ReverseResp, err error) {
	defer func() {
		req := reverseReq{
			ReverseReq: rReq,
			id:         ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.Reverse(rReq)
}

func (pm *postgresMiddleware) TransactionStatus(tsReq mpesa.TransactionStatusReq) (resp mpesa.TransactionStatusResp, err error) {
	defer func() {
		req := transactionStatusReq{
			TransactionStatusReq: tsReq,
			id:                   ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.TransactionStatus(tsReq)
}

func (pm *postgresMiddleware) RemitTax(rtReq mpesa.RemitTaxReq) (resp mpesa.RemitTaxResp, err error) {
	defer func() {
		req := remitTaxReq{
			RemitTaxReq: rtReq,
			id:          ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.RemitTax(rtReq)
}

func (pm *postgresMiddleware) BusinessPayBill(bpbReq mpesa.BusinessPayBillReq) (resp mpesa.BusinessPayBillResp, err error) {
	defer func() {
		req := businessPayBillReq{
			BusinessPayBillReq: bpbReq,
			id:                 ulid.Make().String(),
		}
		_ = pm.db.Create(&req)
	}()

	return pm.sdk.BusinessPayBill(bpbReq)
}
