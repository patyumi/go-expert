package main

import (
	"fmt"
	"sync"
)

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

}

func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	// esse n precisa rodar em background pq quando o canal fecha ele ja sai e n le mais
	reader(ch, &wg)
	wg.Wait()
}
