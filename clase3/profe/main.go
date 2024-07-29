package main

import "fmt"

type Person struct {
	Id          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Id       int
	Position string
	Person   Person
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
	fmt.Println(e1.PrintEmployee())
	// Llamamos al método que no le asignamos que devuelve (no usamos return)
	e1.PrintEmployee2()
}
