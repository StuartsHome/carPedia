package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stuartshome/carpedia/http_client"
	"github.com/stuartshome/carpedia/store"
)

func main() {
	err := godotenv.Load("script_config.env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", user, pass, dbname)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	store.InitStore(&store.DbStore{Db: db})

	fmt.Println("Service starting...")
	http_client.Router()
}
