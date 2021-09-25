package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/stuartshome/carpedia/cache"
)

type DescService interface {
	Validate(desc *cache.Desc) error
	Create(desc *cache.Desc) (*cache.Desc, error)
	FindAll() ([]*cache.Desc, error)
}

func NewDescService() DescService {
	return &service{}
}

func AddDescCache(desc *cache.Desc) {
	// Stores in mysqlDB - Don't think this is needed
	// result, err := descService.Create(&desc)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(errors.ServiceError{Message: "error saving description"})
	// }

	descCache.Set(strconv.FormatInt(desc.Id, 10), desc)
}

// First step, not currently working. More to do!
func GetDescByIDCache(w http.ResponseWriter, r *http.Request) {
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

func GetDescsCache() ([]*cache.Desc, error) {
	descs, err := descService.FindAll()
	if err != nil {
		return []*cache.Desc{}, err
	}
	return descs, nil
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
