package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stuartshome/carpedia/http_client"
	"github.com/stuartshome/carpedia/logging"
	"github.com/stuartshome/carpedia/store"
)

func main() {
	logging.InitLogger()
	store.DbStartup()
	http_client.Router()
	fmt.Println("Service starting...")
}
