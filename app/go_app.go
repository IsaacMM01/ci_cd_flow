package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola desde mi proyecto")
	})

	http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Saludos, espero esten teniendo un buen día")
	})

	http.HandleFunc("/otros", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Estos son otros asuntos")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
