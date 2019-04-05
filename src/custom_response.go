package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a greeter")
	})

	http.HandleFunc("/greeting/CH", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a greeter for Switzerland")
	})

	http.HandleFunc("/greeting/CH/de", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Grüezi")
	})

	http.HandleFunc("/greeting/CH/fr", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bonjour")
	})

	http.HandleFunc("/greeting/CH/it", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ciao")
	})

	http.ListenAndServe(":8088", nil)
}
