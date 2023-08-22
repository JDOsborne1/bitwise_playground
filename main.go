package main

import (
	"log"
	"net/http"
)


func main() {
	handler := new(generic_handler)
	log.Fatal(http.ListenAndServe(":8081", handler))
}