package main

import "fmt"

// canais auxiliam nas comunicações entre threads
// quando as duas precisam alterar a mesma variável por exemplo
func main() {
	// canal vazio
	canal := make(chan string)

	go func() {
		//preenchendo o canal
		canal <- "Olá mundo"
		//se eu precisar de add mais no canal, ele so vai conseguir depois de liberar o canal 1x
	}()

	msg := <-canal //esvazia canal
	fmt.Println(msg)
}
