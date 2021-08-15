package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stuartshome/carpedia/model"
	"github.com/stuartshome/carpedia/store"
)

func GetCarHandler(w http.ResponseWriter, r *http.Request) {
	// Call store GetCars()
	// Marshal the value into JSON
	// Write the JSON to the Response

	cars, err := store.PackStore.GetCars()
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carListBytes, err := json.Marshal(cars)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(carListBytes)
}

func CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	// The HTML information is sent to us in HTML form
	// ParseForm parses the data
	// Query the form for the two fields
	// Append the values to the list of cars
	// Redirect the user to the original HTML page
	car := model.Car{}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	car.Make = r.Form.Get("make")
	car.Model = r.Form.Get("model")
	err = store.PackStore.CreateCar(&car)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}
