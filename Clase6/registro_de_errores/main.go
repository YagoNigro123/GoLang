package main

import (
	"log"
	"os"
)

func main() {
	// os.OpenFile recibe:
	// 1. "info.log"  -> Nombre del archivo a abrir o crear
	// 2. os.O_CREATE -> Crea el archivo si no existe
	// 3. os.O_APPEND -> Escribe al final del archivo (no sobreescribe)
	// 4. os.O_WRONLY -> Abre el archivo solo para escritura
	// 5. 0644        -> Permisos del archivo (lectura/escritura para el dueño, solo lectura para otros)
	log.SetPrefix("main()")
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)
	log.Print("soy un log.")
}
