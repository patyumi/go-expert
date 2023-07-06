package main

import (
	"fmt"

	"github.com/patyumi/go-expert/5packaging/1/math"
)

func main() {

	// iniciar modulo
	// go mod init github.com/patyumi/go-expert/5packaging

	// recomendado sempre utilizar um url nos módulos

	// objetos com letra maíusculas são públicos e minúsculas são privados

	fmt.Println("hello")

	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
