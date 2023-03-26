package main

import "fmt"

// o go não é Orientado a Objetos
// seu principal diferencial, modo são definidos pelos structs
// o struct não é uma classe : (

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	wesley := Cliente{
		Nome:  "wesley",
		Idade: 33,
		Ativo: true,
	}
	wesley.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)

}
