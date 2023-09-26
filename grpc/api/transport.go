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
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ grpc.ServiceServer = (*grpcServer)(nil)

// grpcServer implements the gRPC ServiceServer interface.
type grpcServer struct {
	token             kitgrpc.Handler
	expressQuery      kitgrpc.Handler
	expressSimulate   kitgrpc.Handler
	b2c               kitgrpc.Handler
	accountBalance    kitgrpc.Handler
	c2bRegisterURL    kitgrpc.Handler
	c2bSimulate       kitgrpc.Handler
	generateQR        kitgrpc.Handler
	reverse           kitgrpc.Handler
	transactionStatus kitgrpc.Handler
	remitTax          kitgrpc.Handler
	grpc.UnimplementedServiceServer
}

// NewServer returns a new instance of the grpc server.
// The grpc server is responsible for the grpc api.
func NewServer(svc grpc.Service) grpc.ServiceServer {
	return &grpcServer{
		token: kitgrpc.NewServer(
			tokenEndpoint(svc),
			decodeTokenRequest,
			encodeTokenResponse,
		),
		expressQuery: kitgrpc.NewServer(
			expressQueryEndpoint(svc),
			decodeExpressQueryRequest,
			encodeExpressQueryResponse,
		),
		expressSimulate: kitgrpc.NewServer(
			expressSimulateEndpoint(svc),
			decodeExpressSimulateRequest,
			encodeExpressSimulateResponse,
		),
		b2c: kitgrpc.NewServer(
			b2cEndpoint(svc),
			decodeB2CRequest,
			encodeB2CResponse,
		),
		accountBalance: kitgrpc.NewServer(
			accountBalanceEndpoint(svc),
			decodeAccountBalanceRequest,
			encodeAccountBalanceResponse,
		),
		c2bRegisterURL: kitgrpc.NewServer(
			c2bRegisterURLEndpoint(svc),
			decodeC2BRegisterURLRequest,
			encodeC2BRegisterURLResponse,
		),
		c2bSimulate: kitgrpc.NewServer(
			c2bSimulateEndpoint(svc),
			decodeC2BSimulateRequest,
			encodeC2BSimulateResponse,
		),
		generateQR: kitgrpc.NewServer(
			generateQREndpoint(svc),
			decodeGenerateQRRequest,
			encodeGenerateQRResponse,
		),
		reverse: kitgrpc.NewServer(
			reverseEndpoint(svc),
			decodeReverseRequest,
			encodeReverseResponse,
		),
		transactionStatus: kitgrpc.NewServer(
			transactionStatusEndpoint(svc),
			decodeTransactionStatusRequest,
			encodeTransactionStatusResponse,
		),
		remitTax: kitgrpc.NewServer(
			remitTaxEndpoint(svc),
			decodeRemitTaxRequest,
			encodeRemitTaxResponse,
		),
	}
}

func (s *grpcServer) Token(ctx context.Context, req *grpc.Empty) (*grpc.TokenResp, error) {
	_, res, err := s.token.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.TokenResp), nil
}

func decodeTokenRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return tokenReq{}, nil
}

func encodeTokenResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(tokenResp)

	return &grpc.TokenResp{AccessToken: res.AccessToken, Expiry: res.Expiry}, nil
}

func (s *grpcServer) ExpressQuery(ctx context.Context, req *grpc.ExpressQueryReq) (*grpc.ExpressQueryResp, error) {
	_, res, err := s.expressQuery.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.ExpressQueryResp), nil
}

func decodeExpressQueryRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.ExpressQueryReq)

	return expressQueryReq{
		ExpressQueryReq: mpesa.ExpressQueryReq{
			PassKey:           req.PassKey,
			BusinessShortCode: req.BusinessShortCode,
			Password:          req.Password,
			Timestamp:         req.Timestamp,
			CheckoutRequestID: req.CheckoutRequestID,
		},
	}, nil
}

func encodeExpressQueryResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(expressQueryResp)

	return &grpc.ExpressQueryResp{
		MerchantRequestID:   res.MerchantRequestID,
		CheckoutRequestID:   res.CheckoutRequestID,
		ResponseCode:        res.ResponseCode,
		ResponseDescription: res.ResponseDescription,
		CustomerMessage:     res.CustomerMessage,
		ResultCode:          res.ResultCode,
		ResultDesc:          res.ResultDesc,
	}, nil
}

