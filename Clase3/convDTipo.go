package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var integer16 int16 = 50
	// var integer32 int32 = 100
	s := "100"
	i, _ := strconv.Atoi(s)

	fmt.Println(i)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s)

	var name string
	var name2 string
	var age int

	fmt.Print("Ingrese su nombre:")
	fmt.Scanln(&name, &name2)
	fmt.Printf("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola me llamo %s y tengo %d \n", name, age)

	// greeting := fmt.Sprintf("Hola me llamo %s y tengo %d ", name, age)
	// una chanchada, mejor usar \n
	// fmt.Println(greeting)

	fmt.Printf("el tipo de dato de name, es %t y el de edad es %t \n", name, age)
}
