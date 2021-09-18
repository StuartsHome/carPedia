package concurrent_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stuartshome/carpedia/model"
)

type ConcurrentService interface {
	GetDetails() UserDetails
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

func (*ConService) GetDetails() UserDetails {
	// goroutine to get data from http://localhost:8100/usercar
	go carService.FetchData()
	// goroutine to get data from http://localhost:8100/user
	go userService.FetchData()
	// create car channel to get data
	// create user channel to get data

	car, _ := getCarData()
	user, _ := getUserData()

	// return UserDetails{user: user, car: car}
	return UserDetails{
		User: User{
			Id:   user.Id,
			Name: user.Name,
		},
		Car: model.Car{
			Make:  car.Make,
			Model: car.Model,
			ID:    car.ID,
		},
	}
}

func getCarData() (model.Car, error) {
	r1 := <-carDataChannel
	var car model.Car
	fmt.Println(json.NewDecoder(r1.Body))
	if err := json.NewDecoder(r1.Body).Decode(&car); err != nil {
		fmt.Println(err.Error())
		return car, err
	}
	return car, nil
}
func getUserData() (User, error) {
	r1 := <-userDataChannel
	var user User
	fmt.Println(json.NewDecoder(r1.Body))
	if err := json.NewDecoder(r1.Body).Decode(&user); err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}
