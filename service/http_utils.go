package service

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

func HtmlResponse(response http.ResponseWriter, value interface{}) {
	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(http.StatusOK)
	fp, _ := filepath.Abs(".")
	t := template.Must(template.ParseFiles(fp + "/assets/index.html"))
	fmt.Println(fp + "/assets/index.html")
	t.Execute(response, value)
}
