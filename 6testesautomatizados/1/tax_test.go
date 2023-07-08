package tax

import "testing"

// rodar o teste go test .
// go test -v
// go test --coverprofile=coverage.out
// para ver onde o teste n ta indo 100 %
// go tool cover --html=coverage.out

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{0, 0.0},
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

// benchmark
// usa a letra b, ele é diferente do teste normal
// go test --bench=.
// go test --bench=. --run=^#   roda somente o benchmark
//go test --bench=. --run=^# --count=10 --benchtime=3s --benchmem

// mostra quantas operacoes vc consegue rodar por nanosegundo
// BenchmarkCalculateTax-12        1000000000               0.2563 ns/op
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// fuzzing
// gera coisas aleatórias pra testar pra mim, tenta varias vezes para quebrar meu teste
// go test --fuzz=. --run=^# --fuzztime=5s
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1500.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if amount >= 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})

}
