package main

import "fmt"

type Cliente struct {
	saldo int
}

func NewCliente() *Cliente {
	return &Cliente{saldo: 0}
}

func (c *Cliente) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	pat := Cliente{
		saldo: 200,
	}
	pat.simular(100)
	fmt.Printf("\nO valor da struct com nome %v", pat.saldo)
}
