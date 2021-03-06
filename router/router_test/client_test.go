package router_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuartshome/carpedia/http_client"
	"github.com/stuartshome/carpedia/router"
	"github.com/stuartshome/carpedia/service"
)

// type dbcreds struct {
// 	dbusername string "test"
// 	dbpass     string "123456"
// 	dbname     string "car_pedia"
// }

func TestHealthCheckHandler(t *testing.T) {
	// A new http request, this request is passed to the handler
	// This only tests the handler, not the route to the handler
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a http recorder. This recorder will
	// act as the target of our http request.
	// This will accept the request that we make!
	recorder := httptest.NewRecorder()

	// This is the handler func we want to test
	hf := http.HandlerFunc(service.HealthCheckHandler)

	// Serve the HTTP request to the recorder
	// This line executes the handler
	hf.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is correct
	expected := "{\"status\":\"OK\",\"host\":\"aa\"}\n"
	actual := recorder.Body.String()
	assert.Equal(t, expected, actual)
}

func TestRouter(t *testing.T) {
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	r := router.NewRouter()
	mockServer := httptest.NewServer(r)
	response, err := http.Get(mockServer.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", response.StatusCode)
	}

	defer response.Body.Close()
	// Read the body into bytes
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	// Convert the bytes to a string
	responseString := string(b)
	expected := "{\"status\":\"OK\",\"host\":\"aa\"}\n"

	assert.Equal(t, expected, responseString)
}

func TestBadRoute(t *testing.T) {
	r := router.NewRouter()
	mockServer := httptest.NewServer(r)
	response, err := http.Post(mockServer.URL+"/health", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 405 (method not allowed)
	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", response.StatusCode)
	}

	defer response.Body.Close()
	// Read the body into bytes
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	// Convert the bytes to a string
	responseString := string(b)
	expected := ""

	assert.Equal(t, expected, responseString)
}

func TestInvalidRequestGet(t *testing.T) {
	apiClient := http_client.New()

	actual, err := apiClient.Get("http://google.com")
	assert.Empty(t, err)
	assert.Equal(t, actual.StatusCode, 200)
}
