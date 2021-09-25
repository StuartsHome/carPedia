package store

import (
	"database/sql"
	"fmt"

	"github.com/stuartshome/carpedia/cache"
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

func (store *DbStore) GetCar(car *model.Car) (*model.Car, error) {
	row := store.Db.QueryRow("SELECT make, model FROM cars WHERE make = ?", car.Make)
	result := model.Car{}
	switch err := row.Scan(&result.Make, &result.Model); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(result.Make, result.Model)
	default:
		panic(err)
	}
	return &result, nil
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

func (store *DbStore) DeleteCar(car *model.Car) error {
	_, err := store.Db.Exec("DELETE FROM cars WHERE make=$1", car.Make)
	return err
}
func (store *DbStore) UpdateCar(car *model.Car) error {
	_, err :=
		store.Db.Exec("UPDATE cars SET make=$1, model=$2 WHERE id=$3", car.Make, car.Model)
	return err
}

// description methods
func (store *DbStore) CreateDesc(desc *cache.Desc) error {
	_, err := store.Db.Query("INSERT INTO descs (title, text) VALUES (?,?)", desc.Text, desc.Title)
	logging.Logf("storing in db: %v & %v", desc.Title, desc.Text)
	return err
}
func (store *DbStore) DeleteDesc(desc *cache.Desc) error {
	_, err := store.Db.Exec("DELETE FROM descs WHERE id=$1", desc.Id)
	return err
}

// Package level variable that will be available for use
// throughout our application code.
var PackStore Store

// We call this method to initialise the store
func InitStore(s Store) {
	PackStore = s
}
