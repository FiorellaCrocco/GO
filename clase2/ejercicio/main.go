package main

import (
	"errors"
	"fmt"
)

func main() {

	/* EJERCICIO 1 */

	var sueldo float64 = 150000

	impuesto := impuestosSueldos(sueldo)

	fmt.Print("Impuesto a descontar: ", impuesto)

	/* EJERCICIO 2 */

	//promedio, err := promedioAlumnos(2, 4, -5, 8, 2)
	promedio, err := promedioAlumnos(2, 4, 5, 8, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio del alumno es:", promedio)
	}

}

/* Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo.
Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario, teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).
*/

func impuestosSueldos(sueldo float64) float64 {

	var impuesto float64

	if sueldo > 50000 && sueldo <= 149000 {
		impuesto = (sueldo / 100) * 17
	} else if sueldo >= 150000 {
		impuesto = (sueldo / 100) * 27
	}

	return impuesto
}

/* Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/

func promedioAlumnos(notas ...int) (int, error) {

	var resultado int
	var promedio int

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("Hay numeros negativos en las calificaciones.")
		} else {
			resultado += value
		}
	}

	promedio = resultado / len(notas)

	return promedio, nil

}
