package main

import (
	"errors"
	"log"
)

const (
	a = "Categoría A"
	b = "Categoría B"
	c = "Categoría C"
)

func main() {
	// Se debe enviar la cantidad de minutos como argumento a la función calculo
	aFun, err := calculo(a)
	if err != nil {
		log.Fatal(err)
	}

	// Ejecutamos la función devuelta para calcular el salario basado en 120 minutos trabajados
	salarioA := aFun(120)
	log.Printf("El salario calculado es: $%d", salarioA)

	// Se debe enviar la cantidad de minutos como argumento a la función calculo
	bFun, err := calculo(b)
	if err != nil {
		log.Fatal(err)
	}

	// Ejecutamos la función devuelta para calcular el salario basado en 120 minutos trabajados
	salarioB := bFun(120)
	log.Printf("El salario calculado es: $%d", salarioB)

	// Se debe enviar la cantidad de minutos como argumento a la función calculo
	cFun, err := calculo(c)
	if err != nil {
		log.Fatal(err)
	}

	// Ejecutamos la función devuelta para calcular el salario basado en 120 minutos trabajados
	salarioC := cFun(120)
	log.Printf("El salario calculado es: $%d", salarioC)

	errorFun, err := calculo("x")
	if err != nil {
		log.Fatal(err)
	}

	// Ejecutamos la función devuelta para calcular el salario basado en 120 minutos trabajados
	errorV := errorFun(120)
	log.Printf("El salario calculado es: $%d", errorV)

}

/* Ejercicio 1 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C: su salario es de $1.000 por hora.
Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

func calculo(categoria string) (func(int) int, error) {
	switch categoria {
	case a:
		return sueldoA, nil
	case b:
		return sueldoB, nil
	case c:
		return sueldoC, nil
	default:
		return nil, errors.New("Categoría no válida")
	}
}

func sueldoA(min int) int {
	salario := 3000
	salarioBase := min / 60 * salario
	extraMens := salarioBase * 50 / 100
	salarioFinal := salarioBase + extraMens
	return salarioFinal
}

func sueldoB(min int) int {
	salario := 1500
	salarioBase := min / 60 * salario
	extraMens := salarioBase * 20 / 100
	salarioFinal := salarioBase + extraMens
	return salarioFinal
}

func sueldoC(min int) int {
	salario := 1000
	salarioBase := min / 60 * salario
	return salarioBase
}
