package http_client

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stuartshome/carpedia/service"
)

func newRouter() *mux.Router {
	// config := settings.Get()
	r := mux.NewRouter()

	//Html page
	staticFileDirectory := http.Dir("assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.HandleFunc("/home", service.CreateCarHandler).Methods("GET")
	r.PathPrefix("/home").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/car", service.GetCarHandler).Methods("GET")
	r.HandleFunc("/car", service.CreateCarHandler).Methods("POST")

	// Html page for all cars from db
	staticFileDirectoryAll := http.Dir("assets/")
	staticFileHandlerAll := http.StripPrefix("/assets/", http.FileServer(staticFileDirectoryAll))
	r.HandleFunc("/allcars", service.AllCarsHandler).Methods("GET")
	r.PathPrefix("/allcars").Handler(staticFileHandlerAll).Methods("GET")

	//Healthcheck
	r.HandleFunc("/health", service.HealthCheckHandler).Methods("GET")

	// http.ListenAndServe(*config.HttpSettings.ListenAddress, r)
	return r

}

func Router() {

	// config := settings.Get()
	r := newRouter()

	http.ListenAndServe(":8100", r)

}
