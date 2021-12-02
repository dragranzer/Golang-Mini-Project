// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	reviews "github.com/dragranzer/Golang-Mini-Project/features/reviews"
	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// InsertData provides a mock function with given fields: data
func (_m *Data) InsertData(data reviews.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(reviews.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectDatabyBookID provides a mock function with given fields: id
func (_m *Data) SelectDatabyBookID(id int) ([]reviews.Core, error) {
	ret := _m.Called(id)

	var r0 []reviews.Core
	if rf, ok := ret.Get(0).(func(int) []reviews.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reviews.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
