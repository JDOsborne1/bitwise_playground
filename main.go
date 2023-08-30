package main

import (
	"embed"
	"log"
	"net/http"
	"github.com/pkg/browser"
)

//go:embed *.html
var files embed.FS

func main() {
	exit := make(chan int, 1)
	go func() {
		handler := new(handler)
		log.Fatal(http.ListenAndServe(":8081", handler))
		exit <- 1
	}()

	err := browser.OpenURL("http://localhost:8081")
	if err != nil {
		log.Print(err)
	}

	<- exit
}
