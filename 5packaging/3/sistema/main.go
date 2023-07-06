package main

import (
	"fmt"

	"github.com/patyumi/go-expert/5packaging/3/math"
)

func main() {
	// so e possivel usar pacotes publicados
	// mas tem como contornar isso

	// passando a url relativa
	// go mod edit -replace github.com/patyumi/go-expert/5packaging/3/math=../math
	// isso não funciona bem após publicado ... gambis

	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
