package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stuartshome/carpedia/logging"
	"github.com/stuartshome/carpedia/router"
	"github.com/stuartshome/carpedia/store"
)

func main() {

	logging.InitLogger()
	store.DbStartup()
	router.Router()
	fmt.Println("Service starting...")

	/*

		DESCS table in DB doesn't exist, we need to fix this to get all the descs in the db so we know what to query the cahce for when we make
		GetDescByID() line 103


		struct in lowercase, factory creator with New that returns an interface in uppercase of the struct.
		The struct is dependency injected into subsequent funcs
	*/
}
