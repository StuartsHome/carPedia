package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuartshome/carpedia/model"
	"github.com/stuartshome/carpedia/service"
	"github.com/stuartshome/carpedia/store"
)

func newCreateCarForm() *url.Values {
	form := url.Values{}
	form.Set("make", "ford")
	form.Set("model", "mustang")
	return &form
}

func TestCreateCar(t *testing.T) {
	os.Chdir("../../")
	// Given
	mockStore := store.InitMockStore()
	testData := model.Car{
		Make:  "ford",
		Model: "mustang",
	}

	// When
	mockStore.On("CreateCar", &testData).Return(nil)

	form := newCreateCarForm()
	req2, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(service.CreateCarHandler)
	hf.ServeHTTP(recorder, req2)

	// Then
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

func TestCreateCarHandlerError(t *testing.T) {
	// Given
	mockStore := store.InitMockStore()
	testData := model.Car{
		Make:  "ford",
		Model: "mustang",
	}

	// When
	mockStore.On("CreateCar", &testData).Return(errors.New(""))

	form := newCreateCarForm()
	req2, _ := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(service.CreateCarHandler)
	hf.ServeHTTP(recorder, req2)

	// Then
	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetCarsHandler(t *testing.T) {
	// Given
	// Initialise the mock store
	mockStore := store.InitMockStore()
	mockStore.On("GetCars").Return([]*model.Car{
		{
			Make:  "Citroen",
			Model: "c3",
		},
	}, nil).Once()

	// When
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(service.GetCarHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Then
	expected := model.Car{Make: "Citroen", Model: "c3"}
	c := []model.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&c)
	if err != nil {
		t.Fatal(err)
	}
	actual := c[0]
	assert.Equal(t, expected, actual)
}

func TestGetCarsHandlerError(t *testing.T) {
	// Given
	mockStore := store.InitMockStore()
	mockStore.On("GetCars").Return([]*model.Car{}, errors.New("Error")).Once()

	// When
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(service.GetCarHandler)
	hf.ServeHTTP(recorder, req)

	// Then
	// resp := recorder.Result()
	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
