package main

import "fmt"

func main() {
	//rebanadas
	//arrays dinamicos
	//slice
	var rebanada []int
	//otra forma
	diasSemana := []string{
		"lunes\n",
		"martes\n",
		"miercoles\n",
		"jueves\n",
		"Viernes\n",
		"sabado\n",
		"domingo\n"}
	fmt.Println(rebanada)
	fmt.Println(diasSemana)
	diasLaborales := diasSemana[0:5]
	fmt.Println(diasLaborales)
	//para ver la cantidad de elementos
	// funciona como len()
	fmt.Println(cap(diasSemana))
	//Para agregar elementos usamos append
	diasNecesarios := append(diasSemana, "El dia 8\n", "El dia 9\n")
	fmt.Println(diasNecesarios)

	necesarioNoEsReal := append(diasNecesarios[7:8], diasNecesarios[0:7]...)

	fmt.Println(necesarioNoEsReal)

	//Creamos un slice
	//Tipo : string
	//Longitud inicial: 5
	//Capacidad maxima: 10
	nombres := make([]string, 5, 10)
	nombres[0] = "Alex"
	nombres[1] = "jose"
	nombres[2] = "tito"
	nombres[3] = "martin"
	nombres[4] = "kiko"
	//tira error -> nombres[5] = "jorge"
	// como la longitud inicial es 5, a partir de este numero deja de ser estatico y arranca a ser dinamico
	nombres = append(nombres, "Ruben")
	fmt.Println(nombres)
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	//REBANADA2 copia a rebanada1
	copy(rebanada1, rebanada2)

	fmt.Println(rebanada1, rebanada2)
	//Cantidad de elementos copiados
	fmt.Println(copy(rebanada1, rebanada2))

}
