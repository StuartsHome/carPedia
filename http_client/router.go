package http_client

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stuartshome/carpedia/service"
)

func Router() {
	r := newRouter()

	http.ListenAndServe(":8100", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	//Html page
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/home", service.HomeHandler)
	r.HandleFunc("/car", service.GetCarHandler).Methods("GET")
	r.HandleFunc("/car", service.CreateBirdHandler).Methods("POST")

	//Healthcheck
	r.HandleFunc("/health", service.HealthCheckHandler)

	return r
}
