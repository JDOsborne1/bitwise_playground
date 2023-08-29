package main

import (
	"html/template"
	"net/http"
)

type bitwise_combo struct {
	First_bit    int
	First_label  string
	Second_bit   int
	Second_label string
	Result_bit   int
	Result_label string
}

type set_of_combos struct {
	Set []bitwise_combo
}

var excluded_bits = make(map[int]bool)

// TODO - Allow and exclusions function, to prevent using output labels as input
func combinations_of_elements(_set_of_bitwise map[int]string, _excluded_bits map[int]bool) []bitwise_combo {
	var resp []bitwise_combo
	for bit, label := range _set_of_bitwise {
		if _excluded_bits[bit] {
			continue
		}
		for bit2, label2 := range _set_of_bitwise {
			if bit == bit2 {
				continue
			}
			if _excluded_bits[bit2] {
				continue
			}
			comb_bit := bit ^ bit2
			new_comb := bitwise_combo{
				First_bit:    bit,
				First_label:  label,
				Second_bit:   bit2,
				Second_label: label2,
				Result_bit:   comb_bit,
				Result_label: bitwise_map[comb_bit],
			}
			resp = append(resp, new_comb)
		}
	}
	return resp
}

func comb_handle(w http.ResponseWriter, r *http.Request) error {
	var err error

	out := combinations_of_elements(bitwise_map, excluded_bits)
	set := set_of_combos{
		Set: out,
	}

	comp_tmpl, err := template.ParseFS(files, "combinations.html")
	if err != nil {
		return err
	}

	err = comp_tmpl.Execute(w, set)
	if err != nil {
		return err
	}

	return nil
}
