package main

import (
	"fmt"
)

func main() {
	mensajeBienvenida()
	mensajeBienvenida("Jason")
	mensajeBienvenida("Jessica", "Marcos", "Jonathan")
}
func mensajeBienvenida(invitado ...string) {
	for _, nombre := range invitado {
		fmt.Printf("¡Bienvenido, %s\n", nombre+"!")
	}
}
