package service

import

type user struct {
}

type UserService interface {
	GetDetails() user
}

// var _ =

func NewUserService() UserService {
	return &service{}
}

func (*service) GetDetails() user {
	// goroutine to get data from http://localhost:8100/car/1
	// goroutine to get data from http://localhost:8100/users/1

	// create car channel to get data
	// create user channel to get data

	client := http_client.New()

	return user{}
}
