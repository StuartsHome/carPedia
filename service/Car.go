package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stuartshome/carpedia/model"
	// "github.com/stuartshome/carpedia/store"
	aa "github.com/stuartshome/carpedia/store"
)

func GetCarHandler(w http.ResponseWriter, r *http.Request) {

	cars, err := aa.PackStore.GetCars()
	carListBytes, err := json.Marshal(cars)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(carListBytes)
}

func CreateBirdHandler(w http.ResponseWriter, r *http.Request) {
	car := model.Car{}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	car.Make = r.Form.Get("make")
	car.Model = r.Form.Get("model")
	err = aa.PackStore.CreateCar(&car)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/assets", http.StatusFound)
}
