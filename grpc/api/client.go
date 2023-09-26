// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"time"

	grpcadapter "github.com/0x6flab/mpesaoverlay/grpc"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

const svcName = "mpesaoverlay.overlay.Service"

var _ grpcadapter.ServiceClient = (*grpcClient)(nil)

// grpcClient implements the gRPC ServiceClient interface.
type grpcClient struct {
	token             endpoint.Endpoint
	expressQuery      endpoint.Endpoint
	expressSimulate   endpoint.Endpoint
	b2c               endpoint.Endpoint
	accountBalance    endpoint.Endpoint
	c2bRegisterURL    endpoint.Endpoint
	c2bSimulate       endpoint.Endpoint
	generateQR        endpoint.Endpoint
	reverse           endpoint.Endpoint
	transactionStatus endpoint.Endpoint
	remitTax          endpoint.Endpoint
	timeout           time.Duration
}

// NewClient returns new gRPC client instance.
// The client is responsible for communicating with the mpesaoverlay service.
func NewClient(conn *grpc.ClientConn, timeout time.Duration) grpcadapter.ServiceClient {
	return &grpcClient{
		token: kitgrpc.NewClient(
			conn,
			svcName,
			"GetToken",
			encodeTokenRequest,
			decodeTokenResponse,
			grpcadapter.TokenResp{},
		).Endpoint(),
		expressQuery: kitgrpc.NewClient(
			conn,
			svcName,
			"ExpressQuery",
			encodeExpressQueryRequest,
			decodeExpressQueryResponse,
			grpcadapter.ExpressQueryResp{},
		).Endpoint(),
		expressSimulate: kitgrpc.NewClient(
			conn,
			svcName,
			"ExpressSimulate",
			encodeExpressSimulateRequest,
			decodeExpressSimulateResponse,
			grpcadapter.ExpressSimulateResp{},
		).Endpoint(),
		b2c: kitgrpc.NewClient(
			conn,
			svcName,
			"B2C",
			encodeB2CRequest,
			decodeB2CResponse,
			grpcadapter.B2CPaymentResp{},
		).Endpoint(),
		accountBalance: kitgrpc.NewClient(
			conn,
			svcName,
			"AccountBalance",
			encodeAccountBalanceRequest,
			decodeAccountBalanceResponse,
			grpcadapter.AccountBalanceResp{},
		).Endpoint(),
		c2bRegisterURL: kitgrpc.NewClient(
			conn,
			svcName,
			"C2BRegisterURL",
			encodeC2BRegisterURLRequest,
			decodeC2BRegisterURLResponse,
			grpcadapter.C2BRegisterURLResp{},
		).Endpoint(),
		c2bSimulate: kitgrpc.NewClient(
			conn,
			svcName,
			"C2BSimulate",
			encodeC2BSimulateRequest,
			decodeC2BSimulateResponse,
			grpcadapter.C2BSimulateResp{},
		).Endpoint(),
		generateQR: kitgrpc.NewClient(
			conn,
			svcName,
			"GenerateQR",
			encodeGenerateQRRequest,
			decodeGenerateQRResponse,
			grpcadapter.GenerateQRResp{},
		).Endpoint(),
		reverse: kitgrpc.NewClient(
			conn,
			svcName,
			"Reverse",
			encodeReverseRequest,
			decodeReverseResponse,
			grpcadapter.ReverseResp{},
		).Endpoint(),
		transactionStatus: kitgrpc.NewClient(
			conn,
			svcName,
			"TransactionStatus",
			encodeTransactionStatusRequest,
			decodeTransactionStatusResponse,
			grpcadapter.TransactionStatusResp{},
		).Endpoint(),
		remitTax: kitgrpc.NewClient(
			conn,
			svcName,
			"RemitTax",
			encodeRemitTaxRequest,
			decodeRemitTaxResponse,
			grpcadapter.RemitTaxResp{},
		).Endpoint(),

		timeout: timeout,
	}
}

