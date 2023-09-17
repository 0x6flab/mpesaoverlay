package grpc

import (
	"context"
	"errors"

	"github.com/0x6flab/mpesaoverlay/overlay"
	"github.com/0x6flab/mpesaoverlay/pkg"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ overlay.ServiceServer = (*grpcServer)(nil)

type grpcServer struct {
	getToken          kitgrpc.Handler
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
	overlay.UnimplementedServiceServer
}

func NewServer(svc overlay.Service) overlay.ServiceServer {
	return &grpcServer{
		getToken: kitgrpc.NewServer(
			getTokenEndpoint(svc),
			decodeGetTokenRequest,
			encodeGetTokenResponse,
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

func (s *grpcServer) GetToken(ctx context.Context, req *overlay.Empty) (*overlay.TokenResp, error) {
	_, res, err := s.getToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.TokenResp), nil
}

func decodeGetTokenRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return getTokenReq{}, nil
}

func encodeGetTokenResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(getTokenResp)

	return &overlay.TokenResp{AccessToken: res.AccessToken, Expiry: res.Expiry}, nil
}

func (s *grpcServer) ExpressQuery(ctx context.Context, req *overlay.ExpressQueryReq) (*overlay.ExpressQueryResp, error) {
	_, res, err := s.expressQuery.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.ExpressQueryResp), nil
}

func decodeExpressQueryRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.ExpressQueryReq)

	return expressQueryReq{
		ExpressQueryReq: pkg.ExpressQueryReq{
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

	return &overlay.ExpressQueryResp{
		MerchantRequestID:   res.MerchantRequestID,
		CheckoutRequestID:   res.CheckoutRequestID,
		ResponseCode:        res.ResponseCode,
		ResponseDescription: res.ResponseDescription,
		CustomerMessage:     res.CustomerMessage,
		ResultCode:          res.ResultCode,
		ResultDesc:          res.ResultDesc,
	}, nil
}

func (s *grpcServer) ExpressSimulate(ctx context.Context, req *overlay.ExpressSimulateReq) (*overlay.ExpressSimulateResp, error) {
	_, res, err := s.expressSimulate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.ExpressSimulateResp), nil
}

func decodeExpressSimulateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.ExpressSimulateReq)

	return expressSimulateReq{
		ExpressSimulateReq: pkg.ExpressSimulateReq{
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

	return &overlay.ExpressSimulateResp{
		MerchantRequestID:   res.MerchantRequestID,
		CheckoutRequestID:   res.CheckoutRequestID,
		ResponseCode:        res.ResponseCode,
		ResponseDescription: res.ResponseDescription,
		CustomerMessage:     res.CustomerMessage,
	}, nil
}

func (s *grpcServer) B2CPayment(ctx context.Context, req *overlay.B2CPaymentReq) (*overlay.B2CPaymentResp, error) {
	_, res, err := s.b2c.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.B2CPaymentResp), nil
}

func decodeB2CRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.B2CPaymentReq)

	return b2cReq{B2CPaymentReq: pkg.B2CPaymentReq{
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

	return &overlay.B2CPaymentResp{
		ValidResp: &overlay.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) AccountBalance(ctx context.Context, req *overlay.AccountBalanceReq) (*overlay.AccountBalanceResp, error) {
	_, res, err := s.accountBalance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.AccountBalanceResp), nil
}

func decodeAccountBalanceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.AccountBalanceReq)

	return accountBalanceReq{AccountBalanceReq: pkg.AccountBalanceReq{
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

	return &overlay.AccountBalanceResp{
		ValidResp: &overlay.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) C2BRegisterURL(ctx context.Context, req *overlay.C2BRegisterURLReq) (*overlay.C2BRegisterURLResp, error) {
	_, res, err := s.c2bRegisterURL.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.C2BRegisterURLResp), nil
}

func decodeC2BRegisterURLRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.C2BRegisterURLReq)

	return c2bRegisterURLReq{C2BRegisterURLReq: pkg.C2BRegisterURLReq{
		ShortCode:       req.ShortCode,
		ResponseType:    req.ResponseType,
		ConfirmationURL: req.ConfirmationURL,
		ValidationURL:   req.ValidationURL,
	}}, nil
}

func encodeC2BRegisterURLResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(c2bRegisterURLResp)

	return &overlay.C2BRegisterURLResp{
		ValidResp: &overlay.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) C2BSimulate(ctx context.Context, req *overlay.C2BSimulateReq) (*overlay.C2BSimulateResp, error) {
	_, res, err := s.c2bSimulate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.C2BSimulateResp), nil
}

func decodeC2BSimulateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.C2BSimulateReq)

	return c2bSimulateReq{C2BSimulateReq: pkg.C2BSimulateReq{
		ShortCode:     req.ShortCode,
		CommandID:     req.CommandID,
		Amount:        req.Amount,
		Msisdn:        req.Msisdn,
		BillRefNumber: req.BillRefNumber,
	}}, nil
}

func encodeC2BSimulateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(c2bSimulateResp)

	return &overlay.C2BSimulateResp{
		ValidResp: &overlay.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) GenerateQR(ctx context.Context, req *overlay.GenerateQRReq) (*overlay.GenerateQRResp, error) {
	_, res, err := s.generateQR.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.GenerateQRResp), nil
}

func decodeGenerateQRRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.GenerateQRReq)

	return generateQRReq{GenerateQRReq: pkg.GenerateQRReq{
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

	return &overlay.GenerateQRResp{
		RequestID:           res.RequestID,
		QRCode:              res.QRCode,
		ResponseCode:        res.ResponseCode,
		ResponseDescription: res.ResponseDescription,
	}, nil
}

func (s *grpcServer) Reverse(ctx context.Context, req *overlay.ReverseReq) (*overlay.ReverseResp, error) {
	_, res, err := s.reverse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.ReverseResp), nil
}

func decodeReverseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.ReverseReq)

	return reversalReq{ReverseReq: pkg.ReverseReq{
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

	return &overlay.ReverseResp{
		ValidResp: &overlay.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) TransactionStatus(ctx context.Context, req *overlay.TransactionStatusReq) (*overlay.TransactionStatusResp, error) {
	_, res, err := s.transactionStatus.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.TransactionStatusResp), nil
}

func decodeTransactionStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.TransactionStatusReq)

	return transactionReq{TransactionStatusReq: pkg.TransactionStatusReq{
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

	return &overlay.TransactionStatusResp{
		ValidResp: &overlay.ValidResp{
			ConversationID:           res.ConversationID,
			OriginatorConversationID: res.OriginatorConversationID,
			ResponseCode:             res.ResponseCode,
			ResponseDescription:      res.ResponseDescription,
		},
	}, nil
}

func (s *grpcServer) RemitTax(ctx context.Context, req *overlay.RemitTaxReq) (*overlay.RemitTaxResp, error) {
	_, res, err := s.remitTax.ServeGRPC(ctx, req)
	if err != nil {
		return nil, encodeError(err)
	}

	return res.(*overlay.RemitTaxResp), nil
}

func decodeRemitTaxRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*overlay.RemitTaxReq)

	return remitTaxReq{RemitTaxReq: pkg.RemitTaxReq{
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

	return &overlay.RemitTaxResp{
		ValidResp: &overlay.ValidResp{
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
