package main

import (
	"fmt"

	"github.com/stuartshome/carpedia/http_client"
)

func main() {
	fmt.Println("Service starting...")
	http_client.Router()
}
