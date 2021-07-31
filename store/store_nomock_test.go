package store

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/stuartshome/carpedia/model"
)

type StoreSuite struct {
	suite.Suite
	store *DbStore
	db    *sql.DB
}

// var _ StoreSuite = &StoreSuite{}

func (s *StoreSuite) SetupSuite() {
	/*
		The database connection is opened in the setup,
		and stored as an instance variable,
		as is the higher level "store", that wraps the "db"
	*/
	connString := "dbname<youre test db name> sslmode=disable"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &DbStore{Db: db}
}

func (s *StoreSuite) SetupTest() {
	/*
		We delete all entries from the table before each test runs, to
		ensure a consistent state before our tests run. In more complex apps,
		this is sometimes achieved in the form of migrations.
	*/
	_, err := s.db.Query("DELETE FROM cars")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	// Close the connection after all the tests
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateCar() {
	// Create a car
	s.store.CreateCar(&model.Car{
		Make:  "test make",
		Model: "test model",
	})

	// Query the db for the entry just created
	res, err := s.db.Query("SELECT COUNT(*) FROM cars WHERE make='test make' AND MODEL='test model'")
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	// Assert that there must be one entry with the properties of the car that was just inserted
	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1 car, go %d", count)
	}
}

func (s *StoreSuite) TestGetCar() {
	// Insert a sample car
	_, err := s.db.Query("INSERT INTO cars (make, model) VALUES('car', 'model')")
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the list of cars
	cars, err := s.store.GetCars()
	if err != nil {
		s.T().Fatal(err)
	}

	// Asser that the count of cars received must be 1
	nCars := len(cars)
	if nCars != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nCars)
	}

	// Assert that the details of the bird is the same as the inserted
	expectedCar := model.Car{Make: "car", Model: "model"}
	if *cars[0] != expectedCar {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedCar, *cars[0])
	}
}
