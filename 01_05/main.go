package main

import (
	"fmt"
)

func main() {

	numerosRifa := []int{24, 33, 46, 55, 62}
	fmt.Printf("Números: %v\n", numerosRifa)

	ultimoElemento := numerosRifa[len(numerosRifa)-1]
	fmt.Printf("Número ganador: %v\n", ultimoElemento)
}
