package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	// handle func
	// o mux me permite rodar varios servidores ao mesmo tempo
	// onde cada mux implementa um server que vc quer
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)

	// handle \/

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

// Esse método ServeHTTP é algo obrigatório para implementar interfaces
// Quando uso o handle
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
