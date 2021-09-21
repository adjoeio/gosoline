// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, qb, model
func (_m *Repository) Count(ctx context.Context, qb *db_repo.QueryBuilder, model db_repo.ModelBased) (int, error) {
	ret := _m.Called(ctx, qb, model)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) int); ok {
		r0 = rf(ctx, qb, model)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) error); ok {
		r1 = rf(ctx, qb, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, value
func (_m *Repository) Create(ctx context.Context, value db_repo.ModelBased) error {
	ret := _m.Called(ctx, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, value
func (_m *Repository) Delete(ctx context.Context, value db_repo.ModelBased) error {
	ret := _m.Called(ctx, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMetadata provides a mock function with given fields:
func (_m *Repository) GetMetadata() db_repo.Metadata {
	ret := _m.Called()

	var r0 db_repo.Metadata
	if rf, ok := ret.Get(0).(func() db_repo.Metadata); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(db_repo.Metadata)
	}

	return r0
}

// Query provides a mock function with given fields: ctx, qb, result
func (_m *Repository) Query(ctx context.Context, qb *db_repo.QueryBuilder, result interface{}) error {
	ret := _m.Called(ctx, qb, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, interface{}) error); ok {
		r0 = rf(ctx, qb, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields: ctx, id, out
func (_m *Repository) Read(ctx context.Context, id *uint, out db_repo.ModelBased) error {
	ret := _m.Called(ctx, id, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *uint, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, id, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, value
func (_m *Repository) Update(ctx context.Context, value db_repo.ModelBased) error {
	ret := _m.Called(ctx, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
