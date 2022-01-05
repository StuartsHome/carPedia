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
	descCache.Set(strconv.FormatInt(desc.Id, 10), desc)
}

// TODO: First step, not currently working. More to do!
func GetDescByIDCache(w http.ResponseWriter, r *http.Request) {
	// If desc is in the cache we'll retrieve it
	// else, we'll use the sqlStoreDB.

	w.Header().Set("Content-Type", "application/json")
	descID := strings.Split(r.URL.Path, "/")[2]
	var desc = descCache.Get(descID)
	// if no item in cache
	if desc == nil {
		// TODO: create function to get desc from sql db.

		// If id not in cache, set id in cache.
		descCache.Set(descID, nil)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(desc)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(desc)
	}
}

func GetDescsCache() ([]*cache.Desc, error) {
	descs, err := descCache.GetAll()
	if err != nil {
		return []*cache.Desc{}, err
	}
	return descs, nil
}

func (*service) Validate(desc *cache.Desc) error {
	if desc == nil {
		err := fmt.Errorf("the desc is empty")
		return err
	}
	if desc.Title == "" {
		err := fmt.Errorf("the desc title is empty")
		return err
	}
	return nil

}

func (*service) Create(desc *cache.Desc) (*cache.Desc, error) {
	desc.Id = rand.Int63()

	return nil, nil
}

// TODO:
func (*service) FindAll() ([]*cache.Desc, error) {
	return nil, nil
}
