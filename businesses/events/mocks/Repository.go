// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	events "kemahin/businesses/events"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) (string, error) {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id int) (*events.Domain, error) {
	ret := _m.Called(id)

	var r0 *events.Domain
	if rf, ok := ret.Get(0).(func(int) *events.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*events.Domain)
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

// GetByJudul provides a mock function with given fields: judul
func (_m *Repository) GetByJudul(judul string) ([]events.Domain, error) {
	ret := _m.Called(judul)

	var r0 []events.Domain
	if rf, ok := ret.Get(0).(func(string) []events.Domain); ok {
		r0 = rf(judul)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(judul)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: data
func (_m *Repository) Register(data *events.Domain) (events.Domain, error) {
	ret := _m.Called(data)

	var r0 events.Domain
	if rf, ok := ret.Get(0).(func(*events.Domain) events.Domain); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(events.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*events.Domain) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpcomingEvent provides a mock function with given fields: date
func (_m *Repository) UpcomingEvent(date time.Time) ([]events.Domain, error) {
	ret := _m.Called(date)

	var r0 []events.Domain
	if rf, ok := ret.Get(0).(func(time.Time) []events.Domain); ok {
		r0 = rf(date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, data
func (_m *Repository) Update(id int, data *events.Domain) (events.Domain, error) {
	ret := _m.Called(id, data)

	var r0 events.Domain
	if rf, ok := ret.Get(0).(func(int, *events.Domain) events.Domain); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Get(0).(events.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *events.Domain) error); ok {
		r1 = rf(id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
