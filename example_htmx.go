package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var ajax_tester *test_struct = &test_struct{Author: "John Doe", Body: "This is a test", Counter: 1}

func ajax_handle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("Test")
	comp_tmpl := template.Must(tmpl.Parse(`<div><p>{{.Author}}</p><p>{{.Body}} for the {{.Counter}}th time</p></div>`))

	ajax_tester.Counter += 1
	err := comp_tmpl.Execute(w, ajax_tester)
	if err != nil {
		fmt.Println("Issues with Template: ", err)
	}
}

type test_struct struct {
	Author  string
	Body    string
	Counter int
}

type test_struct2 struct {
	Title  string
	Header string
	Body   string
}

func landing_handle(w http.ResponseWriter, r *http.Request) {

	starting_struct := test_struct2{
		Title: "Landing Page", 
		Header: "Entrypoint to bitwise excellence", 
		Body: "Get ready to start your journey into Bitwise excellence. There are two sections, the definitons and the combinations",
	}

	layout_test := `<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title>
		<link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.16/dist/tailwind.min.css" rel="stylesheet">
		<script src="https://cdn.jsdelivr.net/npm/htmx.org@1.5.0/dist/htmx.min.js"></script>
	</head>
	<body class="bg-gray-100">
		<div class="container mx-auto p-8">
			<h1 class="text-3xl font-semibold">{{.Header}}</h1>
			<p class="my-4">{{.Body}}</p>
			</div>
			<div>
			<button hx-get="/bitwise_list" hx-target="#bitwise_list">List Available Bitwise</button>
			<div id="bitwise_list" class="mt-4"></div>
			</div>

<br>
<br>
<div>
			<button hx-get="/bitwise_test_post" hx-target="#bitwise_test_post">Add new bitwise</button>
			<div id="bitwise_test_post" class="mt-4"></div>
		</div>
	</body>
	</html>
	<!-- END -->
	
	`

	tmpl := template.New("Test")
	comp_tmpl := template.Must(tmpl.Parse(layout_test))

	err := comp_tmpl.Execute(w, starting_struct)
	if err != nil {
		fmt.Println("Issues with Template: ", err)
	}
}
