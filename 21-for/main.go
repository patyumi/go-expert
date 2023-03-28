package main

func main() {
	// no go só tem uma estrutura de loop, que é o for

	for i := 0; i < 10; i++ {
		println(i + 1)
	}

	numeros := []string{"um", "dois", "tres"}

	for indice, valor := range numeros {
		println(indice, valor)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// loop infinito, supondo que vc queira trabalhar com uma fila
	for {
		println("Hello World")
	}
}
