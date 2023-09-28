package main

import "fmt"

func main() {
	clientes := map[string]string{
		"Ronny":     "Cartago",
		"Lincy":     "Heredia",
		"Alejandra": "Guanacaste",
	}

	ciudad, existe := clientes["Ronny"]
	if existe {
		fmt.Println("Ronny vive en", ciudad)
	} else {
		fmt.Println("Ronny no está en el mapa.")
	}

	ciudad, existe = clientes["José"]
	if existe {
		fmt.Println("José vive en", ciudad)
	} else {
		fmt.Println("José no está en el mapa.")
	}
}
