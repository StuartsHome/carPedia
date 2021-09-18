package concurrent_service

import (
	"fmt"

	"github.com/stuartshome/carpedia/http_client"
)

type CarService interface {
	FetchData()
}

const (
	carServiceUrl = "http://localhost:8100/caruser"
)

type fetchCarDataService struct {
}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http_client.New()

	// call the external api
	response, _ := client.Get(carServiceUrl)

	fmt.Println(response.StatusCode)
	carDataChannel <- response
	// return err

}
