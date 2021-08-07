package store

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateCar(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected", err)
	}
	defer db.Close()

}
