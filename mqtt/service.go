package mqtt

import (
	"encoding/json"

	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/mochi-mqtt/server/v2/packets"
)

type Service interface {
	GetToken(pk packets.Packet) (mpesa.TokenResp, error)

	ExpressQuery(pk packets.Packet) (mpesa.ExpressQueryResp, error)

	ExpressSimulate(pk packets.Packet) (mpesa.ExpressSimulateResp, error)

	B2CPayment(pk packets.Packet) (mpesa.B2CPaymentResp, error)

	AccountBalance(pk packets.Packet) (mpesa.AccountBalanceResp, error)

	C2BRegisterURL(pk packets.Packet) (mpesa.C2BRegisterURLResp, error)

	C2BSimulate(pk packets.Packet) (mpesa.C2BSimulateResp, error)

	GenerateQR(pk packets.Packet) (mpesa.GenerateQRResp, error)

	Reverse(pk packets.Packet) (mpesa.ReverseResp, error)

	TransactionStatus(pk packets.Packet) (mpesa.TransactionStatusResp, error)

	RemitTax(pk packets.Packet) (mpesa.RemitTaxResp, error)
}

type service struct {
	sdk mpesa.SDK
}

var _ Service = (*service)(nil)

func NewService(sdk mpesa.SDK) Service {
	return &service{sdk: sdk}
}

func (s *service) GetToken(_ packets.Packet) (mpesa.TokenResp, error) {

	return s.sdk.GetToken()
}

func (s *service) ExpressQuery(pk packets.Packet) (mpesa.ExpressQueryResp, error) {
	var req mpesa.ExpressQueryReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.ExpressQueryResp{}, err
	}

	return s.sdk.ExpressQuery(req)
}

func (s *service) ExpressSimulate(pk packets.Packet) (mpesa.ExpressSimulateResp, error) {
	var req mpesa.ExpressSimulateReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.ExpressSimulateResp{}, err
	}

	return s.sdk.ExpressSimulate(req)
}

func (s *service) B2CPayment(pk packets.Packet) (mpesa.B2CPaymentResp, error) {
	var req mpesa.B2CPaymentReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.B2CPaymentResp{}, err
	}

	return s.sdk.B2CPayment(req)
}

func (s *service) AccountBalance(pk packets.Packet) (mpesa.AccountBalanceResp, error) {
	var req mpesa.AccountBalanceReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.AccountBalanceResp{}, err
	}

	return s.sdk.AccountBalance(req)
}

func (s *service) C2BRegisterURL(pk packets.Packet) (mpesa.C2BRegisterURLResp, error) {
	var req mpesa.C2BRegisterURLReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.C2BRegisterURLResp{}, err
	}

	return s.sdk.C2BRegisterURL(req)
}

func (s *service) C2BSimulate(pk packets.Packet) (mpesa.C2BSimulateResp, error) {
	var req mpesa.C2BSimulateReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.C2BSimulateResp{}, err
	}

	return s.sdk.C2BSimulate(req)
}

func (s *service) GenerateQR(pk packets.Packet) (mpesa.GenerateQRResp, error) {
	var req mpesa.GenerateQRReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.GenerateQRResp{}, err
	}

	return s.sdk.GenerateQR(req)
}

func (s *service) Reverse(pk packets.Packet) (mpesa.ReverseResp, error) {
	var req mpesa.ReverseReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.ReverseResp{}, err
	}

	return s.sdk.Reverse(req)
}

func (s *service) TransactionStatus(pk packets.Packet) (mpesa.TransactionStatusResp, error) {
	var req mpesa.TransactionStatusReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.TransactionStatusResp{}, err
	}

	return s.sdk.TransactionStatus(req)
}

func (s *service) RemitTax(pk packets.Packet) (mpesa.RemitTaxResp, error) {
	var req mpesa.RemitTaxReq
	if err := json.Unmarshal(pk.Payload, &req); err != nil {
		return mpesa.RemitTaxResp{}, err
	}

	return s.sdk.RemitTax(req)
}
