package main

func main() {
	// ele segura a execução
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever // esperando o canal ficar cheio para esvaziar
}
