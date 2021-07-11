package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/stuartshome/carpedia/model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := JsonResponse()
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	w.Header().Set("Content-Type", "application/json")
	hmm, _ := json.Marshal(data)
	fmt.Fprintf(w, string(hmm))

	file, err := os.OpenFile("./tmp/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "prefix", log.LstdFlags)
	logger.Println("text to append")
	logger.Println("more text to append")
	write := io.MultiWriter(os.Stdout, file)
	log.SetOutput(write)
	log.Printf("Home handler hit")

	// json.NewEncoder(w).Encode(&data)

}

func JsonResponse() (model.Response, error) {
	data := model.Response{
		Cars: []model.Car{
			model.Car{
				Model: "Ford",
				Make:  "Fiesta",
				Reg:   1980,
			},
			model.Car{
				Model: "Ford",
				Make:  "Mondeo",
				Reg:   1995,
			},
		},
	}
	return data, nil

}
