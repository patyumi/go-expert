package main

import "fmt"

func main() {
	// A interface vazia suporta qualquer coisa
	var x interface{} = 10
	var y interface{} = "Hello, World"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v", t, t)
}
