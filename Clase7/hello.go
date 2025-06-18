package main

//En este ejercicio, use un repositorio de mi cuenta institucional para entender el manejo de los modulos
import (
	"fmt"
	"github.com/YagoGNC/ModuloGoLang/greetings"
)

func main() {
	fmt.Println()
	message := greetings.Hello("pedro")
	fmt.Println(message)
}
