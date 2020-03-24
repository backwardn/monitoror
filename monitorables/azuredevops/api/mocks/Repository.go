// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock

package mocks

import (
	models "github.com/monitoror/monitoror/monitorables/azuredevops/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetBuild provides a mock function with given fields: project, definition, branch
func (_m *Repository) GetBuild(project string, definition int, branch *string) (*models.Build, error) {
	ret := _m.Called(project, definition, branch)

	var r0 *models.Build
	if rf, ok := ret.Get(0).(func(string, int, *string) *models.Build); ok {
		r0 = rf(project, definition, branch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, *string) error); ok {
		r1 = rf(project, definition, branch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRelease provides a mock function with given fields: project, definition
func (_m *Repository) GetRelease(project string, definition int) (*models.Release, error) {
	ret := _m.Called(project, definition)

	var r0 *models.Release
	if rf, ok := ret.Get(0).(func(string, int) *models.Release); ok {
		r0 = rf(project, definition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(project, definition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