func (s *grpcServer) ExpressSimulate(ctx context.Context, req *grpc.ExpressSimulateReq) (*grpc.ExpressSimulateResp, error) {
	_, res, err := s.expressSimulate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.ExpressSimulateResp), nil
}

func decodeExpressSimulateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.ExpressSimulateReq)

	return expressSimulateReq{
		ExpressSimulateReq: mpesa.ExpressSimulateReq{
			PassKey:           req.PassKey,
			BusinessShortCode: req.BusinessShortCode,
			Password:          req.Password,
			Timestamp:         req.Timestamp,
			TransactionType:   req.TransactionType,
			PhoneNumber:       req.PhoneNumber,
			Amount:            req.Amount,
			PartyA:            req.PartyA,
			PartyB:            req.PartyB,
			AccountReference:  req.AccountReference,
			TransactionDesc:   req.TransactionDesc,
			CallBackURL:       req.CallBackURL,
		},
	}, nil
}

func encodeExpressSimulateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(expressSimulateResp)

	return &grpc.ExpressSimulateResp{
		MerchantRequestID:   res.MerchantRequestID,
		CheckoutRequestID:   res.CheckoutRequestID,
		ResponseCode:        res.ResponseCode,
		ResponseDescription: res.ResponseDescription,
		CustomerMessage:     res.CustomerMessage,
	}, nil
}

func (s *grpcServer) B2CPayment(ctx context.Context, req *grpc.B2CPaymentReq) (*grpc.B2CPaymentResp, error) {
	_, res, err := s.b2c.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.B2CPaymentResp), nil
}

func decodeB2CRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.B2CPaymentReq)

	return b2cReq{B2CPaymentReq: mpesa.B2CPaymentReq{
		InitiatorName:      req.InitiatorName,
		InitiatorPassword:  req.InitiatorPassword,
		SecurityCredential: req.SecurityCredential,
		CommandID:          req.CommandID,
		Amount:             req.Amount,
		PartyA:             req.PartyA,
		PartyB:             req.PartyB,
		Remarks:            req.Remarks,
		QueueTimeOutURL:    req.QueueTimeOutURL,
		ResultURL:          req.ResultURL,
		Occasion:           req.Occasion,
	}}, nil
}

func encodeB2CResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(b2cResp)

	return &grpc.B2CPaymentResp{
		ValidResp: &grpc.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) AccountBalance(ctx context.Context, req *grpc.AccountBalanceReq) (*grpc.AccountBalanceResp, error) {
	_, res, err := s.accountBalance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.AccountBalanceResp), nil
}

func decodeAccountBalanceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.AccountBalanceReq)

	return accountBalanceReq{AccountBalanceReq: mpesa.AccountBalanceReq{
		InitiatorName:      req.InitiatorName,
		InitiatorPassword:  req.InitiatorPassword,
		SecurityCredential: req.SecurityCredential,
		CommandID:          req.CommandID,
		PartyA:             req.PartyA,
		IdentifierType:     uint8(req.IdentifierType),
		Remarks:            req.Remarks,
		QueueTimeOutURL:    req.QueueTimeOutURL,
		ResultURL:          req.ResultURL,
	}}, nil
}

func encodeAccountBalanceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(accountBalanceResp)

	return &grpc.AccountBalanceResp{
		ValidResp: &grpc.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) C2BRegisterURL(ctx context.Context, req *grpc.C2BRegisterURLReq) (*grpc.C2BRegisterURLResp, error) {
	_, res, err := s.c2bRegisterURL.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.C2BRegisterURLResp), nil
}

func decodeC2BRegisterURLRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.C2BRegisterURLReq)

	return c2bRegisterURLReq{C2BRegisterURLReq: mpesa.C2BRegisterURLReq{
		ShortCode:       req.ShortCode,
		ResponseType:    req.ResponseType,
		ConfirmationURL: req.ConfirmationURL,
		ValidationURL:   req.ValidationURL,
	}}, nil
}

func encodeC2BRegisterURLResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(c2bRegisterURLResp)

	return &grpc.C2BRegisterURLResp{
		ValidResp: &grpc.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) C2BSimulate(ctx context.Context, req *grpc.C2BSimulateReq) (*grpc.C2BSimulateResp, error) {
	_, res, err := s.c2bSimulate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.C2BSimulateResp), nil
}

func decodeC2BSimulateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.C2BSimulateReq)

	return c2bSimulateReq{C2BSimulateReq: mpesa.C2BSimulateReq{
		ShortCode:     req.ShortCode,
		CommandID:     req.CommandID,
		Amount:        req.Amount,
		Msisdn:        req.Msisdn,
		BillRefNumber: req.BillRefNumber,
	}}, nil
}

