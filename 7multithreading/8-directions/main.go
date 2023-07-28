package main

import "fmt"

// canal<- canal apenas recebe infos
func recebe(nome string, hello chan<- string) {
	hello <- nome

}

// <-canal esvazia o canal
func ler(data <-chan string) {
	fmt.Println(<-data)

}

func main() {

	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
