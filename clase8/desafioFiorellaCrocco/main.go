package main

import (
	"desafioFiorellaCrocco/internal/tickets"
	"fmt"
	"sync"
)

// Requerimiento 4: Ejecutar al menos una vez cada requerimiento en la función main.
// Las ejecuciones deben realizarse de manera concurrente (utilizando diferentes goroutines).
func main() {

	var wg sync.WaitGroup

	// Requerimiento 1
	wg.Add(1)
	go func(destino string) {
		defer wg.Done()
		resultado, err := tickets.GetTotalTicketsDestination(destino)
		if err != nil {
			fmt.Printf("Error obteniendo total de tickets para %v: %v\n", destino, err)
		} else {
			fmt.Printf("Total de tickets a %v: %d\n", destino, resultado)
		}
	}("Brazil")

	// Requerimiento 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		madrugada, manana, tarde, noche, err := tickets.GetTravelPersonHours()
		if err != nil {
			fmt.Println("Error obteniendo horas de viaje: ", err)
		} else {
			fmt.Printf("Total de viajes del día: Mañana: %d, Tarde: %d, Noche: %d, Madrugada: %d\n", manana, tarde, noche, madrugada)
		}
	}()

	// Requerimiento 3
	wg.Add(1)
	go func(destino string) {
		defer wg.Done()
		resultado, err := tickets.PorcentajePasajerosSegunDestino(destino)
		if err != nil {
			fmt.Printf("Error obteniendo el porcentaje de pasajeros a %v: %v\n", destino, err)
		} else {
			fmt.Printf("Porcentaje de pasajeros a %v: %.2f%%\n", destino, resultado)
		}
	}("Brazil")

	// Esperar a que todas las goroutines terminen
	wg.Wait()
}
