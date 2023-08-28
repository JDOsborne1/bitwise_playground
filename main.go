package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed *.html
var files embed.FS

func main() {

	handler := new(generic_handler)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
