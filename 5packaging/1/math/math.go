package math

type Math struct {
	a int
	b int
}

// dessa maneira é possível utilizar propriedades privadas
// previne a modificação das props
func NewMath(a, b int) Math {
	return Math{a: a, b: b}
}

func (m Math) Add() int {
	return m.a + m.b
}
