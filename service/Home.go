package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stuartshome/carpedia/model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := JsonResponse()
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	w.Header().Set("Content-Type", "application/json")
	hmm, _ := json.Marshal(data)
	fmt.Fprintf(w, string(hmm))

	// json.NewEncoder(w).Encode(&data)

}

func JsonResponse() (model.Response, error) {
	data := model.Response{
		Cars: []model.Car{
			model.Car{
				Model: "Ford",
				Make:  "Fiesta",
				Reg:   1980,
			},
			model.Car{
				Model: "Ford",
				Make:  "Mondeo",
				Reg:   1995,
			},
		},
	}
	return data, nil

}
