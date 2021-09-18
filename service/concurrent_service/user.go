package concurrent_service

import (
	"encoding/json"
	"net/http"

	"github.com/stuartshome/carpedia/model"
	"github.com/stuartshome/carpedia/ref"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	user := User{
		Id:   100,
		Name: "Trevor",
	}

	json.NewEncoder(w).Encode(&user)
}
func CarHandler(w http.ResponseWriter, r *http.Request) {

	car := model.Car{
		Make:  "Dacia",
		Model: "Sandero",
		ID:    ref.Int(100),
	}

	json.NewEncoder(w).Encode(&car)
}
