// Code generated by mockery v1.0.0. DO NOT EDIT.

package rpcbackendmocks

import (
	context "context"

	rpcbackend "github.com/hyperledger/firefly-signer/pkg/rpcbackend"
	mock "github.com/stretchr/testify/mock"
)

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// CallRPC provides a mock function with given fields: ctx, result, method, params
func (_m *Backend) CallRPC(ctx context.Context, result interface{}, method string, params ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, result, method)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, result, method, params...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SyncRequest provides a mock function with given fields: ctx, rpcReq
func (_m *Backend) SyncRequest(ctx context.Context, rpcReq *rpcbackend.RPCRequest) (*rpcbackend.RPCResponse, error) {
	ret := _m.Called(ctx, rpcReq)

	var r0 *rpcbackend.RPCResponse
	if rf, ok := ret.Get(0).(func(context.Context, *rpcbackend.RPCRequest) *rpcbackend.RPCResponse); ok {
		r0 = rf(ctx, rpcReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcbackend.RPCResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *rpcbackend.RPCRequest) error); ok {
		r1 = rf(ctx, rpcReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
