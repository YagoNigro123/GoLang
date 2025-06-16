package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//Guardar en un archivo los contactos
//Agregar contactos
//Editar contactos
//Ver todos los contactos
//Salir del programa
//Guardamos los contactos en un archivo

type Contact struct {
	Name  string `json:name`
	Email string `json:email`
	Phone string `json:phone`
}

// Guarda contactos en un archivo de tipo json
func saveContactsToFile(contacts []Contact) error {
	// 1. Crea (o sobreescribe) el archivo "contacts.json"
	file, err := os.Create("contacts.json")
	if err != nil {
		// 2. Si ocurre un error al crear el archivo, lo retorna
		return err
	}
	// 3. Asegura que el archivo se cierre al terminar la función
	defer file.Close()

	// 4. Crea un codificador JSON que escribirá en el archivo
	encoder := json.NewEncoder(file)

	// 5. Codifica (serializa) el slice de contactos y lo escribe en el archivo
	err = encoder.Encode(contacts)
	if err != nil {
		// 6. Si ocurre un error al codificar/escribir, lo retorna
		return err
	}

	// 7. Si todo sale bien, retorna nil (sin error)
	return nil
}

// Carga contactos desde un arhcivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	//Slice de contactos
	var contacts []Contact

	//Cargar contactos existentes desde el archivo

	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar el contacto", err)
	}

	//Crear instancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {
		//Muestra las opciones al usuario
		fmt.Print(
			"==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar un contacto \n",
			"2. Mostrar todos los contactos \n",
			"3. Salir\n",
			"Elige una opcion:")
		//Leer la opcion del usuario
		var option int

		_, err := fmt.Scanln(&option)

		if err != nil {
			fmt.Println("Error al leer la opcion", err)
			return
		}

		//Manejar la opcion del usuario

		switch option {
		case 1:
			//Ingresar y crear contacto
			var c Contact
			fmt.Print("Name: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			//Agregar un contacto a slice
			contacts = append(contacts, c)
			//Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto: ", err)
			}
		case 2:
			//Mostrar todos los datos de nuestro contacto
			for index, contact := range contacts {
				fmt.Println("===============================")
				fmt.Printf("%d.Nombre:%s Email:%s Numero:%s\n", index+1, contact.Name, contact.Email, contact.Phone)
				fmt.Println("===============================")
			}
		case 3:
			// salir del programa
			return
		default:
		}
	}
}
