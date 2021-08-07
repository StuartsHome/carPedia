package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuartshome/carpedia/model"
	"github.com/stuartshome/carpedia/store"
)

func newCreateCarForm() *url.Values {
	form := url.Values{}
	form.Set("make", "ford")
	form.Set("model", "mustang")
	return &form
}

func TestCreateCar(t *testing.T) {
	// Given
	// var mockStore = new(mock_store.Store)
	// var mockStore mock_store.Store = mock_store.Store{}
	// mockStore2 := store.InitStore(&mockStore)
	mockStore := store.InitMockStore()

	testData := model.Car{
		Make:  "ford",
		Model: "mustang",
	}

	// json := `{"car": {"make": "ford", "model": "mustang"}}`
	// r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	// req, err := http.NewRequest("POST", "", r)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	mockStore.On("CreateCar", &testData).Return(nil)

	form := newCreateCarForm()
	req2, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(CreateCarHandler)
	hf.ServeHTTP(recorder, req2)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)

}

func TestGetCarsHandler(t *testing.T) {
	// Initialise the mock store
	// var mockStore mock_store.Store = mock_store.Store{}
	mockStore := store.InitMockStore()
	mockStore.On("GetCars").Return([]*model.Car{
		{
			Make:  "Citroen",
			Model: "c3",
		},
	}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(GetCarHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := model.Car{Make: "Citroen", Model: "c3"}

	c := []model.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&c)
	if err != nil {
		t.Fatal(err)
	}
	actual := c[0]
	assert.Equal(t, expected, actual)

}

// json := `{}`
// r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
// response := &http.Response{Body: r, StatusCode: 200}
// mockClient.On("GET", "").Return(response)

// req, err := http.NewRequest("GET", "", nil)

// if err != nil {
// 	t.Fatal(err)
// }
// hf := http.HandlerFunc(aa.GetCarHandler)
// hf.ServeHTTP(recorder, req)

// recorder := httptest.NewRecorder()
// expected := model.Car{Make: "citroen", Model: "c3"}
// c := []model.Car{}
// json.NewDecoder(recorder.Body).Decode(&c)

// actual := c[0]
// if actual != expected {
// 	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
// }
// mockStore.AssertExpectations(t)
