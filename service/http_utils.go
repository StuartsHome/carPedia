package service

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

func DisplayHTMLResponse(response http.ResponseWriter, value interface{}) {
	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(http.StatusOK)
	fp, _ := filepath.Abs(".")
	t := template.Must(template.ParseFiles(fp + "/assets/index.html"))
	fmt.Println(fp + "/assets/index.html")
	t.Execute(response, value)
}
func DisplayAllHTMLResponse(response http.ResponseWriter, value interface{}) {
	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(http.StatusOK)
	fp, _ := filepath.Abs(".")
	t := template.Must(template.ParseFiles(fp + "/assets/index.html"))
	fmt.Println(fp + "/assets/index.html")
	t.Execute(response, value)
}

// ctrl + a or s 	- cursor line start/end
// ctrl + d 		- delete char forward
// shift + alt + a - multiple line comment
// cmd + p
// opt + cmd + l or r (move tabs l or right)
// opt + l or r
