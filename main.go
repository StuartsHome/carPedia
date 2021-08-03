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

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")

	// connString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", user, pass, dbname)
	connString := fmt.Sprintf("%v:%v@tcp(docker.for.mac.localhost:3306)/%v", user, pass, dbname)
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
