package main

import "fmt"

func main() {
	Descuento()
	Prestamo()
}

/* Ejercicio 1 - Descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola. */

func Descuento() {
	var precio float64
	var descuento float64

	// Pedir al usuario que ingrese el precio original del producto
	fmt.Print("Ingrese el precio del producto: ")
	fmt.Scan(&precio)

	// Pedir al usuario que ingrese el porcentaje de descuento
	fmt.Print("Ingrese el porcentaje de descuento: ")
	fmt.Scan(&descuento)

	// Calcular el valor del descuento
	valorDescuento := precio * (descuento / 100)

	// Calcular el precio final con descuento
	precioFinal := precio - valorDescuento

	// Imprimir el precio final
	fmt.Printf("El precio final después del descuento es: $%.2f\n", precioFinal)
}

/* Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes. */

func Prestamo() {
	var edad int
	var empleo bool
	var antiguedad int
	var sueldo float64

	// Pedir al usuario que ingrese la edad
	fmt.Print("Ingrese su edad: ")
	fmt.Scan(&edad)

	// Pedir al usuario que indique si está empleado (true/false)
	fmt.Print("¿Está usted actualmente empleado? (true/false): ")
	fmt.Scan(&empleo)

	// Pedir al usuario que ingrese la antigüedad en su trabajo (en años)
	fmt.Print("Ingrese su antigüedad en el trabajo (en años): ")
	fmt.Scan(&antiguedad)

	// Pedir al usuario que ingrese su sueldo
	fmt.Print("Ingrese su sueldo: ")
	fmt.Scan(&sueldo)

	// Verificar condiciones para otorgar el préstamo
	if edad > 22 && empleo && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Felicidades, usted califica para un préstamo sin interés.")
		} else {
			fmt.Println("Usted califica para un préstamo con interés.")
		}
	} else {
		fmt.Println("Lo sentimos, usted no califica para un préstamo.")
	}
}
