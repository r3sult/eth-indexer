// Code generated by mockery v1.0.0
package mocks

import big "math/big"
import common "github.com/ethereum/go-ethereum/common"
import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/maichain/eth-indexer/model"

// ServiceManager is an autogenerated mock type for the ServiceManager type
type ServiceManager struct {
	mock.Mock
}

// FindBlockByHash provides a mock function with given fields: hash
func (_m *ServiceManager) FindBlockByHash(hash []byte) (*model.Header, error) {
	ret := _m.Called(hash)

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func([]byte) *model.Header); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBlockByNumber provides a mock function with given fields: blockNumber
func (_m *ServiceManager) FindBlockByNumber(blockNumber int64) (*model.Header, error) {
	ret := _m.Called(blockNumber)

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func(int64) *model.Header); ok {
		r0 = rf(blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLatestBlock provides a mock function with given fields:
func (_m *ServiceManager) FindLatestBlock() (*model.Header, error) {
	ret := _m.Called()

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func() *model.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
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

// FindTransaction provides a mock function with given fields: hash
func (_m *ServiceManager) FindTransaction(hash []byte) (*model.Transaction, error) {
	ret := _m.Called(hash)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func([]byte) *model.Transaction); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransactionsByBlockHash provides a mock function with given fields: blockHash
func (_m *ServiceManager) FindTransactionsByBlockHash(blockHash []byte) ([]*model.Transaction, error) {
	ret := _m.Called(blockHash)

	var r0 []*model.Transaction
	if rf, ok := ret.Get(0).(func([]byte) []*model.Transaction); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: ctx, address, blockNr
func (_m *ServiceManager) GetBalance(ctx context.Context, address common.Address, blockNr int64) (*big.Int, *big.Int, error) {
	ret := _m.Called(ctx, address, blockNr)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, int64) *big.Int); ok {
		r0 = rf(ctx, address, blockNr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, int64) *big.Int); ok {
		r1 = rf(ctx, address, blockNr)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Address, int64) error); ok {
		r2 = rf(ctx, address, blockNr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetERC20Balance provides a mock function with given fields: ctx, contractAddress, address, blockNr
func (_m *ServiceManager) GetERC20Balance(ctx context.Context, contractAddress common.Address, address common.Address, blockNr int64) (*big.Int, *big.Int, error) {
	ret := _m.Called(ctx, contractAddress, address, blockNr)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Address, int64) *big.Int); ok {
		r0 = rf(ctx, contractAddress, address, blockNr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Address, int64) *big.Int); ok {
		r1 = rf(ctx, contractAddress, address, blockNr)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Address, common.Address, int64) error); ok {
		r2 = rf(ctx, contractAddress, address, blockNr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}