package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func combinations_of_elements(_set_of_bitwise map[int]string) string {
	var resp strings.Builder
	for bit, label := range _set_of_bitwise {
		for bit2, label2 := range _set_of_bitwise {
			if bit == bit2 {
				continue
			}
			comb_bit := bit ^ bit2 
			if bitwise_map[comb_bit] == "" {
				resp.WriteString("Combination of " + label + " and " + label2 + "produces an unlabelled value: " + fmt.Sprint(comb_bit) + "\n")
				continue
			}
			
			resp.WriteString("Combination of " + label + " and " + label2 + "produces: " + bitwise_map[comb_bit] + "\n")
		}
	}
	return resp.String()
}

func comb_handle(w http.ResponseWriter, r *http.Request) {
	out := combinations_of_elements(bitwise_map)
	fmt.Fprint(w, out)
}

func main() {

	bitwise_map[1] = "test"
	bitwise_map[2] = "test2"

	fmt.Println(bitwise_map)

	handler := new(generic_handler)
	log.Fatal(http.ListenAndServe(":8081", handler))
}