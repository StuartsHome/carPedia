package service

import (
	"encoding/json"
	"math/rand"
	"net/http"
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
)

type service struct{}
type controller struct{}

type DescService interface {
	Validate(desc *cache.Desc) error
	Create(desc *cache.Desc) (*cache.Desc, error)
	FindAll() ([]*cache.Desc, error)
}

func NewDescService() DescService {
	return &service{}
}

type DescController interface {
	GetDesc(w http.ResponseWriter, r *http.Request)
	AddDesc(w http.ResponseWriter, r *http.Request)
}

// constructor service
func NewDescController(service DescService) DescController {
	descService = service
	return &controller{}
}

func (*controller) AddDesc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	result, err := descService.Create(&desc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error saving description"})

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

// First step, not currently working. More to do!
func (*controller) GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postID := strings.Split(r.URL.Path, "/")[2]
	post, err := descService.FindById(postID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "no posts found!"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}

}

func (*controller) GetDesc(w http.ResponseWriter, r *http.Request) {

}

func (*service) Validate(desc *cache.Desc) error {
	if desc == nil {
		err := errors.New("The desc is empty")
		return err
	}
	if desc.Title == "" {
		err := errors.New("The desc title is empty")
		return err
	}
	return nil

}

func (*service) Create(desc *cache.Desc) (*cache.Desc, error) {
	desc.Id = rand.Int63()
	return repo.Save(desc)
}
func (*service) FindAll() ([]*cache.Desc, error) {
	return repo.FindAll()

}
