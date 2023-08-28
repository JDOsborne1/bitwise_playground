package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type test_struct2 struct {
	Title  string
	Header string
	Body   string
}

func landing_handle(w http.ResponseWriter, r *http.Request) {

	starting_struct := test_struct2{
		Title:  "Landing Page",
		Header: "Entrypoint to bitwise excellence",
		Body:   "Get ready to start your journey into Bitwise excellence. There are two sections, the definitons and the combinations",
	}

	comp_tmpl, err := template.ParseFS(files, "landing.html")

	if err != nil {
		fmt.Println("Issue with template: ", err)
	}

	err = comp_tmpl.Execute(w, starting_struct)
	if err != nil {
		fmt.Println("Issues with Template: ", err)
	}
}
