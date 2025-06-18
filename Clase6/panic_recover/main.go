package main

import "fmt"

func divide(dividendo, divisor int) {
	defer func() {
		// si ocurre un panico, se almacena en r, y no interrumpe la ejecucion de nuestra app
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero")
	}
}

func main() {
	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(100, 4)
}
