package main

import (
	"fmt"
	"sync"
	"time"
)

// O WaitGroup é separado em 3 partes
// 1. Adicionar qtd de tarefas e operações
// 2. Informar que você terminou uma operação
// 3. Esperar até que as operações sejam finalizadas

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d : Task %s is running\n", i+1, name)
		time.Sleep(1 * time.Second)
		// passo 2
		wg.Done()
	}
}

// thread main já é 1 thread
func main() {
	waitGroup := sync.WaitGroup{}
	// passo 1
	waitGroup.Add(25)

	go task("A", &waitGroup)
	go task("B", &waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d : Task %s is running\n", i+1, "anonymous")
			time.Sleep(1 * time.Second)
			// passo 2
			waitGroup.Done()
		}
	}()

	// passo 3
	waitGroup.Wait()
}
