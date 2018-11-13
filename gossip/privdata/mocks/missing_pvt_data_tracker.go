// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import ledger "github.com/hyperledger/udo/core/ledger"
import mock "github.com/stretchr/testify/mock"

// MissingPvtDataTracker is an autogenerated mock type for the MissingPvtDataTracker type
type MissingPvtDataTracker struct {
	mock.Mock
}

// GetMissingPvtDataInfoForMostRecentBlocks provides a mock function with given fields: maxBlocks
func (_m *MissingPvtDataTracker) GetMissingPvtDataInfoForMostRecentBlocks(maxBlocks int) (ledger.MissingPvtDataInfo, error) {
	ret := _m.Called(maxBlocks)

	var r0 ledger.MissingPvtDataInfo
	if rf, ok := ret.Get(0).(func(int) ledger.MissingPvtDataInfo); ok {
		r0 = rf(maxBlocks)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ledger.MissingPvtDataInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(maxBlocks)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}