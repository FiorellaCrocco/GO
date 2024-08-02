package main

import (
	"fmt"
	"os"
)

/* Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto con la información de productos comprados, separados por punto y coma (CSV). Este archivo debe tener el ID del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable. */

func main() {
	// Datos de productos comprados
	products := []struct {
		ID       string
		Precio   float64
		Cantidad int
	}{
		{"P001", 150.00, 2},
		{"P002", 80.50, 5},
		{"P003", 220.25, 3},
	}

	// Crear un archivo CSV
	file, err := os.Create("productos_comprados.csv")
	if err != nil {
		panic("no se pudo crear el archivo")
	}
	defer file.Close()

	// Escribir datos en el archivo CSV
	for _, product := range products {
		line := fmt.Sprintf("%s;%.2f;%d\n", product.ID, product.Precio, product.Cantidad)
		_, err := file.WriteString(line)
		if err != nil {
			panic("error al escribir en el archivo")
		}
	}

	fmt.Println("Archivo guardado exitosamente")
}
