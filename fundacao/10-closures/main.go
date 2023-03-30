package main

import (
	"fmt"
)

func main() {

	total := func() int {
		return sum(2+2) * 2
	}()

	fmt.Println(total)
}

// esses 3 pontinhos significa que eu vou fazer um loop com todos os parametros desse mesmo tipo
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
