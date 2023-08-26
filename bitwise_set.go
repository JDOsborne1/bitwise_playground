package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var bitwise_post_form string = `<form hx-post="/bitwise_test_post">
	<input type="text" name="bitwise" value="1">
	<br>
	<input type="text" name="label" value="test">
	<br>
	<input type="submit" value="Submit">	
</form>`

func bitwise_post_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.New("Test")
		comp_tmpl := template.Must(tmpl.Parse(bitwise_post_form))
		
		err := comp_tmpl.Execute(w, nil)
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
	fmt.Println(bitwise_map)
	}
}


var bitwise_template string = `<div>
	{{range $key, $value := .Set}} 
	<p> Key: {{$key}} </p>
	<p> Label: {{$value}} </p>
	{{end}}
	</div>`
type bitwise_struct struct {
	Label string
	Bitwise int
}

type bitwise_set_container struct {
	Set map[int]string
}

func bitwise_handle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("Bitwise List")
	comp_tmpl := template.Must(tmpl.Parse(bitwise_template))	

	bitwise_tester := bitwise_set_container{
		Set: bitwise_map,
	}

	err := comp_tmpl.Execute(w, bitwise_tester)
	if err != nil {
		fmt.Println("Issues with Template: ", err)
	}
}

var bitwise_map map[int]string = make(map[int]string)