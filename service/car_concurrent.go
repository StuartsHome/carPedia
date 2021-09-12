package service

type CarService interface {
	FetchData()
}

const (
	carServiceUrl = ""
)

type fetchCarDataService struct {
}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {

}
