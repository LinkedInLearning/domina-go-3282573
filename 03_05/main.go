package main

import (
	"fmt"
	"net/http"

	"github.com/tealeg/xlsx"
)

func generarXLSX() (*xlsx.File, error) {
	archivo := xlsx.NewFile()
	hoja, err := archivo.AddSheet("Hoja 1")
	if err != nil {
		return nil, err
	}

	fila := hoja.AddRow()
	celda := fila.AddCell()
	celda.Value = "Ana"
	celda = fila.AddCell()
	celda.Value = "Heredia"

	return archivo, nil
}

func main() {
	http.HandleFunc("/descarga", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			archivoXLSX, error := generarXLSX()
			if error != nil {
				http.Error(w, "Error creaando el archivo XLSX", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Disposition", "attachment; filename=archivo.xlsx")
			w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
			error = archivoXLSX.Write(w)
			if error != nil {
				http.Error(w, "Error escribiendo el arhivo XLSX", http.StatusInternalServerError)
				return
			}

		} else {
			http.Error(w, "Recurso no encontrado.", http.StatusNotFound)
			return
		}
	})

	fmt.Println("La aplicación se está ejecutando en 127.0.0.1:8080/descarga")
	http.ListenAndServe(":8080", nil)
}
