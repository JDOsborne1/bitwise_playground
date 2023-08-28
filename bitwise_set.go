package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func bitwise_post_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		comp_tmpl, err := template.ParseFS(files, "bitwise_post_form.html")

		if err != nil {
			fmt.Println("Issues with Template: ", err)
		}

		err = comp_tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println("Issues with Template: ", err)
		}

	} else if r.Method == "POST" {

		r.ParseForm()
		bitwise_int, err := strconv.Atoi(r.FormValue("bitwise"))
		if err != nil {
			fmt.Println("Issues with Atoi: ", err)
		}
		bitwise_map[bitwise_int] = r.FormValue("label")
	}
}

type bitwise_struct struct {
	Label   string
	Bitwise int
}

type bitwise_set_container struct {
	Set map[int]string
}

func bitwise_handle(w http.ResponseWriter, r *http.Request) {
	comp_tmpl, err := template.ParseFS(files, "bitwise_list.html")

	if err != nil {
		fmt.Println("Issues with Template: ", err)
	}

	bitwise_tester := bitwise_set_container{
		Set: bitwise_map,
	}

	err = comp_tmpl.Execute(w, bitwise_tester)
	if err != nil {
		fmt.Println("Issues with Template: ", err)
	}
}

var bitwise_map map[int]string = make(map[int]string)
