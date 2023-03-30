package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int `json:"s"` //quando usa o - , quer dizer que está omitindo esse campo
}

func main() {
	conta := Conta{Numero: 1, Saldo: 200}
	// transformar para json
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	// ele sempre me retorna o valor em byte
	// para ler um byte é só transformar em string
	println(res)
	println(string(res))

	// o encoder ja traz o json parseado em string para vc ler
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"Numero": 2, "s": 200}`)
	var contax Conta
	// aqui eu faço o inverso, pego um json e faço voltar em struct
	// nesse caso como as variaveis sao locais e necessario alterar sempre no local da memoria
	err = json.Unmarshal(jsonPuro, &contax)
	if err != nil {
		panic(err)
	}
	println(contax.Saldo)
}
