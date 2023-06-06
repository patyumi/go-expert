package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// criar um arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// escrever qualquer dado vc manda em formato de byte
	//tamanho, err := f.Write([]byte()"Hello world"))

	// escrever string, quando eu sei que é uma string de fato
	tamanho, err := f.WriteString("Hello world")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	// leitura de arquivos
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// é necessario transformar em arquivo porque ele vem em formato de slice de byte
	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo | casos onde o arquivo é muito grande
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	// pra isso a gente define um buffer que tem o tamanho que eu consigo ler
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// remover o arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
