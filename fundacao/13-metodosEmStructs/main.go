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

func (c Cliente) DesativarCliente() {
	c.Ativo = false
	fmt.Printf("%s foi desativado", c.Nome)
}

func main() {

	wesley := Cliente{
		Nome:  "wesley",
		Idade: 33,
		Ativo: true,
	}
	wesley.Cidade = "SP"
	wesley.Endereco.Numero = 230
	wesley.DesativarCliente()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)

}
