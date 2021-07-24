package http_client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuartshome/carpedia/service"
)

func TestHealthCheckHandler(t *testing.T) {
	// A new http request, this request is passed to the handler
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
