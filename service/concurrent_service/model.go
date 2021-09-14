package service

import "github.com/stuartshome/carpedia/model"

type User struct {
	Id int
}

type UserDetails struct {
	user User
	car  model.Car
}
