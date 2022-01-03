// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	users "kemahin/businesses/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id int) (users.Domain, error) {
	ret := _m.Called(id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(int) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNIM provides a mock function with given fields: nim
func (_m *Repository) GetByNIM(nim string) (users.Domain, error) {
	ret := _m.Called(nim)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(string) users.Domain); ok {
		r0 = rf(nim)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nim)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: data
func (_m *Repository) Register(data *users.Domain) (users.Domain, error) {
	ret := _m.Called(data)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: Data
func (_m *Repository) Update(Data users.Domain) (users.Domain, error) {
	ret := _m.Called(Data)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(users.Domain) users.Domain); ok {
		r0 = rf(Data)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Domain) error); ok {
		r1 = rf(Data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
