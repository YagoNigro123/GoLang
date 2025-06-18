package main

import "fmt"

func main() {
	var firstName, lastName string
	var age int
	firstName = "pedro"
	lastName = "suarez"
	age = 28

	var (
		nombre string = "juan"
		edad   int    = 20
	)

	var dom, num, xin = "hola", 19, "adios"

	nick := "pedro"

	fmt.Println(firstName, lastName, age)
	fmt.Println(nombre, edad)
	fmt.Println(dom, num, xin)
	fmt.Println(nick)
}
