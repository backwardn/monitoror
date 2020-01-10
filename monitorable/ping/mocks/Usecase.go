// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-monitorable

package mocks

import (
	monitorormodels "github.com/monitoror/monitoror/models"
	models "github.com/monitoror/monitoror/monitorable/ping/models"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Ping provides a mock function with given fields: params
func (_m *Usecase) Ping(params *models.PingParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.PingParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.PingParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
