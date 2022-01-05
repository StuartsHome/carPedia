package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	HostHeaderKey = "carPedia"
	HostName      = "aa"
)

type healthcheck struct {
	Status string `json:"status"`
	Host   string `json:"host"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	var statusCode int
	switch r.Method {
	case http.MethodPost:
		statusCode = http.StatusCreated
	case http.MethodDelete:
		statusCode = http.StatusNoContent
	case http.MethodGet, http.MethodPut:
		fallthrough
	default:
		statusCode = http.StatusOK
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	temp := healthcheck{
		Status: "OK",
		Host:   HostName,
	}

	// _ = json.NewEncoder(w).Encode(temp)
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(&temp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// write the json payload to the body
	w.Write(buffer.Bytes())
}

// TODO: add new rate struct, create new func to process the available rate limits
func RateLimitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "rate limit endpoint hit")
}
