package concurrent_service

import "github.com/stuartshome/carpedia/model"

type User struct {
	Id   int
	Name string
}

type UserDetails struct {
	User User
	Car  model.Car
}