func (client grpcClient) Token(ctx context.Context, _ *grpcadapter.Empty, _ ...grpc.CallOption) (r *grpcadapter.TokenResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	res, err := client.token(ctx, tokenReq{})
	if err != nil {
		return &grpcadapter.TokenResp{}, err
	}

	ares := res.(tokenResp)

	return &grpcadapter.TokenResp{
		AccessToken: ares.AccessToken,
		Expiry:      ares.Expiry,
	}, err
}

func decodeTokenResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.TokenResp)

	return grpcadapter.TokenResp{
		AccessToken: res.GetAccessToken(),
		Expiry:      res.GetExpiry(),
	}, nil
}

func encodeTokenRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return &grpcadapter.Empty{}, nil
}

func (client grpcClient) ExpressQuery(ctx context.Context, req *grpcadapter.ExpressQueryReq, _ ...grpc.CallOption) (r *grpcadapter.ExpressQueryResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	expressQueryReq := expressQueryReq{
		ExpressQueryReq: mpesa.ExpressQueryReq{
			PassKey:           req.PassKey,
			BusinessShortCode: req.BusinessShortCode,
			Password:          req.Password,
			Timestamp:         req.Timestamp,
			CheckoutRequestID: req.CheckoutRequestID,
		},
	}
	res, err := client.expressQuery(ctx, expressQueryReq)
	if err != nil {
		return &grpcadapter.ExpressQueryResp{}, err
	}

	ares := res.(expressQueryResp)

	return &grpcadapter.ExpressQueryResp{
		ResponseCode:        ares.ResponseCode,
		ResponseDescription: ares.ResponseDescription,
		MerchantRequestID:   ares.MerchantRequestID,
		CheckoutRequestID:   ares.CheckoutRequestID,
		CustomerMessage:     ares.CustomerMessage,
		ResultCode:          ares.ResultCode,
		ResultDesc:          ares.ResultDesc,
	}, err
}

func decodeExpressQueryResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.ExpressQueryResp)

	return expressQueryResp{
		ExpressQueryResp: mpesa.ExpressQueryResp{
			ResponseCode:        res.GetResponseCode(),
			ResponseDescription: res.GetResponseDescription(),
			MerchantRequestID:   res.GetMerchantRequestID(),
			CheckoutRequestID:   res.GetCheckoutRequestID(),
			CustomerMessage:     res.GetCustomerMessage(),
			ResultCode:          res.GetResultCode(),
			ResultDesc:          res.GetResultDesc(),
		},
	}, nil
}

func encodeExpressQueryRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(expressQueryReq)

	return &mpesa.ExpressQueryReq{
		PassKey:           req.PassKey,
		BusinessShortCode: req.BusinessShortCode,
		Password:          req.Password,
		Timestamp:         req.Timestamp,
		CheckoutRequestID: req.CheckoutRequestID,
	}, nil
}

func (client grpcClient) ExpressSimulate(ctx context.Context, req *grpcadapter.ExpressSimulateReq, _ ...grpc.CallOption) (r *grpcadapter.ExpressSimulateResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	expressSimulateReq := expressSimulateReq{
		ExpressSimulateReq: mpesa.ExpressSimulateReq{
			PassKey:           req.PassKey,
			BusinessShortCode: req.BusinessShortCode,
			Password:          req.Password,
			Timestamp:         req.Timestamp,
			Amount:            req.Amount,
			PartyA:            req.PartyA,
			PartyB:            req.PartyB,
			PhoneNumber:       req.PhoneNumber,
			CallBackURL:       req.CallBackURL,
			AccountReference:  req.AccountReference,
			TransactionDesc:   req.TransactionDesc,
		},
	}
	res, err := client.expressSimulate(ctx, expressSimulateReq)
	if err != nil {
		return &grpcadapter.ExpressSimulateResp{}, err
	}

	ares := res.(expressSimulateResp)

	return &grpcadapter.ExpressSimulateResp{
		MerchantRequestID:   ares.MerchantRequestID,
		CheckoutRequestID:   ares.CheckoutRequestID,
		ResponseCode:        ares.ResponseCode,
		ResponseDescription: ares.ResponseDescription,
		CustomerMessage:     ares.CustomerMessage,
	}, err
}

func decodeExpressSimulateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.ExpressSimulateResp)

	return expressSimulateResp{
		ExpressSimulateResp: mpesa.ExpressSimulateResp{
			MerchantRequestID:   res.GetMerchantRequestID(),
			CheckoutRequestID:   res.GetCheckoutRequestID(),
			ResponseCode:        res.GetResponseCode(),
			ResponseDescription: res.GetResponseDescription(),
			CustomerMessage:     res.GetCustomerMessage(),
		},
	}, nil
}

func encodeExpressSimulateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(expressSimulateReq)

	return &mpesa.ExpressSimulateReq{
		PassKey:           req.PassKey,
		BusinessShortCode: req.BusinessShortCode,
		Password:          req.Password,
		Timestamp:         req.Timestamp,
		Amount:            req.Amount,
		PartyA:            req.PartyA,
		PartyB:            req.PartyB,
		PhoneNumber:       req.PhoneNumber,
		CallBackURL:       req.CallBackURL,
		AccountReference:  req.AccountReference,
		TransactionDesc:   req.TransactionDesc,
	}, nil
}

func (client grpcClient) B2CPayment(ctx context.Context, req *grpcadapter.B2CPaymentReq, _ ...grpc.CallOption) (r *grpcadapter.B2CPaymentResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	b2cReq := b2cReq{
		B2CPaymentReq: mpesa.B2CPaymentReq{
			InitiatorName:      req.InitiatorName,
			SecurityCredential: req.SecurityCredential,
			CommandID:          req.CommandID,
			Amount:             req.Amount,
			PartyA:             req.PartyA,
			PartyB:             req.PartyB,
			Remarks:            req.Remarks,
			QueueTimeOutURL:    req.QueueTimeOutURL,
			ResultURL:          req.ResultURL,
			Occasion:           req.Occasion,
		},
	}
	res, err := client.b2c(ctx, b2cReq)
	if err != nil {
		return &grpcadapter.B2CPaymentResp{}, err
	}

	ares := res.(b2cResp)

	return &grpcadapter.B2CPaymentResp{
		ValidResp: &grpcadapter.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeB2CResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.B2CPaymentResp)

	return b2cResp{
		B2CPaymentResp: mpesa.B2CPaymentResp{
			ValidResp: mpesa.ValidResp{
				OriginatorConversationID: res.ValidResp.GetOriginatorConversationID(),
				ResponseCode:             res.ValidResp.GetResponseCode(),
				ResponseDescription:      res.ValidResp.GetResponseDescription(),
				ConversationID:           res.ValidResp.GetConversationID(),
			},
		},
	}, nil
}

func encodeB2CRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(b2cReq)

	return &mpesa.B2CPaymentReq{
		InitiatorName:      req.InitiatorName,
		SecurityCredential: req.SecurityCredential,
		CommandID:          req.CommandID,
		Amount:             req.Amount,
		PartyA:             req.PartyA,
		PartyB:             req.PartyB,
		Remarks:            req.Remarks,
		QueueTimeOutURL:    req.QueueTimeOutURL,
		ResultURL:          req.ResultURL,
		Occasion:           req.Occasion,
	}, nil
}

func (client grpcClient) AccountBalance(ctx context.Context, req *grpcadapter.AccountBalanceReq, _ ...grpc.CallOption) (r *grpcadapter.AccountBalanceResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	accountBalanceReq := accountBalanceReq{
		AccountBalanceReq: mpesa.AccountBalanceReq{
			InitiatorName:      req.InitiatorName,
			SecurityCredential: req.SecurityCredential,
			CommandID:          req.CommandID,
			PartyA:             req.PartyA,
			IdentifierType:     uint8(req.IdentifierType),
			Remarks:            req.Remarks,
			QueueTimeOutURL:    req.QueueTimeOutURL,
			ResultURL:          req.ResultURL,
		},
	}
	res, err := client.accountBalance(ctx, accountBalanceReq)
	if err != nil {
		return &grpcadapter.AccountBalanceResp{}, err
	}

	ares := res.(accountBalanceResp)

	return &grpcadapter.AccountBalanceResp{
		ValidResp: &grpcadapter.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeAccountBalanceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.AccountBalanceResp)

	return accountBalanceResp{
		AccountBalanceResp: mpesa.AccountBalanceResp{
			ValidResp: mpesa.ValidResp{
				OriginatorConversationID: res.ValidResp.GetOriginatorConversationID(),
				ResponseCode:             res.ValidResp.GetResponseCode(),
				ResponseDescription:      res.ValidResp.GetResponseDescription(),
				ConversationID:           res.ValidResp.GetConversationID(),
			},
		},
	}, nil
}

func encodeAccountBalanceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(accountBalanceReq)

	return &mpesa.AccountBalanceReq{
		InitiatorName:      req.InitiatorName,
		SecurityCredential: req.SecurityCredential,
		CommandID:          req.CommandID,
		PartyA:             req.PartyA,
		IdentifierType:     req.IdentifierType,
		Remarks:            req.Remarks,
		QueueTimeOutURL:    req.QueueTimeOutURL,
		ResultURL:          req.ResultURL,
	}, nil
}

func (client grpcClient) C2BRegisterURL(ctx context.Context, req *grpcadapter.C2BRegisterURLReq, _ ...grpc.CallOption) (r *grpcadapter.C2BRegisterURLResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	c2bRegisterURLReq := c2bRegisterURLReq{
		C2BRegisterURLReq: mpesa.C2BRegisterURLReq{
			ShortCode:       req.ShortCode,
			ResponseType:    req.ResponseType,
			ConfirmationURL: req.ConfirmationURL,
			ValidationURL:   req.ValidationURL,
		},
	}
	res, err := client.c2bRegisterURL(ctx, c2bRegisterURLReq)
	if err != nil {
		return &grpcadapter.C2BRegisterURLResp{}, err
	}

	ares := res.(c2bRegisterURLResp)

	return &grpcadapter.C2BRegisterURLResp{
		ValidResp: &grpcadapter.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeC2BRegisterURLResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.C2BRegisterURLResp)

	return c2bRegisterURLResp{
		C2BRegisterURLResp: mpesa.C2BRegisterURLResp{
			ValidResp: mpesa.ValidResp{
				OriginatorConversationID: res.ValidResp.GetOriginatorConversationID(),
				ResponseCode:             res.ValidResp.GetResponseCode(),
				ResponseDescription:      res.ValidResp.GetResponseDescription(),
				ConversationID:           res.ValidResp.GetConversationID(),
			},
		},
	}, nil
}

func encodeC2BRegisterURLRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(c2bRegisterURLReq)

	return &mpesa.C2BRegisterURLReq{
		ShortCode:       req.ShortCode,
		ResponseType:    req.ResponseType,
		ConfirmationURL: req.ConfirmationURL,
		ValidationURL:   req.ValidationURL,
	}, nil
}

func (client grpcClient) C2BSimulate(ctx context.Context, req *grpcadapter.C2BSimulateReq, _ ...grpc.CallOption) (r *grpcadapter.C2BSimulateResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	c2bSimulateReq := c2bSimulateReq{
		C2BSimulateReq: mpesa.C2BSimulateReq{
			ShortCode:     req.ShortCode,
			CommandID:     req.CommandID,
			Amount:        req.Amount,
			Msisdn:        req.Msisdn,
			BillRefNumber: req.BillRefNumber,
		},
	}
	res, err := client.c2bSimulate(ctx, c2bSimulateReq)
	if err != nil {
		return &grpcadapter.C2BSimulateResp{}, err
	}

	ares := res.(c2bSimulateResp)

	return &grpcadapter.C2BSimulateResp{
		ValidResp: &grpcadapter.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeC2BSimulateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.C2BSimulateResp)

	return c2bSimulateResp{
		C2BSimulateResp: mpesa.C2BSimulateResp{
			ValidResp: mpesa.ValidResp{
				OriginatorConversationID: res.ValidResp.GetOriginatorConversationID(),
				ResponseCode:             res.ValidResp.GetResponseCode(),
				ResponseDescription:      res.ValidResp.GetResponseDescription(),
				ConversationID:           res.ValidResp.GetConversationID(),
			},
		},
	}, nil
}

