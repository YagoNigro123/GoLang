package main

import "fmt"

func main() {

	// Se usa una clave y un valor
	colors := map[string]string{
		//clave    valor
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF"}
	fmt.Println(colors)
	fmt.Println(colors["rojo"])

	colors["negro"] = "#000000"
	fmt.Println(colors)

	valor, ok := colors["blanco"]
	fmt.Println(valor, ok)

	if valor, ok := colors["verde"]; ok {
		fmt.Println("Si existe la clave: ", valor)
	} else {
		fmt.Println("No existe la clave: ")
	}
	//para eliminar un elemento de un mapa
	delete(colors, "azul")
	fmt.Println("Sin azul:\n", colors)

	for clave, valor := range colors {
		fmt.Printf("Clave: %s, valor: %s\n", clave, valor)
	}
}
