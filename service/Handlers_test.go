package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuartshome/carpedia/mock/mock_store"
	"github.com/stuartshome/carpedia/model"
	// "github.com/stuartshome/carpedia/store"
)

func TestGetCarsHandler(t *testing.T) {
	// Initialise the mock store
	var mockStore mock_store.Store = mock_store.Store{}
	// var mockClient mock_client.Client = mock_client.Client{}

	mockStore.On("GetCars").Return([]*model.Car{
		{Make: "citroen",
			Model: "c3",
			Reg:   123},
	}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	// var store12 Store
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(GetCarHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := []model.Car{
		{"citroen", "c3", 123},
		// model.Car{"ford", "fiesta", 234},
	}
	c := []model.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&c)
	if err != nil {
		t.Fatal(err)
	}
	actual := c[0]
	assert.Equal(t, expected, actual)

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
}
