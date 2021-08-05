package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stuartshome/carpedia/http_client"
	"github.com/stuartshome/carpedia/logging"
	"github.com/stuartshome/carpedia/store"
)

func main() {

	store.DbStartup()
	logging.InitLogger()
	fmt.Println("Service starting...")
	http_client.Router()
}
