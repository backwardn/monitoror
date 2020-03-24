// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock

package mocks

import (
	monitorormodels "github.com/monitoror/monitoror/models"
	models "github.com/monitoror/monitoror/monitorables/azuredevops/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Build provides a mock function with given fields: params
func (_m *Usecase) Build(params *models.BuildParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.BuildParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.BuildParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Release provides a mock function with given fields: params
func (_m *Usecase) Release(params *models.ReleaseParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.ReleaseParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ReleaseParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