func encodeC2BSimulateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(c2bSimulateResp)

	return &grpc.C2BSimulateResp{
		ValidResp: &grpc.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) GenerateQR(ctx context.Context, req *grpc.GenerateQRReq) (*grpc.GenerateQRResp, error) {
	_, res, err := s.generateQR.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.GenerateQRResp), nil
}

func decodeGenerateQRRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.GenerateQRReq)

	return generateQRReq{GenerateQRReq: mpesa.GenerateQRReq{
		MerchantName: req.MerchantName,
		RefNo:        req.RefNo,
		Amount:       req.Amount,
		TrxCode:      req.TrxCode,
		CPI:          req.CPI,
		Size:         req.Size,
	}}, nil
}

func encodeGenerateQRResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(generateQRResp)

	return &grpc.GenerateQRResp{
		RequestID:           res.RequestID,
		QRCode:              res.QRCode,
		ResponseCode:        res.ResponseCode,
		ResponseDescription: res.ResponseDescription,
	}, nil
}

func (s *grpcServer) Reverse(ctx context.Context, req *grpc.ReverseReq) (*grpc.ReverseResp, error) {
	_, res, err := s.reverse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.ReverseResp), nil
}

func decodeReverseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.ReverseReq)

	return reversalReq{ReverseReq: mpesa.ReverseReq{
		CommandID:              req.CommandID,
		ReceiverParty:          req.ReceiverParty,
		RecieverIdentifierType: uint8(req.RecieverIdentifierType),
		Remarks:                req.Remarks,
		InitiatorName:          req.InitiatorName,
		InitiatorPassword:      req.InitiatorPassword,
		SecurityCredential:     req.SecurityCredential,
		QueueTimeOutURL:        req.QueueTimeOutURL,
		ResultURL:              req.ResultURL,
		TransactionID:          req.TransactionID,
		Occasion:               req.Occasion,
		Amount:                 req.Amount,
	}}, nil
}

func encodeReverseResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(reverseResp)

	return &grpc.ReverseResp{
		ValidResp: &grpc.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) TransactionStatus(ctx context.Context, req *grpc.TransactionStatusReq) (*grpc.TransactionStatusResp, error) {
	_, res, err := s.transactionStatus.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.TransactionStatusResp), nil
}

func decodeTransactionStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.TransactionStatusReq)

	return transactionReq{TransactionStatusReq: mpesa.TransactionStatusReq{
		CommandID:          req.CommandID,
		PartyA:             req.PartyA,
		IdentifierType:     uint8(req.IdentifierType),
		Remarks:            req.Remarks,
		InitiatorName:      req.InitiatorName,
		InitiatorPassword:  req.InitiatorPassword,
		SecurityCredential: req.SecurityCredential,
		QueueTimeOutURL:    req.QueueTimeOutURL,
		ResultURL:          req.ResultURL,
		TransactionID:      req.TransactionID,
		Occasion:           req.Occasion,
	}}, nil
}

func encodeTransactionStatusResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(transactionStatusResp)

	return &grpc.TransactionStatusResp{
		ValidResp: &grpc.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) RemitTax(ctx context.Context, req *grpc.RemitTaxReq) (*grpc.RemitTaxResp, error) {
	_, res, err := s.remitTax.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*grpc.RemitTaxResp), nil
}

func decodeRemitTaxRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*grpc.RemitTaxReq)

	return remitTaxReq{RemitTaxReq: mpesa.RemitTaxReq{
		InitiatorName:          req.InitiatorName,
		InitiatorPassword:      req.InitiatorPassword,
		SecurityCredential:     req.SecurityCredential,
		CommandID:              req.CommandID,
		Amount:                 req.Amount,
		PartyA:                 req.PartyA,
		PartyB:                 req.PartyB,
		Remarks:                req.Remarks,
		SenderIdentifierType:   uint8(req.SenderIdentifierType),
		RecieverIdentifierType: uint8(req.RecieverIdentifierType),
		AccountReference:       req.AccountReference,
		QueueTimeOutURL:        req.QueueTimeOutURL,
		ResultURL:              req.ResultURL,
	}}, nil
}

func encodeRemitTaxResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(remitTaxResp)

	return &grpc.RemitTaxResp{
		ValidResp: &grpc.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func encodeError(err error) error {
	switch {
	case errors.Is(err, nil):
		return nil
	case errors.Is(err, errValidation):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
