package main

import "fmt"

/* Calcular precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento. Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda, y para optimizar la velocidad, requieren que el cálculo de la sumatoria se realice en paralelo mediante tres goroutines.
Se requieren tres estructuras:
- Productos: nombre, precio, cantidad.
- Servicios: nombre, precio, minutos trabajados.
- Mantenimiento: nombre, precio. */

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

/* Y se requieren tres funciones:
- Sumar Productos: recibe un array de producto y devuelve el precio total (precio por cantidad).
- Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio por media hora trabajada. Si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
- Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total. */

func sumarProductos(productos []Producto, ch chan float64) {
	var total float64
	for _, p := range productos {
		total += p.Precio * float64(p.Cantidad)
	}
	ch <- total
}

func sumarServicios(servicios []Servicio, ch chan float64) {
	var total float64
	for _, s := range servicios {
		// Redondear hacia arriba a la media hora más cercana
		horasTrabajadas := float64(s.MinutosTrabajados) / 30
		if horasTrabajadas < 1 {
			horasTrabajadas = 1
		}
		total += s.Precio * horasTrabajadas
	}
	ch <- total
}

func sumarMantenimiento(mantenimientos []Mantenimiento, ch chan float64) {
	var total float64
	for _, m := range mantenimientos {
		total += m.Precio
	}
	ch <- total
}

// Las tres se deben ejecutar concurrentemente y, al final, se debe mostrar por pantalla el monto final (sumando el total de las 3).

func main() {
	// Datos de ejemplo
	productos := []Producto{{"Producto A", 10.0, 2}, {"Producto B", 5.0, 3}}
	servicios := []Servicio{{"Servicio X", 50.0, 45}, {"Servicio Y", 80.0, 20}}
	mantenimientos := []Mantenimiento{{"Mantenimiento Z", 100.0}}

	ch := make(chan float64)

	go sumarProductos(productos, ch)
	go sumarServicios(servicios, ch)
	go sumarMantenimiento(mantenimientos, ch)

	var total float64
	for i := 0; i < 3; i++ {
		total += <-ch
	}

	fmt.Println("El precio total es:", total)
}
