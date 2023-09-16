package pkg

// import "go.uber.org/zap"

// type loggingMiddleware struct {
// 	logger *zap.Logger
// 	sdk    SDK
// }

// func LoggingMiddleware(logger *zap.Logger, sdk SDK) SDK {
// 	return &loggingMiddleware{
// 		logger: logger,
// 		sdk:    sdk,
// 	}
// }

// func (lm *loggingMiddleware) GetToken() (TokenResp, error) {
// 	lm.logger.Info("getting token")
// 	return lm.sdk.GetToken()
// }

// func (lm *loggingMiddleware) ExpressQuery(eqReq ExpressQueryReq) (ExpressQueryResp, error) {
// 	lm.logger.Info("express query")
// 	return lm.sdk.ExpressQuery(eqReq)
// }

// func (lm *loggingMiddleware) ExpressSimulate(esReq ExpressSimulateReq) (ExpressSimulateResp, error) {
// 	lm.logger.Info("express simulate")
// 	return lm.sdk.ExpressSimulate(esReq)
// }

// func (lm *loggingMiddleware) B2CPayment(b2cReq B2CPaymentReq) (B2CPaymentResp, error) {
// 	lm.logger.Info("b2c payment")
// 	return lm.sdk.B2CPayment(b2cReq)
// }

// func (lm *loggingMiddleware) AccountBalance(abReq AccountBalanceReq) (AccountBalanceResp, error) {
// 	lm.logger.Info("account balance")
// 	return lm.sdk.AccountBalance(abReq)
// }

// // C2BRegisterURL
// func (lm *loggingMiddleware) C2BRegisterURL(c2bReq C2BRegisterURLReq) (C2BRegisterURLResp, error) {
// 	lm.logger.Info("c2b register url")
// 	return lm.sdk.C2BRegisterURL(c2bReq)
// }

// // C2BSimulate
// func (lm *loggingMiddleware) C2BSimulate(c2bReq C2BSimulateReq) (C2BSimulateResp, error) {
// 	lm.logger.Info("c2b simulate")
// 	return lm.sdk.C2BSimulate(c2bReq)
// }

// // GenerateQR
// func (lm *loggingMiddleware) GenerateQR(gqrReq GenerateQRReq) (GenerateQRResp, error) {
// 	lm.logger.Info("generate qr")
// 	return lm.sdk.GenerateQR(gqrReq)
// }

// // Reverse
// func (lm *loggingMiddleware) Reverse(rReq ReverseReq) (ReverseResp, error) {
// 	lm.logger.Info("reverse")
// 	return lm.sdk.Reverse(rReq)
// }

// // TransactionStatus
// func (lm *loggingMiddleware) TransactionStatus(tsReq TransactionStatusReq) (TransactionStatusResp, error) {
// 	lm.logger.Info("transaction status")
// 	return lm.sdk.TransactionStatus(tsReq)
// }

// // RemitTax
// func (lm *loggingMiddleware) RemitTax(rtReq RemitTaxReq) (RemitTaxResp, error) {
// 	lm.logger.Info("remit tax")
// 	return lm.sdk.RemitTax(rtReq)
// }
