// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Requester2 is an autogenerated mock type for the Requester2 type
type Requester2 struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *Requester2) Get(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRequester2 creates a new instance of Requester2. It also registers a cleanup function to assert the mocks expectations.
func NewRequester2(t testing.TB) *Requester2 {
	mock := &Requester2{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
