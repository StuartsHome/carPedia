package store

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stuartshome/carpedia/mock/mock_store"
	"github.com/stuartshome/carpedia/model"
	aa "github.com/stuartshome/carpedia/service"
)

func TestGetCarsHandler(t *testing.T) {
	// Initialise the mock store
	var mockStore mock_store.Store = mock_store.Store{}
	// var mockClient mock_client.Client = mock_client.Client{}

	mockStore.On("GetCars").Return([]*model.Car{
		{
			Make:  "citroen",
			Model: "c3",
		},
		{
			Make:  "citroen",
			Model: "c3",
		},
	}, nil).Once()

	// json := `{}`
	// r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	// response := &http.Response{Body: r, StatusCode: 200}
	// mockClient.On("GET", "").Return(response)

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(aa.GetCarHandler)
	hf.ServeHTTP(recorder, req)

	expected := model.Car{Make: "citroen", Model: "c3"}
	c := []model.Car{}
	json.NewDecoder(recorder.Body).Decode(&c)

	actual := c[0]
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
	mockStore.AssertExpectations(t)
}
