package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/stuartshome/carpedia/model"
	"github.com/stuartshome/carpedia/settings"
)

//go:generate go run "github.com/vektra/mockery/cmd/mockery" -case=underscore -outpkg mock_store -output ../mock/mock_store -name=Store
type Store interface {
	CreateCar(car *model.Car) error
	GetCars() ([]*model.Car, error)
	GetCar(*model.Car) (*model.Car, error)
	UpdateCar(car *model.Car) error
	DeleteCar(car *model.Car) error
}

func DbStartup() {
	config := settings.Get()
	// err := godotenv.Load("script_config.env")
	// if err != nil {
	// 	log.Fatalf("error loading .env file")
	// }

	// user := os.Getenv("MYSQL_USER")
	// pass := os.Getenv("PASS")
	// dbname := os.Getenv("MYSQL_DATABASE")
	dbname := config.DatabaseCreds.DBName
	user := config.DatabaseCreds.DBUser
	pass := config.DatabaseCreds.DBPass

	// connString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", user, pass, dbname)
	// connString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", user, pass, dbname)
	connString := fmt.Sprintf("%v:%v@tcp(docker.for.mac.localhost:3306)/%v", user, pass, dbname)
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
