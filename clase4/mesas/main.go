package main

import "fmt"

/* Productos
Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total. La empresa tiene tres tipos de productos: Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
Los costos adicionales para cada uno son:
Pequeño: solo tiene el costo del producto.
Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional de $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio, y retorne una interfaz Producto que tenga el método Precio.
Se debe poder ejecutar el método Precio y que el método devuelva el precio total basándose en el costo del producto y los adicionales, en caso de que los tenga.*/

// INTERFAZ PRODUCTO CON METODO PRECIO
type Producto interface {
	Precio() float64
}

//TIPO CONCRETOS PARA CADA PRODUCTO
type Pequeño struct {
	precio float64
}

type Mediano struct {
	precio float64
}

type Grande struct {
	precio float64
}

//METODO PRECIO DE CADA PRODUCTO
func (p Pequeño) Precio() float64 {
	return p.precio
}

func (m Mediano) Precio() float64 {
	return m.precio * 1.03
}

func (g Grande) Precio() float64 {
	return g.precio*1.06 + 2500
}

//IMPLEMENTAMOS FUCION FACTORY PARA CREAR PRODUCTO
func factoryNuevoProducto(tipoProducto string, precio float64) Producto {
	switch tipoProducto {
	case "Pequeño":
		return Pequeño{precio: precio}
	case "Mediano":
		return Mediano{precio: precio}
	case "Grande":
		return Grande{precio: precio}
	default:
		return nil
	}

}

func main() {
	tipos := []string{"Pequeño", "Mediano", "Grande"}

	for _, tipo := range tipos {
		producto := factoryNuevoProducto(tipo, 100) // Precio base de ejemplo
		fmt.Printf("El precio total del producto %s es: $%.2f\n", tipo, producto.Precio())
	}
}

/* GUÍA PARA LA RESOLUCIÓN

Este ejercicio de programación en Go es un ejemplo clásico de cómo implementar patrones de diseño, específicamente el patrón Factory, en conjunto con el uso de interfaces para manejar un sistema de objetos con comportamientos variados pero relacionados. Vamos a desglosar la consigna y entender cada parte:

1. **Objetivo general**: Crear un sistema para administrar productos en una tienda de e-commerce, donde se pueden calcular los precios totales de los productos que tienen diferentes costos adicionales basados en su tamaño.

2. **Tipos de productos**:
   - **Pequeño**: Solo tiene el costo del producto sin adicionales.
   - **Mediano**: Incluye el precio del producto más un 3% adicional por el costo de mantenerlo en la tienda.
   - **Grande**: Incluye el precio del producto más un 6% adicional por el costo de mantenerlo en la tienda y un costo fijo adicional de $2500 por el envío.

3. **Cálculo de costos adicionales**:
   - Para los productos medianos y grandes, el porcentaje adicional se calcula sobre el precio base del producto. Esto significa que necesitas una forma de tomar el precio del producto, calcular el porcentaje adicional, y luego sumarlo al precio base.
   - Para los productos grandes, además del porcentaje, se añade un costo fijo de envío al total.

4. **Interfaz Producto**:
   - Esta interfaz debe incluir al menos un método, `Precio`, que cuando sea llamado, deberá retornar el precio total del producto considerando todos los costos adicionales.
   - La interfaz `Producto` permite manejar diferentes tipos de productos de una manera uniforme. Cada tipo de producto implementará esta interfaz de manera que calcule su precio de forma específica.

5. **Función Factory**:
   - Se debe implementar una función factory que cree instancias de estos productos. Esta función recibirá el tipo de producto y su precio base como parámetros y retornará una instancia del tipo correcto que implemente la interfaz `Producto`.
   - La función factory es responsable de determinar qué tipo de objeto de producto crear y retornar, basado en el tipo de producto especificado. Por ejemplo, si se solicita un producto "Grande", la factory debería retornar una instancia de un tipo que calcule el precio como el precio base más un 6% y $2500 adicionales.

6. **Ejecución del método Precio**:
   - Al obtener un objeto producto de la función factory, podrás llamar al método `Precio` para obtener el precio final del producto.
   - Esto permite que el cálculo específico del precio esté encapsulado dentro de cada clase de producto, y el código que utiliza estos productos puede permanecer simple y desacoplado de los detalles de cómo se calculan estos precios.

Este ejercicio te da la oportunidad de practicar con conceptos importantes en Go como son las interfaces, métodos, y el patrón de diseño Factory, todos los cuales son fundamentales para escribir código limpio y bien organizado en proyectos más grandes.

*/
