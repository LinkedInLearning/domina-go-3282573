package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "La programación es el arte de crear soluciones con código"
	textoBuscar := "arte"
	
	if(strings.Contains(frase, textoBuscar)){
		fmt.Printf("La palabra %s existe en la frase.\n", textoBuscar)
	} else {
		fmt.Printf("La palabra %s no existe en la frase.\n", textoBuscar)
	}
}
