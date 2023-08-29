package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed *.html
var files embed.FS

func main() {
	excluded_bits[3] = true
	excluded_bits[5] = true
	handler := new(handler)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
