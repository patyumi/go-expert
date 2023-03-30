package main

// o generics nesse exemplo parte do principio de que vc consegue reaproveitar uma função que usa valores similares
// por exemplo, se eu quiser usar float dentro da função abaixo, teria que criar outra
// e ai voce usa essa notação de [] antes dos parametros da func

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// constraints especiais que já existem: any, comparable
// https://pkg.go.dev/golang.org/x/exp/constraints
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Pat": 10, "Lipe": 5}
	n := map[string]float64{"Pat": 10.5, "Lipe": 5.5}
	println(Soma(m))
	println(Soma(n))
	println(Compara(10, 10))
}

// voce tambem pode usar constraint que sao tipos especificos

// esse ~ significa que vc vai usar um tipo criado por voce mas que no fim das contas tem esse mesmo tipo
// exemplo type MyNumber int

type Number interface {
	~int | float64
}

func Soma2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}
