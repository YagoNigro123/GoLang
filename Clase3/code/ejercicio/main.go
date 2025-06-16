package main

import (
	"fmt"
	"math"
)

const MAX_PRECICION = 2

func areaTriangulo(a float32, b float32) {
	fmt.Printf("El area vale: %.*f \n", MAX_PRECICION, (b*a)/2)
}

func permiteroTriangulo(a float32, b float32, c float32) {
	fmt.Printf("El perimetro vale: %.*f \n", MAX_PRECICION, a+b+c)
}
func hipotenusaTriangulo(a float32, b float32, c *float32) {
	*c = float32(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("La hipotenusa vale:%.*f \n", MAX_PRECICION, *c)
}
func main() {
	var ca float32
	var cb float32
	var hip float32

	fmt.Println("Bienvenido a CALCULE EL PERIMETRO DE SU TRIANGULO \n ingrese el lo siguientes datos")
	fmt.Println("Cateto A")
	fmt.Scan(&ca)
	fmt.Println("Cateto B")
	fmt.Scan(&cb)

	hipotenusaTriangulo(ca, cb, &hip)
	areaTriangulo(ca, cb)
	permiteroTriangulo(ca, cb, hip)
}
