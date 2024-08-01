package main

import "fmt"

func main() {
	// ejercicio1()
	ejercicio2()
}

//Ejercicio 1 - Listado de nombres
//Una profesora de la universidad quiere tener un listado con todos sus estudiantes.
//Es necesario crear una aplicación que contenga dicha lista. Sus estudiantes son:
//Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.
//Luego de dos clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar el código que escribiste inicialmente.
// El nombre de la nueva estudiante es Gabriela.

func ejercicio1() {
	var nombres = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println("Los nombres de los estudianres son", nombres)
	nombres = append(nombres, "Gabriela")
	fmt.Println("Agregada la nueva estudiante, los nombres son ", nombres)
}

//Ejercicio 2 - Qué edad tiene...
//Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente map, debemos imprimir la edad de Benjamin.
// var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
//Por otro lado, también es necesario:
//Saber cuántos de sus empleados son mayores de 21 años.
//Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
//Eliminar a Pedro del map.

func ejercicio2() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	nombreBuscar := "Benjamin"

	// Obtener y imprimir la edad del empleado
	if edad, existe := employees[nombreBuscar]; existe {
		fmt.Printf("La edad de %s es %d años.\n", nombreBuscar, edad)
	} else {
		fmt.Println("Empleado no encontrado.")
	}

	// Inicialización de la variable para contar empleados mayores de 21 años
	var mayoresEdad int = 0

	// Recorrer el mapa para contar los empleados mayores de 21 años
	for _, edad := range employees {

		if edad > 21 {
			mayoresEdad++
		}

		// Imprimir el total de empleados mayores de 21 años
		fmt.Printf("Número de empleados mayores de 21 años: %d\n", mayoresEdad)

	}

	// Agregar un nuevo empleado llamado Federico con edad de 25 años
	employees["Federico"] = 25

	// Eliminar a Pedro del mapa
	delete(employees, "Pedro")

}
