package main

// Ao declarar uma variável o go infere uma variável dentro dela
// Essa declaração é global por que está fora de funções, ou seja, em escopo global
var (
	// Tipagem forte: Depois que uma variável nasce com um tipo você nunca mais consegue mudar ela
	// Entretando, em uma variável, é possível alterar o valor dela durante sua vida
	b bool = true
	c int
	d string
	e float64 = 1.2
)

func main() {
	// Dentro do escopo local, se você declara uma variável é obrigatório utilizá-la
	println(b)
	println(c)
	println(d)
	println(e)

	// ShortHand  ( := ) : você ao invés de declarar o tipo, você faz por inferência
	// mas ele é feito somente na primeira vez que a variável é criada, depois vc usa um =
	a := "Olá, Mundo!"
	println(a)
}
