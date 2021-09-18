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
}
