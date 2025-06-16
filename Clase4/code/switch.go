package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// En que sistema operativo se ejecuta go
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "Linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> MacOs")
	default:
		fmt.Println("Go run -> Otro OS")
	}
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println(t.Hour(), "¡Mañana!")
	case t.Hour() < 17:
		fmt.Println(t.Hour(), "¡Tarde!")
	default:
		fmt.Println(t.Hour(), "¡Noche!")
	}
}
