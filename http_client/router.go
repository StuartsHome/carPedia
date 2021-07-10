package http_client

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stuartshome/carpedia/service"
)

func Router() {

	r := mux.NewRouter()
	r.HandleFunc("/home", service.HomeHandler)
	http.ListenAndServe(":8100", r)
}
