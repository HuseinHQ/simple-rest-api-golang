package controllers

import (
	"io"
	"net/http"
)

var apiUrl string = "https://jsonplaceholder.typicode.com/posts"

func GetPosts(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(apiUrl)
	if err != nil {
		http.Error(w, "Error making request", http.StatusBadRequest)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, "Unexpected status code", response.StatusCode)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
