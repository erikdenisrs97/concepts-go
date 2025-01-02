package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexer(escrever("Olá Mundo 1"), escrever("Olá Mundo 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// multiplexer pattern
func multiplexer(canal1, canal2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalDeSaida <- mensagem
			case mensagem := <-canal2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

// generator pattern
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
