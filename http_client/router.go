package http_client

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stuartshome/carpedia/service"
	"github.com/stuartshome/carpedia/settings"
	// "github.com/stuartshome/carpedia/settings"
)

func newRouter() *mux.Router {
	config := settings.Get()
	r := mux.NewRouter()

	//Html page
	staticFileDirectory := http.Dir("assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/home", service.HomeHandler)
	r.HandleFunc("/car", service.GetCarHandler).Methods("GET")
	r.HandleFunc("/car", service.CreateCarHandler).Methods("POST")

	//Healthcheck
	r.HandleFunc("/health", service.HealthCheckHandler).Methods("GET")

	http.ListenAndServe(*config.HttpSettings.ListenAddress, r)
	return r

}

func Router() {

	// config := settings.Get()
	r := newRouter()

	http.ListenAndServe(":8100", r)

}
