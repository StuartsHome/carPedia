package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func DisplayHTMLResponse(response http.ResponseWriter, value interface{}) {
	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(http.StatusOK)
	fp, _ := filepath.Abs(".")
	t := template.Must(template.ParseFiles(fp + "/assets/index.html"))
	// fmt.Println(fp + "/assets/index.html")
	t.Execute(response, value)
}
func DisplayAllHTMLResponse(response http.ResponseWriter, value interface{}) {
	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(http.StatusOK)
	fp, _ := filepath.Abs(".")
	funcs := template.FuncMap{"add": add}
	t := template.Must(template.New("allcars.html").Funcs(funcs).ParseFiles(fp + "/assets/allcars.html"))
	// fmt.Println(fp + "/assets/allCars.html")
	t.Execute(response, value)
}

// log an error and return it in the specified HTTP response
func ErrorResponse(response *http.ResponseWriter, statusCode int, err error) {
	message := err.Error()
	log.Print(fmt.Sprintf("HTTP error returned: %s - %s", http.StatusText(statusCode), message))
	http.Error(*response, message, statusCode)
}

func FromJSON(data io.Reader, value interface{}) error {
	decoder := json.NewDecoder(data)
	decoder.DisallowUnknownFields()
	return decoder.Decode(value)
}

func ToJSON(value interface{}) []byte {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)

	err := encoder.Encode(value)
	if err != nil {
		log.Panicf("Failure converting entity to JSON: %v", err)
	}
	return buffer.Bytes()
}

// add function to count rows in allcars html template
func add(x, y int) int {
	return x + y
}
