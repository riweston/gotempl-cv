package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/riweston/riweston/views/layout"
)

func main() {
	// Serve the ./public directory.
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/", templ.Handler(layout.Layout()))
	http.ListenAndServe("localhost:8080", nil)
}
