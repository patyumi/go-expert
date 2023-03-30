package main

func main() {

	// memoria -> endereco -> valor
	// &endereco
	// *valor que esta guardado dentro da memoria
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20

	b := &a
	println(b)
	println(*b)

}

// quando usar os ponteiros
// quando vc quiser alterar o valor daquela variavel em outros contextos, outros lugares
// vc vai mudar o valor do endereço dela e não apenas da varzinha

// nào usar somente quando vc quiser fazer uma cópia dos dados, por exemplo fazer uma soma
// se eu quiser tornar os valores mutaveis (let) vc usa os ponteiros
