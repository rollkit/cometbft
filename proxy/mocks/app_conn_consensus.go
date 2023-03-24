// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	abcicli "github.com/cometbft/cometbft/abci/client"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cometbft/cometbft/abci/types"
)

// AppConnConsensus is an autogenerated mock type for the AppConnConsensus type
type AppConnConsensus struct {
	mock.Mock
}

// BeginBlockSync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) BeginBlockSync(_a0 types.RequestBeginBlock) (*types.ResponseBeginBlock, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseBeginBlock
	if rf, ok := ret.Get(0).(func(types.RequestBeginBlock) *types.ResponseBeginBlock); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseBeginBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestBeginBlock) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommitSync provides a mock function with given fields:
func (_m *AppConnConsensus) CommitSync() (*types.ResponseCommit, error) {
	ret := _m.Called()

	var r0 *types.ResponseCommit
	if rf, ok := ret.Get(0).(func() *types.ResponseCommit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseCommit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeliverTxAsync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) DeliverTxAsync(_a0 types.RequestDeliverTx) *abcicli.ReqRes {
	ret := _m.Called(_a0)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(types.RequestDeliverTx) *abcicli.ReqRes); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// EndBlockSync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) EndBlockSync(_a0 types.RequestEndBlock) (*types.ResponseEndBlock, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseEndBlock
	if rf, ok := ret.Get(0).(func(types.RequestEndBlock) *types.ResponseEndBlock); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseEndBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestEndBlock) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Error provides a mock function with given fields:
func (_m *AppConnConsensus) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateFraudProofSync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) GenerateFraudProofSync(_a0 types.RequestGenerateFraudProof) (*types.ResponseGenerateFraudProof, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseGenerateFraudProof
	if rf, ok := ret.Get(0).(func(types.RequestGenerateFraudProof) *types.ResponseGenerateFraudProof); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseGenerateFraudProof)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestGenerateFraudProof) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppHashSync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) GetAppHashSync(_a0 types.RequestGetAppHash) (*types.ResponseGetAppHash, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseGetAppHash
	if rf, ok := ret.Get(0).(func(types.RequestGetAppHash) *types.ResponseGetAppHash); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseGetAppHash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestGetAppHash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitChainSync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) InitChainSync(_a0 types.RequestInitChain) (*types.ResponseInitChain, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseInitChain
	if rf, ok := ret.Get(0).(func(types.RequestInitChain) *types.ResponseInitChain); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseInitChain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestInitChain) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareProposalSync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) PrepareProposalSync(_a0 types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponsePrepareProposal
	if rf, ok := ret.Get(0).(func(types.RequestPrepareProposal) *types.ResponsePrepareProposal); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponsePrepareProposal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestPrepareProposal) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessProposalSync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) ProcessProposalSync(_a0 types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseProcessProposal
	if rf, ok := ret.Get(0).(func(types.RequestProcessProposal) *types.ResponseProcessProposal); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseProcessProposal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestProcessProposal) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetResponseCallback provides a mock function with given fields: _a0
func (_m *AppConnConsensus) SetResponseCallback(_a0 abcicli.Callback) {
	_m.Called(_a0)
}

// VerifyFraudProofSync provides a mock function with given fields: _a0
func (_m *AppConnConsensus) VerifyFraudProofSync(_a0 types.RequestVerifyFraudProof) (*types.ResponseVerifyFraudProof, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseVerifyFraudProof
	if rf, ok := ret.Get(0).(func(types.RequestVerifyFraudProof) *types.ResponseVerifyFraudProof); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseVerifyFraudProof)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestVerifyFraudProof) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAppConnConsensus interface {
	mock.TestingT
	Cleanup(func())
}

// NewAppConnConsensus creates a new instance of AppConnConsensus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAppConnConsensus(t mockConstructorTestingTNewAppConnConsensus) *AppConnConsensus {
	mock := &AppConnConsensus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
