package main

import (
	"POO/animal"
	"POO/book"
	"fmt"
)

func main() {
	fmt.Println("===================")
	fmt.Println("Uso basico de un struct")
	fmt.Println("===================")
	var myBook = book.Book{
		Title:  "Moby Dick",
		Author: "Herman Melville",
		Pages:  300,
	}
	myBook.PrintInfo()

	myNewBook := book.NewBook("Titulo", "Autor", 0)

	fmt.Println()
	fmt.Println("===================")
	fmt.Println(`(Constructores, Estructuras, funciones) 
Publicas y privadas`)
	fmt.Println("===================")
	myPrivateBook := book.NewPrivateBook("Titulo Privado", "Autor Privado", 100)
	myPrivateBook.PrintPrivateInfo()
	myBook.Title = "Moby Dick (Edicion especial)"

	fmt.Println()
	fmt.Println(myBook.Title)

	fmt.Println()
	fmt.Println(myNewBook)
	fmt.Println()

	myTextBook := book.NewTextBook("Comunicación", "Jaime Gamarra", 261, "Santillana SAC", "secundaria")
	myTextBook.PrintPrivateInfo()

	fmt.Println()
	fmt.Println("===================")
	fmt.Println("Composición")
	fmt.Println("===================")
	myBook.PrintInfo()
	fmt.Println()
	myTextBook.PrintInfo()
	fmt.Println("===================")
	fmt.Println("Polimorfismo")
	fmt.Println("===================")
	book.Print(myBook)
	book.Print(myTextBook)
	fmt.Println("===================")
	fmt.Println("Interfaces")
	miPerro := animal.Perro{Nombre: "Tobias"}
	miGato := animal.Gato{Nombre: "Tom"}
	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Tobias"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"}}
	for _, animal := range animales {
		animal.Sonido()
	}
}
