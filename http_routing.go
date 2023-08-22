package main

import (
	"net/http"
	"path"
	"strings"
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