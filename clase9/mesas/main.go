package main

/* Lista de productos
Vamos a crear una aplicación web con el framework Gin que tenga un endpoint /productos que responda con una lista de productos.
Crear una estructura Productos con los valores:
Id
Nombre
Precio
Stock
Codigo
Publicado
FechaDeCreacion
Crear un archivo productos.json con al menos seis productos (deben seguir la misma estructura ya mencionada).
Leer el archivo productos.json.
Imprimir la lista de productos por consola.
Imprimir la lista de productos a través del endpoint en formato JSON. El endpoint deberá ser de método GET. */

/*
Para utilizar Gin se requiere la versión 1.13+ de Go, una vez
instalada, utilizamos el siguiente comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/

// go mod init actividad/mesas

// go mod tidy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id              int
	Nombre          string
	Precio          float64
	Stock           int
	Codigo          string
	Publicado       bool
	FechaDeCreacion string
}

func main() {
	router := gin.Default()

	//Captura la solicitud GET “/productos"
	router.GET("/productos", func(c *gin.Context) {

		datosBytes, err := os.ReadFile("./productos.json")
		if err != nil {
			log.Fatal(err)
		}
		datosString := string(datosBytes)

		//Imprimo la lista de personas por consola
		fmt.Println(datosString)

		var listProducts []Productos

		if err := json.Unmarshal([]byte(datosString), &listProducts); err != nil {
			log.Fatal(err)
		}

		//Devolver Lista de Productos
		c.JSON(200, listProducts)
	})

	router.Run(":8080")

	//Ejecutar con: go run main.go
	//Para para el servidor hacer: Ctrl + C
}
