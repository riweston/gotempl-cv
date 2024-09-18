//+build static
//go:build static
// +build static

package main

import (
	"context"
	"log"
	"os"

	"github.com/riweston/gotempl-cv/views/layout"
)

func entrypoint() {
	log.Println("Creating index.html")
	indexHTML, err := os.Create("index.html")
	if err != nil {
		log.Println("Failed to create index.html")
		log.Fatal(err)
	}
	defer indexHTML.Close()

	log.Println("Rendering index.html")
	if err := layout.Layout().Render(context.Background(), indexHTML); err != nil {
		log.Println("Failed to render index.html")
		log.Fatal(err)
	}
	log.Println("Rendered index.html")
}
