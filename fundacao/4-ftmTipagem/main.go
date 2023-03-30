package main

// O pacote fmt ou format vai ser usado para exibir resultados de acordo com o tipo certo
import "fmt"

var (
	ej float64 = 1.2
)

func main() {
	// %T exibe o tipo da variável

	fmt.Printf("O tipo é %T", ej)
}
