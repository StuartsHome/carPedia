package service

import (
	"encoding/json"
	"net/http"

	"github.com/stuartshome/carpedia/cache"
	"github.com/stuartshome/carpedia/errors"
)

func (*controller) AddDescCache(w http.ResponseWriter, r *http.Request) {
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

}
