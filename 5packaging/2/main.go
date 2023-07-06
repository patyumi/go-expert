package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello")

	// go mod tidy para buscar todos os imports
	// go get para adicionar pacotes de maneira individual

	println(uuid.New().String())
}
