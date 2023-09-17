package grpc

import (
	"context"
	"time"

	"github.com/0x6flab/mpesaoverlay/overlay"
	"github.com/0x6flab/mpesaoverlay/pkg"
	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

const svcName = "mpesaoverlay.overlay.Service"

var _ overlay.ServiceClient = (*grpcClient)(nil)

type grpcClient struct {
	getToken          endpoint.Endpoint
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
func NewClient(conn *grpc.ClientConn, timeout time.Duration) overlay.ServiceClient {
	return &grpcClient{
		getToken: kitgrpc.NewClient(
			conn,
			svcName,
			"GetToken",
			encodeGetTokenRequest,
			decodeGetTokenResponse,
			overlay.TokenResp{},
		).Endpoint(),
		expressQuery: kitgrpc.NewClient(
			conn,
			svcName,
			"ExpressQuery",
			encodeExpressQueryRequest,
			decodeExpressQueryResponse,
			overlay.ExpressQueryResp{},
		).Endpoint(),
		expressSimulate: kitgrpc.NewClient(
			conn,
			svcName,
			"ExpressSimulate",
			encodeExpressSimulateRequest,
			decodeExpressSimulateResponse,
			overlay.ExpressSimulateResp{},
		).Endpoint(),
		b2c: kitgrpc.NewClient(
			conn,
			svcName,
			"B2C",
			encodeB2CRequest,
			decodeB2CResponse,
			overlay.B2CPaymentResp{},
		).Endpoint(),
		accountBalance: kitgrpc.NewClient(
			conn,
			svcName,
			"AccountBalance",
			encodeAccountBalanceRequest,
			decodeAccountBalanceResponse,
			overlay.AccountBalanceResp{},
		).Endpoint(),
		c2bRegisterURL: kitgrpc.NewClient(
			conn,
			svcName,
			"C2BRegisterURL",
			encodeC2BRegisterURLRequest,
			decodeC2BRegisterURLResponse,
			overlay.C2BRegisterURLResp{},
		).Endpoint(),
		c2bSimulate: kitgrpc.NewClient(
			conn,
			svcName,
			"C2BSimulate",
			encodeC2BSimulateRequest,
			decodeC2BSimulateResponse,
			overlay.C2BSimulateResp{},
		).Endpoint(),
		generateQR: kitgrpc.NewClient(
			conn,
			svcName,
			"GenerateQR",
			encodeGenerateQRRequest,
			decodeGenerateQRResponse,
			overlay.GenerateQRResp{},
		).Endpoint(),
		reverse: kitgrpc.NewClient(
			conn,
			svcName,
			"Reverse",
			encodeReverseRequest,
			decodeReverseResponse,
			overlay.ReverseResp{},
		).Endpoint(),
		transactionStatus: kitgrpc.NewClient(
			conn,
			svcName,
			"TransactionStatus",
			encodeTransactionStatusRequest,
			decodeTransactionStatusResponse,
			overlay.TransactionStatusResp{},
		).Endpoint(),
		remitTax: kitgrpc.NewClient(
			conn,
			svcName,
			"RemitTax",
			encodeRemitTaxRequest,
			decodeRemitTaxResponse,
			overlay.RemitTaxResp{},
		).Endpoint(),

		timeout: timeout,
	}
}

func (client grpcClient) GetToken(ctx context.Context, _ *overlay.Empty, _ ...grpc.CallOption) (r *overlay.TokenResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	res, err := client.getToken(ctx, getTokenReq{})
	if err != nil {
		return &overlay.TokenResp{}, err
	}

	ares := res.(getTokenResp)

	return &overlay.TokenResp{
		AccessToken: ares.AccessToken,
		Expiry:      ares.Expiry,
	}, err
}

func decodeGetTokenResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.TokenResp)

	return overlay.TokenResp{
		AccessToken: res.GetAccessToken(),
		Expiry:      res.GetExpiry(),
	}, nil
}

func encodeGetTokenRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return &overlay.Empty{}, nil
}

