package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

	// Serve the ./public directory.
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.Handle("/", templ.Handler(showAll()))

	fmt.Println("Listening on :8081")
	http.ListenAndServe(":8081", nil)
}
