package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola desde mi proyecto, test case 1")
	})

	http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola, esta es la ruta /saludo")
	})

	http.HandleFunc("/otros", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola, esta es la ruta para otros asuntos")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
