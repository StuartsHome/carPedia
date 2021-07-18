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
