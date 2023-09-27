package main

import (
	"fmt"
	"strconv"
)

func main() {

	var numeroEntero int
	var numeroFlotante float64
	var numeroString string

	// Convertir de int a float
	numeroEntero = 12
	numeroFlotante = float64(numeroEntero)
	fmt.Printf("%f\n", numeroFlotante)

	// Convertir de float a int
	numeroFlotante = 32.9
	numeroEntero = int(numeroFlotante)
	fmt.Printf("%d\n", numeroEntero)

	// Convertir de int a string
	numeroEntero = 39
	numeroString = strconv.Itoa(numeroEntero)
	fmt.Printf("%q\n", numeroString)

	// Convertir de string a int
	numeroString = "50"
	numeroEntero, err := strconv.Atoi(numeroString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", numeroEntero)
}
