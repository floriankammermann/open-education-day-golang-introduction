package main

import (
	"fmt"
	"net/http"
)

func greet(greeting string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, greeting)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/greeting", greet("This is a greeter in different languages"))
	mux.Handle("/greeting/CH", greet("This is a greeter Switzerland"))
	mux.Handle("/greeting/CH/de", greet("Grüezi"))
	mux.Handle("/greeting/CH/fr", greet("Bonjour"))
	mux.Handle("/greeting/CH/it", greet("Ciao"))
	http.ListenAndServe(":8088", mux)
}
