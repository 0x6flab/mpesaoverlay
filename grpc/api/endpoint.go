// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"errors"

	"github.com/0x6flab/mpesaoverlay/grpc"
	"github.com/go-kit/kit/endpoint"
)

// errValidation is returned when a request validation has failed.
var errValidation = errors.New("validation error")

func tokenEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(tokenReq)
		if err := req.validate(); err != nil {
			return tokenResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.Token()
		if err != nil {
			return tokenResp{}, err
		}

		return tokenResp{resp}, nil
	}
}

func expressQueryEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(expressQueryReq)
		if err := req.validate(); err != nil {
			return expressQueryResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.ExpressQuery(req.ExpressQueryReq)
		if err != nil {
			return expressQueryResp{}, err
		}

		return expressQueryResp{resp}, nil
	}
}

func expressSimulateEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(expressSimulateReq)
		if err := req.validate(); err != nil {
			return expressSimulateResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.ExpressSimulate(req.ExpressSimulateReq)
		if err != nil {
			return expressSimulateResp{}, err
		}

		return expressSimulateResp{resp}, nil
	}
}

func b2cEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(b2cReq)
		if err := req.validate(); err != nil {
			return b2cResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.B2CPayment(req.B2CPaymentReq)
		if err != nil {
			return b2cResp{}, err
		}

		return b2cResp{resp}, nil
	}
}

func accountBalanceEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(accountBalanceReq)
		if err := req.validate(); err != nil {
			return accountBalanceResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.AccountBalance(req.AccountBalanceReq)
		if err != nil {
			return accountBalanceResp{}, err
		}

		return accountBalanceResp{resp}, nil
	}
}

func c2bRegisterURLEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(c2bRegisterURLReq)
		if err := req.validate(); err != nil {
			return c2bRegisterURLResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.C2BRegisterURL(req.C2BRegisterURLReq)
		if err != nil {
			return c2bRegisterURLResp{}, err
		}

		return c2bRegisterURLResp{resp}, nil
	}
}

func c2bSimulateEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(c2bSimulateReq)
		if err := req.validate(); err != nil {
			return c2bSimulateResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.C2BSimulate(req.C2BSimulateReq)
		if err != nil {
			return c2bSimulateResp{}, err
		}

		return c2bSimulateResp{resp}, nil
	}
}

func generateQREndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(generateQRReq)
		if err := req.validate(); err != nil {
			return generateQRResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.GenerateQR(req.GenerateQRReq)
		if err != nil {
			return generateQRResp{}, err
		}

		return generateQRResp{resp}, nil
	}
}

func reverseEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(reversalReq)
		if err := req.validate(); err != nil {
			return reverseResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.Reverse(req.ReverseReq)
		if err != nil {
			return reverseResp{}, err
		}

		return reverseResp{resp}, nil
	}
}

func transactionStatusEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transactionReq)
		if err := req.validate(); err != nil {
			return transactionStatusResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.TransactionStatus(req.TransactionStatusReq)
		if err != nil {
			return transactionStatusResp{}, err
		}

		return transactionStatusResp{resp}, nil
	}
}

func remitTaxEndpoint(svc grpc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(remitTaxReq)
		if err := req.validate(); err != nil {
			return remitTaxResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.RemitTax(req.RemitTaxReq)
		if err != nil {
			return remitTaxResp{}, err
		}

		return remitTaxResp{resp}, nil
	}
}
