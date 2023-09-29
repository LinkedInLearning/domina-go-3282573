package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/imagen", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			archivo := "volcan.jpeg"

		} else {
			http.Error(w, "Recurso no encontrado.", http.StatusNotFound)
			return
		}
	})

	fmt.Println("La aplicación se está ejecutando en 127.0.0.1:8080/imagen")
	http.ListenAndServe(":8080", nil)
}
