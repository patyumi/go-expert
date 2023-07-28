package main

import "time"

func main() {
	// o select é como se fosse um switch para canais

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		c2 <- 2
	}()

	for i := 0; i < 2; i++ {
		// Aqui quem chega primeiro é executado
		select {
		case msg1 := <-c1:
			println("received", msg1)

		case msg2 := <-c2:
			println("received", msg2)

		case <-time.After(time.Second * 3):
			println("timeout")

		default:
			println("default")
		}
	}

}
