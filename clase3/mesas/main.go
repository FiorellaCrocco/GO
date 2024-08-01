package main

import "fmt"

/* Crear un programa que cumpla los siguientes puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products, instanciado con valores.
Dos métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definidos desde main().
*/

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product = []Product{
	Product{1, "Producto1", 100.50, "Descripcion1", "Arte"},
	Product{2, "Producto2", 200.50, "Descripcion2", "Arte"},
	Product{3, "Producto3", 300.50, "Descripcion3", "Arte"},
}

func (product Product) Save() {
	Products = append(Products, product)
}

func (product Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product)
	}
}

func getById(id int) Product {
	for _, product := range Products {
		if product.Id == id {
			return product
		}
	}
	return Product{}
}
func main() {
	product := Product{4, "Producto4", 400.99, "Descripcion4", "Categoria4"}
	product.Save()
	product.GetAll()
	fmt.Println(getById(4))
}
