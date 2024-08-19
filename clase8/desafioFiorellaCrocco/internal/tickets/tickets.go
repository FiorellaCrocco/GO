package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// /* Planteo
// Una aerolínea pequeña tiene un sistema de reservas de pasajes a diferentes países. Este
// retorna un archivo con la información de los pasajes sacados en las últimas 24 horas. La
// aerolínea necesita un programa para extraer información de las ventas del día y, así,
// analizar las tendencias de compra.
// El archivo en cuestión es del tipo valores separados por coma (CSV), donde los campos
// están compuestos por: id, nombre, email, país de destino, hora del vuelo y precio. */

// Id 1 del tickets.csv: 1,Tait Mc Caughan,tmc0@scribd.com,Finland,17:11,785
type Ticket struct {
	id           int
	nombre       string
	email        string
	pais_destino string
	hora_vuelo   string
	precio       float64
}

// Abre el archivo CSV y retorna una lista de tickets o un error
func obtenerTickets() ([]Ticket, error) {
	file, err := os.Open("tickets.csv")
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rawRecords, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo CSV: %v", err)
	}

	var tickets []Ticket
	for _, record := range rawRecords[1:] {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Printf("error al convertir el id: %v, ignorando el registro\n", err)
			continue
		}
		precio, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			fmt.Printf("error al convertir el precio: %v, ignorando el registro\n", err)
			continue
		}
		tickets = append(tickets, Ticket{
			id:           id,
			nombre:       record[1],
			email:        record[2],
			pais_destino: record[3],
			hora_vuelo:   record[4],
			precio:       precio,
		})
	}

	return tickets, nil
}

// Requerimiento 1: Una función que calcule cuántas personas viajan a un país determinado.
func GetTotalTicketsDestination(destination string) (int, error) {
	tickets, err := obtenerTickets()
	if err != nil {
		return 0, err
	}

	cantidadTickets := 0
	for _, ticket := range tickets {
		if ticket.pais_destino == destination {
			cantidadTickets++
		}
	}

	return cantidadTickets, nil
}

// Requerimiento 2: Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).
func GetTravelPersonHours() (int, int, int, int, error) {
	tickets, err := obtenerTickets()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	var viajanEnMadrugada, viajanDeMañana, viajanDeTarde, viajanDeNoche int
	for _, ticket := range tickets {
		hora, err := strconv.Atoi(strings.Split(ticket.hora_vuelo, ":")[0])
		if err != nil {
			fmt.Printf("error al convertir la hora del vuelo: %v, ignorando el registro\n", err)
			continue
		}

		switch {
		case hora >= 0 && hora <= 6:
			viajanEnMadrugada++
		case hora >= 7 && hora <= 12:
			viajanDeMañana++
		case hora >= 13 && hora <= 19:
			viajanDeTarde++
		case hora >= 20 && hora <= 23:
			viajanDeNoche++
		}
	}

	return viajanEnMadrugada, viajanDeMañana, viajanDeTarde, viajanDeNoche, nil
}

// Requerimiento 3: Calcular el porcentaje de personas que viajan a un país determinado en un día.
func PorcentajePasajerosSegunDestino(destination string) (float64, error) {
	tickets, err := obtenerTickets()
	if err != nil {
		return 0, err
	}

	if len(tickets) == 0 {
		panic("no hay registros de tickets para calcular el porcentaje")
	}

	viajerosDestino := 0
	for _, ticket := range tickets {
		if ticket.pais_destino == destination {
			viajerosDestino++
		}
	}

	porcentaje := (float64(viajerosDestino) / float64(len(tickets))) * 100
	return porcentaje, nil
}
