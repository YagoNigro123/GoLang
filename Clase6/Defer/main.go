package main

import (
	"fmt"
	"os"
)

func main() {
	// Pospone la ejecucion de una funcion, hasta que la funcion que la contine haya finalizado
	defer fmt.Println("4@ en pila: fin") // se hace al final
	fmt.Println(1)
	fmt.Println(2)

	// 1,2,3

	defer fmt.Println("3@ en pila")
	//se agrega a una pila de ejecucion
	defer fmt.Println("2@ en pila")
	defer fmt.Println("\n \n \n1@ en pila")
	//1,2,3

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Manejo de errores eficiente y eficaz, ademas de ahorranos 2 lineas.
	defer file.Close() //ventajas comentadas 1 2

	cantBytes, err := file.Write([]byte("Hola, bro"))
	if err != nil {
		fmt.Println(err)
		// file.Close() <- 1
		return
	}
	fmt.Println("Cantidad de bytes", cantBytes)
	// file.Close() // cerramos el flujo <- 2
}
