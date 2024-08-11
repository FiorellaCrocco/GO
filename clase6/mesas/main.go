package main

import (
	"fmt"
	"os"
)

/* Ejercicio - Mayorista
Se necesita un software desarrollado en Go que imprima las estimaciones de lo que generarían, los productos de cada categoría en un mayorista, en ventas de productos para el hogar. Para eso se detallarán los pasos para conseguirlo:
1-Funcionalidad para generar un archivo CSV, llamado categorías.csv.
2-Cargar el archivo con las categorías. Ejemplo:
		001	Electrodomésticos	ListaProductos
		002	Muebles		        ListaProductos
		003	Herramientas	    ListaProductos
		004	Pinturas		    ListaProductos
		005	Aberturas		    ListaProductos
		006	Construcción	    ListaProductos
		007	Automotor 		    ListaProductos
	Elegir al menos tres categorías.
3-Generar para cada una de estas categorías los productos. Estos tendrán como información:
-Código
-Nombre
-PrecioActual
-CantidadActual
	Insertar al menos cuatro productos.

4-Generar otro archivo CSV, llamado estimaciones.csv. Este tendrá los resultados de la suma de todos los productos de cada una de las categorías.
5-Imprimir todos los estimativos por consola. Ejemplo:
Categoría			EstimativoPorCategoría
Construcción			60.700
Pinturas 				40.500
Aberturas				55.300
TotalEstimativo 		156.500  */

type Producto struct {
	Codigo         string
	Nombre         string
	PrecioActual   float64
	CantidadActual int
}

type Categoria struct {
	ID        string
	Nombre    string
	Productos []Producto
}

func crearArchivo(name string, data string) {
	file, err := os.Create(name)
	if err != nil {
		panic("no se pudo crear el archivo: " + err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		panic("error al escribir en el archivo: " + err.Error())
	}

	fmt.Println("Archivo", name, "guardado exitosamente")
}

func calcularEstimaciones(categorias []Categoria) (string, float64) {
	var total float64
	var data string

	for _, categoria := range categorias {
		var suma float64
		for _, producto := range categoria.Productos {
			suma += producto.PrecioActual * float64(producto.CantidadActual)
		}
		total += suma
		data += fmt.Sprintf("%s\t%.2f\n", categoria.Nombre, suma)
	}
	data += fmt.Sprintf("TotalEstimativo\t%.2f\n", total)

	return data, total
}

func main() {
	categorias := []Categoria{
		{
			ID:     "001",
			Nombre: "Electrodomésticos",
			Productos: []Producto{
				{"E1001", "Licuadora", 1200.50, 30},
				{"E1002", "Microondas", 4500.00, 15},
			},
		},
		{
			ID:     "002",
			Nombre: "Muebles",
			Productos: []Producto{
				{"M2001", "Silla", 800.00, 50},
				{"M2002", "Mesa", 1500.00, 20},
			},
		},
		{
			ID:     "003",
			Nombre: "Herramientas",
			Productos: []Producto{
				{"H3001", "Taladro", 2500.00, 25},
				{"H3002", "Destornillador", 200.00, 40},
			},
		},
	}

	// Creación y guardado del archivo CSV con los datos de las categorías y sus productos
	dataCSV, _ := calcularEstimaciones(categorias) // Calcula las estimaciones y descarta el valor total retornado usando '_'
	crearArchivo("categorías.csv", dataCSV)        // Crea y guarda el archivo 'categorías.csv' con los datos de estimaciones de ventas

	// Creación y guardado del archivo de estimaciones
	estimacionesCSV, total := calcularEstimaciones(categorias)           // Calcula las estimaciones y guarda el total
	fmt.Println(estimacionesCSV)                                         // Imprime las estimaciones en la consola
	fmt.Printf("Total estimado de todas las categorías: $%.2f\n", total) // Imprime el total de todas las categorías
	crearArchivo("estimaciones.csv", estimacionesCSV)                    // Crea y guarda el archivo con las estimaciones

}
