// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"
	mongo "task7/mongo"

	mock "github.com/stretchr/testify/mock"

	mongo_drivermongo "go.mongodb.org/mongo-driver/mongo"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Connect provides a mock function with given fields: _a0
func (_m *Client) Connect(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database provides a mock function with given fields: _a0
func (_m *Client) Database(_a0 string) mongo.Database {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Database")
	}

	var r0 mongo.Database
	if rf, ok := ret.Get(0).(func(string) mongo.Database); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Database)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields: _a0
func (_m *Client) Disconnect(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Disconnect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields: _a0
func (_m *Client) Ping(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartSession provides a mock function with given fields:
func (_m *Client) StartSession() (mongo_drivermongo.Session, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StartSession")
	}

	var r0 mongo_drivermongo.Session
	var r1 error
	if rf, ok := ret.Get(0).(func() (mongo_drivermongo.Session, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() mongo_drivermongo.Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo_drivermongo.Session)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseSession provides a mock function with given fields: ctx, fn
func (_m *Client) UseSession(ctx context.Context, fn func(mongo_drivermongo.SessionContext) error) error {
	ret := _m.Called(ctx, fn)

	if len(ret) == 0 {
		panic("no return value specified for UseSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(mongo_drivermongo.SessionContext) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
