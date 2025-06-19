package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//canla := make(chan int) // Se declara
	//canla <- 15             // Se envia datos al canal
	//valor := <-canla        // Recibe datos del canal

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com"}

	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
		fmt.Println(<-ch)
	}

	//time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: ¡%s está caído!\n", api)
	}
	ch <- fmt.Sprintf("SUCCESS: ¡%s está en funcionamiento!\n", api)
}
