package main

import (
	"html/template"
	"net/http"
)

type landing_data struct {
	Title  string
	Header string
	Body   string
}

func landing_handle(w http.ResponseWriter, r *http.Request) error {

	starting_struct := landing_data{
		Title:  "Landing Page",
		Header: "Entrypoint to bitwise excellence",
		Body:   "Get ready to start your journey into Bitwise excellence. There are two sections, the definitons and the combinations",
	}

	comp_tmpl, err := template.ParseFS(files, "landing.html")

	if err != nil {
		return err
	}

	err = comp_tmpl.Execute(w, starting_struct)
	if err != nil {
		return err
	}

	return nil
}
