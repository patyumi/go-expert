package main

import (
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	// o defer é um statement
	// mas quando ele existe ele nào roda na hora, ele só funciona de fato quando o código acaba
	// entao se vc quer que algo seja executado por ultimo, vc usa o defer
	// ele sempre executa por ultimo
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))
}
