package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/bar", FooHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