func (client grpcClient) ExpressQuery(ctx context.Context, req *overlay.ExpressQueryReq, _ ...grpc.CallOption) (r *overlay.ExpressQueryResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	expressQueryReq := expressQueryReq{
		ExpressQueryReq: pkg.ExpressQueryReq{
			PassKey:           req.PassKey,
			BusinessShortCode: req.BusinessShortCode,
			Password:          req.Password,
			Timestamp:         req.Timestamp,
			CheckoutRequestID: req.CheckoutRequestID,
		},
	}
	res, err := client.expressQuery(ctx, expressQueryReq)
	if err != nil {
		return &overlay.ExpressQueryResp{}, err
	}

	ares := res.(expressQueryResp)

	return &overlay.ExpressQueryResp{
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
	res := grpcRes.(*overlay.ExpressQueryResp)

	return expressQueryResp{
		ExpressQueryResp: pkg.ExpressQueryResp{
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

	return &pkg.ExpressQueryReq{
		PassKey:           req.PassKey,
		BusinessShortCode: req.BusinessShortCode,
		Password:          req.Password,
		Timestamp:         req.Timestamp,
		CheckoutRequestID: req.CheckoutRequestID,
	}, nil
}

func (client grpcClient) ExpressSimulate(ctx context.Context, req *overlay.ExpressSimulateReq, _ ...grpc.CallOption) (r *overlay.ExpressSimulateResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	expressSimulateReq := expressSimulateReq{
		ExpressSimulateReq: pkg.ExpressSimulateReq{
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
		return &overlay.ExpressSimulateResp{}, err
	}

	ares := res.(expressSimulateResp)

	return &overlay.ExpressSimulateResp{
		MerchantRequestID:   ares.MerchantRequestID,
		CheckoutRequestID:   ares.CheckoutRequestID,
		ResponseCode:        ares.ResponseCode,
		ResponseDescription: ares.ResponseDescription,
		CustomerMessage:     ares.CustomerMessage,
	}, err
}

func decodeExpressSimulateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.ExpressSimulateResp)

	return expressSimulateResp{
		ExpressSimulateResp: pkg.ExpressSimulateResp{
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

	return &pkg.ExpressSimulateReq{
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

func (client grpcClient) B2CPayment(ctx context.Context, req *overlay.B2CPaymentReq, _ ...grpc.CallOption) (r *overlay.B2CPaymentResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	b2cReq := b2cReq{
		B2CPaymentReq: pkg.B2CPaymentReq{
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
		return &overlay.B2CPaymentResp{}, err
	}

	ares := res.(b2cResp)

	return &overlay.B2CPaymentResp{
		ValidResp: &overlay.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeB2CResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.B2CPaymentResp)

	return b2cResp{
		B2CPaymentResp: pkg.B2CPaymentResp{
			ValidResp: pkg.ValidResp{
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

	return &pkg.B2CPaymentReq{
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

func (client grpcClient) AccountBalance(ctx context.Context, req *overlay.AccountBalanceReq, _ ...grpc.CallOption) (r *overlay.AccountBalanceResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	accountBalanceReq := accountBalanceReq{
		AccountBalanceReq: pkg.AccountBalanceReq{
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
		return &overlay.AccountBalanceResp{}, err
	}

	ares := res.(accountBalanceResp)

	return &overlay.AccountBalanceResp{
		ValidResp: &overlay.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeAccountBalanceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.AccountBalanceResp)

	return accountBalanceResp{
		AccountBalanceResp: pkg.AccountBalanceResp{
			ValidResp: pkg.ValidResp{
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

	return &pkg.AccountBalanceReq{
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

func (client grpcClient) C2BRegisterURL(ctx context.Context, req *overlay.C2BRegisterURLReq, _ ...grpc.CallOption) (r *overlay.C2BRegisterURLResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	c2bRegisterURLReq := c2bRegisterURLReq{
		C2BRegisterURLReq: pkg.C2BRegisterURLReq{
			ShortCode:       req.ShortCode,
			ResponseType:    req.ResponseType,
			ConfirmationURL: req.ConfirmationURL,
			ValidationURL:   req.ValidationURL,
		},
	}
	res, err := client.c2bRegisterURL(ctx, c2bRegisterURLReq)
	if err != nil {
		return &overlay.C2BRegisterURLResp{}, err
	}

	ares := res.(c2bRegisterURLResp)

	return &overlay.C2BRegisterURLResp{
		ValidResp: &overlay.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeC2BRegisterURLResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.C2BRegisterURLResp)

	return c2bRegisterURLResp{
		C2BRegisterURLResp: pkg.C2BRegisterURLResp{
			ValidResp: pkg.ValidResp{
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

	return &pkg.C2BRegisterURLReq{
		ShortCode:       req.ShortCode,
		ResponseType:    req.ResponseType,
		ConfirmationURL: req.ConfirmationURL,
		ValidationURL:   req.ValidationURL,
	}, nil
}

func (client grpcClient) C2BSimulate(ctx context.Context, req *overlay.C2BSimulateReq, _ ...grpc.CallOption) (r *overlay.C2BSimulateResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	c2bSimulateReq := c2bSimulateReq{
		C2BSimulateReq: pkg.C2BSimulateReq{
			ShortCode:     req.ShortCode,
			CommandID:     req.CommandID,
			Amount:        req.Amount,
			Msisdn:        req.Msisdn,
			BillRefNumber: req.BillRefNumber,
		},
	}
	res, err := client.c2bSimulate(ctx, c2bSimulateReq)
	if err != nil {
		return &overlay.C2BSimulateResp{}, err
	}

	ares := res.(c2bSimulateResp)

	return &overlay.C2BSimulateResp{
		ValidResp: &overlay.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeC2BSimulateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.C2BSimulateResp)

	return c2bSimulateResp{
		C2BSimulateResp: pkg.C2BSimulateResp{
			ValidResp: pkg.ValidResp{
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

	return &pkg.C2BSimulateReq{
		ShortCode:     req.ShortCode,
		CommandID:     req.CommandID,
		Amount:        req.Amount,
		Msisdn:        req.Msisdn,
		BillRefNumber: req.BillRefNumber,
	}, nil
}

func (client grpcClient) GenerateQR(ctx context.Context, req *overlay.GenerateQRReq, _ ...grpc.CallOption) (r *overlay.GenerateQRResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	generateQRReq := generateQRReq{
		GenerateQRReq: pkg.GenerateQRReq{
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
		return &overlay.GenerateQRResp{}, err
	}

	ares := res.(generateQRResp)

	return &overlay.GenerateQRResp{
		ResponseDescription: ares.ResponseDescription,
		ResponseCode:        ares.ResponseCode,
		RequestID:           ares.RequestID,
		QRCode:              ares.QRCode,
	}, err
}

func decodeGenerateQRResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.GenerateQRResp)

	return generateQRResp{
		GenerateQRResp: pkg.GenerateQRResp{
			ResponseDescription: res.GetResponseDescription(),
			ResponseCode:        res.GetResponseCode(),
			RequestID:           res.GetRequestID(),
			QRCode:              res.GetQRCode(),
		},
	}, nil
}

func encodeGenerateQRRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(generateQRReq)

	return &pkg.GenerateQRReq{
		MerchantName: req.MerchantName,
		RefNo:        req.RefNo,
		Amount:       req.Amount,
		TrxCode:      req.TrxCode,
		CPI:          req.CPI,
		Size:         req.Size,
	}, nil
}

func (client grpcClient) Reverse(ctx context.Context, req *overlay.ReverseReq, _ ...grpc.CallOption) (r *overlay.ReverseResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	reversalReq := reversalReq{
		ReverseReq: pkg.ReverseReq{
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
		return &overlay.ReverseResp{}, err
	}

	ares := res.(reverseResp)

	return &overlay.ReverseResp{
		ValidResp: &overlay.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeReverseResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.ReverseResp)

	return reverseResp{
		ReverseResp: pkg.ReverseResp{
			ValidResp: pkg.ValidResp{
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

	return &pkg.ReverseReq{
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

func (client grpcClient) TransactionStatus(ctx context.Context, req *overlay.TransactionStatusReq, _ ...grpc.CallOption) (r *overlay.TransactionStatusResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	transactionReq := transactionReq{
		TransactionStatusReq: pkg.TransactionStatusReq{
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
		return &overlay.TransactionStatusResp{}, err
	}

	ares := res.(transactionStatusResp)

	return &overlay.TransactionStatusResp{
		ValidResp: &overlay.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeTransactionStatusResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.TransactionStatusResp)

	return transactionStatusResp{
		TransactionStatusResp: pkg.TransactionStatusResp{
			ValidResp: pkg.ValidResp{
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

	return &pkg.TransactionStatusReq{
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

func (client grpcClient) RemitTax(ctx context.Context, req *overlay.RemitTaxReq, _ ...grpc.CallOption) (r *overlay.RemitTaxResp, err error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	remitTaxReq := remitTaxReq{
		RemitTaxReq: pkg.RemitTaxReq{
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
		return &overlay.RemitTaxResp{}, err
	}

	ares := res.(remitTaxResp)

	return &overlay.RemitTaxResp{
		ValidResp: &overlay.ValidResp{
			OriginatorConversationID: ares.OriginatorConversationID,
			ResponseCode:             ares.ResponseCode,
			ResponseDescription:      ares.ResponseDescription,
			ConversationID:           ares.ConversationID,
		},
	}, err
}

func decodeRemitTaxResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*overlay.RemitTaxResp)

	return remitTaxResp{
		RemitTaxResp: pkg.RemitTaxResp{
			ValidResp: pkg.ValidResp{
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

	return &pkg.RemitTaxReq{
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
