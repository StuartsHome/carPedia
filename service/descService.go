package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/stuartshome/carpedia/cache"
	"github.com/stuartshome/carpedia/errors"
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
	GetDesc(w http.ResponseWriter, r *http.Request)
	AddDesc(w http.ResponseWriter, r *http.Request)
	GetDescByID(w http.ResponseWriter, r *http.Request)
}

// constructor service
func NewDescController(service DescService, cache cache.RedisCache) DescController {
	descService = service
	descCache = cache
	return &controller{}
}

type DescService interface {
	Validate(desc *cache.Desc) error
	Create(desc *cache.Desc) (*cache.Desc, error)
	FindAll() ([]*cache.Desc, error)
}

func NewDescService() DescService {
	return &service{}
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
	// Stores in mysqlDB - Don't think this is needed
	// result, err := descService.Create(&desc)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(errors.ServiceError{Message: "error saving description"})
	// }

	descCache.Set(strconv.FormatInt(desc.Id, 10), &desc)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&desc)
}

// First step, not currently working. More to do!
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

func (*controller) GetDesc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	descs, err := descService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error getting the descs"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&descs)
}

func (*service) Validate(desc *cache.Desc) error {
	if desc == nil {
		// err := errors.New("The desc is empty")
		err := fmt.Errorf("The desc is empty")
		return err
	}
	if desc.Title == "" {
		// err := errors.New("The desc title is empty")
		err := fmt.Errorf("The desc title is empty")
		return err
	}
	return nil

}

func (*service) Create(desc *cache.Desc) (*cache.Desc, error) {
	desc.Id = rand.Int63()

	return nil, nil
}
func (*service) FindAll() ([]*cache.Desc, error) {
	// return repo.FindAll()
	return nil, nil

}
