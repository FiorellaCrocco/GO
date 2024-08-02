package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	readFile()
}

func readFile() {
	data, err := os.ReadFile("./ctdGO.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)
}
