package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Para usar etiquetas no debe haber espacios entre `json:""`
type Person struct {
	Id          int    `json:"id"`
	Name        string `json:"name,omitempty"` // Usamos "omitempty" para que si no se pasa el valor al instanciar se omita
	DateOfBirth string `json:"dateofbirth"`
}

type Employee struct {
	Id       int    `json:"-"` // El "-" se usa para no mostrar ese valor, es decir, ocultarlo
	Position string `json:"position"`
	Person   Person `json:"person"`
}

// metodo asociado a employee, estructura que devuelve string
func (e Employee) PrintEmployee() string {
	return fmt.Sprintf("Employee ID: %d\n Name: %s\n Date of Birth %s\n Position: %s", e.Id, e.Person.Name, e.Person.DateOfBirth, e.Position)
}

// metodo asociado a employee, estructura que no asignamos que devuelve - no usa return -
func (e Employee) PrintEmployee2() {
	fmt.Printf("Employee ID: %d\n Name: %s\n Date of Birth %s\n Position: %s", e.Id, e.Person.Name, e.Person.DateOfBirth, e.Position)
}

func main() {
	p1 := Person{1, "Juan", "10/10/2000"}
	e1 := Employee{1, "Gerente", p1}

	//Llamamos al método que retorna string
	fmt.Println("\n ")
	fmt.Print("PrintEmployee en string: \n")
	fmt.Println(e1.PrintEmployee())
	// Llamamos al método que no le asignamos que devuelve (no usamos return)
	fmt.Println("\n ")
	fmt.Print("PrintEmployee sin return: \n")
	e1.PrintEmployee2()

	// Código de ejemplo para ver el uso de etiquetas, convirtiendo a formato json (Para eso uso json.Marchal)
	employee, err := json.Marshal(e1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n ")
	fmt.Print("PrintEmployee con Marshal: \n")
	fmt.Println(string(employee))

	// Código de ejemplo para convertir de json a string (formato definido en PrintEmployee, para eso uso json.Unmarchal)
	var e2 Employee

	employee2 := `{"Id":1,"Position":"Gerente","Person":{"Id":1,"Name":"Juan","DateOfBirth":"10/10/2000"}}`

	if err2 := json.Unmarshal([]byte(employee2), &e2); err2 != nil {
		log.Fatal(err)
	}

	fmt.Println("\n ")
	fmt.Print("PrintEmployee con Unmarchal: \n")
	fmt.Println(e2.PrintEmployee())
}
