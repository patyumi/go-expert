package main

import "fmt"

func main() {

	// o slice tem 3 partes
	// um ponteiro
	// tamanho - saber ate onde ir
	// capacidade - saber o quanto ele consegue receber

	// quando o colchete fica em branco, ou seja, ele não tem um valor definido ele é um slice
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// tudo que estiver a direita dos dois pontos : deve desaparecer, ser apagado
	// mesmo sem dados a capacidade dele não diminui
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// quando eu coloco os dois pontos na frente ele ignora tudo que ta no começo
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 12)

	// quando vc precisa dar um append no slice ele automaticamente dobra a capacidade dele, independente da quantidade que vc precisa
	// então se vc trabalha com algo muito grande, o ideal é inicializar ele com o valor próximo que vc vai usar
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
