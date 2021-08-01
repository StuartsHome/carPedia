package reader

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
The io Reader interface represents an entity from which you can
read a stream of bytes.

*/

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

}

func dummy() (*http.Response, error) {
	// use a request to populate NewReader to return an io Reader
	r := strings.NewReader("my request")
	resp, err := http.Post("http://foo.bar",
		"application/x-www-form-urlencoded", r)

	return resp, err

	// Read directly from a byte stream
	// Not a very common use-case
	/*
		a := strings.NewReader("abcde")

		buf := make([]byte, 4)
		for {
			n, err := r.Read(buf)
			fmt.Println(n, err, buf[:n])
			if err == io.EOF {
				break
			}
		}
	*/
}

func TestDummy(t *testing.T) {
	// create a request to pass to the handler
	reader := strings.NewReader("number=2")
	req, _ := http.NewRequest("POST", "/my_url", reader)

	// Records the response from the handler
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(w, req)

	// check the status code
	if status := w.Code; status != http.StatusOK {
		t.Error()
	}

	// check the response body is what we expect
	expected := `{Status: ok}`
	if w.Body.String() != expected {
		t.Error()
	}
}
