package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ConcurrentService interface {
	GetDetails() (interface{}, interface{})
}

type ConService struct {
}

var (
	carService      CarService  = NewCarService()
	userService     UserService = NewUserService()
	userDataChannel             = make(chan *http.Response)
	carDataChannel              = make(chan *http.Response)
)

func NewConcurrentService() ConcurrentService {
	return &ConService{}
}

func (*ConService) GetDetails() (interface{}, interface{}) {
	// goroutine to get data from http://localhost:8100/car/1
	go carService.FetchData()
	// goroutine to get data from http://localhost:8100/users/1
	go userService.FetchData()
	// create car channel to get data
	// create user channel to get data

	car, _ := getCarData()
	user, _ := getUserData()

	// return UserDetails{user: user, car: car}
	return car, user
}

func getCarData() (interface{}, error) {
	r1 := <-carDataChannel
	// var car model.Car
	var car interface{}
	fmt.Println(json.NewDecoder(r1.Body))
	if err := json.NewDecoder(r1.Body).Decode(&car); err != nil {
		fmt.Println(err.Error())
		return car, err
	}
	return car, nil
}
func getUserData() (interface{}, error) {
	r1 := <-carDataChannel
	var user interface{}
	fmt.Println(json.NewDecoder(r1.Body))
	if err := json.NewDecoder(r1.Body).Decode(&user); err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}
