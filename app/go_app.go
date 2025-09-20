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
		fmt.Fprintf(w, "Hola, esta es la ruta /saludo")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
