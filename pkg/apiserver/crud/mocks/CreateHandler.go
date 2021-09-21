// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	crud "github.com/justtrackio/gosoline/pkg/apiserver/crud"
	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"

	mock "github.com/stretchr/testify/mock"
)

// CreateHandler is an autogenerated mock type for the CreateHandler type
type CreateHandler struct {
	mock.Mock
}

// GetCreateInput provides a mock function with given fields:
func (_m *CreateHandler) GetCreateInput() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetModel provides a mock function with given fields:
func (_m *CreateHandler) GetModel() db_repo.ModelBased {
	ret := _m.Called()

	var r0 db_repo.ModelBased
	if rf, ok := ret.Get(0).(func() db_repo.ModelBased); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db_repo.ModelBased)
		}
	}

	return r0
}

// GetRepository provides a mock function with given fields:
func (_m *CreateHandler) GetRepository() crud.Repository {
	ret := _m.Called()

	var r0 crud.Repository
	if rf, ok := ret.Get(0).(func() crud.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crud.Repository)
		}
	}

	return r0
}

// TransformCreate provides a mock function with given fields: input, model
func (_m *CreateHandler) TransformCreate(input interface{}, model db_repo.ModelBased) error {
	ret := _m.Called(input, model)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, db_repo.ModelBased) error); ok {
		r0 = rf(input, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransformOutput provides a mock function with given fields: model, apiView
func (_m *CreateHandler) TransformOutput(model db_repo.ModelBased, apiView string) (interface{}, error) {
	ret := _m.Called(model, apiView)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(db_repo.ModelBased, string) interface{}); ok {
		r0 = rf(model, apiView)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db_repo.ModelBased, string) error); ok {
		r1 = rf(model, apiView)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