func encodeC2BSimulateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(c2bSimulateReq)

	return &mpesa.C2BSimulateReq{
		ShortCode:     req.ShortCode,
		CommandID:     req.CommandID,
		Amount:        req.Amount,
		Msisdn:        req.Msisdn,
		BillRefNumber: req.BillRefNumber,
	}, nil
}

func (client grpcClient) GenerateQR(ctx context.Context, req *grpcadapter.GenerateQRReq, _ ...grpc.CallOption) (r *grpcadapter.GenerateQRResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	generateQRReq := generateQRReq{
		GenerateQRReq: mpesa.GenerateQRReq{
			MerchantName: req.MerchantName,
			RefNo:        req.RefNo,
			Amount:       req.Amount,
			TrxCode:      req.TrxCode,
			CPI:          req.CPI,
			Size:         req.Size,
		},
	}
	res, err := client.generateQR(ctx, generateQRReq)
	if err != nil {
		return &grpcadapter.GenerateQRResp{}, err
	}

	ares := res.(generateQRResp)

	return &grpcadapter.GenerateQRResp{
		ResponseDescription: ares.ResponseDescription,
		ResponseCode:        ares.ResponseCode,
		RequestID:           ares.RequestID,
		QRCode:              ares.QRCode,
	}, err
}

func decodeGenerateQRResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.GenerateQRResp)

	return generateQRResp{
		GenerateQRResp: mpesa.GenerateQRResp{
			ResponseDescription: res.GetResponseDescription(),
			ResponseCode:        res.GetResponseCode(),
			RequestID:           res.GetRequestID(),
			QRCode:              res.GetQRCode(),
		},
	}, nil
}

func encodeGenerateQRRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(generateQRReq)

	return &mpesa.GenerateQRReq{
		MerchantName: req.MerchantName,
		RefNo:        req.RefNo,
		Amount:       req.Amount,
		TrxCode:      req.TrxCode,
		CPI:          req.CPI,
		Size:         req.Size,
	}, nil
}

func (client grpcClient) Reverse(ctx context.Context, req *grpcadapter.ReverseReq, _ ...grpc.CallOption) (r *grpcadapter.ReverseResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	reversalReq := reversalReq{
		ReverseReq: mpesa.ReverseReq{
			InitiatorName:          req.InitiatorName,
			SecurityCredential:     req.SecurityCredential,
			CommandID:              req.CommandID,
			TransactionID:          req.TransactionID,
			Amount:                 req.Amount,
			ReceiverParty:          req.ReceiverParty,
			RecieverIdentifierType: uint8(req.RecieverIdentifierType),
			ResultURL:              req.ResultURL,
			QueueTimeOutURL:        req.QueueTimeOutURL,
			Remarks:                req.Remarks,
			Occasion:               req.Occasion,
		},
	}
	res, err := client.reverse(ctx, reversalReq)
	if err != nil {
		return &grpcadapter.ReverseResp{}, err
	}

	ares := res.(reverseResp)

	return &grpcadapter.ReverseResp{
		ValidResp: &grpcadapter.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeReverseResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.ReverseResp)

	return reverseResp{
		ReverseResp: mpesa.ReverseResp{
			ValidResp: mpesa.ValidResp{
				OriginatorConversationID: res.ValidResp.GetOriginatorConversationID(),
				ResponseCode:             res.ValidResp.GetResponseCode(),
				ResponseDescription:      res.ValidResp.GetResponseDescription(),
				ConversationID:           res.ValidResp.GetConversationID(),
			},
		},
	}, nil
}

func encodeReverseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(reversalReq)

	return &mpesa.ReverseReq{
		InitiatorName:          req.InitiatorName,
		SecurityCredential:     req.SecurityCredential,
		CommandID:              req.CommandID,
		TransactionID:          req.TransactionID,
		Amount:                 req.Amount,
		ReceiverParty:          req.ReceiverParty,
		RecieverIdentifierType: req.RecieverIdentifierType,
		ResultURL:              req.ResultURL,
		QueueTimeOutURL:        req.QueueTimeOutURL,
		Remarks:                req.Remarks,
		Occasion:               req.Occasion,
	}, nil
}

