//+build dev
//go:build dev
// +build dev

package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/riweston/gotempl-cv/views/layout"
)

func entrypoint() {
	// Serve the ./public directory.
	log.Println("Running development server")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/", templ.Handler(layout.Layout()))
	http.Handle("/pdf", templ.Handler(layout.PDFLayout()))
	http.ListenAndServe("localhost:8080", nil)
}
