package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string
	Author string
	Launch string
}

func saveBooksToFile(books []Book) error {
	//guardamos en file el archivo book.json, en caso de errores los guardamos en err
	file, err := os.Create("book.json")
	//si hay algun error, osea que err no esta vacio
	if err != nil {
		return err
		//retornamos el error
	}
	//decimos que file se cierre al final del codigo
	defer file.Close()

	//codificamos en json file
	encoder := json.NewEncoder(file)
	//Convierte el slice books a JSON y lo escribe en book.json
	err = encoder.Encode(books)
	if err != nil {
		return err
	}
	return nil
}

func loadBooksFromFile(books *[]Book) error {
	// Creamos la varibales file y error
	// file almacena el book.json, en caso de errores se guardan en error
	file, err := os.Open("book.json")
	// si los errores son distintos de vacio los retorna
	if err != nil {
		return err
	}
	// cerramos file al final del codigo
	defer file.Close()
	// decodificamos el json que esta en file
	decoder := json.NewDecoder(file)
	// capturamos el error si no puede decodificar los libros
	// conviuere JSON a books y los imprime por conosla
	err = decoder.Decode(books)
	// si hay errores los muestra
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var books []Book
	err := loadBooksFromFile(&books)
	if err != nil {
		fmt.Println("Error al cargar el contacto", err)
	}

	reader := bufio.NewReader(os.Stdin) // para leer la entrada de el teclado

	for {
		fmt.Println(`
		=== Libreria Pedrido ===
		1.Cargar libro
		2.Mostar libro
		3.Salir
		========================
		`)
		var option int
		_, err := fmt.Scanln(&option)

		if err != nil {
			fmt.Println("Error al leer la opcion")
			return
		}
		switch option {
		case 1:
			var b Book
			fmt.Println("Titulo: ")
			b.Title, _ = reader.ReadString('\n')
			fmt.Println("Autor: ")
			b.Author, _ = reader.ReadString('\n')
			fmt.Println("Año de estreno: ")
			b.Launch, _ = reader.ReadString('\n')

			books = append(books, b)
			if err := saveBooksToFile(books); err != nil {
				fmt.Println("Error al guardar el libro: ", err)
			}
		case 2:
			for index, book := range books {
				fmt.Println("\n======================================================")
				fmt.Printf("N':%d \n Titulo: %s\n Autor: %s\n año de lanzamiento: %s\n", index+1, book.Title, book.Author, book.Launch)
				fmt.Println("========================================================\n")
			}
		case 3:
			return
		default:
		}
	}
}
