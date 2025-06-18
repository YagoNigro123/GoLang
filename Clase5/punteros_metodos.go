package main

import "fmt"

type Perro struct {
	nombre string
	edad   int
}

func (p *Perro) ladra() {
	fmt.Println("Guau")
}

func main() {
	var x int = 10
	fmt.Println(x)
	editar(&x)

	var p *int = &x
	fmt.Println(&x) // direccion de memoria
	fmt.Println(p)  // puntero a esa direccion de memoria
	editar(&x)
	fmt.Println(x)

	leon := Perro{
		nombre: "Leon",
		edad:   10}
	leon.ladra()
}

func editar(x *int) {
	*x = 20
}
