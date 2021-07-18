package store

import (
	"database/sql"

	"github.com/stuartshome/carpedia/model"
)

//go:generate go run "github.com/vektra/mockery/cmd/mockery" -case=underscore -outpkg mock_store -output ../mock/mock_store -name=Store

type Store interface {
	CreateCar(car *model.Car) error
	GetCars() ([]*model.Car, error)
}

type dbStore struct {
	db *sql.DB
}

// The dbStore struct will implement the Store interface
// var _ Store = &dbStore{}

func (store *dbStore) CreateCar(car *model.Car) error {
	_, err := store.db.Query("INSERT INTO cars(make, model) VALUES ($1, $2)", car.Make, car.Model)
	return err
}

func (store *dbStore) GetCars() ([]*model.Car, error) {
	rows, err := store.db.Query("SELECT make, model from cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []*model.Car{}
	for rows.Next() {
		car := &model.Car{}
		if err := rows.Scan(&car.Make, &car.Model); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

// Package level variable that will be available for use
// throughout our application code.
var PackStore Store

// We call this method to initialise the store
func InitStore(s Store) {
	PackStore = s
}
