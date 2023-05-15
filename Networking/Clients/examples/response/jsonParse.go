package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct { // defines struct Status, which contains expected elements from server response
	Message string
	Status string
}

func main() {
	res, err := http.Post( //sends POST request
		"http://IP:PORT/ping",
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	// Decode response body
	if err := json.NewDocer(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	// Query Status struct normally - accessing exported data types status and message
	log.Printf("%s -> %s\n", status.Status, status.Message)
}