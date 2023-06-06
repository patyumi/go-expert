package main

import "fmt"

func main() {
	// um array é uma variável que tem tamanho fixo e voce pode percorrer ele
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 10

	fmt.Println(len(meuArray) - 1)
	fmt.Println(meuArray[len(meuArray)-1])

	for indice, value := range meuArray {
		fmt.Printf("O valor do indice %d é %d\n", indice, value)
	}
}
