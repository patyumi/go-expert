package main

import (
	"fmt"

	"github.com/patyumi/go-expert/20-pacotesModulos1/matematica"
)

// Essa aula aqui fala sobre os modulos para importar pacotes dentro da pasta
// então ele fala do go mod init - go mod tidy

// na terceira aula ele fala sobre o nome das funções, onde as que iniciam com letra maiuscula é visível globalmente
// e as minusculas apenas no escopo atual
// e isso nào é só pra funções, se aplica as variaveis tambem, qualquer uma, seja dentro de um struct etc

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v", s)
}
