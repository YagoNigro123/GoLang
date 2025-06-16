package main

import (
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		//Creamos un nuevo mensaje de error
		// return 0, errors.New("No se puede dividir por cero")
		return 0, fmt.Errorf("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	str := "123"
	//Convierte de string a numero
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Hubo error", err)
		return
	}

	fmt.Println("Str a Int:", num)

	//Manejamos el error
	result, err := divide(12, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}
