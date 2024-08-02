package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	createFile()
}

func createFile() {

	file := "./ctdGO.txt"
	f, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("¿Es un directorio?:", f.IsDir())
	fmt.Println("Nombre del archivo/directorio:", f.Name())
	fmt.Println("Tamaño del archivo en Bytes:", f.Size())
	fmt.Println("Fecha y hora del archivo:", f.ModTime())
	fmt.Println("Permisos del archivo:", f.Mode())
}
