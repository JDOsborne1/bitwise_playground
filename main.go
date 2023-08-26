package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {

	bitwise_map[1] = "test"
	bitwise_map[2] = "test2"

	fmt.Println(bitwise_map)

	handler := new(generic_handler)
	log.Fatal(http.ListenAndServe(":8081", handler))
}