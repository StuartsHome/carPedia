package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stuartshome/carpedia/cache"
	"github.com/stuartshome/carpedia/model"
)

//go:generate go run "github.com/vektra/mockery/cmd/mockery" -case=underscore -outpkg mock_store -output ../mock/mock_store -name=Store
type Store interface {
	CreateCar(car *model.Car) error
	GetCars() ([]*model.Car, error)
	GetCar(*model.Car) (*model.Car, error)
	UpdateCar(car *model.Car) error
	DeleteCar(car *model.Car) error

	CreateDesc(desc *cache.Desc) error
	DeleteDesc(desc *cache.Desc) error
	GetAllDescs() ([]*cache.Desc, error)
}

func DbStartup() {
	err := godotenv.Load("script_config.env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")

	connString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", user, pass, dbname)
	// connString := fmt.Sprintf("%v:%v@tcp(docker.for.mac.localhost:3306)/%v", user, pass, dbname)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	InitStore(&DbStore{Db: db})
}
