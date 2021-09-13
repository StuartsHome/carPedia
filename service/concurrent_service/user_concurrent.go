package service

import (
	"fmt"

	"github.com/stuartshome/carpedia/http_client"
)

type UserService interface {
	FetchData() error
}

const (
	userServiceUrl = "http://google.com"
)

type fetchUserDataService struct {
}

func NewUserService() UserService {
	return &fetchUserDataService{}
}

func (*fetchUserDataService) FetchData() error {
	client := http_client.New()

	// call the external api
	response, err := client.Get(userServiceUrl)
	if err != nil {
		return err
	}
	userDataChannel <- response
	fmt.Println(response.StatusCode)
	return err

}
