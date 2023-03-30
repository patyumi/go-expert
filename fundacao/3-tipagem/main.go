package main

type ID int

var (
	f ID = 1
)

func main() {
	// O go tem suporte aos tipos básicos
	/*

		INTEGERS
		unsigned integers = contém apenas números positivos, incluindo o zero
		uint8 - byte (alias)
		uint16
		uint32
		uint64

		signed integers
		int8
		int16
		int32 - rune (alias)
		int64

		FLOAT - FLOATING POINT NUMBERS
		números que contém componentes decimais
		esses números são inexatos
		float32 (single precision)
		float64 (double precision) - quanto maior o tamanho do float, maior vai ser sua exatidão
		complex64
		complex128

		geralmente se usa o float64

		STRINGS
		"string" - não aceita novas linhas, aceita escapes \n \t
		`string` - aceita novas linhas

		é possível acessar o byte da string usando notação de array [1], toda string começa no 0 também
		+ concatenação

		BOOLEANS
		No Go você pode usar os operadores com valores booleanos:
		||
		&&
		!
	*/

	// Mas é possível voce criar o próprio tipo
	println(f)

}
