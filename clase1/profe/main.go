package main

import "fmt"

func main() {
	/* ------ EJERCICIO 1 ------ */

	palabra := "palabra" // := es para asignar un valor y que GO autodetecte el tipo de dato
	cant := len(palabra) // autodetecta que el tipo es INT y len() es el equivalente a lenght

	//	fmt.Println("La palabra tiene ", cant, " letras.") // Imprimo en consola la cantidad de letras que tiene la palabra

	//  %s indica string, %d indica decimal (entero), %f indica float, %v indica boolean
	//	fmt.Printf("La palabra %s tiene  %d letras.", palabra, cant)

	for i := 0; i < cant; i++ {
		fmt.Println(string(palabra[i]))
		// acá casteamos el valor a string, de no hacerlo den vez de imprimir las letras figurarian sus respectivos números asci
	}

	/* ------ EJERCICIO 2 ------ */

	var meses = []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Setiembre", "Octubre", "Noviembre", "Diciembre"}

	mes := 12

	// Las variables que sean declaradas con minúscula son scope local, pero si la definimos en mayúscula son globales.
	// var palabra (local) y var Palabra (global)

	// Las validaciones se plantean al inicio

	// Aqui se valida si el mes existe, de no ser así se muestra mensaje de error
	if mes < 1 || mes > 12 {
		fmt.Println("Error mes no existe")
	}

	// Aqui se valida si el mes es mayor o igual a 1 y menor o igual a 12
	// Esto implica que como el slice funciona como array la primer posición es 0
	// por este motivo se imprime con i-1
	if mes >= 1 && mes <= 12 {
		for i := 0; i <= len(meses); i++ {
			if i == mes {
				fmt.Println(meses[i-1])
			}
		}
	}

	// Aqui se valida si el mes es mayor a 1 y menor a 12
	// Esto implica que como el slice funciona como array la primer posición es 0 pero como ya no usamos el igual
	// no necesitamos restar haciendo i-1; se imprime directamente i
	/* ///// LO COMENTO PORQUE ES INNECESARIO AL TENER LA VALIDACIÓN INICIAL \\\\\  */
	/*
		if mes > 1 && mes < 12 {
			for i := 0; i <= len(meses); i++ {
				if i == mes {
					fmt.Println(meses[i])
				}
			}
		}
	*/
}
