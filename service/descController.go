package service

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"

	"github.com/stuartshome/carpedia/cache"
	"github.com/stuartshome/carpedia/errors"
	"github.com/stuartshome/carpedia/store"
)

var (
	/*
		// before
		// using this to create a new reference to DescService
		// this is bad as the controller should be injected into the service,
		// not instantiated by itself
		descService DescService = NewDescService()
	*/
	descService DescService
	descCache   cache.RedisCache
)

type service struct{}
type controller struct{}
type DescController interface {
	GetDescs(w http.ResponseWriter, r *http.Request)
	AddDesc(w http.ResponseWriter, r *http.Request)
	AddDescs(w http.ResponseWriter, r *http.Request)
	GetDescByID(w http.ResponseWriter, r *http.Request)
}

// constructor service
func NewDescController(service DescService, cache cache.RedisCache) DescController {
	descService = service
	descCache = cache
	return &controller{}
}
func (*controller) AddDesc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	var desc cache.Desc
	err := json.NewDecoder(r.Body).Decode(&desc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte(`{"error": "Error unmarshalling data"}`))
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error getting the posts"})
		return
	}
	if err := descService.Validate(&desc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
	}
	desc.Id = rand.Int63()

	// Add desc to cache, may need a method to check if desc is already in cache
	AddDescCache(&desc) // update the cache methods to return errors

	// Add desc to mysqlStore
	store.PackStore.CreateDesc(&desc)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&desc)
}
func (*controller) AddDescs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	var descs []*cache.Desc
	err := json.NewDecoder(r.Body).Decode(&descs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte(`{"error": "Error unmarshalling data"}`))
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error getting the posts"})
		return
	}
	// if err := descService.Validate(&descs); err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
	// }

	for _, desc := range descs {
		desc.Id = rand.Int63()
		// Add desc to cache, may need a method to check if desc is already in cache
		AddDescCache(desc)
		// Add desc to mysqlStore
		store.PackStore.CreateDesc(desc)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(descs)
}

func (*controller) GetDescByID(w http.ResponseWriter, r *http.Request) {
	// If desc is in the cache we'll retrieve it
	// if not, we'll use the sqlStoreDB

	w.Header().Set("Content-Type", "application/json")
	descID := strings.Split(r.URL.Path, "/")[2]
	var desc = descCache.Get(descID)
	// if no item in cache
	if desc == nil {
		// create function to get desc from sql db
		// if  != nil {

		// }
		descCache.Set(descID, nil)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(nil)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(nil)
	}
}

func (*controller) GetDescs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	descs, err := GetDescsCache()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error getting the descs"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&descs)
}
