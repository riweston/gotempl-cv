//go:build !dev
// +build !dev

package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/riweston/gotempl-cv/views/layout"
)

func httpHandler() {
	// Serve the ./public directory.
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/", templ.Handler(layout.Layout()))
	http.ListenAndServe(":8080", nil)
}
