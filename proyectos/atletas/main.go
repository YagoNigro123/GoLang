package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//Crear un sistema de registro de atletas de alto rendimineto, se solicita:
// 1. Nombre, 2. Edad, 3. Peso
//Create
//Read
//Update
//Delate

type Athlete struct {
	Id     int8    `json:"id"`
	Name   string  `json:"name"`
	Yeards int8    `json:"yeards"`
	Meters float32 `json:"meters"`
}

func saveAthleteToFile(athletes []Athlete) error {
	file, err := os.Create("athletes.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(athletes)
	if err != nil {
		return err
	}
	return nil
}

func loadAthleteToFile(athletes *[]Athlete) error {
	file, err := os.Open("athletes.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(athletes)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var list []Athlete
	err := loadAthleteToFile(&list)
	if err != nil {
		fmt.Println("Error al cargar el atleta", err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println(`
		===== GYM ULTRA VIP =====
		¿Qué operación desea realizar?
		1.Cargar 
		2.Mostar 
		3.Editar 
		4.Eliminar
		5.Salir
		=========================
		`)
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al cargar la opcion")
		}
		switch option {
		case 1:
			var athlete Athlete
			athlete.Id = int8(len(list) + 1)

			fmt.Println("Ingrese el nombre del atleta")
			name, _ := reader.ReadString('\n')
			athlete.Name = name[:len(name)-1]

			fmt.Println("Ingrese la edad del atleta")
			fmt.Scanln(&athlete.Yeards)

			fmt.Println("Ingrese la altura del atleta")
			fmt.Scanln(&athlete.Meters)

			list = append(list, athlete)

			err := saveAthleteToFile(list)

			if err != nil {
				fmt.Println("Erro al guardar: ", err)
			}

			fmt.Println("Atleta creado correctamente")
		case 2:
			if len(list) == 0 {
				fmt.Println("No hay atletas cargados.")
				continue
			}
			for _, athlete := range list {
				fmt.Println("========================")
				fmt.Printf("ID: %d\n", athlete.Id)
				fmt.Printf("Nombre: %s\n", athlete.Name)
				fmt.Printf("Edad: %d\n", athlete.Yeards)
				fmt.Printf("Altura: %.2f\n", athlete.Meters)
				fmt.Println("========================")
			}
		case 3:
			fmt.Println("Selecione el atleta que quiere modificar")
			var id int8
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			found := false
			for i := range list {
				if list[i].Id == id {
					var newName string
					fmt.Print("Ingrese el nuevo nombre: ")
					newName, _ = reader.ReadString('\n')
					fmt.Print("Ingrese la nueva edad: ")
					var newYeards int8
					fmt.Scanln(&newYeards)
					fmt.Print("Ingrese la nueva altura: ")
					var newMeters float32
					fmt.Scanln(&newMeters)
					list[i].Name = newName
					list[i].Yeards = newYeards
					list[i].Meters = newMeters
					found = true
					fmt.Println("Atleta actualizado con exito")
					saveAthleteToFile(list)
					break
				}
			}
			if !found {
				fmt.Println("No se encontro ningún atleta con ese ID.")
			}
		case 4:
			fmt.Println("Selecione el atleta que desea eliminar")
			var id int8
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
				return
			}
			found := false
			for i, atlheta := range list {
				if atlheta.Id == id {
					list = append(list[:i], list[i+1:]...)
					found = true
					fmt.Println("Atleta eliminado correctamente")
					saveAthleteToFile(list)
				}
			}
			if !found {
				fmt.Println("No se encontro ningun atleta con ese ID.")
			}
		case 5:
			return
		}

	}
}
