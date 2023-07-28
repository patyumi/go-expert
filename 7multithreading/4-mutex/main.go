package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

// go run -race main.go
// verifica cenários de race condition

// como alternativa ao mutex vc pode usar somas atomicas tbm
// com pacote atomic
func main() {
	//m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// opção com mutex
		//m.Lock()
		//number++
		//m.Unlock()

		// opção com atomic
		atomic.AddUint64(&number, 1)

		w.Write([]byte(fmt.Sprintf("Você é o visitante #%d ", number)))
	})
	http.ListenAndServe(":3000", nil)
}
