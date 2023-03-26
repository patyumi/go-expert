package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // isso é uma composição, semelhante a herança, mas não é uma herança
}

func main() {

	wesley := Cliente{
		Nome:  "wesley",
		Idade: 33,
		Ativo: true,
	}
	wesley.Ativo = false
	wesley.Cidade = "SP"
	wesley.Endereco.Numero = 230

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)

}
