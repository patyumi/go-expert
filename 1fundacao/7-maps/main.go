package main

import "fmt"

func main() {
	// maps sao hashttables
	salarios := map[string]int{"wesley": 100, "pat": 200}
	fmt.Println(salarios["wesley"])
	delete(salarios, "wesley")
	salarios["wes"] = 5000
	fmt.Println(salarios["wes"])

	sal := make(map[string]int)
	sal["pat"] = 198
	fmt.Println(sal["pat"])

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}
}
