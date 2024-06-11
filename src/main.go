package main

import (
	"log"
	"net/http"
	"strconv"
)

const (
	SERVER_PORT = 9001
)

func CheckRateLimit(w http.ResponseWriter, req *http.Request) {
	log.Printf("Rate limiter called")
}

func main() {
	http.HandleFunc("/rate-limit", CheckRateLimit)
	log.Printf("run server on port %d", SERVER_PORT)
	err := http.ListenAndServe(":"+strconv.Itoa(SERVER_PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
