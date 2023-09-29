package main

import (
	"fmt"
	"net/http"
)

type Usuario struct {
	Usuario      string
	ClaveSecreta string
}

func inicioHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("¡Bienvenido(a)!"))
}

func main() {
	mux := http.NewServeMux()

	fmt.Println("La aplicación se está ejecutando en 127.0.0.1:8080/inicio")
	http.ListenAndServe(":8080", mux)
}
