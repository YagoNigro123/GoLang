package main

import (
	"fmt"
	"math"
)

func main() {
	//alamcena 8 bits
	//-128 hasta 127
	fmt.Println(math.MinInt64, math.MaxInt64)
	fmt.Println(math.MaxFloat32)

	fullName := "Oso roxo \t (alias \"roxo\")\n"
	fmt.Println(fullName)

	var a byte = 'a'
	//Imprime el valor de codigo ASII
	fmt.Println(a)
	s := "hola"
	fmt.Println(s[0]) //imprime el primer valor de la cadena

	var r rune = '♥'
	fmt.Println(r) //impre el valor unicode del caracter
}
