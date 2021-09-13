package service

import (
	"fmt"

	"github.com/stuartshome/carpedia/http_client"
)

type CarService interface {
	FetchData() error
}

const (
	carServiceUrl = "http://google.com"
)

type fetchCarDataService struct {
}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() error {
	client := http_client.New()

	// call the external api
	response, err := client.Get(carServiceUrl)
	if err != nil {
		return err
	}

	carDataChannel <- response
	fmt.Println(response.StatusCode)
	return err

}
