package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func tao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "tao")
}

func art(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "art")
}

func main() {
	http.HandleFunc("/tao", logging(tao))
	http.HandleFunc("/art", logging(art))

	http.ListenAndServe(":8080", nil)
}
