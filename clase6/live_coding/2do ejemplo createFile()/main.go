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
	f, err := os.Create("ctdGO.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString("Este texto proviene de nuestro programa GO.")
	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
