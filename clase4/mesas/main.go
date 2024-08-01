package main

/* Productos
Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total. La empresa tiene tres tipos de productos: Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
Los costos adicionales para cada uno son:
Pequeño: solo tiene el costo del producto.
Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional de $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio, y retorne una interfaz Producto que tenga el método Precio.
Se debe poder ejecutar el método Precio y que el método devuelva el precio total basándose en el costo del producto y los adicionales, en caso de que los tenga.
*/