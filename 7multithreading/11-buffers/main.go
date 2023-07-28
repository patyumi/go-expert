package main

func main() {
	// evitar isso para nao usar mais memoria
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"

	println(<-ch)
	println(<-ch)
}