func (client grpcClient) TransactionStatus(ctx context.Context, req *grpcadapter.TransactionStatusReq, _ ...grpc.CallOption) (r *grpcadapter.TransactionStatusResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	transactionReq := transactionReq{
		TransactionStatusReq: mpesa.TransactionStatusReq{
			InitiatorName:      req.InitiatorName,
			SecurityCredential: req.SecurityCredential,
			CommandID:          req.CommandID,
			TransactionID:      req.TransactionID,
			PartyA:             req.PartyA,
			IdentifierType:     uint8(req.IdentifierType),
			ResultURL:          req.ResultURL,
			QueueTimeOutURL:    req.QueueTimeOutURL,
			Remarks:            req.Remarks,
			Occasion:           req.Occasion,
		},
	}
	res, err := client.transactionStatus(ctx, transactionReq)
	if err != nil {
		return &grpcadapter.TransactionStatusResp{}, err
	}

	ares := res.(transactionStatusResp)

	return &grpcadapter.TransactionStatusResp{
		ValidResp: &grpcadapter.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeTransactionStatusResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.TransactionStatusResp)

	return transactionStatusResp{
		TransactionStatusResp: mpesa.TransactionStatusResp{
			ValidResp: mpesa.ValidResp{
				OriginatorConversationID: res.ValidResp.GetOriginatorConversationID(),
				ResponseCode:             res.ValidResp.GetResponseCode(),
				ResponseDescription:      res.ValidResp.GetResponseDescription(),
				ConversationID:           res.ValidResp.GetConversationID(),
			},
		},
	}, nil
}

func encodeTransactionStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(transactionReq)

	return &mpesa.TransactionStatusReq{
		InitiatorName:      req.InitiatorName,
		SecurityCredential: req.SecurityCredential,
		CommandID:          req.CommandID,
		TransactionID:      req.TransactionID,
		PartyA:             req.PartyA,
		IdentifierType:     req.IdentifierType,
		ResultURL:          req.ResultURL,
		QueueTimeOutURL:    req.QueueTimeOutURL,
		Remarks:            req.Remarks,
		Occasion:           req.Occasion,
	}, nil
}

func (client grpcClient) RemitTax(ctx context.Context, req *grpcadapter.RemitTaxReq, _ ...grpc.CallOption) (r *grpcadapter.RemitTaxResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	remitTaxReq := remitTaxReq{
		RemitTaxReq: mpesa.RemitTaxReq{
			InitiatorName:      req.InitiatorName,
			SecurityCredential: req.SecurityCredential,
			CommandID:          req.CommandID,
			Amount:             req.Amount,
			PartyA:             req.PartyA,
			PartyB:             req.PartyB,
			Remarks:            req.Remarks,
			QueueTimeOutURL:    req.QueueTimeOutURL,
			ResultURL:          req.ResultURL,
		},
	}
	res, err := client.remitTax(ctx, remitTaxReq)
	if err != nil {
		return &grpcadapter.RemitTaxResp{}, err
	}

	ares := res.(remitTaxResp)

	return &grpcadapter.RemitTaxResp{
		ValidResp: &grpcadapter.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeRemitTaxResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*grpcadapter.RemitTaxResp)

	return remitTaxResp{
		RemitTaxResp: mpesa.RemitTaxResp{
			ValidResp: mpesa.ValidResp{
				OriginatorConversationID: res.ValidResp.GetOriginatorConversationID(),
				ResponseCode:             res.ValidResp.GetResponseCode(),
				ResponseDescription:      res.ValidResp.GetResponseDescription(),
				ConversationID:           res.ValidResp.GetConversationID(),
			},
		},
	}, nil
}

func encodeRemitTaxRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(remitTaxReq)

	return &mpesa.RemitTaxReq{
		InitiatorName:      req.InitiatorName,
		SecurityCredential: req.SecurityCredential,
		CommandID:          req.CommandID,
		Amount:             req.Amount,
		PartyA:             req.PartyA,
		PartyB:             req.PartyB,
		Remarks:            req.Remarks,
		QueueTimeOutURL:    req.QueueTimeOutURL,
		ResultURL:          req.ResultURL,
	}, nil
}
