// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	users "github.com/dragranzer/Golang-Mini-Project/features/users"
	mock "github.com/stretchr/testify/mock"
)

// Bussiness is an autogenerated mock type for the Bussiness type
type Bussiness struct {
	mock.Mock
}

// ChangeDatabyID provides a mock function with given fields: id, newData
func (_m *Bussiness) ChangeDatabyID(id int, newData users.Core) error {
	ret := _m.Called(id, newData)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, users.Core) error); ok {
		r0 = rf(id, newData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateData provides a mock function with given fields: data
func (_m *Bussiness) CreateData(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllData provides a mock function with given fields:
func (_m *Bussiness) GetAllData() []users.Core {
	ret := _m.Called()

	var r0 []users.Core
	if rf, ok := ret.Get(0).(func() []users.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
		}
	}

	return r0
}

// GetDatabyID provides a mock function with given fields: id
func (_m *Bussiness) GetDatabyID(id int) (users.Core, error) {
	ret := _m.Called(id)

	var r0 users.Core
	if rf, ok := ret.Get(0).(func(int) users.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatabyName provides a mock function with given fields: name
func (_m *Bussiness) GetDatabyName(name string) (users.Core, error) {
	ret := _m.Called(name)

	var r0 users.Core
	if rf, ok := ret.Get(0).(func(string) users.Core); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, pass
func (_m *Bussiness) Login(email string, pass string) (users.UserResp, bool, error) {
	ret := _m.Called(email, pass)

	var r0 users.UserResp
	if rf, ok := ret.Get(0).(func(string, string) users.UserResp); ok {
		r0 = rf(email, pass)
	} else {
		r0 = ret.Get(0).(users.UserResp)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(email, pass)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, pass)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}