package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// o contexto controla o que a aplicaco ta fazendo
	// ele permite que a gente guarde informações no contexto
	// para que a gente consiga resgatar em todas as partes da aplicação

	// o context por convencao deve ser sempre o primeiro parametro a ser passado

	// o contexto sempre inicia vazio
	ctx := context.Background()

	// quando eu quero colocar uma regra nele é assim
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	// select é um case que funciona assincrono
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
