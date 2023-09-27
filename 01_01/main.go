package main

import (
	"fmt"
)

func main() {
	var asunto string
	var mensaje string
	asunto = "Beneficio solo por abrir este correo."
	mensaje = "Hace mucho tiempo no nos visitás. Preparamos una sorpresa exclusiva para vos."
	correo := asunto + " " + mensaje
	fmt.Println(correo)
}
