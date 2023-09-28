package main

import (
	"fmt"
	"os"
)

func main() {
	ruta_archivo := "texto.txt"
	datos, error := os.ReadFile(ruta_archivo)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println("Contenido del archivo de texto:", string(datos))
}
