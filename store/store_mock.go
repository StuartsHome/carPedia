package store

import (
	"github.com/stretchr/testify/mock"
	"github.com/stuartshome/carpedia/cache"
	"github.com/stuartshome/carpedia/model"
)

// The mock store contains additonal methods for inspection
type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateCar(car *model.Car) error {
	/*
		When this method is called, `m.Called` records the call, and also
		returns the result that we pass to it (which you will see in the
		handler tests)
	*/
	rets := m.Called(car)
	return rets.Error(0)
}

func (m *MockStore) GetCars() ([]*model.Car, error) {
	/*
		Since `rets.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to the type we expect, which in this case is []*Car
	*/
	rets := m.Called()
	return rets.Get(0).([]*model.Car), rets.Error(1)
}
func (m *MockStore) GetCar(car *model.Car) (*model.Car, error) {
	/*
		Since `rets.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to the type we expect, which in this case is []*Car
	*/
	rets := m.Called()
	return rets.Get(0).(*model.Car), rets.Error(1)
}
func (m *MockStore) DeleteCar(car *model.Car) error {
	/*
		Since `rets.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to the type we expect, which in this case is []*Car
	*/
	rets := m.Called(car)
	return rets.Error(0)
}
func (m *MockStore) UpdateCar(car *model.Car) error {
	/*
		Since `rets.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to the type we expect, which in this case is []*Car
	*/
	rets := m.Called()
	return rets.Error(0)
}

func (m *MockStore) CreateDesc(desc *cache.Desc) error {
	/*
		When this method is called, `m.Called` records the call, and also
		returns the result that we pass to it (which you will see in the
		handler tests)
	*/
	rets := m.Called(desc)
	return rets.Error(0)
}
func (m *MockStore) DeleteDesc(desc *cache.Desc) error {
	/*
		Since `rets.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to the type we expect, which in this case is []*Car
	*/
	rets := m.Called(desc)
	return rets.Error(0)
}

func InitMockStore() *MockStore {
	/*
		Like the InitStore function we defined earlier, this function
		also initializes the store variable, but this time, it assigns
		a new MockStore instance to it, instead of an actual store
	*/
	s := new(MockStore)
	PackStore = s
	return s
}
