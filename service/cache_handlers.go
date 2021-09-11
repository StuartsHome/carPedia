package service

import (
	"encoding/json"
	"net/http"
	"strings"
)

// First step, not currently working. More to do!
func GetPostByID(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	postID := strings.Split(r.URL.Path, "/")[2]
	post, err := postService.FindById(postID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "No posts found!"})
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	}

}
