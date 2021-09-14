package service

import (
	"fmt"

	"github.com/stuartshome/carpedia/http_client"
)

type UserService interface {
	FetchData()
}

const (
	userServiceUrl = "http://google.com"
)

type fetchUserDataService struct {
}

func NewUserService() UserService {
	return &fetchUserDataService{}
}

func (*fetchUserDataService) FetchData() {
	client := http_client.New()

	// call the external api
	response, _ := client.Get(userServiceUrl)

	fmt.Println(response.StatusCode)
	userDataChannel <- response

}
