package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func bitwise_post_handler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {

		comp_tmpl, err := template.ParseFS(files, "bitwise_post_form.html")

		if err != nil {
			return err
		}

		err = comp_tmpl.Execute(w, nil)
		if err != nil {
			return err
		}

	} else if r.Method == "POST" {

		r.ParseForm()
		bitwise_int, err := strconv.Atoi(r.FormValue("bitwise"))
		if err != nil {
			return err
		}
		bitwise_map[bitwise_int] = r.FormValue("label")
	}
	return nil
}

type bitwise_struct struct {
	Label   string
	Bitwise int
}

type bitwise_set_container struct {
	Set map[int]string
}

func bitwise_handle(w http.ResponseWriter, r *http.Request) error {
	comp_tmpl, err := template.ParseFS(files, "bitwise_list.html")

	if err != nil {
		return err
	}

	bitwise_tester := bitwise_set_container{
		Set: bitwise_map,
	}

	err = comp_tmpl.Execute(w, bitwise_tester)
	if err != nil {
		return err
	}
	return nil
}

var bitwise_map map[int]string = make(map[int]string)
