package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var alex Persona

	alex.nombre = "Alex"
	alex.edad = 28
	alex.correo = "alex@gmail.com"

	fmt.Println(alex)

	pedro := Persona{
		nombre: "pedro",
		edad:   25,
		correo: "pedro@gmail.com"}
	pedro.edad = 26
	fmt.Println(pedro)

	juan := Persona{
		nombre: "Juan",
		edad:   34,
		correo: "juan@gmail.com"}

	fmt.Println(juan)
}
