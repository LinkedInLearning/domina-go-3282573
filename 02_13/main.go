package main

import (
	"fmt"
	"os"
)

func main() {
	archivo := "texto.txt"
	f, error := os.Create(archivo)
	if error != nil {
		fmt.Println(error)
		return
	}

	_, error = f.WriteString("La programación es el arte de crear soluciones con código.")
	if error != nil {
		fmt.Println(error)
		return
	}

	defer f.Close()
	fmt.Printf("El archivo %s ha sido exitosamente creado", archivo)
}
