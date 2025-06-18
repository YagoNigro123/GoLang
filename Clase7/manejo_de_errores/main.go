package main

import (
	"fmt"
	"github.com/YagoGNC/ModuloGoLang/greetings"
	"log"
)

func main() {
	//Creamos un prefijo en los mensajes de registro
	log.SetPrefix("greetings: ")
	//Quita la fecha y la hora
	log.SetFlags(0)

	names := []string{"Alex", "Juan", "Pedro"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	//message, err := greetings.Hello("juan")
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(messages)
}
