package main

import "fmt"

// vc pode forçar o tipo de uma variável de interface vazia e isso é chamado de type assertion

func main() {
	var minhaVar interface{} = "paty"

	println(minhaVar)
	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("o valor de res é %v e o valor de ok é %v", res, ok)

}
