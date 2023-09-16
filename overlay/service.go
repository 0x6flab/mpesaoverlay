package overlay

import (
	"context"

	"github.com/0x6flab/mpesaoverlay/pkg"
)

type Service interface {
	// GetToken Gives you a time bound access token to call allowed APIs.
	//
	// The token is valid for the specified time duration, which is usually an hour.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/Authorization
	GetToken(ctx context.Context) (pkg.TokenResp, error)

	// ExpressQuery Check the status of a Lipa Na M-Pesa Online Payment.
	//
	// Query the payment status of a Lipa Na M-Pesa Online Payment using the M-Pesa transaction ID.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/MpesaExpressQuery
	ExpressQuery(ctx context.Context, eqReq pkg.ExpressQueryReq) (pkg.ExpressQueryResp, error)

	// ExpressSimulate Initiates online payment on behalf of a customer.
	//
	// Sends a USSD push to the customer’s phone to prompt them to enter their PIN to authorize the payment.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/MpesaExpressSimulate
	ExpressSimulate(ctx context.Context, eReq pkg.ExpressSimulateReq) (pkg.ExpressSimulateResp, error)

	// B2CPayment Transact between an M-Pesa short code to a phone number registered on M-Pesa
	//
	// B2C API is an API used to make payments from a Business to Customers (Pay Outs), also known as Bulk Disbursements.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/BusinessToCustomer
	B2CPayment(ctx context.Context, b2cReq pkg.B2CPaymentReq) (pkg.B2CPaymentResp, error)

	// AccountBalance Enquire the balance on an M-Pesa BuyGoods (Till Number)
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/AccountBalance
	AccountBalance(ctx context.Context, abReq pkg.AccountBalanceReq) (pkg.AccountBalanceResp, error)

	// C2BRegisterURL Register validation and confirmation URLs on M-Pesa
	//
	// Register URL API works hand in hand with Customer to Business (C2B) APIs and allows receiving payment notifications to your paybill.
	//
	// This API enables you to register the callback URLs via which you shall receive notifications for payments to your pay bill/till number.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/CustomerToBusinessRegisterURL
	C2BRegisterURL(ctx context.Context, c2bReq pkg.C2BRegisterURLReq) (pkg.C2BRegisterURLResp, error)

	// C2BSimulate Make payment requests from Client to Business (C2B)
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/CustomerToBusinessRegisterURL
	C2BSimulate(ctx context.Context, c2bReq pkg.C2BSimulateReq) (pkg.C2BSimulateResp, error)

	// GenerateQR Generates a dynamic M-PESA QR Code.
	//
	// Documentation: https://developer.safaricom.co.ke/APIs/DynamicQRCode

	GenerateQR(ctx context.Context, qReq pkg.GenerateQRReq) (pkg.GenerateQRResp, error)

	// Reverse Reverses an M-Pesa transaction.
	Reverse(ctx context.Context, rReq pkg.ReverseReq) (pkg.ReverseResp, error)

	// TransactionStatus Check the status of a transaction
	//
	// Check the status of a transaction.
	TransactionStatus(ctx context.Context, tReq pkg.TransactionStatusReq) (pkg.TransactionStatusResp, error)

	// RemitTax enables businesses to remit tax to Kenya Revenue Authority (KRA).
	RemitTax(ctx context.Context, rReq pkg.RemitTaxReq) (pkg.RemitTaxResp, error)
}

type service struct {
	sdk pkg.SDK
}

var _ Service = (*service)(nil)

func NewService(sdk pkg.SDK) Service {
	return &service{sdk: sdk}
}

func (s *service) GetToken(_ context.Context) (pkg.TokenResp, error) {
	return s.sdk.GetToken()
}

func (s *service) ExpressQuery(_ context.Context, eqReq pkg.ExpressQueryReq) (pkg.ExpressQueryResp, error) {
	return s.sdk.ExpressQuery(eqReq)
}

func (s *service) ExpressSimulate(_ context.Context, eReq pkg.ExpressSimulateReq) (pkg.ExpressSimulateResp, error) {
	return s.sdk.ExpressSimulate(eReq)
}

func (s *service) B2CPayment(_ context.Context, b2cReq pkg.B2CPaymentReq) (pkg.B2CPaymentResp, error) {
	return s.sdk.B2CPayment(b2cReq)
}

func (s *service) AccountBalance(_ context.Context, abReq pkg.AccountBalanceReq) (pkg.AccountBalanceResp, error) {
	return s.sdk.AccountBalance(abReq)
}

func (s *service) C2BRegisterURL(_ context.Context, c2bReq pkg.C2BRegisterURLReq) (pkg.C2BRegisterURLResp, error) {
	return s.sdk.C2BRegisterURL(c2bReq)
}

func (s *service) C2BSimulate(_ context.Context, c2bReq pkg.C2BSimulateReq) (pkg.C2BSimulateResp, error) {
	return s.sdk.C2BSimulate(c2bReq)
}

func (s *service) GenerateQR(_ context.Context, qReq pkg.GenerateQRReq) (pkg.GenerateQRResp, error) {
	return s.sdk.GenerateQR(qReq)
}

func (s *service) Reverse(_ context.Context, rReq pkg.ReverseReq) (pkg.ReverseResp, error) {
	return s.sdk.Reverse(rReq)
}

func (s *service) TransactionStatus(_ context.Context, tReq pkg.TransactionStatusReq) (pkg.TransactionStatusResp, error) {
	return s.sdk.TransactionStatus(tReq)
}

func (s *service) RemitTax(_ context.Context, rReq pkg.RemitTaxReq) (pkg.RemitTaxResp, error) {
	return s.sdk.RemitTax(rReq)
}
