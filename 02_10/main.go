package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

		} else {
			http.Error(w, "Recurso no encontrado.", http.StatusNotFound)
			return
		}
	})

	fmt.Println("La aplicación se está ejecutando en 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
