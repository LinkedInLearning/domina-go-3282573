package main

import (
	"fmt"
	"net/http"
)

type Cliente struct {
	Nombre string `json:"Nombre"`
	Edad   int    `json:"Edad"`
}

func main() {
	http.HandleFunc("/cliente", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

		} else {
			http.Error(w, "Recurso no encontrado.", http.StatusNotFound)
			return
		}
	})

	fmt.Println("La aplicación se está ejecutando en http://127.0.0.1:8080/cliente")
	http.ListenAndServe(":8080", nil)
}
