package model

type Response struct {
	Cars []Car `json:"cars"`
}

type MultipleCars struct {
	Cars []Car `json:"cars"`
}

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Reg   *int   `json:"reg"`
	ID    *int   `json:"id"`
}

var Cars []Car
