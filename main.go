package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type bitwise_combo struct {
	first_bit int
	First_label string
	second_bit int
	Second_label string
	result_bit int
	Result_label string
}

type set_of_combos struct {
	Set []bitwise_combo
}

func combinations_of_elements(_set_of_bitwise map[int]string) []bitwise_combo {
	var resp []bitwise_combo
	for bit, label := range _set_of_bitwise {
		for bit2, label2 := range _set_of_bitwise {
			if bit == bit2 {
				continue
			}
			comb_bit := bit ^ bit2 
			new_comb := bitwise_combo{
				first_bit: bit,
				First_label: label,
				second_bit: bit2,
				Second_label: label2,
				result_bit: comb_bit,
				Result_label: bitwise_map[comb_bit],
			}
			resp = append(resp, new_comb)
		}
	}
	return resp
}

func (b bitwise_combo) First() string {
	return b.First_label
}

var combo_template string = `<div>
{{range $val := .Set}}
<p> Combine {{$val.First_label}} and {{$val.Second_label}} into: {{$val.Result_label}} </p>
{{end}}
</div>`

func comb_handle(w http.ResponseWriter, r *http.Request) {
	out := combinations_of_elements(bitwise_map)
	set := set_of_combos{
		Set: out,
	}
	tmpl := template.New("Combos")
	comp_tmpl := template.Must(tmpl.Parse(combo_template))

	comp_tmpl.Execute(w, set)

}

func main() {

	bitwise_map[1] = "test"
	bitwise_map[2] = "test2"

	fmt.Println(bitwise_map)

	handler := new(generic_handler)
	log.Fatal(http.ListenAndServe(":8081", handler))
}