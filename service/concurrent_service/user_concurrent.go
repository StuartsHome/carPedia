package concurrent_service

import (
	"fmt"

	"github.com/stuartshome/carpedia/http_client"
)

type UserService interface {
	FetchData()
}

const (
	userServiceUrl = "http://localhost:8100/user"
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
