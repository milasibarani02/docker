package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>“We are one”</h1>")
	fmt.Println("Someone hit me!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :80")
	http.ListenAndServe(":80", nil)
}