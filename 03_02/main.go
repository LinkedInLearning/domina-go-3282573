package main

import (
	"03_02/operaciones"
	"fmt"
)

func aplicarOperacion(numero1, numero2 int, operacion func(int, int) int) int {
	return operacion(numero1, numero2)
}

func main() {
	var resultado int
	resultado = aplicarOperacion(3, 4, operaciones.Sumar)
	fmt.Println("El resultado de la suma es:", resultado)

	resultado = aplicarOperacion(2, 7, operaciones.Restar)
	fmt.Println("El resultado de la resta es:", resultado)

	resultado = aplicarOperacion(2, 3, operaciones.Multiplicar)
	fmt.Println("El resultado la multiplicación es:", resultado)
}
