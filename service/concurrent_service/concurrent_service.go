package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stuartshome/carpedia/model"
)

type ConcurrentService interface {
	GetDetails() User
}

type concurrentService struct {
}

var (
	carService      CarService  = NewCarService()
	userService     UserService = NewUserService()
	userDataChannel             = make(chan *http.Response)
	carDataChannel              = make(chan *http.Response)
)

func NewConcurrentService() ConcurrentService {
	return &concurrentService{}
}

func (*concurrentService) GetDetails() User {
	// goroutine to get data from http://localhost:8100/car/1
	go carService.FetchData()
	// goroutine to get data from http://localhost:8100/users/1
	go userService.FetchData()
	// create car channel to get data
	// create user channel to get data

	car, _ := getCarData()
	user, _ := getUserData()

	return User{}
}

func getCarData() (model.Car, error) {
	r1 := <-carDataChannel
	var car model.Car
	if err := json.NewDecoder(r1.Body).Decode(&car); err != nil {
		fmt.Println(err.Error())
		return car, err
	}
	return car, nil
}
func getUserData() (User, error) {
	r1 := <-carDataChannel
	var user User
	if err := json.NewDecoder(r1.Body).Decode(&user); err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}
