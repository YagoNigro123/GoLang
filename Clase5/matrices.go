package main

import "fmt"

func main() {
	var a [5]int
	//matriz de 5 elementos

	fmt.Println(a)
	// [0 ,0,0,0,0]
	// [10,0,0,0,0]
	a[0] = 10
	// [10,5,0,0,0]
	a[1] = 5
	fmt.Println(a)
	//Declaramos e inicializamos
	var b = [...]int{10, 20, 30, 40, 50}
	// ... -> no sabemos la cantidad de datos
	fmt.Println(b)
	//¿Como accedo a los elementos?
	fmt.Println(a[1]) // posición 1
	//¿Como recorro todos los elementos?
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	//Otra forma -> range
	for index, value := range b {
		fmt.Printf("índece:%d, valor:%d\n", index, value)
	}
	for _, value := range b {
		fmt.Printf("valor:%d\n", value)
	}
	//Matrices multidimensionales
	//1@ filas
	//2@ columnas
	var matriz = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println(matriz)
	
}
