package main

import (
	"fmt"
	"os" // Lectura de archivos.
)

func main() {

	/* EJERCICIO 1 */
	var salary int
	//salary = 20
	salary = 200

	if salary < 150 {
		panic(&MyError{})
	}

	fmt.Println("Debe pagar impuesto")

	/* EJERCICIO 2 */

	fmt.Println("\nIniciando...")

	readFile("customers.txt")

	fmt.Println("\nEjecución finalizada")

}

/* Ejercicio 1 - Impuestos de salario
En la función main, definir una variable salary y asignarle un valor de tipo “int”. Luego, crear un error personalizado con un struct que implemente Error() con el mensaje “Error: el salario ingresado no alcanza el mínimo imponible” y lanzarlo en caso de que salary sea menor a 150.000. De lo contrario, imprimir por consola el mensaje “Debe pagar impuesto”. */

// Definición de la estructura MyError para el error personalizado
type MyError struct{}

// Implementación del método Error() para MyError
func (e *MyError) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}

/* Ejercicio 2 - Datos de clientes
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo TXT.
Desarrollar el código necesario para leer los datos de un archivo llamado “customers.txt”. Sin embargo, debemos tener en cuenta que la empresa no nos ha pasado el archivo a leer por el programa. Dado que no contamos con el archivo necesario, se obtendrá un error. En tal caso, el programa deberá arrojar un panic al intentar leer un archivo que no existe, mostrando el mensaje “El archivo indicado no fue encontrado o está dañado”. Más allá de eso, deberá siempre imprimirse por consola “Ejecución finalizada”. */

func readFile(name string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.Open(name)
	if err != nil {
		panic("\nEl archivo indicado no fue encontrado o está dañado")
	}
}
