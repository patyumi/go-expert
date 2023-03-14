// ----------------------------------------------------- 1 Aula teórica e definições --------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------
// o main é o principal package do sistema
// package de entrada do programa

// todos os outros pacotes devem acompanhar o nome do diretório em que ele está

// dentro do mesmo pacote, nesse caso o main, todas as variáveis e funções são visíveis para todos

// ----------------------------------------------------- 2 Testes ---------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

// Para testar de maneira mais fácil, é necessário utilizar funções
// Então separa-se o código principal do mundo de fora e seus efeitos colaterais, como por exemplo, o PrintLn escreve um stdout

// Criando testes
// O arquivo de teste deve se chamar xxx_test.go
// As funcões de teste sempre começam com Test e tem o argumento t *testing.T

// TODO : Finalizar testes https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

func Hello() string {
	return "Hello, my name is Patrícia"
}

func main() {
	fmt.Println(Hello())
	os.Exit(3)
}

// ----------------------------------------------------- 3 Exercícios  ----------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------
/*
Problems
What is whitespace?
Novas linhas, espaços e tabs entre as linhas são consideradas whitespaces
Elas deixam o programa mais legível mas você não vê elas, apesar disso, se não forem utilizadas, o programa ainda vai funcionar

What is a comment? What are the two ways of writing a comment?
O compilador do Go ignora os comentários.
É possível declarar comentários de uma linha e de múltiplas linhas.
// -> uma linha
/ * * / -> múltiplas linhas

Our program began with package main. What would the files in the fmt package begin with?
É necessário importar os pacotes, então começaria com import "fmt".

We used the Println function defined in the fmt package. If we wanted to use the Exit function from the os package what would we need to do?
Seria necessário chamar o pacote "os" e chamar a função os.Exit(3)

Modify the program we wrote so that instead of printing Hello World it prints Hello, my name is followed by your name.
Ok.

func main() {
	fmt.Println("Hello, my name is Patrícia")
	os.Exit(3)
}

*/
