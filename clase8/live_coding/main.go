package main

import "fmt"

type Bootcamp struct {
	Organizacion string
	Lenguaje     string
	Alumnos      []Alumno
}

type Alumno struct {
	Nombre       string
	FechaIngreso string
}

func crearBootcamp(org, lng string) Bootcamp {
	boot := Bootcamp{
		Organizacion: org,
		Lenguaje:     lng,
	}
	return boot
}

func (b *Bootcamp) guardarAlumno(alumno Alumno) {
	b.Alumnos = append(b.Alumnos, alumno)
}

func main() {
	bootcampGO := crearBootcamp("MELI", "GO")

	alm1 := Alumno {
		Nombre: "Agustin",
		FechaIngreso: "18/02/2022",
	}
	alm2 := Alumno {
		Nombre: "Vicent",
		FechaIngreso: "18/02/2022",
	}
	alm3 := Alumno {
		Nombre: "Sol",
		FechaIngreso: "18/02/2022",
	}

	bootcampGO.guardarAlumno(alm1)
	bootcampGO.guardarAlumno(alm2)
	bootcampGO.guardarAlumno(alm3)

	fmt.Println(bootcampGO)
}
