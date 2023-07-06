package main

import (
	"fmt"

	"github.com/patyumi/go-expert/5packaging/3/math"
)

func main() {
	// workspaces
	// go work init ./math ./sistema
	// isso fica apenas localmente e voce pode adicionar a go.work no git ignore

	// go mod tidy -e baixa tudo que não tiver erro

	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
