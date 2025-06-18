package main

import "fmt"

const pi float32 = 3.14
const (
	x = 100
	y = 0b1010 //binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println(x, y, z, w)
}
