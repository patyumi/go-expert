package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		// go build -o cep main.go
		// url = go run main.go http://viacep.com.br/ws/32675750/json/
		req, err := http.Get("http://viacep.com.br/ws/" + url + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição %v", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			panic(err)
		}

		file, err := os.Create("cidade.txt")
		if err != nil {
			panic(err)
		}

		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("Cep: %s", data.Cep))
		if err != nil {
			panic(err)
		}

	}
}
