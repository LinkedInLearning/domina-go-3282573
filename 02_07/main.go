package main

import "fmt"

func main() {
	numeros := [...]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(numeros)

	for indice, numero := range numeros {
		fmt.Println("Índice: ", indice, "Número:", numero)
	}
}
