// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	mpesa "github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	mock "github.com/stretchr/testify/mock"
)

// SDK is an autogenerated mock type for the SDK type
type SDK struct {
	mock.Mock
}

// AccountBalance provides a mock function with given fields: abReq
func (_m *SDK) AccountBalance(abReq mpesa.AccountBalanceReq) (mpesa.AccountBalanceResp, error) {
	ret := _m.Called(abReq)

	var r0 mpesa.AccountBalanceResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.AccountBalanceReq) (mpesa.AccountBalanceResp, error)); ok {
		return rf(abReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.AccountBalanceReq) mpesa.AccountBalanceResp); ok {
		r0 = rf(abReq)
	} else {
		r0 = ret.Get(0).(mpesa.AccountBalanceResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.AccountBalanceReq) error); ok {
		r1 = rf(abReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// B2CPayment provides a mock function with given fields: b2cReq
func (_m *SDK) B2CPayment(b2cReq mpesa.B2CPaymentReq) (mpesa.B2CPaymentResp, error) {
	ret := _m.Called(b2cReq)

	var r0 mpesa.B2CPaymentResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.B2CPaymentReq) (mpesa.B2CPaymentResp, error)); ok {
		return rf(b2cReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.B2CPaymentReq) mpesa.B2CPaymentResp); ok {
		r0 = rf(b2cReq)
	} else {
		r0 = ret.Get(0).(mpesa.B2CPaymentResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.B2CPaymentReq) error); ok {
		r1 = rf(b2cReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// C2BRegisterURL provides a mock function with given fields: c2bReq
func (_m *SDK) C2BRegisterURL(c2bReq mpesa.C2BRegisterURLReq) (mpesa.C2BRegisterURLResp, error) {
	ret := _m.Called(c2bReq)

	var r0 mpesa.C2BRegisterURLResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.C2BRegisterURLReq) (mpesa.C2BRegisterURLResp, error)); ok {
		return rf(c2bReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.C2BRegisterURLReq) mpesa.C2BRegisterURLResp); ok {
		r0 = rf(c2bReq)
	} else {
		r0 = ret.Get(0).(mpesa.C2BRegisterURLResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.C2BRegisterURLReq) error); ok {
		r1 = rf(c2bReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// C2BSimulate provides a mock function with given fields: c2bReq
func (_m *SDK) C2BSimulate(c2bReq mpesa.C2BSimulateReq) (mpesa.C2BSimulateResp, error) {
	ret := _m.Called(c2bReq)

	var r0 mpesa.C2BSimulateResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.C2BSimulateReq) (mpesa.C2BSimulateResp, error)); ok {
		return rf(c2bReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.C2BSimulateReq) mpesa.C2BSimulateResp); ok {
		r0 = rf(c2bReq)
	} else {
		r0 = ret.Get(0).(mpesa.C2BSimulateResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.C2BSimulateReq) error); ok {
		r1 = rf(c2bReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpressQuery provides a mock function with given fields: eqReq
func (_m *SDK) ExpressQuery(eqReq mpesa.ExpressQueryReq) (mpesa.ExpressQueryResp, error) {
	ret := _m.Called(eqReq)

	var r0 mpesa.ExpressQueryResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.ExpressQueryReq) (mpesa.ExpressQueryResp, error)); ok {
		return rf(eqReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.ExpressQueryReq) mpesa.ExpressQueryResp); ok {
		r0 = rf(eqReq)
	} else {
		r0 = ret.Get(0).(mpesa.ExpressQueryResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.ExpressQueryReq) error); ok {
		r1 = rf(eqReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpressSimulate provides a mock function with given fields: eReq
func (_m *SDK) ExpressSimulate(eReq mpesa.ExpressSimulateReq) (mpesa.ExpressSimulateResp, error) {
	ret := _m.Called(eReq)

	var r0 mpesa.ExpressSimulateResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.ExpressSimulateReq) (mpesa.ExpressSimulateResp, error)); ok {
		return rf(eReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.ExpressSimulateReq) mpesa.ExpressSimulateResp); ok {
		r0 = rf(eReq)
	} else {
		r0 = ret.Get(0).(mpesa.ExpressSimulateResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.ExpressSimulateReq) error); ok {
		r1 = rf(eReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateQR provides a mock function with given fields: qReq
func (_m *SDK) GenerateQR(qReq mpesa.GenerateQRReq) (mpesa.GenerateQRResp, error) {
	ret := _m.Called(qReq)

	var r0 mpesa.GenerateQRResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.GenerateQRReq) (mpesa.GenerateQRResp, error)); ok {
		return rf(qReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.GenerateQRReq) mpesa.GenerateQRResp); ok {
		r0 = rf(qReq)
	} else {
		r0 = ret.Get(0).(mpesa.GenerateQRResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.GenerateQRReq) error); ok {
		r1 = rf(qReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemitTax provides a mock function with given fields: rReq
func (_m *SDK) RemitTax(rReq mpesa.RemitTaxReq) (mpesa.RemitTaxResp, error) {
	ret := _m.Called(rReq)

	var r0 mpesa.RemitTaxResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.RemitTaxReq) (mpesa.RemitTaxResp, error)); ok {
		return rf(rReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.RemitTaxReq) mpesa.RemitTaxResp); ok {
		r0 = rf(rReq)
	} else {
		r0 = ret.Get(0).(mpesa.RemitTaxResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.RemitTaxReq) error); ok {
		r1 = rf(rReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reverse provides a mock function with given fields: rReq
func (_m *SDK) Reverse(rReq mpesa.ReverseReq) (mpesa.ReverseResp, error) {
	ret := _m.Called(rReq)

	var r0 mpesa.ReverseResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.ReverseReq) (mpesa.ReverseResp, error)); ok {
		return rf(rReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.ReverseReq) mpesa.ReverseResp); ok {
		r0 = rf(rReq)
	} else {
		r0 = ret.Get(0).(mpesa.ReverseResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.ReverseReq) error); ok {
		r1 = rf(rReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Token provides a mock function with given fields:
func (_m *SDK) Token() (mpesa.TokenResp, error) {
	ret := _m.Called()

	var r0 mpesa.TokenResp
	var r1 error
	if rf, ok := ret.Get(0).(func() (mpesa.TokenResp, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() mpesa.TokenResp); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mpesa.TokenResp)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionStatus provides a mock function with given fields: tReq
func (_m *SDK) TransactionStatus(tReq mpesa.TransactionStatusReq) (mpesa.TransactionStatusResp, error) {
	ret := _m.Called(tReq)

	var r0 mpesa.TransactionStatusResp
	var r1 error
	if rf, ok := ret.Get(0).(func(mpesa.TransactionStatusReq) (mpesa.TransactionStatusResp, error)); ok {
		return rf(tReq)
	}
	if rf, ok := ret.Get(0).(func(mpesa.TransactionStatusReq) mpesa.TransactionStatusResp); ok {
		r0 = rf(tReq)
	} else {
		r0 = ret.Get(0).(mpesa.TransactionStatusResp)
	}

	if rf, ok := ret.Get(1).(func(mpesa.TransactionStatusReq) error); ok {
		r1 = rf(tReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSDK creates a new instance of SDK. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSDK(t interface {
	mock.TestingT
	Cleanup(func())
}) *SDK {
	mock := &SDK{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
