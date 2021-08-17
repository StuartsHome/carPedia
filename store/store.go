package store

import (
	"database/sql"

	"github.com/stuartshome/carpedia/logging"
	"github.com/stuartshome/carpedia/model"
)

type DbStore struct {
	Db *sql.DB
}

// The DbStore struct will implement the Store interface
var _ Store = &DbStore{}

func (store *DbStore) CreateCar(car *model.Car) error {
	_, err := store.Db.Query("INSERT INTO cars(make, model) VALUES (?,?)", car.Make, car.Model)
	logging.Logf("storing in db: %v & %v", car.Make, car.Model)
	return err
}

func (store *DbStore) GetCars() ([]*model.Car, error) {
	rows, err := store.Db.Query("SELECT make, model from cars")
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

func (store *DbStore) GetCar(car *model.Car) (*model.Car, error) {
	rows, err := store.Db.Query("SELECT make, model from cars WHERE make, model = (?, ?)", car.Make, car.Model)
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
	return car, nil
}

func (store *DbStore) DeleteCar(car *model.Car) error {
	return nil
}
func (store *DbStore) UpdateCar(car *model.Car) error {
	return nil
}

// Package level variable that will be available for use
// throughout our application code.
var PackStore Store

// We call this method to initialise the store
func InitStore(s Store) {
	PackStore = s
}
