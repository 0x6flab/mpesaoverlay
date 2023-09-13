package grpc

import (
	"context"
	"errors"

	"github.com/0x6flab/mpesaoverlay/overlay"
	"github.com/go-kit/kit/endpoint"
)

// errValidation is returned when a request validation has failed.
var errValidation = errors.New("validation error")

func getTokenEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getTokenReq)
		if err := req.validate(); err != nil {
			return getTokenResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.GetToken(ctx)
		if err != nil {
			return getTokenResp{}, err
		}

		return getTokenResp{resp}, nil
	}
}

func expressQueryEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(expressQueryReq)
		if err := req.validate(); err != nil {
			return expressQueryResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.ExpressQuery(ctx, req.ExpressQueryReq)
		if err != nil {
			return expressQueryResp{}, err
		}

		return expressQueryResp{resp}, nil
	}
}

func expressSimulateEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(expressSimulateReq)
		if err := req.validate(); err != nil {
			return expressSimulateResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.ExpressSimulate(ctx, req.ExpressSimulateReq)
		if err != nil {
			return expressSimulateResp{}, err
		}

		return expressSimulateResp{resp}, nil
	}
}

func b2cEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(b2cReq)
		if err := req.validate(); err != nil {
			return b2cResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.B2CPayment(ctx, req.B2Creq)
		if err != nil {
			return b2cResp{}, err
		}

		return b2cResp{resp}, nil
	}
}

func accountBalanceEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(accountBalanceReq)
		if err := req.validate(); err != nil {
			return accountBalanceResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.AccountBalance(ctx, req.AccBalanceReq)
		if err != nil {
			return accountBalanceResp{}, err
		}

		return accountBalanceResp{resp}, nil
	}
}

func c2bRegisterURLEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(c2bRegisterURLReq)
		if err := req.validate(); err != nil {
			return c2bRegisterURLResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.C2BRegisterURL(ctx, req.C2BRegisterURLReq)
		if err != nil {
			return c2bRegisterURLResp{}, err
		}

		return c2bRegisterURLResp{resp}, nil
	}
}

func c2bSimulateEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(c2bSimulateReq)
		if err := req.validate(); err != nil {
			return c2bSimulateResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.C2BSimulate(ctx, req.C2BSimulateReq)
		if err != nil {
			return c2bSimulateResp{}, err
		}

		return c2bSimulateResp{resp}, nil
	}
}

func generateQREndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(generateQRReq)
		if err := req.validate(); err != nil {
			return generateQRResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.GenerateQR(ctx, req.QRReq)
		if err != nil {
			return generateQRResp{}, err
		}

		return generateQRResp{resp}, nil
	}
}

func reverseEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(reversalReq)
		if err := req.validate(); err != nil {
			return reverseResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.Reverse(ctx, req.ReversalReq)
		if err != nil {
			return reverseResp{}, err
		}

		return reverseResp{resp}, nil
	}
}

func transactionStatusEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transactionReq)
		if err := req.validate(); err != nil {
			return transactionStatusResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.TransactionStatus(ctx, req.TransactionReq)
		if err != nil {
			return transactionStatusResp{}, err
		}

		return transactionStatusResp{resp}, nil
	}
}

func remitTaxEndpoint(svc overlay.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(remitTaxReq)
		if err := req.validate(); err != nil {
			return remitTaxResp{}, errors.Join(errValidation, err)
		}

		resp, err := svc.RemitTax(ctx, req.RemitTax)
		if err != nil {
			return remitTaxResp{}, err
		}

		return remitTaxResp{resp}, nil
	}
}
