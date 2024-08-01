package main

import "fmt"

/* Consigna
Una empresa de redes sociales requiere implementar una estructura usuarios con funciones que vayan agregando información a la misma. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones. La estructura debe tener los campos: nombre, apellido, edad, correo y contraseña. Y deben implementarse las funciones:
cambiarNombre: permite cambiar el nombre y apellido.
cambiarEdad: permite cambiar la edad.
cambiarCorreo: permite cambiar el correo.
cambiarContraseña: permite cambiar la contraseña. */

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

// Función para cambiar nombre y apellido
func (u *Usuario) cambiarNombre(nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

// Función para cambiar la edad
func (u *Usuario) cambiarEdad(edad int) {
	u.Edad = edad
}

// Función para cambiar el correo
func (u *Usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

// Función para cambiar la contraseña
func (u *Usuario) cambiarContrasena(contrasena string) {
	u.Contrasena = contrasena
}

func main() {
	usuario := Usuario{
		Nombre:     "Juan",
		Apellido:   "Pérez",
		Edad:       30,
		Correo:     "juan.perez@example.com",
		Contrasena: "123456",
	}

	// Cambiando nombre y apellido
	usuario.cambiarNombre("Carlos", "Gómez")

	// Cambiando la edad
	usuario.cambiarEdad(35)

	// Cambiando el correo
	usuario.cambiarCorreo("carlos.gomez@example.com")

	// Cambiando la contraseña
	usuario.cambiarContrasena("nueva1234")

	// Mostramos los datos del usuario después de los cambios
	fmt.Printf("Usuario: %+v \n", usuario)
	// %+v: Es un especificador de formato en Go que se usa para imprimir valores en un formato verboso, adecuado para depuración. El signo + indica que se deben incluir tanto los nombres de los campos de la estructura como sus valores. Esto es útil cuando deseas ver claramente cuáles son los campos y sus valores asociados, lo que facilita entender la salida, especialmente cuando trabajas con estructuras complejas.
}
