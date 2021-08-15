// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock_store

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/stuartshome/carpedia/model"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CreateCar provides a mock function with given fields: car
func (_m *Store) CreateCar(car *model.Car) error {
	ret := _m.Called(car)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Car) error); ok {
		r0 = rf(car)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCar provides a mock function with given fields: car
func (_m *Store) DeleteCar(car *model.Car) error {
	ret := _m.Called(car)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Car) error); ok {
		r0 = rf(car)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCar provides a mock function with given fields: _a0
func (_m *Store) GetCar(_a0 *model.Car) (*model.Car, error) {
	ret := _m.Called(_a0)

	var r0 *model.Car
	if rf, ok := ret.Get(0).(func(*model.Car) *model.Car); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Car)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Car) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCars provides a mock function with given fields:
func (_m *Store) GetCars() ([]*model.Car, error) {
	ret := _m.Called()

	var r0 []*model.Car
	if rf, ok := ret.Get(0).(func() []*model.Car); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Car)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCar provides a mock function with given fields: car
func (_m *Store) UpdateCar(car *model.Car) error {
	ret := _m.Called(car)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Car) error); ok {
		r0 = rf(car)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
