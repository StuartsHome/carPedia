package router

import (
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/stuartshome/carpedia/cache"
	"github.com/stuartshome/carpedia/service"
	"github.com/stuartshome/carpedia/service/concurrent_service"
)

var (
	descService    service.DescService    = service.NewDescService()
	descCache      cache.RedisCache       = cache.NewRedisCache("localhost:6379", 0, 10)
	descController service.DescController = service.NewDescController(descService, descCache)
	// descRepository reposi
)

func NewRouter() *mux.Router {
	// config := settings.Get()
	r := mux.NewRouter()

	// use defaultMux for debug routes for pprof
	r.PathPrefix("/debug/").Handler(http.DefaultServeMux)

	// html page
	staticFileDirectory := http.Dir("assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.HandleFunc("/home", service.CreateCarHandler).Methods("GET")
	r.PathPrefix("/home").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/home", service.CreateCarHandler).Methods("POST")
	r.HandleFunc("/car", service.GetCarHandler).Methods("GET")
	r.HandleFunc("/car/{id:[0-9]+}", service.UpdateCar).Methods("PUT")
	r.HandleFunc("/car/{id:[0-9]+}", service.DeleteCar).Methods("DELETE")
	r.HandleFunc("/results", service.GetSingleCarHandler).Methods("POST")

	// description handlers used by cache
	r.HandleFunc("/desc", descController.GetDesc).Methods("GET")
	r.HandleFunc("/desc/{id}", descController.GetDescByID).Methods("GET")
	r.HandleFunc("/desc", descController.AddDesc).Methods("POST")

	// user
	r.HandleFunc("/user", concurrent_service.UserHandler).Methods("GET")
	r.HandleFunc("/caruser", concurrent_service.CarHandler).Methods("GET")

	// html page for all cars from db
	staticFileDirectoryAll := http.Dir("assets/")
	staticFileHandlerAll := http.StripPrefix("/assets/", http.FileServer(staticFileDirectoryAll))
	r.HandleFunc("/allcars", service.AllCarsHandler).Methods("GET")
	r.PathPrefix("/allcars").Handler(staticFileHandlerAll).Methods("GET")

	// healthcheck
	r.HandleFunc("/health", service.HealthCheckHandler).Methods("GET")

	// http.ListenAndServe(*config.HttpSettings.ListenAddress, r)
	return r

}

func Router() {

	// config := settings.Get()
	r := NewRouter()

	// http.ListenAndServe(":8100", r)
	log.Println(http.ListenAndServe(":8100", r))

}
