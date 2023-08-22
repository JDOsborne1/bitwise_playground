package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"text/template"
)

// shift_path splits the given path into the first segment (head) and
// the rest (tail). For example, "/foo/bar/baz" gives "foo", "/bar/baz".
func shift_path(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type generic_handler struct {
}

// ServeHTTP is a custom replacement for the default handler from the http package.
// It makes use of the shift path strategy to walk through the route and then delegate
// the processing to the appropriate sub handler or sub strategy.
func (generic_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)
 if head == "test" {
		test_handle(w, r)
	} else if head == "ajax-example" {
		ajax_handle(w, r)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

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
	Author string
	Body string
	Counter int
}

type test_struct2 struct {
	Title string
	Header string
	Body string
}

func test_handle(w http.ResponseWriter, r *http.Request) {

	starting_struct := test_struct2{Title: "Test", Header: "Test Header", Body: "Test Body"}

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
			<button class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded" hx-get="/ajax-example" hx-target="#ajax-result">Load Content via AJAX</button>
			<div id="ajax-result" class="mt-4"></div>
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

func main() {
	handler := new(generic_handler)
	log.Fatal(http.ListenAndServe(":8081", handler))
}