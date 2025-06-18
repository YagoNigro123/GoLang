package main

import "fmt"

func main() {
	saludo := hello("papucho")
	fmt.Println(saludo)
	sum, mul := calc(4, 5)
	fmt.Println("La suma es", sum)
	fmt.Println("La multiplicacion es", mul)
}

func hello(name string) string {
	return "Hola " + name
}

func calc(a, b int) (int, int) {
	suma := a + b
	mul := a * b
	return suma, mul
}
