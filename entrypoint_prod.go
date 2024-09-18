//go:build !dev && !static
// +build !dev,!static

package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/riweston/gotempl-cv/views/layout"
)

func entrypoint() {
	// Serve the ./public directory.
	log.Println("Running production server")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/", templ.Handler(layout.Layout()))
	http.Handle("/pdf", templ.Handler(layout.PDFLayout()))
	http.ListenAndServe(":8080", nil)
}
