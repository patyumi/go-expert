package main

import "fmt"

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
	}

}

func main() {
	data := make(chan int)

	qtdWorkers := 3

	// inicializa workers
	for i := 0; i < qtdWorkers; i++ {
		// esvaziando meu canal
		go worker(i, data)
	}

	// enviando infos pro meu canal
	for i := 0; i < 100; i++ {
		data <- i
	}

}
