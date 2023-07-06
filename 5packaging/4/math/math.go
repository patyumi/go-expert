package math

// se esse pacote vai ser uma aplicação voce pode criar um modulo somente para ele
// isso o torna um sistema independente

type Math struct {
	a int
	b int
}

func NewMath(a, b int) Math {
	return Math{a: a, b: b}
}

func (m Math) Add() int {
	return m.a + m.b
}
